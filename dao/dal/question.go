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
