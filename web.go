package goseed

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
)

type response map[string]interface{}

type HttpError struct {
	Code int
	Msg  string
}

func (err HttpError) Error() string {
	return err.Msg
}

// 捕获异常，并返回 HTTP Error
func catch(ctx *gin.Context) {
	if e := recover(); e != nil {
		var code int
		var msg string

		switch e.(type) {
		// 如果是自定义 HttpError ，则获取状态码和错误信息
		case HttpError:
			msg = e.(HttpError).Msg
			code = e.(HttpError).Code
			break
		// 如果是普通的 error， 则转换为 error
		case error:
			msg = e.(error).Error()
			code = 500
			break
		}
		resp := getResponse(code, msg, nil)
		ctx.JSON(code, resp)
	}
}

// 用于包装用户定义的 Handler，提供通过 panic 返回 HTTP Error 的功能
func WrapHandler(f func(ctx *gin.Context) interface{}) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		defer catch(ctx)
		result := f(ctx)
		resp := getResponse(200, nil, result)
		ctx.JSON(200, resp)
	}
}

// 处理跨域的中间件
func CorsHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		// 设置跨域响应头
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Headers", "Content-Type,AccessToken,X-CSRF-Token, Authorization, Token")
		c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS")
		c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Content-Type")
		c.Header("Access-Control-Allow-Credentials", "true")
		// 放行所有的 OPTIONS 方法的请求
		if method == "OPTIONS" {
			c.JSON(http.StatusOK, "Options Request!")
		}
		c.Next()
	}
}

// 抛出一个 HTTP 错误
func AbortErr(ctx *gin.Context, code int, msg string) {
	resp := getResponse(code, msg, nil)
	ctx.AbortWithStatusJSON(code, resp)
}

// 抛出一个 HTTP 错误
func ThrowsErr(code int, err error) {
	panic(HttpError{Code: code, Msg: err.Error()})
}

// 将 POST 方法的 JSON Body 解析为相对应的结构体
func ParseJSON(ctx *gin.Context, obj interface{}) {
	err := ctx.BindJSON(obj)
	if err != nil {
		ThrowsErr(400, err)
	}
}

// 获取 POST 方法中 JSON Body 里的数据
func GetJsonBody(ctx *gin.Context, key string) (interface{}, error) {
	data, _ := ioutil.ReadAll(ctx.Request.Body)
	m := make(map[string]interface{})
	err := json.Unmarshal(data, &m)
	if err != nil {
		return nil, err
	}
	return m[key], nil
}

func getResponse(code int, msg interface{}, data interface{}) response {
	return response{
		"code":    code,
		"message": msg,
		"data":    data,
	}
}
