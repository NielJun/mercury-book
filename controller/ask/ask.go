package ask

import (
	"github.com/NielJun/go-logger"
	"github.com/daniel/AnserBlock/common/utils"
	"github.com/daniel/AnserBlock/dao/model"
	"github.com/gin-gonic/gin"
)

func QuestionSubmitHandle(ctx *gin.Context) {

	var question model.Question

	err := ctx.BindJSON(&question)

	if err != nil {
		utils.ResponseError(ctx, utils.ErrCodeParamter)
		return
	}

	logger.Debug("Ask请求成功，接收到的数据: %#v", question)

}
