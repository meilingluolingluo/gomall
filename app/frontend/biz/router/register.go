// Code generated by hertz generator. DO NOT EDIT.

package router

import (
	"github.com/cloudwego/hertz/pkg/app/server"
	auth "github.com/meilingluolingluo/gomall/app/frontend/biz/router/auth"
	cart "github.com/meilingluolingluo/gomall/app/frontend/biz/router/cart"
	category "github.com/meilingluolingluo/gomall/app/frontend/biz/router/category"
	checkout "github.com/meilingluolingluo/gomall/app/frontend/biz/router/checkout"
	home "github.com/meilingluolingluo/gomall/app/frontend/biz/router/home"
	product "github.com/meilingluolingluo/gomall/app/frontend/biz/router/product"
)

// GeneratedRegister registers routers generated by IDL.
func GeneratedRegister(r *server.Hertz) {
	//INSERT_POINT: DO NOT DELETE THIS LINE!
	checkout.Register(r)

	cart.Register(r)

	category.Register(r)

	product.Register(r)

	auth.Register(r)

	home.Register(r)
}
