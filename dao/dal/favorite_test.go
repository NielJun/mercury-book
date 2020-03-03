package dal

import (
	"fmt"
	"github.com/NielJun/go-logger"
	"github.com/daniel/AnserBlock/dao"
	"github.com/daniel/AnserBlock/dao/model"
	"github.com/daniel/AnserBlock/generateid"
	"testing"
)

func init() {
	dns := "root:12345678@tcp(localhost:3306)/mercury?parseTime=true"
	err := dao.Init(dns)
	if err != nil {
		fmt.Printf("Init data base Err ! %v", err)
		return
	}

}

func TestCreateFavoriteDir(t *testing.T) {

	var fd model.FavoriteDir

	fd.UserId = 290520267376033793
	fd.DirName = "叉叉叉"

	id, _ := generateid.GetId()

	dId := int64(id)
	fd.DirId = dId
	fd.Count = 1

	err := CreateFavoriteDir(&fd)
	if err != nil {
		t.Error(err)
		return
	}
	t.Logf("success")
}

func TestCreateFavorite(t *testing.T) {
	var fd model.Favorite

	fd.UserId = 290520267376033793
	fd.DirId = 291225146855784449
	fd.AnswerId = 1

	err := CreateFavorite(&fd)
	if err != nil {
		t.Error(err)
		return
	}
	t.Logf("success")
}

func TestGetFavoriteDirList(t *testing.T) {
	var userId int64 = 290520267376033793
	list, err := GetFavoriteDirList(userId)
	if err != nil {
		logger.Error("%#v", err)
		return
	}
	t.Logf("%#v", list)

}

func TestGetFavoriteList(t *testing.T) {
	var userId int64 = 290520267376033793
	var dirId int64 = 291225146855784449

	favoriteList,err:= GetFavoriteList(userId,dirId,0,10)
	if err != nil {
		t.Errorf("faild  err: %#v",err)
		return
	}

	t.Logf("%#v",favoriteList)

}
