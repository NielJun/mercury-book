package model

import "time"

// 结构体model
type Question struct {
	QuestionId    int64     `json:"question_id_number" db:"question_id"`
	Caption       string    `json:"caption" db:"caption"`
	Content       string    `json:"content" db:"content"`
	AuthorId      int64     `json:"author_id_number" db:"author_id"`
	CategoryId    int64     `json:"category_id" db:"category_id"`
	Status        int32     `json:"status" db:"status"`
	CreateTime    time.Time `json:"-" db:"create_time"`
	CreateTimeStr string    `json:"create_time"`
	QuestionIdStr string    `json:"question_id"`
	AuthorIdStr   string    `json:"author_id"`
}

// 回应给客户端的列表数据结构单元
// 通过查询和聚合作者名字返回给客户端
type ResponseQuestion struct {
	Question
	AuthorName string `json:"author_name"`
}

// 回应给客户端问题细节
type QuestionDetail struct {
	Question
	CategoryName string `json:"name" db:"category_name"`
	AuthorName string `json:"author_name"`
}