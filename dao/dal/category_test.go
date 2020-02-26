package dal

import (
	"fmt"
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

}

func TestGetCategoryList(t *testing.T) {

	categoryList, err := GetCategoryList()
	if err != nil {
		t.Errorf("%#v", err)
		return
	}

	for _, value := range categoryList {
		t.Logf("CategoryName: %s  \n", value.CategoryName)
	}
}

func TestInsertCategory(t *testing.T) {

	categoryList := make([]model.Category, 5)

	categoryList[0] = model.Category{
		CategoryName: "运动",
		CategoryId:   0,
	}
	categoryList[1] = model.Category{
		CategoryName: "健身",
		CategoryId:   1,
	}
	categoryList[2] = model.Category{
		CategoryName: "撩妹",
		CategoryId:   2,
	}
	categoryList[3] = model.Category{
		CategoryName: "看电影",
		CategoryId:   3,
	}
	categoryList[4] = model.Category{
		CategoryName: "爬山",
		CategoryId:   4,
	}

	for _, value := range categoryList {
		err := InsertCategory(value)

		if err != nil {
			t.Errorf("%#v", err)
		}
	}

}

func TestGetCategoryMap(t *testing.T) {
	 categoryIds  := []int64{1,2,3,4}

	categoryMap,err:= GetCategoryMap(categoryIds)
	if err != nil {
		t.Errorf("%#v",err)
		return
	}
	for _, value := range categoryMap {

		t.Logf("categoryId is %v , categoryName is %v",value.CategoryId,value.CategoryName)
	}
}