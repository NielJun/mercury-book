package account

import (
	"errors"
	"github.com/NielJun/go-logger"
	"github.com/daniel/AnserBlock/session"
	"github.com/gin-gonic/gin"
	"net/http"
)

// 处理 账号登陆的时候相关的session的请求
func ProcessRequest(ctx *gin.Context) {

	// 定义出来两个变量 userSession存储从cookies里面取出来的sessionId对应的session
	var userSession session.Session
	// err对应在取得过程中是否报错 如果报错 defer的时候创建一个新的session过去
	var err error

	defer func() {
		if userSession == nil {
			userSession, err = session.CreateSession()
		}
		ctx.Set(MercurySessionName, userSession)
	}()

	// 直接拿到cookies里面关于用户登录的cookie
	cookie, err := ctx.Request.Cookie(CookieSessionId)
	if err != nil {
		//如果对应的cookie不存在 返回相应的逻辑码到逻辑层
		ctx.Set(MercuryUserId, int64(0))
		ctx.Set(MercuryUserLoginStatus, int64(0))
		return
	}
	// 判断sessionId的合法性
	sessionId := cookie.Value
	if len(sessionId) == 0 {
		// SessionID 异常
		ctx.Set(MercuryUserId, int64(0))
		ctx.Set(MercuryUserLoginStatus, int64(0))
		return
	}
	// 当前获取成功 获得对应的session
	userSession, err = session.Get(sessionId)
	if err != nil {
		ctx.Set(MercuryUserId, int64(0))
		ctx.Set(MercuryUserLoginStatus, int64(0))
		return
	}

	// 从session里面取得userid
	userIdInterface, err := userSession.Get(MercuryUserId)
	if err != nil {
		ctx.Set(MercuryUserId, int64(0))
		ctx.Set(MercuryUserLoginStatus, int64(0))
		return
	}

	// 吧interface的userId类型转换为int64位
	userId, ok := userIdInterface.(int64)
	if !ok || userId == 0 {
		ctx.Set(MercuryUserId, int64(0))
		ctx.Set(MercuryUserLoginStatus, int64(0))
		logger.Debug("")
		return
	}

	// 登录成功
	ctx.Set(MercuryUserId, int64(userId))
	ctx.Set(MercuryUserLoginStatus, int64(1))
	logger.Debug("------> 登录状态设置成功")
	return

}

// 回应账号的登录相关的请求
func ProcessResponse(ctx *gin.Context) {
	var userSession session.Session
	tempSession, exists := ctx.Get(MercurySessionName)
	if !exists {
		return
	}

	userSession, ok := tempSession.(session.Session)
	if !ok {
		return
	}

	if userSession == nil {
		return
	}

	if userSession.IsModify() == false {
		return
	}
	err := userSession.Save()
	if err != nil {
		return
	}

	// 成功了以后 种cookie
	sessionId := userSession.Id()
	cookie := &http.Cookie{
		Name:     CookieSessionId,
		Value:    sessionId,
		MaxAge:   CookieMaxAge,
		HttpOnly: true,
		Path:     "/",
	}
	http.SetCookie(ctx.Writer, cookie)
	return
}

// 设置user_id到gin框架上下文
func SetUserId(userId int64, ctx *gin.Context) {
	var userSession session.Session
	tempSession, exist := ctx.Get(MercurySessionName)
	if !exist {
		return
	}
	userSession, ok := tempSession.(session.Session)
	if !ok {
		return
	}

	userSession.Set(MercuryUserId, userId)
}

// 从gin上下文获取UserId
func GetUserId(ctx *gin.Context) (userId int64, err error) {

	userIdInterface, exists := ctx.Get(MercuryUserId)
	if !exists {
		err = errors.New(" user id not exists ")
		return
	}

	userId, ok := userIdInterface.(int64)

	if !ok {
		err = errors.New(" user id corvered to int64 faild ")
		return
	}
	return
}

//当前是否已经登陆成功
func IsLogin(ctx *gin.Context) (islogin bool) {
	islogin = false

	loginStatus, exists := ctx.Get(MercuryUserLoginStatus)
	if !exists {
		logger.Debug("不存在登录状态码")
		return
	}

	loginCode, ok := loginStatus.(int64)

	if !ok {
		logger.Debug("转换登录状态码失败")
		return
	}
	if loginCode == 0 {

		logger.Debug("登录状态码为0")

		return
	}
	islogin = true

	return

}
