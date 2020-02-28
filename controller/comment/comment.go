package comment

import (
	"github.com/NielJun/go-logger"
	"github.com/daniel/AnserBlock/common/utils"
	"github.com/daniel/AnserBlock/dao/dal"
	"github.com/daniel/AnserBlock/dao/model"
	"github.com/daniel/AnserBlock/generateid"
	"github.com/gin-gonic/gin"
	"html"
)

const MinCommentSize = 10

// 发表评论
func PostCommentHandle(ctx *gin.Context) {

	var comment model.Comment

	err := ctx.BindJSON(&comment)
	if err != nil {

		logger.Error("paprams err,bind json err ,err :%#v", err)
		utils.ResponseError(ctx, utils.ErrCodeParamter)
		return
	}
	logger.Debug("json is : %#v", comment)

	// 做一个合法性判断
	if len(comment.Content) < MinCommentSize {
		utils.ResponseError(ctx, utils.ErrCodeParamter)
		return
	}

	//userId, err := account.GetUserId(ctx)
	userId := int64(290520267376033793)

	if err != nil {
		logger.Error("gin get user id faild. err: %#v", err)
		utils.ResponseError(ctx, utils.ErrCodeParamter)
		return
	}

	// 对内容进行zhuang yi
	comment.Content = html.EscapeString(comment.Content)

	id, err := generateid.GetId()

	if err != nil {
		logger.Error("generate id faild,err :%#v", err)
		utils.ResponseError(ctx, utils.ErrCodeServerBusy)
		return
	}

	comment.CommentId = int64(id)
	comment.AuthorId = userId

	err = dal.CreatePostComment(&comment)

	if err != nil {
		utils.ResponseError(ctx, utils.ErrCodeServerBusy)
		return
	}

	// 否则成功
	utils.ResponseSuccess(ctx, comment)

}

// 评论回复
func PostCommentReplyHandle(ctx *gin.Context) {

	var comment model.Comment

	err := ctx.BindJSON(&comment)
	if err != nil {

		logger.Error("paprams err,bind json err ,err :%#v", err)
		utils.ResponseError(ctx, utils.ErrCodeParamter)
		return
	}
	logger.Debug("json is : %#v", comment)

	// 做一个合法性判断
	if len(comment.Content) < MinCommentSize || comment.QuestionId == 0 || comment.ReplyCommentId == 0 || comment.ParentId == 0 {
		utils.ResponseError(ctx, utils.ErrCodeParamter)
		logger.Error("len(commentContent) <minSize, comment.QuestionId = %v,comment.ReplyCommentId = %v, comment.ParentId = %v.", comment.QuestionId, comment.ReplyCommentId, comment.ParentId)
		return
	}

	//userId, err := account.GetUserId(ctx)
	userId := int64(290520267376033793)

	if err != nil {
		logger.Error("gin get user id faild. err: %#v", err)
		utils.ResponseError(ctx, utils.ErrCodeParamter)
		return
	}

	// 对内容进行zhuang yi
	comment.Content = html.EscapeString(comment.Content)

	id, err := generateid.GetId()

	if err != nil {
		logger.Error("generate id faild,err :%#v", err)
		utils.ResponseError(ctx, utils.ErrCodeServerBusy)
		return
	}

	comment.CommentId = int64(id)
	comment.AuthorId = userId

	err = dal.CreatePostReplyComment(&comment)

	if err != nil {
		utils.ResponseError(ctx, utils.ErrCodeServerBusy)
		logger.Error("create post reply coment failed ,err :%#v", err)
		return
	}

	// 否则成功
	utils.ResponseSuccess(ctx, comment)
}
