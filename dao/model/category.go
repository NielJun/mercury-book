package model

// 提交页面分类标签的数据结构
type Category struct {
	CategoryName string `json:"name" db:"category_name"`
	CategoryId   int64  `json:"id" db:"category_id"`
}
