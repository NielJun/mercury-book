package account

import (
	"github.com/daniel/AnserBlock/common/utils"
	"github.com/gin-gonic/gin"
)

// 登陆处理中间件
func AuthMiddleware(ctx *gin.Context) {
	ProcessRequest(ctx)
	isLogin := IsLogin(ctx)
	if isLogin==false{
		utils.ResponseError(ctx,utils.ErrCodeNotLogin)
		// 中断当前请求
		ctx.Abort()
		return
	}

	ctx.Next()
}
