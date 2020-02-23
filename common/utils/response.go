package utils

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

/*
Json回述串：{
				code: 表示当前的状态码
				message: 用来秒时失败的原因
				data: {} 返回的数据

			}
*/

// 错误回复的数据结构
type ResponseData struct {
	Code    int         `json:"code"`
	Message string      `json:"message"	`
	Data    interface{} `json:"data"`
}

// 回应错误
func ResponseError(ctx *gin.Context, code int) {

	// 1、生成响应数据对象

	responseData := &ResponseData{
		Code:    code,
		Message: GetMessage(code),
		Data:    make(map[string]interface{}),
	}

	// 2、gin框架进行回应
	ctx.JSON(http.StatusOK, responseData)

}

// 回应成功
func ResponseSuccess(ctx *gin.Context, data interface{}) {
	// 1、生成响应数据对象

	responseData := &ResponseData{
		Code:    ErrCodeSuccess,
		Message: GetMessage(ErrCodeSuccess),
		Data:    data,
	}

	// 3、gin框架进行回应
	ctx.JSON(http.StatusOK, responseData)
}
