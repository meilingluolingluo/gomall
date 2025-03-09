package service

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/hertz-contrib/sessions"
	auth "github.com/meilingluolingluo/gomall/app/frontend/hertz_gen/frontend/auth"
	"github.com/meilingluolingluo/gomall/app/frontend/infra/rpc"
	frontendutils "github.com/meilingluolingluo/gomall/app/frontend/utils"
	"github.com/meilingluolingluo/gomall/rpc_gen/kitex_gen/user"
)

type LoginService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewLoginService(Context context.Context, RequestContext *app.RequestContext) *LoginService {
	return &LoginService{RequestContext: RequestContext, Context: Context}
}

func (h *LoginService) Run(req *auth.LoginReq) (resp string, err error) {
	res, err := rpc.UserClient.Login(h.Context, &user.LoginReq{Email: req.Email, Password: req.Password})
	if err != nil {
		return
	}

	session := sessions.Default(h.RequestContext)
	session.Set("user_id", res.UserId)
	session.Set("username", res.Username)
	err = session.Save()
	frontendutils.MustHandleError(err)
	redirect := "/"
	if frontendutils.ValidateNext(req.Next) {
		redirect = req.Next
	}
	if err != nil {
		return "", err
	}

	return redirect, nil
}
