package main

import (
	"fmt"

	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/cloudwego/hertz/pkg/app/server/registry"
	"github.com/cloudwego/hertz/pkg/common/utils"
	"github.com/hertz-contrib/registry/etcd"
	"github.com/hertz-contrib/swagger"
	_ "github.com/murInJ/go-pic-bed/docs"
	m "github.com/murInJ/go-pic-bed/middleware"
	r "github.com/murInJ/go-pic-bed/router"
	u "github.com/murInJ/go-pic-bed/utils"
	swaggerFiles "github.com/swaggo/files"
)

// @title go-pic-bed
// @version 0.0.1
// @description go pic bed

// @contact.name murInJ
// @contact.url https://github.com/murInJ/go-pic-bed

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8080
// @BasePath /
// @schemes http
func main() {
	var h *server.Hertz
	_, err := u.LoadConfigJSON("config.json")
	if err != nil {
		panic(err)
	}
	addr := fmt.Sprintf("%s:%d", u.Conf.Server.Host, u.Conf.Server.Port)
	if u.Conf.Registry.Able {
		r, err := etcd.NewEtcdRegistry([]string{"127.0.0.1:2379"},
			etcd.WithAuthOpt("root", "123456"),
		)
		if err != nil {
			panic(err)
		}
		h = server.Default(
			server.WithHostPorts(addr),
			server.WithRegistry(r, &registry.Info{
				ServiceName: "hertz.test.demo",
				Addr:        utils.NewNetAddr("tcp", addr),
				Weight:      10,
				Tags:        nil,
			}))
	} else {
		h = server.Default(server.WithHostPorts(addr), server.WithMaxRequestBodySize(20<<20))
	}
	r.InitRouter(h)
	m.InitMiddleware(h)

	url := swagger.URL(fmt.Sprintf("http://%s/swagger/doc.json", addr))
	h.GET("/swagger/*any", swagger.WrapHandler(swaggerFiles.Handler, url))

	h.Spin()
}
