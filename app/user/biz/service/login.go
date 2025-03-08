package service

import (
	"context"
	"errors"
	"fmt"

	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/meilingluolingluo/gomall/app/user/biz/dal/mysql"
	"github.com/meilingluolingluo/gomall/app/user/biz/model"
	user "github.com/meilingluolingluo/gomall/rpc_gen/kitex_gen/user"
	"golang.org/x/crypto/bcrypt"
)

type LoginService struct {
	ctx context.Context
} // NewLoginService new LoginService
func NewLoginService(ctx context.Context) *LoginService {
	return &LoginService{ctx: ctx}
}

// Run create note info
func (s *LoginService) Run(req *user.LoginReq) (resp *user.LoginResp, err error) {
	// 邮箱非空校验
	if req.Email == "" {
		return nil, errors.New("邮箱不能为空")
	}

	// 密码非空校验
	if req.Password == "" {
		return nil, errors.New("密码不能为空")
	}

	klog.Infof("LoginReq:%+v", req)
	userRow, err := model.GetByEmail(mysql.DB, s.ctx, req.Email)
	if err != nil {
		// 如果用户不存在或数据库错误，返回相应的错误信息
		return nil, fmt.Errorf("用户不存在: %v", err)
	}

	err = bcrypt.CompareHashAndPassword([]byte(userRow.PasswordHashed), []byte(req.Password))
	if err != nil {
		// 密码不匹配的情况
		return nil, errors.New("密码错误")
	}

	// fmt.Printf("[DEBUG] Login Success: user_id=%v, user_name=%v\n", int32(userRow.ID), userRow.Username)
	// 返回新增的 username 字段
	return &user.LoginResp{
		UserId:   uint32(userRow.ID),
		Username: userRow.Username,
	}, nil

}
