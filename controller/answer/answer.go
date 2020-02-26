package answer

import (
	"github.com/daniel/AnserBlock/common/utils"
	"github.com/daniel/AnserBlock/dao/dal"
	"github.com/daniel/AnserBlock/dao/model"
	"github.com/gin-gonic/gin"
)

// 通过问题的Id获取回答的列表
// 参数:
// question_id :	问题的id
// offset:			回答的第几个
// limit:			允许多少个评论展示
func AnswerListHandle(c *gin.Context) {

	// 1.通过query得到客户端传入的数据参数
	questionId, err := utils.GetQueryInt64(c, "question_id")
	if err != nil {
		utils.ResponseError(c, utils.ErrCodeParamter)
		return
	}

	offset, err := utils.GetQueryInt64(c, "offset")
	if err != nil {
		utils.ResponseError(c, utils.ErrCodeParamter)
		return
	}

	limit, err := utils.GetQueryInt64(c, "limit")
	if err != nil {
		utils.ResponseError(c, utils.ErrCodeParamter)
		return
	}

	// 2.通过参数查询评论的列表id
	answerIdList, err := dal.GetAnswerIdList(questionId, offset, limit)
	if err != nil {
		utils.ResponseError(c, utils.ErrCodeServerBusy)
		return
	}

	// 3.通过评论的列表Id查处评论列表
	answerList, err := dal.GetAnwerList(answerIdList)
	if err != nil {
		utils.ResponseError(c, utils.ErrCodeServerBusy)
		return
	}

	// 4.通过评论列表 构造作者的id表并且查询出作者的信息
	var userIdList [] int64
	for _, v := range answerList {
		userIdList = append(userIdList, v.AuthorId)
	}
	userInfoList, err := dal.GetUserInfoList(userIdList)
	if err != nil {
		utils.ResponseError(c, utils.ErrCodeServerBusy)
		return
	}

	//5. 此时已经具有 答案的列表信息 答案列表对应的作者信息 就可以构造返回的数据结构
	respoAnswerList := &model.ResponseAnswerList{}
	for _, answer := range answerList {

		// 构造一条数据
		responAnswer := &model.ResponseAnswer{}
		// 赋予 answer内容
		responAnswer.Answer = *answer
		// 赋予
		for _, userInfo := range userInfoList {
			if userInfo.UserId == answer.AuthorId {
				responAnswer.AuthorName = userInfo.NickName
				break
			}
		}
		// 把元素加入到返回的数组里面
		respoAnswerList.AnswerList = append(respoAnswerList.AnswerList,responAnswer)
	}

	// 最后在获取问题对应的回答的总数
	totalAnswerCount,err := dal.GetTotalAnswerCount(questionId)
	if err != nil {
		utils.ResponseError(c, utils.ErrCodeServerBusy)
		return
	}
	// 聚合总数
	respoAnswerList.TotalCount = totalAnswerCount

	// 返回
	utils.ResponseSuccess(c, respoAnswerList)

}
