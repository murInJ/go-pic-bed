package router

import (
	"context"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/app/server"
	handler "github.com/murInJ/go-pic-bed/handler"
)

func IinitUploadRouter(h *server.Hertz) {
	h.POST("/singleFile", func(ctx context.Context, c *app.RequestContext) {
		handler.SingleFile(c)
	})

	h.POST("/multiFile", func(ctx context.Context, c *app.RequestContext) {
		handler.MultiFile(c)
	})
}
