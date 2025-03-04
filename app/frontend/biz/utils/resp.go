package utils

import (
	"context"
	"github.com/hertz-contrib/sessions"
	frontendutils "github.com/meilingluolingluo/gomall/app/frontend/utils"

	"github.com/cloudwego/hertz/pkg/app"
)

// SendErrResponse  pack error response
func SendErrResponse(ctx context.Context, c *app.RequestContext, code int, err error) {
	// todo edit custom code
	c.String(code, err.Error())
}

// SendSuccessResponse  pack success response
func SendSuccessResponse(ctx context.Context, c *app.RequestContext, code int, data interface{}) {
	// todo edit custom code
	c.JSON(code, data)
}

func WarpResponse(ctx context.Context, c *app.RequestContext, content map[string]any) map[string]any {
	var cartNum int
	session := sessions.Default(c)
	//userId := frontendutils.GetUserIdFromCtx(ctx)

	content["user_id"] = ctx.Value(frontendutils.UserIdKey)
	Username := session.Get("user_name")
	content["user_name"] = Username
	content["cart_num"] = cartNum
	return content
}
