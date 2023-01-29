package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type ResponseData struct {
	Code    MyCode      `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"` //omitempty表示如果这字段为空,则不展示该字段
}

func ResponseError(ctx *gin.Context, c MyCode) {
	rd := &ResponseData{
		Code:    c,
		Message: c.Msg(),
		Data:    nil,
	}
	ctx.JSON(http.StatusOK, rd)
}
func ResponseErrorWithMsg(ctx *gin.Context, code MyCode, data interface{}) {
	rd := &ResponseData{
		Code:    code,
		Message: code.Msg(),
		Data:    data,
	}
	ctx.JSON(http.StatusOK, rd)
}
func ResponseSuccess(ctx *gin.Context, data interface{}) {

	rd := &ResponseData{
		Code:    CodeSuccess,
		Message: CodeSuccess.Msg(),
		Data:    data,
	}
	ctx.JSON(http.StatusOK, rd)
}
