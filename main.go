package main

import (
	"fmt"
	"github.com/DeanThompson/ginpprof"
	"github.com/NielJun/go-logger"
	"github.com/daniel/AnserBlock/controller/account"
	"github.com/daniel/AnserBlock/controller/ask"
	"github.com/daniel/AnserBlock/controller/category"
	"github.com/daniel/AnserBlock/dao"
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

func initSession() (err error) {
	err = middleware.Init("memory", "localhost:6379")
	return
}

func initLogger() (err error) {
	config := make(map[string]string)
	config["log_level"] = "debug"
	err = logger.InitLogger("console", config)
	return
}

func RegisterAPI(router *gin.Engine) {

	router.POST("/api/user/register", account.RegisterHandle)
	router.POST("/api/user/login", account.LoginHandle)
	router.GET("/api/category/list", category.CategoryListHandle)
	router.POST("/api/ask/submit", ask.QuestionSubmitHandle)
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

	// 初始化ID生成器
	generateid.Init(1)
	RegisterAPI(router)
	// 启动gin监听
	router.Run(":9090")
}
