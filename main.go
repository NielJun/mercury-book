package main

import (
	"fmt"
	"github.com/DeanThompson/ginpprof"
	"github.com/NielJun/go-logger"
	"github.com/daniel/AnserBlock/controller/account"
	"github.com/daniel/AnserBlock/controller/answer"
	"github.com/daniel/AnserBlock/controller/category"
	"github.com/daniel/AnserBlock/controller/comment"
	"github.com/daniel/AnserBlock/controller/question"
	"github.com/daniel/AnserBlock/dao"
	"github.com/daniel/AnserBlock/filter"
	"github.com/daniel/AnserBlock/generateid"
	middleware "github.com/daniel/AnserBlock/middleware/account"
	"github.com/gin-gonic/gin"
)

func initWebFront(router *gin.Engine) {
	router.StaticFile("/", "./static/index.html")
	router.StaticFile("/favicon.ico/", "./static/favicon.ico")
	router.Static("/css/", "./static/css")
	router.Static("/fonts/", "./static/fonts")
	router.Static("/img/", "./static/img")
	router.Static("/js/", "./static/js")
}

func initDataBase() (err error) {
	dns := "root:12345678@tcp(localhost:3306)/mercury?parseTime=true"
	err = dao.Init(dns)
	if err != nil {
		fmt.Printf("Init data base Err ! %v", err)
		return
	}
	return
}

// 初始化session
func initSession() (err error) {
	err = middleware.Init("memory", "localhost:6379")
	return
}

//	 初始化logger日志系统
func initLogger() (err error) {
	config := make(map[string]string)
	config["log_level"] = "debug"
	err = logger.InitLogger("console", config)
	return
}

// 初始化敏感词过滤模块
func initFiltter() (err error) {

	err = filter.Init("./filter/data/filter.dat.txt")
	if err != nil {
		logger.Debug("敏感词库文件加载失败 %#v", err)
	}
	return
}

//注册对外的API接口
func RegisterAPI(router *gin.Engine) {

	router.POST("/api/user/register", account.RegisterHandle)
	router.POST("/api/user/login", account.LoginHandle)
	router.GET("/api/category/list", category.CategoryListHandle)
	// 问题发布页面  第一个参数是中间件的操作
	router.POST("/api/question/submit", middleware.AuthMiddleware, question.QuestionSubmitHandle)
	router.GET("/api/question/list", question.GetQuestionListHandle)
	router.GET("/api/question/detail", question.QuestionDetailHandle)
	router.GET("/api/answer/list", answer.AnswerListHandle)

	// 评论相关 作为一个组
	commentGroup := router.Group("/api/comment/")//, middleware.AuthMiddleware
	// 评论
	commentGroup.POST("/post_comment", comment.PostCommentHandle)
	// 回复评论
	commentGroup.POST("/post_comment_reply", comment.PostCommentReplyHandle)
}

func main() {
	router := gin.Default()
	ginpprof.Wrapper(router)

	initWebFront(router)
	err := initDataBase()
	if err != nil {
		panic(err)
	}

	err = initSession()
	if err != nil {
		panic(err)
	}

	err = initLogger()
	if err != nil {
		panic(err)
	}

	err = initFiltter()
	if err != nil {
		panic(err)
	}

	// 初始化ID生成器
	generateid.Init(1)
	RegisterAPI(router)
	// 启动gin监听
	router.Run(":9090")
}
