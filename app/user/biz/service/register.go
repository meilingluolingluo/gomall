package service

import (
	"context"
	"errors"
	"github.com/meilingluolingluo/gomall/app/user/biz/dal/mysql"
	"github.com/meilingluolingluo/gomall/app/user/biz/model"
	user "github.com/meilingluolingluo/gomall/rpc_gen/kitex_gen/user"
	"golang.org/x/crypto/bcrypt"
)

type RegisterService struct {
	ctx context.Context
} // NewRegisterService new RegisterService
func NewRegisterService(ctx context.Context) *RegisterService {
	return &RegisterService{ctx: ctx}
}

// Run create note info
func (s *RegisterService) Run(req *user.RegisterReq) (resp *user.RegisterResp, err error) {
	// Finish your business logic.
	// 参数校验
	if req.Email == "" || req.Password == "" || req.ConfirmPassword == "" || req.Username == "" {
		return nil, errors.New("email or password is empty")
	}
	println("username", req.Username)
	// 验证密码正确性
	if req.Password != req.ConfirmPassword {
		return nil, errors.New("password not match")
	}
	// 对密码进行哈希加密
	passwordHashed, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}
	// 创建用户信息
	newUser := &model.User{
		Username:       req.Username,
		Email:          req.Email,
		PasswordHashed: string(passwordHashed),
	}
	err = model.Create(mysql.DB, newUser)
	if err != nil {
		return nil, err
	}

	return &user.RegisterResp{UserId: uint32(newUser.ID), Username: newUser.Username}, nil
}
