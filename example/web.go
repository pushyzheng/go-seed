package example

import (
	"github.com/gin-gonic/gin"
	goseed "powhole.com/go-seed"
)

func main() {
	app := gin.Default()
	// 必须用 WrapHandle 包装方法
	app.GET("/", goseed.WrapHandle(Index))
}

func Index(ctx *gin.Context) interface{} {
	if ctx.Query("name") == "" {
		// 抛出异常，返回客户端 HTTP 400 错误
		panic(goseed.HttpError{Code: 400, Msg: "Error"})
	}
	// 正常返回数据
	return "Hello World"
}
