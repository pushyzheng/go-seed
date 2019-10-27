package goseed

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

type response map[string]interface{}

type HttpError struct {
	Code int
	Msg  string
}

func (err HttpError) Error() string {
	return fmt.Sprintf("The error code is %d, message is %s", err.Code, err.Msg)
}

// catch panic error and return error response
func catch(ctx *gin.Context) {
	if e := recover(); e != nil {
		msg := e.(HttpError).Msg
		code := e.(HttpError).Code
		resp := response{
			"code":    code,
			"message": msg,
			"data":    nil,
		}
		ctx.JSON(code, resp)
	}
}

// wrap to gin.HandlerFunc
func WrapHandle(f func(ctx *gin.Context) interface{}) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		defer catch(ctx)
		result := f(ctx)
		resp := response{
			"code":    200,
			"message": nil,
			"data":    result,
		}
		ctx.JSON(200, resp)
	}
}
