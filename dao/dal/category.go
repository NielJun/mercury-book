package dal

import (
	"database/sql"
	"fmt"
	"github.com/NielJun/go-logger"
	"github.com/daniel/AnserBlock/dao"
	"github.com/daniel/AnserBlock/dao/model"
	"github.com/jmoiron/sqlx"
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

// 插入category到数据库
func InsertCategory(category model.Category) (err error) {

	sqlStr := "insert into category (category_id,category_name) values(?,?)"

	_, err = dao.DB.Exec(sqlStr, category.CategoryId, category.CategoryName)

	if err != nil {
		fmt.Printf("插入标签数据失败，err: %#v", err)
	}

	return
}

// 通过指定标签id切片查询对应的标签map
func GetCategoryMap(categoryIds []int64) (categoryMap map[int64]*model.Category, err error) {

	// 首先把发挥的map进行初始化
	categoryMap = make(map[int64]*model.Category, len(categoryIds))

	// 对sqlIn语句进行处理
	sqlStr := "select category_id,category_name from category where category_id in (?)"
	var interfaceArr []  interface{}
	for _, c := range categoryIds {
		interfaceArr = append(interfaceArr, c)
	}
	insqlStr, params, err := sqlx.In(sqlStr, interfaceArr)
	if err != nil {
		logger.Error("sqlx faild , sql is %s,err is %#v", sqlStr, err)
		return
	}

	//查询
	var categoryList []*model.Category
	err = dao.DB.Select(&categoryList, insqlStr, params...)
	if err != nil {
		logger.Error("GeCategoryMap faild err : %#v, sqlstr is %s, categoryids are %#v", err, insqlStr, categoryIds)
		return
	}

	// 查询成功  构造map

	for _, v := range categoryList {
		categoryMap[v.CategoryId] = v
	}
	return
}
