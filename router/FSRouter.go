package router

import (
	"context"
	"time"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	u "github.com/murInJ/go-pic-bed/utils"
)

func InitFSRouter(h *server.Hertz) {
	h.StaticFS(u.Conf.FS.Path, &app.FS{
		Root:        u.Conf.FS.Root,
		PathRewrite: app.NewPathSlashesStripper(1),
		PathNotFound: func(_ context.Context, ctx *app.RequestContext) {
			ctx.JSON(consts.StatusNotFound, "The requested resource does not exist")
		},
		CacheDuration:        time.Second * 5,
		Compress:             true,
		CompressedFileSuffix: "picBed",
		AcceptByteRange:      true,
	})
}
