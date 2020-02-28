package dal

import (
	"fmt"
	"github.com/NielJun/go-logger"
	"github.com/daniel/AnserBlock/dao"
	"github.com/daniel/AnserBlock/dao/model"
)

// 对评论做一个数据库的存储操作
// 此操作是一个事务操作
func CreatePostComment(comment *model.Comment) (err error) {

	dbTx, err := dao.DB.Beginx()
	if err != nil {
		logger.Error("create post comment faild,err : %#v,comment: %#v", err, comment)
		return
	}
	sqlStr := "insert into comment (comment_id,content,author_id) values (?,?,?)"

	_, err = dbTx.Exec(sqlStr, comment.CommentId, comment.Content, comment.AuthorId)
	if err != nil {
		logger.Error("insert to comment faild,err : %#v,comment: %#v", err, comment)
		dbTx.Rollback() //回滚操作
		return
	}

	sqlStr = "insert into comment_rel (comment_id,parent_id,level,question_id,reply_author_id) values (?,?,?,?,?)"
	_, err = dbTx.Exec(sqlStr, comment.CommentId, comment.ParentId, 1, comment.QuestionId, comment.ReplyAuthorId)
	if err != nil {
		logger.Error("insert to comment_rel faild,err : %#v,comment: %#v", err, comment)
		dbTx.Rollback() //回滚操作
		return
	}
	// 提交事务
	err = dbTx.Commit()
	if err != nil {
		logger.Error("commit tx err,err : %#v,comment: %#v", err, comment)
		dbTx.Rollback() //回滚操作
		return
	}

	return
}

// 对评论做一个数据库的存储操作
// 此操作是一个事务操作
// 1. 取出评论id对应的作者id信息 2.插入数据到coment表 3.插入数据到关系表 4.提交
// 对评论进行回复操作
func CreatePostReplyComment(comment *model.Comment) (err error) {

	dbTx, err := dao.DB.Beginx()
	if err != nil {
		logger.Error("create post comment faild,err : %#v,comment: %#v", err, comment)
		return
	}
	var replyAuthorId int64
	sqlStr := "select author_id from comment where comment_id = ?"
	err = dbTx.Get(&replyAuthorId, sqlStr, comment.ReplyCommentId)
	if err != nil {
		logger.Error("db get author id faild ,sql string is %s,commment_id is %v,err is:%#v", sqlStr, comment.CommentId, err)
		return
	}

	// 取出来非法的Id
	if replyAuthorId == 0 {
		err = fmt.Errorf("invalid author id.")
		dbTx.Rollback()
		return
	}

	comment.ReplyAuthorId = replyAuthorId

	sqlStr = "insert into comment (comment_id,content,author_id) values (?,?,?)"

	_, err = dbTx.Exec(sqlStr, comment.CommentId, comment.Content, comment.AuthorId)
	if err != nil {
		logger.Error("insert to comment faild,err : %#v,comment: %#v", err, comment)
		dbTx.Rollback() //回滚操作
		return
	}

	sqlStr = "insert into comment_rel (comment_id,parent_id,level,question_id,reply_author_id,reply_comment_id) values (?,?,?,?,?,?)"
	_, err = dbTx.Exec(sqlStr, comment.CommentId, comment.ParentId, 2, comment.QuestionId, comment.ReplyAuthorId, comment.ReplyCommentId)
	if err != nil {
		logger.Error("insert to comment_rel faild,err : %#v,comment: %#v", err, comment)
		dbTx.Rollback() //回滚操作
		return
	}
	// 提交事务
	err = dbTx.Commit()
	if err != nil {
		logger.Error("commit tx err,err : %#v,comment: %#v", err, comment)
		dbTx.Rollback() //回滚操作
		return
	}

	return
}

// 通过评论的Id获得对应的坐着的id
func GetAuthorId(commentId int64) (authorId int64, err error) {

	sqlStr := "select author_id from comment where comment_id = ?"
	err = dao.DB.Get(&authorId, sqlStr, commentId)
	if err != nil {
		logger.Error("db get author id faild ,sql string is %s,commment_id is %v,err is:%#v", sqlStr, commentId, err)
		return
	}
	return
}
