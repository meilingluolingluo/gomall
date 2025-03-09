package service

import (
	"context"
	"fmt"
	"log"
	"time"

	jwt "github.com/golang-jwt/jwt/v5"
	"github.com/hertz-contrib/sessions"
	"github.com/meilingluolingluo/gomall/app/frontend/hertz_gen/frontend/common"
	"github.com/meilingluolingluo/gomall/app/frontend/infra/rpc"
	"github.com/meilingluolingluo/gomall/rpc_gen/kitex_gen/user"

	"github.com/cloudwego/hertz/pkg/app"
	auth "github.com/meilingluolingluo/gomall/app/frontend/hertz_gen/frontend/auth"
)

type RegisterService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewRegisterService(Context context.Context, RequestContext *app.RequestContext) *RegisterService {
	return &RegisterService{RequestContext: RequestContext, Context: Context}
}

func (h *RegisterService) Run(req *auth.RegisterReq) (resp *common.Empty, err error) {
	/*
		// 检查邮箱是否已注册
		_, err = model.GetByEmail(mysql.DB, req.Email)
		if err == nil {
			return nil, errors.New("email already registered")
		}
	*/
	/*
		if req.Password != req.ConfirmPassword {
			return nil, errors.New("password and confirm password do not match")
		}
	*/
	// 创建用户
	userResp, err := rpc.UserClient.Register(h.Context, &user.RegisterReq{
		Username:        req.Username,
		Email:           req.Email,
		Password:        req.Password,
		ConfirmPassword: req.ConfirmPassword,
	})
	if err != nil {
		log.Printf("Error Register: %v", err)
		return nil, err
	}
	println("userResp:", userResp)

	log.Printf("Error")

	// 设置session
	session := sessions.Default(h.RequestContext)
	session.Set("user_id", userResp.UserId)
	session.Set("username", userResp.Username)
	err = session.Save()
	if err != nil {
		log.Printf("Error saving session: %v", err)
		return nil, err
	}

	return
}
func generateJWT(userID int32, secretKey string) (string, error) {
	claims := jwt.RegisteredClaims{
		Subject:   fmt.Sprintf("%d", userID),
		IssuedAt:  jwt.NewNumericDate(time.Now()),
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)), // 设置token过期时间
		// ExpiresAt: jwt.NewNumericDate(time.Now().Add(1 * time.Minute)), // 设置token过期时间
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, err := token.SignedString([]byte(secretKey))
	if err != nil {
		return "", fmt.Errorf("failed to sign token: %v", err)
	}
	return ss, nil
}
