package main

import (
	"fmt"

	"github.com/cloudwego/hertz/pkg/app/server"
	r "github.com/murInJ/go-pic-bed/router"
	u "github.com/murInJ/go-pic-bed/utils"
)

func main() {
	config, err := u.LoadConfigJSON("config.json")
	if err != nil {
		panic(err)
	}
	h := server.Default(server.WithHostPorts(fmt.Sprintf("%s:%d", config.Server.Host, config.Server.Port)), server.WithMaxRequestBodySize(20<<20))
	r.InitRouter(h)

	h.Spin()
}
