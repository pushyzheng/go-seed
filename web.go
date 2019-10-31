package goseed

import (
	"github.com/gin-gonic/gin"
)

type response map[string]interface{}

type HttpError struct {
	Code int
	Msg  string
}

func (err HttpError) Error() string {
	return err.Msg
}

// catch panic error and return error response
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

// wrap to gin.HandlerFunc
func WrapHandle(f func(ctx *gin.Context) interface{}) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		defer catch(ctx)
		result := f(ctx)
		resp := getResponse(200, nil, result)
		ctx.JSON(200, resp)
	}
}

func AbortErr(ctx *gin.Context, code int, msg string) {
	resp := getResponse(code, msg, nil)
	ctx.AbortWithStatusJSON(code, resp)
}

func ThrowsErr(code int, err error) {
	panic(HttpError{Code: code, Msg: err.Error()})
}

func ParseJSON(ctx *gin.Context, obj interface{}) {
	err := ctx.BindJSON(obj)
	if err != nil {
		ThrowsErr(400, err)
	}
}

func getResponse(code int, msg interface{}, data interface{}) response {
	return response{
		"code":    code,
		"message": msg,
		"data":    data,
	}
}
