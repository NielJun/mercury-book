package account

import (
	"fmt"
	"github.com/daniel/AnserBlock/common/utils"
	"github.com/daniel/AnserBlock/dao/dal"
	"github.com/daniel/AnserBlock/dao/model"
	"github.com/daniel/AnserBlock/generateid"
	"github.com/daniel/AnserBlock/middleware/account"
	"github.com/gin-gonic/gin"
)

func LoginHandle(c *gin.Context) {

	// 通过中间件进行 中间件session的过滤
	account.ProcessRequest(c)

	var userInfo model.UserInfo
	var err error
	defer func() {

		if err != nil {
			return
		}

		account.SetUserId(userInfo.UserId, c)
		account.ProcessResponse(c)
		// 成功 处理登录成功请求 并且设置user_id到session里面
		utils.ResponseSuccess(c, nil)
	}()

	err = c.BindJSON(&userInfo)
	if err != nil {
		// 参数错误
		utils.ResponseError(c, utils.ErrCodeParamter)
		fmt.Printf("%#v", userInfo)
		return
	}

	if len(userInfo.UserName) == 0 || len(userInfo.Password) == 0 {
		utils.ResponseError(c, utils.ErrCodeParamter)
		fmt.Printf("%#v", userInfo)
		return
	}

	// dal操作数据库
	err = dal.Login(&userInfo)
	if err == utils.ErrUserNotExisted {
		utils.ResponseError(c, utils.ErrCodeUserExit)
		return
	} else if err == utils.ErrUserPwdWrong {
		utils.ResponseError(c, utils.ErrCodeUserNameOrPwdWrong)
		return
	}

	// 判断是否错误
	if err != nil {
		utils.ResponseError(c, utils.ErrCodeServerBusy)
		return
	}

}

// 注册的Handle
func RegisterHandle(c *gin.Context) {
	var userInfo model.UserInfo
	err := c.BindJSON(&userInfo)
	if err != nil {
		// 参数错误
		utils.ResponseError(c, utils.ErrCodeParamter)
		fmt.Printf("%#v", userInfo)
		return
	}
	if userInfo.Sex != model.Boy && userInfo.Sex != model.Girl {
		utils.ResponseError(c, utils.ErrCodeParamter)
		fmt.Printf("%#v", userInfo)
		return
	}

	if len(userInfo.UserName) == 0 || len(userInfo.Email) == 0 || len(userInfo.Password) == 0 {
		utils.ResponseError(c, utils.ErrCodeParamter)
		fmt.Printf("%#v", userInfo)
		return
	}

	userId, err := generateid.GetId()
	if err != nil {
		panic(err)
		return
	}
	userInfo.UserId = int64(userId)

	// dal操作数据库
	err = dal.Register(&userInfo)
	if err == utils.ErrUserAlreadyExisted {
		utils.ResponseError(c, utils.ErrCodeUserExit)
		return
	}

	// 判断是否错误
	if err != nil {
		utils.ResponseError(c, utils.ErrCodeServerBusy)
		return
	}
	// 成功
	utils.ResponseSuccess(c, nil)

}
