package user

import (
	"context"
	user "github.com/meilingluolingluo/gomall/rpc_gen/kitex_gen/user"

	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/client/callopt"
	"github.com/meilingluolingluo/gomall/rpc_gen/kitex_gen/user/userservice"
)

type RPCClient interface {
	KitexClient() userservice.Client
	Service() string
	Register(ctx context.Context, Req *user.RegisterReq, callOptions ...callopt.Option) (r *user.RegisterResp, err error)
	Login(ctx context.Context, Req *user.LoginReq, callOptions ...callopt.Option) (r *user.LoginResp, err error)
	Delete(ctx context.Context, Req *user.DeleteReq, callOptions ...callopt.Option) (r *user.DeleteResp, err error)
	Modify(ctx context.Context, Req *user.ModifyReq, callOptions ...callopt.Option) (r *user.ModifyResp, err error)
}

func NewRPCClient(dstService string, opts ...client.Option) (RPCClient, error) {
	kitexClient, err := userservice.NewClient(dstService, opts...)
	if err != nil {
		return nil, err
	}
	cli := &clientImpl{
		service:     dstService,
		kitexClient: kitexClient,
	}

	return cli, nil
}

type clientImpl struct {
	service     string
	kitexClient userservice.Client
}

func (c *clientImpl) Service() string {
	return c.service
}

func (c *clientImpl) KitexClient() userservice.Client {
	return c.kitexClient
}

func (c *clientImpl) Register(ctx context.Context, Req *user.RegisterReq, callOptions ...callopt.Option) (r *user.RegisterResp, err error) {
	return c.kitexClient.Register(ctx, Req, callOptions...)
}

func (c *clientImpl) Login(ctx context.Context, Req *user.LoginReq, callOptions ...callopt.Option) (r *user.LoginResp, err error) {
	return c.kitexClient.Login(ctx, Req, callOptions...)
}

func (c *clientImpl) Delete(ctx context.Context, Req *user.DeleteReq, callOptions ...callopt.Option) (r *user.DeleteResp, err error) {
	return c.kitexClient.Delete(ctx, Req, callOptions...)
}

func (c *clientImpl) Modify(ctx context.Context, Req *user.ModifyReq, callOptions ...callopt.Option) (r *user.ModifyResp, err error) {
	return c.kitexClient.Modify(ctx, Req, callOptions...)
}
