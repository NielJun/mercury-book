package category

import (
	"github.com/NielJun/go-logger"
	"github.com/daniel/AnserBlock/common/utils"
	"github.com/daniel/AnserBlock/dao/dal"
	"github.com/daniel/AnserBlock/dao/model"
	"github.com/gin-gonic/gin"
	"strconv"
)

func CategoryListHandle(c *gin.Context) {

	categoryList, err := dal.GetCategoryList()
	if err != nil {
		// 参数错误

		logger.Error("Get category List err,err: %#v", err)

		utils.ResponseError(c, utils.ErrCodeServerBusy)
		return
	}

	utils.ResponseSuccess(c, categoryList)

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
