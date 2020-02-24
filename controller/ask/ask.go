package ask

import (
	"github.com/NielJun/go-logger"
	"github.com/daniel/AnserBlock/common/utils"
	"github.com/daniel/AnserBlock/dao/dal"
	"github.com/daniel/AnserBlock/dao/model"
	"github.com/daniel/AnserBlock/filter"
	"github.com/daniel/AnserBlock/generateid"
	"github.com/daniel/AnserBlock/middleware/account"
	"github.com/gin-gonic/gin"
)

// 请求发送提问提交的Handle
func QuestionSubmitHandle(ctx *gin.Context) {

	var question model.Question

	err := ctx.BindJSON(&question)

	if err != nil {
		utils.ResponseError(ctx, utils.ErrCodeParamter)
		return
	}

	logger.Debug("Ask请求成功，接收到的数据: %#v", question)

	// 标题敏感词过滤
	_, isReplaced := filter.Replace(question.Caption, "***")
	if isReplaced {
		utils.ResponseError(ctx, utils.ErrCodeCaptionSensitive)
		return
	}

	// 内容敏感词过滤
	_, isReplaced = filter.Replace(question.Content, "***")
	if isReplaced {
		utils.ResponseError(ctx, utils.ErrCodeContentSensitive)
		return
	}

	// 生成问题的id
	questionId, err := generateid.GetId()
	if err != nil {
		logger.Error("生成questionId 失败  %#v", err)
		utils.ResponseError(ctx, utils.ErrCodeServerBusy)
		return
	}

	question.QuestionId = (int64)(questionId)

	userId, err := account.GetUserId(ctx)
	if err != nil || userId < 0 {
		logger.Error("还未登陆 %#v", err)
		utils.ResponseError(ctx, utils.ErrCodeNotLogin)
	}

	question.AuthorId = userId

	// 写入数据库
	err = dal.CreateQuestion(&question)

	if err != nil {
		utils.ResponseError(ctx, utils.ErrCodeServerBusy)
		return
	}

	logger.Debug("提交问题成功 %#v ", question)

	utils.ResponseSuccess(ctx, nil)
}
