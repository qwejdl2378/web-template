package httputil

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type HttpSuccess struct {
	Success bool        `json:"success" example:"true"`
	Code    int         `json:"code" example:"200"`
	Message string      `json:"message" example:"create book item success"`
	Data    interface{} `json:"data"`
}

type HTTPError struct {
	Success bool   `json:"success" example:"false"`
	Code    int    `json:"code" example:"400"`
	Message string `json:"message" example:"status bad request"`
}

func success(ctx *gin.Context, msg string, data interface{}) {
	body := HttpSuccess{
		Success: true,
		Code:    200,
		Message: msg,
		Data:    data,
	}
	ctx.JSON(http.StatusOK, body)
	return
}

func ReturnSuccess(ctx *gin.Context, data interface{}) {
	success(ctx, "", data)
	return
}

func ReturnSuccessWithMsg(ctx *gin.Context, msg string, data interface{}) {
	success(ctx, msg, data)
}

func ThrowError(ctx *gin.Context, status int, err error) {
	er := HTTPError{
		Success: false,
		Code:    status,
		Message: err.Error(),
	}
	ctx.AbortWithStatusJSON(status, er)
	return
}
