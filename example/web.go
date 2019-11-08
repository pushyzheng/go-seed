package example

import (
	"fmt"
	"github.com/gin-gonic/gin"
	goseed "powhole.com/go-seed"
)

func main() {
	app := gin.Default()
	app.Use(goseed.CorsHandler())

	// 必须用 WrapHandler 包装方法
	app.GET("/", goseed.WrapHandler(Index))
}

func Index(ctx *gin.Context) interface{} {
	if ctx.Query("name") == "" {
		// 抛出异常，返回客户端 HTTP 400 错误
		goseed.ThrowsErr(400, fmt.Errorf("%s", "Error"))
	}
	// 正常返回数据
	return "Hello World"
}
