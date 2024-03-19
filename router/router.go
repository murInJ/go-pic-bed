package router

import (
	"github.com/cloudwego/hertz/pkg/app/server"
)

func InitRouter(h *server.Hertz) {
	InitPingRouter(h)
	InitFSRouter(h)

}
