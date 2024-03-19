package handler

import (
	"fmt"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/utils"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	s "github.com/murInJ/go-pic-bed/service"
	u "github.com/murInJ/go-pic-bed/utils"
)

func SingleFile(c *app.RequestContext) {
	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(consts.StatusBadRequest, utils.H{
			"msg": err.Error(),
		})
		return
	}
	name := s.GenerateFileIDName(file.Filename)
	c.SaveUploadedFile(file, fmt.Sprintf("%s/%s", u.Conf.FS.Root, name))
	c.JSON(consts.StatusOK, utils.H{
		"msg": fmt.Sprintf("'%s' uploaded!", file.Filename),
		"url": fmt.Sprintf("%s/%s", u.Conf.FS.Path, name),
		"cnt": 1,
	})
}

func MultiFile(c *app.RequestContext) {
	// Multipart form
	form, err := c.MultipartForm()
	if err != nil {
		c.JSON(consts.StatusBadRequest, utils.H{
			"msg": err.Error(),
		})
		return
	}
	files := form.File["file"]
	var urls []string
	var fileNames []string
	for _, file := range files {
		// Upload the file to specific dst.
		name := s.GenerateFileIDName(file.Filename)
		fileNames = append(fileNames, file.Filename)
		c.SaveUploadedFile(file, fmt.Sprintf("%s/%s", u.Conf.FS.Root, name))
		urls = append(urls, fmt.Sprintf("%s/%s", u.Conf.FS.Path, name))
	}
	c.String(consts.StatusOK, fmt.Sprintf("%d files uploaded!", len(files)))
	c.JSON(consts.StatusOK, utils.H{
		"msg":  fmt.Sprintf("'%v' uploaded!", fileNames),
		"urls": urls,
		"cnt":  len(files),
	})
}
