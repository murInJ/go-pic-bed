package middleware

import (
	"github.com/cloudwego/hertz/pkg/app/server"
	_ "github.com/murInJ/go-pic-bed/docs"
)

func InitMiddleware(h *server.Hertz) {
	h.Use(Cors())
}
