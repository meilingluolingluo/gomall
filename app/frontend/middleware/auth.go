package middleware

import (
	"context"
	"fmt"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/hertz-contrib/sessions"
	frontendUtils "github.com/meilingluolingluo/gomall/app/frontend/utils"
	"net/http"
)

type SessionUserIdKey string

const SessionUserId SessionUserIdKey = "user_id"

func GlobalAuth() app.HandlerFunc {
	return func(ctx context.Context, c *app.RequestContext) {
		session := sessions.Default(c)
		userId := session.Get("user_id")
		if userId == nil {
			c.Next(ctx)
			return
		}
		ctx = context.WithValue(ctx, frontendUtils.SessionUserId, userId)
		c.Next(ctx)
	}
}

func Auth() app.HandlerFunc {
	return func(ctx context.Context, c *app.RequestContext) {
		session := sessions.Default(c)
		userId := session.Get("user_id")
		if userId == nil {
			ref := string(c.GetHeader("Referer"))
			next := "/sign-in"
			if ref != "" && frontendUtils.ValidateNext(ref) {
				next = fmt.Sprintf("/sign-in?next=%s", ref)
			}
			c.Redirect(http.StatusFound, []byte(next)) // 使用 http.StatusFound (302)
			c.Abort()
			return
		}
		ctx = context.WithValue(ctx, SessionUserId, userId)
		c.Next(ctx)
	}
}
