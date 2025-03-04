package service

import (
	"context"

	"github.com/cloudwego/hertz/pkg/app"
	common "github.com/meilingluolingluo/gomall/app/frontend/hertz_gen/frontend/common"
)

type HomeService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewHomeService(Context context.Context, RequestContext *app.RequestContext) *HomeService {
	return &HomeService{RequestContext: RequestContext, Context: Context}
}

func (h *HomeService) Run(req *common.Empty) (res map[string]any, err error) {
	var resp = make(map[string]any)
	items := []map[string]any{
		{"Name": "T-shirt-1", "Price": 100, "Picture": "/static/image/Picture001.png"},
		{"Name": "T-shirt-1", "Price": 100, "Picture": "/static/image/Picture006.png"},
		{"Name": "T-shirt-1", "Price": 100, "Picture": "/static/image/Picture009.png"},
		{"Name": "T-shirt-1", "Price": 100, "Picture": "/static/image/Picture012.png"},
	}
	resp["Title"] = "Hot Sales"
	resp["items"] = items
	return resp, nil
}
