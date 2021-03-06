package model

import "time"

// 问题评论的表结构体映射
type Answer struct {
	AnswerId     int64     `json:"answer_id" db:"answer_id"`
	Content      string    `json:"content" db:"content"`
	CommentCount int32     `json:"comment_count" db:"comment_count"`
	VoteupCount  int32     `json:"like_count" db:"like_count"`
	AuthorId     int64     `json:"author_id" db:"author_id"`
	Status       int32     `json:"status" db:"status"`
	CanComment   int32     `json:"can_comment" db:"can_comment"`
	CreateTime   time.Time `json:"create_time" db:"create_time"`
	UpdateTime   time.Time `json:"update_time" db:"update_time"`
	QuestionId   string    `json:"question_id"`
}

// 返回给前端的评论结构体单元
type ResponseAnswer struct {
	Answer
	AuthorName string `json:"author_name"`
}

// 返回各级结构体的数据块
type ResponseAnswerList struct {
	AnswerList [] *ResponseAnswer `json:"answer_list"`
	TotalCount int32              `json:"total_count"`
}
