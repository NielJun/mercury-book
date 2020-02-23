package category

import (
	"github.com/daniel/AnserBlock/common/utils"
	"github.com/daniel/AnserBlock/dao/dal"
	"github.com/gin-gonic/gin"
)

func CategoryListHandle(c *gin.Context) {

	categoryList, err := dal.GetCategoryList()
	if err != nil {
		// 参数错误
		utils.ResponseError(c, utils.ErrCodeServerBusy)
		return
	}

	utils.ResponseSuccess(c, categoryList)

}
