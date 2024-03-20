package middleware

import (
	"time"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/hertz-contrib/cors"
	_ "github.com/murInJ/go-pic-bed/docs"
)

func Cors() app.HandlerFunc {
	return cors.New(cors.Config{
		// 允许所有来源（ "*" 表示任意域名）
		AllowOrigins: []string{"*"},
		// 允许所有HTTP方法
		AllowMethods: []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		// 允许常见的请求头，例如Origin、Content-Type等
		AllowHeaders: []string{"Origin", "Content-Type", "Accept", "Authorization"},
		// 暴露给客户端的响应头
		ExposeHeaders: []string{"Content-Length", "Content-Type"},
		// 允许携带凭据（如Cookies）
		AllowCredentials: true,
		// 预检请求（OPTIONS）的缓存时间，默认可设置为较长时间
		MaxAge: 12 * time.Hour,
	})
}
