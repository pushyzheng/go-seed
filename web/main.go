package web

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

type Map map[string]interface{}

type HttpError struct {
	Code    int
	Message string
}

func (err HttpError) Error() string {
	return fmt.Sprintf("The error code is %d, message is %s", err.Code, err.Message)
}

// cache panic error and return error response
func catch(ctx *gin.Context) {
	if e := recover(); e != nil {
		msg := e.(HttpError).Message
		code := e.(HttpError).Code
		resp := Map{
			"code":    code,
			"message": msg,
			"data":    nil,
		}
		ctx.JSON(code, resp)
	}
}

func WrapHandle(f func(ctx *gin.Context) interface{}) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		defer catch(ctx)
		result := f(ctx)
		ctx.JSON(200, result)
	}
}
