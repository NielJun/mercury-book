package dal

import (
	"github.com/NielJun/go-logger"
	"github.com/daniel/AnserBlock/dao"
	"github.com/daniel/AnserBlock/dao/model"
	"github.com/jmoiron/sqlx"
)

// 根据指定的问题的id和查询的开始位置和查询的条数查询该问题对应的 answer id的列表
func GetAnswerIdList(questionId int64, offset, limit int64) (answerIdList []int64, err error) {

	sqlStr := "select answer_id from question_answer_rel where question_id = ? limit ?,?"
	err = dao.DB.Select(&answerIdList, sqlStr, questionId, offset, limit)
	if err != nil {
		logger.Error("get answer id list faild,sql str is %s ,questionid is %v,offset is %v, limit is %v",
			sqlStr, questionId, offset, limit)
		return
	}

	return
}

/*

	AnswerId     int64     `json:"answer_id" db:"answer_id"`
	Content      string    `json:"content" db:"content"`
	CommentCount int32     `json:"comment_count" db:"comment_count"`
	VoteupCount  int32     `json:"voteup_count" db:"voteup_count"`
	AuthorId     int64     `json:"author_id" db:"author_id"`
	Status       int32     `json:"status" db:"status"`
	CanComment   int32     `json:"can_comment" db:"can_comment"`
	CreateTime   time.Time `json:"create_time" db:"create_time"`
	UpdateTime   time.Time `json:"update_time" db:"update_time"`
	QuestionId   string    `json:"question_id"`
*/

// 根据回答的id切片查询对应的answer列表
func GetAnwerList(answerIds []int64) (answerList [] *model.Answer, err error) {

	// 对sqlIn语句进行处理
	sqlStr := "select answer_id,content,comment_count,voteup_count,author_id,status,can_comment,create_time,update_time from answer where answer_id in (?)"
	var interfaceArr []  interface{}
	for _, c := range answerIds {
		interfaceArr = append(interfaceArr, c)
	}
	insqlStr, params, err := sqlx.In(sqlStr, interfaceArr)
	if err != nil {
		logger.Error("sqlx faild , sql is %s,err is %#v", sqlStr, err)
		return
	}

	//查询
	err = dao.DB.Select(&answerList, insqlStr, params...)
	if err != nil {
		logger.Error("GetAnwerMap faild err : %#v, sqlstr is %s, answerIds are %#v", err, insqlStr, answerIds)
		return
	}

	return
}

//获取总共的评论条数
func GetTotalAnswerCount(questionId int64) (totalCount int32, err error) {

	sqlStr := "select count(answer_id) from question_answer_rel where question_id = ?"

	err = dao.DB.Get(&totalCount, sqlStr, questionId)
	if err != nil {
		logger.Error("get question answer total count faild,err : %#v, sql str is %s,question id is %v", err, sqlStr, questionId)
		return
	}
	return

}
