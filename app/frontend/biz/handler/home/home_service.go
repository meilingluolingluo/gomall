package home

import (
	"context"

	"github.com/meilingluolingluo/gomall/app/frontend/biz/utils"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"github.com/hertz-contrib/sessions"
	"github.com/meilingluolingluo/gomall/app/frontend/biz/service"
	common "github.com/meilingluolingluo/gomall/app/frontend/hertz_gen/frontend/common"
	frontendutils "github.com/meilingluolingluo/gomall/app/frontend/utils"
)

// Home .
// @router / [GET]
func Home(ctx context.Context, c *app.RequestContext) {
	var err error
	var req common.Empty
	err = c.BindAndValidate(&req)

	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	// resp, err :=

	resp, err := service.NewHomeService(ctx, c).Run(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}
	session := sessions.Default(c)

	updateSuccessMsg := session.Get("update_success")
	if updateSuccessMsg != nil {
		resp["update_success"] = true
		session.Delete("update_success") // 清除消息以避免重复显示
		err = session.Save()
		frontendutils.MustHandleError(err)
	}

	Username := session.Get("user_name")
	resp["user_name"] = Username

	c.HTML(consts.StatusOK, "home", utils.WarpResponse(ctx, c, resp))
}
