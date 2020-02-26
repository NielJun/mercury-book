package category

import (
	"github.com/NielJun/go-logger"
	"github.com/daniel/AnserBlock/common/utils"
	"github.com/daniel/AnserBlock/dao/dal"
	"github.com/gin-gonic/gin"
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


