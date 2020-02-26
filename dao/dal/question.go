package dal

import (
	"github.com/NielJun/go-logger"
	"github.com/daniel/AnserBlock/dao"
	"github.com/daniel/AnserBlock/dao/model"
)

// 在数据库中存储question的操作
func CreateQuestion(question *model.Question) (err error) {

	sqlstr := "insert into question (question_id,caption,content,author_id,category_id) values (?,?,?,?,?)"
	_, err = dao.DB.Exec(sqlstr, question.QuestionId, question.Caption, question.Content, question.AuthorId, question.CategoryId)
	if err != nil {
		logger.Error("写入 question 表数据出错 %#v", err)
	}
	return
}

/*
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
*/

// 在数据库中根据categoryId[分类的Id]查询所有的Question
func GetQuestionList(categoryId int64) (questionList []*model.Question, err error) {

	sqlstr := "select question_id,caption,content,author_id,category_id,create_time from question where category_id = ?"
	//select用来做分组查询【查询一个链表】
	err = dao.DB.Select(&questionList, sqlstr, categoryId)
	if err != nil {
		return
	}
	return
}

// 通过questionid查询Question的详情
func GetQuestion(questionId int64) (question *model.Question, err error) {
	question = &model.Question{}
	sqlstr := `select 
							question_id, caption, content, author_id, category_id, create_time
						from 
							question
						where question_id=?`

	err = dao.DB.Get(question, sqlstr, questionId)
	if err != nil {
		//.Error("get question  failed, err:%v", err)
		return
	}

	return
}
