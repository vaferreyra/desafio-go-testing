package web

import "github.com/gin-gonic/gin"

type Response struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func ResponseOk(ctx *gin.Context, code int, message string, data interface{}) {
	ctx.JSON(code, Response{
		Message: message,
		Data:    data,
	})
}

func ResponseErr(ctx *gin.Context, code int, message string) {
	ctx.JSON(code, Response{
		Message: message,
	})
}
