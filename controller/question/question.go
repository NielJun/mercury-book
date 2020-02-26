package question

import (
	"github.com/NielJun/go-logger"
	"github.com/daniel/AnserBlock/common/utils"
	"github.com/daniel/AnserBlock/dao/dal"
	"github.com/daniel/AnserBlock/dao/model"
	"github.com/daniel/AnserBlock/filter"
	"github.com/daniel/AnserBlock/generateid"
	"github.com/daniel/AnserBlock/middleware/account"
	"github.com/gin-gonic/gin"
	"strconv"
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

func GetQuestionListHandle(ctx *gin.Context) {

	// 从网站地址 [ www.rubyboy.cn/ ?categoryId=?? ] 获取category_id
	categoryIdStr, ok := ctx.GetQuery("category_id")

	if !ok {
		logger.Error("invaild category_id, not fount category_id")
		utils.ResponseError(ctx, utils.ErrCodeParamter)
		return
	}

	categoryId, err := strconv.ParseInt(categoryIdStr, 10, 64)

	if err != nil {
		logger.Error("invalid category_id,strconv.ParseInt faild,err: %#v ,str : %v", err, categoryIdStr)
		utils.ResponseError(ctx, utils.ErrCodeParamter)
		return
	}

	// 从数据库查询某个category下面的所有的Question 列表
	questionList, err := dal.GetQuestionList(categoryId)
	if err != nil {
		logger.Error("select  question  faild.  err: %#v", err)
		utils.ResponseError(ctx, utils.ErrCodeServerBusy)
		return
	}

	// 空的列表 表示某个栏目不存在内容
	if len(questionList) == 0 {
		logger.Warn("question list is null")
		utils.ResponseSuccess(ctx, questionList)
		return
	}

	// 相应的查询对应的坐着的相关信息  并且去重
	// 去重方式 用一个map存储当前的作者信息 然后遍历的时候如果map里面包含则continue
	var authorIdList [] int64
	authorIdMap := make(map[int64]bool)
	for _, question := range questionList {
		// 匹配map中间是否包含该作者id
		_, contain := authorIdMap[question.AuthorId]
		if contain {
			continue
		}
		authorIdMap[question.AuthorId] = true
		authorIdList = append(authorIdList, question.AuthorId)
	}

	// 获取对应的 作者信息列表
	authorInfoList, err := dal.GetUserInfoList(authorIdList)
	if err != nil {
		logger.Error("get user info list error,err : %#v", err)
		utils.ResponseError(ctx, utils.ErrCodeServerBusy)
		return
	}

	//
	var respoQuestionList []*model.ResponseQuestion

	for _, question := range questionList {

		// 设置question本体
		respoQuestion := &model.ResponseQuestion{}
		respoQuestion.Question = *question
		respoQuestion.Question.CreateTimeStr = respoQuestion.Question.CreateTime.Format("2016/1/2 15:04:05")
		// 设置作者的信息 进行一个匹配
		for _, userInfo := range authorInfoList {

			if question.AuthorId == userInfo.UserId {
				if len(userInfo.NickName) != 0 {
					respoQuestion.AuthorName = userInfo.NickName
				} else {
					respoQuestion.AuthorName = userInfo.UserName
				}

				break
			}
		}

		// 把结果放进list
		respoQuestionList = append(respoQuestionList, respoQuestion)

	}

	// 返回成功的信息
	utils.ResponseSuccess(ctx, respoQuestionList)

}

// 问题详情页面
func QuestionDetailHandle(c *gin.Context) {

	questionIdStr, ok := c.GetQuery("question_id")
	if !ok {
		logger.Error("question id is not availd,not found question_id")
		utils.ResponseError(c, utils.ErrCodeParamter)
		return
	}

	questionId, err := strconv.ParseInt(questionIdStr, 10, 64)
	if err != nil {
		logger.Error("string conv to int faild,err : %#v ", err)
		utils.ResponseError(c, utils.ErrCodeServerBusy)
		return
	}

	question, err := dal.GetQuestion(questionId)
	if err != nil {
		logger.Error("get question faild where question id is : %#v,err:%#v ", questionId, err)
		utils.ResponseError(c, utils.ErrCodeServerBusy)
		return
	}

	// 查询标签信息
	categoryMap, err := dal.GetCategoryMap([]int64{question.CategoryId})
	if err != nil {
		logger.Error("get category map faild,category id is %v, err:%#v ", questionId, question.CategoryId, err)
		utils.ResponseError(c, utils.ErrCodeServerBusy)
		return
	}

	category, ok := categoryMap[question.CategoryId]
	if !ok {
		logger.Error("get category map faild,question is %#v ,err:%#v ", questionId, question, err)
		utils.ResponseError(c, utils.ErrCodeServerBusy)
		return
	}

	// 获取对应的 作者信息列表
	authorInfoList, err := dal.GetUserInfoList([]int64{question.AuthorId})
	if err != nil || len(authorInfoList) == 0 {
		logger.Error("get user info list error,err : %#v", err)
		utils.ResponseError(c, utils.ErrCodeServerBusy)
		return
	}

	// 把数据进行聚合
	questionDetail := model.QuestionDetail{}

	questionDetail.Question = *question
	questionDetail.AuthorName = authorInfoList[0].NickName
	questionDetail.CategoryName = category.CategoryName

	// 然后在进行返回
	utils.ResponseSuccess(c, questionDetail)

}
