package comment

import (
	"github.com/NielJun/go-logger"
	"github.com/daniel/AnserBlock/common/utils"
	"github.com/daniel/AnserBlock/dao/dal"
	"github.com/daniel/AnserBlock/dao/model"
	"github.com/daniel/AnserBlock/generateid"
	"github.com/gin-gonic/gin"
	"html"
	"strconv"
	"strings"
)

const (
	MinCommentSize = 10
	LEVEL_MASTER   = 1
	LEVEL_SECOND   = 2
)

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

// 请求评论列表handle
// 1. 获得answer_id offset limit
// 2. dao层查取该问题下所有的一级评论
// 3. 通过查得的评论列表获得对应的作者id 通过id列表获取作者信息的列表
// 4. 作者列表和评论列表进行匹配出对应的作者名字
func CommentListHandle(ctx *gin.Context) {

	answerIdStr, ok := ctx.GetQuery("answer_id")
	if !ok {
		logger.Error("get query answer_id faild")
		utils.ResponseError(ctx, utils.ErrCodeParamter)
		return
	}
	if len(answerIdStr) == 0 {
		return
	}
	answerId, err := strconv.ParseInt(strings.TrimSpace(answerIdStr), 10, 64)
	if err != nil {
		logger.Error("strvonv string to int faild")
		utils.ResponseError(ctx, utils.ErrCodeServerBusy)
		return
	}

	limit, err := utils.GetQueryInt64(ctx, "limit")
	offset, err := utils.GetQueryInt64(ctx, "offset")
	if err != nil {
		logger.Error("get query limit & offset faild")
		limit = 10
		offset = 0
	}

	logger.Debug("%v", answerId)

	// 从数据库里查处对应的数据
	commmentList, count, err := dal.GetCommentList(answerId, offset, limit)
	if err != nil {
		utils.ResponseError(ctx, utils.ErrCodeServerBusy)
		return
	}

	respoCommentList := model.ResponseCommentList{
		CommentList: commmentList,
		Count:       count,
	}

	// 查询对应的名字

	var userIdList [] int64
	for _, comment := range commmentList {
		userIdList = append(userIdList, comment.AuthorId)
	}

	userInfoList, err := dal.GetUserInfoList(userIdList)
	if err != nil {
		logger.Error("get user info list faild,err: %#v", err)
		utils.ResponseError(ctx, utils.ErrCodeServerBusy)
		return
	}

	userMap := make(map[int64]*model.UserInfo, len(userIdList))

	// 转化为map 便于查询
	for _, userInfo := range userInfoList {
		userMap[userInfo.UserId] = userInfo
	}
	for _, comment := range commmentList {

		// 评论作者名字
		user, ok := userMap[comment.AuthorId]
		if ok {
			comment.AuthorName = user.NickName
		}

		//被回复作者的名字
		replyUser, oks := userMap[comment.ReplyAuthorId]
		if oks {
			comment.ReplyAuthorName = replyUser.NickName
		}

	}

	utils.ResponseSuccess(ctx, respoCommentList)

}

// 获取二级回复评论列表
func CommentReplyListHandle(ctx *gin.Context) {

	commentIdStr, ok := ctx.GetQuery("comment_id")
	if !ok {
		logger.Error("get query question_id faild")
		utils.ResponseError(ctx, utils.ErrCodeParamter)
		return
	}
	if len(commentIdStr) == 0 {
		return
	}
	commentId, err := strconv.ParseInt(strings.TrimSpace(commentIdStr), 10, 64)
	if err != nil {
		logger.Error("strvonv string to int faild")
		utils.ResponseError(ctx, utils.ErrCodeServerBusy)
		return
	}

	limit, err := utils.GetQueryInt64(ctx, "limit")
	offset, err := utils.GetQueryInt64(ctx, "offset")
	if err != nil {
		logger.Error("get query limit & offset faild")
		limit = 10
		offset = 0
	}

	logger.Debug("%v", commentId)

	// 从数据库里查处对应回复列表的数据
	commmentList, count, err := dal.GetRePlyCommentList(commentId, offset, limit)
	if err != nil {
		utils.ResponseError(ctx, utils.ErrCodeServerBusy)
		return
	}

	respoCommentList := model.ResponseCommentList{
		CommentList: commmentList,
		Count:       count,
	}

	// 查询对应的名字

	var userIdList [] int64
	for _, comment := range commmentList {
		userIdList = append(userIdList, comment.AuthorId)
	}

	userInfoList, err := dal.GetUserInfoList(userIdList)
	if err != nil {
		logger.Error("get user info list faild,err: %#v", err)
		utils.ResponseError(ctx, utils.ErrCodeServerBusy)
		return
	}

	userMap := make(map[int64]*model.UserInfo, len(userIdList))

	// 转化为map 便于查询
	for _, userInfo := range userInfoList {
		userMap[userInfo.UserId] = userInfo
	}
	for _, comment := range commmentList {

		// 评论作者名字
		user, ok := userMap[comment.AuthorId]
		if ok {
			comment.AuthorName = user.NickName
		}

		//被回复作者的名字
		replyUser, oks := userMap[comment.ReplyAuthorId]
		if oks {
			comment.ReplyAuthorName = replyUser.NickName
		}

	}

	utils.ResponseSuccess(ctx, respoCommentList)
}

//点赞功能的实现
func LikeHandle(ctx *gin.Context) {

	var like model.Like

	err := ctx.BindJSON(&like)
	if err != nil {
		logger.Error("bind like json faild,err:%#v", err)
		utils.ResponseError(ctx, utils.ErrCodeParamter)
		return
	}

	// 参数校验
	if like.Id == 0 || (like.LikeType != model.LikeTypeAnswer && like.LikeType != model.LikeTypeComment) {
		logger.Error("params errpr,like is %#v", like)
		utils.ResponseError(ctx, utils.ErrCodeParamter)
		return
	}

	if like.LikeType == model.LikeTypeComment {

		err = dal.UpdateCommentCountCount(like.Id)

		// 调用评论的点赞
	} else {
		//调用回答的点赞
		err = dal.UpdateAnswerCountCount(like.Id)
	}

	if err != nil {
		utils.ResponseError(ctx, utils.ErrCodeServerBusy)
		return
	}

	utils.ResponseSuccess(ctx, "点赞")
}
