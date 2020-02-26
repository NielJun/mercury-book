package dal

import (
	"fmt"
	"github.com/daniel/AnserBlock/dao"
	"github.com/daniel/AnserBlock/dao/model"
	"github.com/daniel/AnserBlock/generateid"
	"testing"
	"time"
)

func init() {
	dns := "root:12345678@tcp(localhost:3306)/mercury?parseTime=true"
	err := dao.Init(dns)
	if err != nil {
		fmt.Printf("Init data base Err ! %v", err)
		return
	}

}

func TestCreateQuestion(t *testing.T) {
	id, _ := generateid.GetId()

	question := model.Question{
		QuestionId:    int64(id),
		Caption:       "老婆我爱你",
		Content:       "阿技术不断环境保护局北京阿斯顿好噶结束的感觉哈是给大家哈告诉大家哈是",
		AuthorId:      7520,
		CategoryId:    1,
		Status:        1,
		CreateTime:    time.Time{},
		CreateTimeStr: "",
		QuestionIdStr: "",
		AuthorIdStr:   "",
	}

	err := CreateQuestion(&question)

	if err != nil {
		t.Error(err)
		return
	}

}

func TestGetQuestionList(t *testing.T) {

	questionList, err := GetQuestionList(1)

	if err != nil {
		t.Error(err)
		return
	}
	for _, q := range questionList {
		t.Logf("%#v", *q)
	}

}
