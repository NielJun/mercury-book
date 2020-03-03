package favorite

import (
	"github.com/NielJun/go-logger"
	"github.com/daniel/AnserBlock/common/utils"
	"github.com/daniel/AnserBlock/dao/dal"
	"github.com/daniel/AnserBlock/dao/model"
	"github.com/daniel/AnserBlock/generateid"
	"github.com/gin-gonic/gin"
)

// 添加收藏家功能
func AddDirHandle(c *gin.Context) {
	var favoriteDir model.FavoriteDir
	err := c.BindJSON(&favoriteDir)
	if err != nil {
		logger.Error("bind json faild,err: %#v", err)
		utils.ResponseError(c, utils.ErrCodeParamter)
		return
	}
	if len(favoriteDir.DirName) == 0 {
		logger.Error("length of the dir_name err")
		utils.ResponseError(c, utils.ErrCodeParamter)
		return
	}

	favoriteDirUnsingId, err := generateid.GetId()
	if err != nil {
		logger.Error("get id faild,err : %#v", err)
		utils.ResponseError(c, utils.ErrCodeServerBusy)
		return
	}
	//userId, err := account.GetUserId(ctx)
	userId := int64(290520267376033793)

	// 设置相关参数
	favoriteDir.DirId = int64(favoriteDirUnsingId)
	favoriteDir.UserId = userId
	err = dal.CreateFavoriteDir(&favoriteDir)
	if err != nil {
		utils.ResponseError(c, utils.ErrCodeServerBusy)
		return
	}
	utils.ResponseSuccess(c, "创建收藏夹成功")

}

// 添加收藏
func AddFavoriteHandle(c *gin.Context) {

	var favorite model.Favorite
	err := c.BindJSON(&favorite)
	if err != nil {
		logger.Error("bind json faild,err: %#v", err)
		utils.ResponseError(c, utils.ErrCodeParamter)
		return
	}
	if favorite.DirId == 0 || favorite.AnswerId == 0 {
		logger.Error("favorite dir_id or answer_id is 0")
		utils.ResponseError(c, utils.ErrCodeParamter)
		return
	}
	//userId, err := account.GetUserId(ctx)
	userId := int64(290520267376033793)

	favorite.UserId = userId

	err = dal.CreateFavorite(&favorite)
	if err == utils.ErrRecordExisted {
		utils.ResponseError(c, utils.ErrCodeRecordExisted)
		return
	}

	if err != nil {
		utils.ResponseError(c, utils.ErrCodeServerBusy)
		return
	}

	utils.ResponseSuccess(c, "收藏成功")

}

// 收藏夹列表
func FavoriteDirListHandle(c *gin.Context) {

	//userId, err := account.GetUserId(ctx)
	userId := int64(290520267376033793)
	favoriteList, err := dal.GetFavoriteDirList(userId)
	if err != nil {
		utils.ResponseError(c, utils.ErrCodeServerBusy)
		return
	}

	utils.ResponseSuccess(c, favoriteList)

}

/// 某个收藏夹的收藏列表
func FavoriteListHandle(c *gin.Context) {

	dir_id, err := utils.GetQueryInt64(c, "dir_id")
	offset, err := utils.GetQueryInt64(c, "offset")
	limit, err := utils.GetQueryInt64(c, "limit")
	//userId, err := account.GetUserId(ctx)
	userId := int64(290520267376033793)
	if err != nil {
		utils.ResponseError(c, utils.ErrCodeParamter)
		return
	}
	fevoriteList, err := dal.GetFavoriteList(userId, dir_id, offset, limit)

	if err != nil {
		utils.ResponseError(c, utils.ErrCodeServerBusy)
		return
	}

	var answerIdList [] int64
	for _, v := range fevoriteList {
		answerIdList = append(answerIdList, v.AnswerId)
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
		respoAnswerList.AnswerList = append(respoAnswerList.AnswerList, responAnswer)
	}
	// 返回
	utils.ResponseSuccess(c, respoAnswerList)

}
