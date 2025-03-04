package user

import (
	"context"
	"github.com/cloudwego/kitex/client/callopt"
	"github.com/cloudwego/kitex/pkg/klog"
	user "github.com/meilingluolingluo/gomall/rpc_gen/kitex_gen/user"
)

func Register(ctx context.Context, req *user.RegisterReq, callOptions ...callopt.Option) (resp *user.RegisterResp, err error) {
	resp, err = defaultClient.Register(ctx, req, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "Register call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}

func Login(ctx context.Context, req *user.LoginReq, callOptions ...callopt.Option) (resp *user.LoginResp, err error) {
	resp, err = defaultClient.Login(ctx, req, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "Login call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}

func GetPersonal(ctx context.Context, req *user.GetPersonalReq, callOptions ...callopt.Option) (resp *user.GetPersonalResp, err error) {
	resp, err = defaultClient.GetPersonal(ctx, req, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "GetPersonal call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}

func UpdatePwd(ctx context.Context, req *user.UpdatePwdReq, callOptions ...callopt.Option) (resp *user.UpdatePwdResp, err error) {
	resp, err = defaultClient.UpdatePwd(ctx, req, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "UpdatePwd call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}

func Delete(ctx context.Context, req *user.DeleteReq, callOptions ...callopt.Option) (resp *user.DeleteResp, err error) {
	resp, err = defaultClient.Delete(ctx, req, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "Delete call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}
