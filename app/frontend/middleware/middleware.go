package middleware

import "github.com/cloudwego/hertz/pkg/app/server"

func RegisterMiddleware(h *server.Hertz, jwtSecret string) {
	h.Use(GlobalAuth())

	h.Use(JWTAuthMiddleware(jwtSecret))
}
