package dal

import (
	"database/sql"
	"fmt"
	"github.com/daniel/AnserBlock/dao"
	"github.com/daniel/AnserBlock/dao/model"
)

// 从数据库获取category的列表
func GetCategoryList() (categoryList [] *model.Category, err error) {

	sqlStr := "select category_name,category_id from category"

	err = dao.DB.Select(&categoryList, sqlStr)
	//当前没有数据行
	if err == sql.ErrNoRows {
		err = nil
		return
	}
	if err != nil {
		return
	}

	return
}

func InsertCategory(category model.Category) (err error) {

	sqlStr := "insert into category (category_id,category_name) values(?,?)"

	_, err = dao.DB.Exec(sqlStr, category.CategoryId,category.CategoryName)

	if err != nil {
		fmt.Printf("插入标签数据失败，err: %#v", err)
	}

	return
}
