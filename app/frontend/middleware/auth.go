// Copyright 2024 CloudWeGo Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package middleware

import (
	"context"
	"fmt"
	"strconv"
	"strings"
	"time"

	resp "github.com/cloudwego/biz-demo/gomall/app/frontend/biz/utils"
	utils "github.com/cloudwego/biz-demo/gomall/app/frontend/utils"
	"github.com/cloudwego/hertz/pkg/app"
	hertzUtils "github.com/cloudwego/hertz/pkg/common/utils"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"github.com/golang-jwt/jwt/v5"
	"github.com/hertz-contrib/sessions"
)

func GlobalAuth() app.HandlerFunc {
	return func(ctx context.Context, c *app.RequestContext) {
		session := sessions.Default(c)
		userId := session.Get("user_id")

		if userId == nil {
			c.Next(ctx)
			return
		}
		ctx = context.WithValue(ctx, utils.UserIdKey, userId)
		c.Next(ctx)
	}
}

func Auth() app.HandlerFunc {
	return func(ctx context.Context, c *app.RequestContext) {
		session := sessions.Default(c)
		userId := session.Get("user_id")
		if userId == nil {
			byteRef := c.GetHeader("Referer")
			ref := string(byteRef)
			next := "/sign-in"
			if ref != "" {
				if utils.ValidateNext(ref) {
					next = fmt.Sprintf("%s?next=%s", next, ref)
				}
			}
			c.Redirect(302, []byte(next))
			c.Abort()
			c.Next(ctx)
			return
		}
		ctx = context.WithValue(ctx, utils.UserIdKey, userId)
		c.Next(ctx)
	}
}

// JWTAuthMiddleware 使用传入的 secretKey 来验证 JWT
func JWTAuthMiddleware(secretKey string) app.HandlerFunc {
	return func(ctx context.Context, c *app.RequestContext) {
		// 将 []byte 类型的路径转换为 string 类型
		path := string(c.Path())

		// 跳过不需要认证的路径
		if shouldSkipAuth(path) {
			c.Next(ctx)
			return
		}

		// 尝试从cookie中获取token
		cookie := c.Request.Header.Cookie("access_token")
		var tokenString string
		if len(cookie) > 0 {
			tokenString = strings.TrimSpace(string(cookie))
		} else {
			// 如果cookie中没有找到token，则尝试从Authorization header中获取
			authHeader := strings.TrimSpace(string(c.GetHeader("Authorization")))
			if authHeader == "" {
				// c.JSON(401, map[string]string{"error": "Missing token"})
				wrappedResp := resp.WarpResponse(ctx, c, hertzUtils.H{"error": "用户未登录！"})
				session := sessions.Default(c)
				session.Clear()
				err := session.Save()
				utils.MustHandleError(err)
				c.HTML(consts.StatusOK, "sign-in", wrappedResp)
				c.Abort()
				return
			}

			parts := strings.SplitN(authHeader, " ", 2)
			if len(parts) != 2 || !strings.EqualFold(parts[0], "Bearer") {
				c.JSON(401, map[string]string{"error": "Invalid Authorization format"})
				c.Abort()
				return
			}
			tokenString = parts[1]
		}

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}
			return []byte(secretKey), nil
		})

		if err != nil || !token.Valid {
			c.JSON(401, map[string]string{"error": "Invalid token"})
			c.Abort()
			return
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok || claims["sub"] == nil {
			c.JSON(401, map[string]string{"error": "Invalid token claims"})
			c.Abort()
			return
		}

		userIdStr, ok := claims["sub"].(string)
		if !ok {
			c.JSON(401, map[string]string{"error": "Invalid user ID in token claims"})
			c.Abort()
			return
		}
		// fmt.Printf("User ID from token: %s\n", userIdStr)
		userId, err := strconv.ParseInt(userIdStr, 10, 32)
		if err != nil {
			// 添加详细错误信息以便于调试
			fmt.Printf("Error parsing userId: %v\n", err)
			c.JSON(401, map[string]string{"error": "Invalid user ID format"})
			c.Abort()
			return
		}

		// 校验 token 有效期
		if exp, ok := claims["exp"].(float64); ok && time.Now().Unix() > int64(exp) {
			// c.JSON(401, map[string]string{"error": "Token expired"})
			// wrappedResp := resp.WarpResponse(ctx, c, hertzUtils.H{"error": "用户身份已过期，请重新登陆！"})
			c.Redirect(consts.StatusFound, []byte("/"))
			c.Abort()
			return
		}

		ctx = context.WithValue(ctx, utils.UserIdKey, float64(userId))
		c.Next(ctx)
	}
}

// shouldSkipAuth 判断是否应该跳过认证
func shouldSkipAuth(path string) bool {
	// 定义不需要认证的路径
	if strings.HasPrefix(path, "/static/") {
		return true
	}
	publicPaths := []string{"/", "/sign-in", "/sign-up", "/ping", "/auth/login", "/auth/register", "/auth/logout", "/robots.txt"}
	for _, p := range publicPaths {
		if p == path {
			return true
		}
	}
	return false
}
