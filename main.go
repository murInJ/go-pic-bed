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

	// h.POST("/singleFile", func(ctx context.Context, c *app.RequestContext) {
	// 	// single file
	// 	file, _ := c.FormFile("file")
	// 	fmt.Println(file.Filename)

	// 	// Upload the file to specific dst
	// 	c.SaveUploadedFile(file, fmt.Sprintf("./file/upload/%s", file.Filename))

	// 	c.String(consts.StatusOK, fmt.Sprintf("'%s' uploaded!", file.Filename))
	// })

	// h.POST("/multiFile", func(ctx context.Context, c *app.RequestContext) {
	// 	// Multipart form
	// 	form, _ := c.MultipartForm()
	// 	files := form.File["file"]

	// 	for _, file := range files {
	// 		fmt.Println(file.Filename)

	// 		// Upload the file to specific dst.
	// 		c.SaveUploadedFile(file, fmt.Sprintf("./file/upload/%s", file.Filename))
	// 	}
	// 	c.String(consts.StatusOK, fmt.Sprintf("%d files uploaded!", len(files)))
	// })

	h.Spin()
}
