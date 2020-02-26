package dal

import (
	"fmt"
	"github.com/NielJun/go-logger"
	"github.com/daniel/AnserBlock/common/utils"
	"github.com/daniel/AnserBlock/dao"
	"github.com/daniel/AnserBlock/dao/model"
	"testing"
)

func init() {
	dns := "root:12345678@tcp(localhost:3306)/mercury?parseTime=true"
	err := dao.Init(dns)
	if err != nil {
		fmt.Printf("Init data base Err ! %v", err)
		return
	}

	config := make(map[string]string)
	config["log_level"] = "debug"
	err = logger.InitLogger("console", config)
	return
}

func TestLogin(t *testing.T) {
	user := &model.UserInfo{
		UserId:   100020201,
		NickName: "烈火讽刺",
		Sex:      1,
		UserName: "廖涛	",
		Email:    "769288695@qq.com",
		Password: "123312",
	}
	err := Login(user)
	if err != nil {
		t.Errorf("登陆失败 ,%#v", err)
	} else {
		t.Errorf("登录成功")
	}
}

func TestRegister(t *testing.T) {
	user := &model.UserInfo{
		UserId:   100020201,
		NickName: "烈火讽刺",
		Sex:      1,
		UserName: "廖涛",
		Email:    "769288695@qq.com",
		Password: "123312",
	}
	err := Register(user)
	if err == utils.ErrUserNotExisted {
		t.Errorf("注册成功 ,%#v", err)
	}
}

func TestGetUserInfoList(t *testing.T) {

	sqlstr := "select user_id from user"
	var userIdList [] int64

	err := dao.DB.Select(&userIdList, sqlstr)

	if err != nil {
		t.Error(err)
	}
	userInfoList, err := GetUserInfoList(userIdList)

	for _, userInfo := range userInfoList {
		t.Logf("%#v", userInfo)
	}
}
