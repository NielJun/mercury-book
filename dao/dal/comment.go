package dal

import (
	"fmt"
	"github.com/NielJun/go-logger"
	"github.com/daniel/AnserBlock/dao"
	"github.com/daniel/AnserBlock/dao/model"
	"github.com/jmoiron/sqlx"
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

	// 在发表评论的时候 评论条数进行增加
	sqlStr = "update answer set comment_count = comment_count+1 where answer_id = ?"
	_, err = dbTx.Exec(sqlStr, comment.QuestionId)
	if err != nil {
		logger.Error("update answer comment_count faild,err : %#v,comment: %#v", err, comment)
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

	// 在二级评论发表时 进行评论条数增加
	sqlStr = "update comment set comment_count = comment_count+1 where  comment_id = ?"
	_, err = dbTx.Exec(sqlStr, comment.ParentId)
	if err != nil {
		logger.Error("iupdate comment_count faild,err : %#v,comment: %#v", err, comment)
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

// 通过评论的Id获得对应的作者的id
func GetAuthorId(commentId int64) (authorId int64, err error) {

	sqlStr := "select author_id from comment where comment_id = ?"
	err = dao.DB.Get(&authorId, sqlStr, commentId)
	if err != nil {
		logger.Error("db get author id faild ,sql string is %s,commment_id is %v,err is:%#v", sqlStr, commentId, err)
		return
	}
	return
}

// 获得一级评论列表
// 1. 通过当前的answer_id 其实也是 question_id 查出当前的评论id列表
// 2. 通过sql.In语句查询出所有的评论信息
func GetCommentList(answerId int64, offset, limit int64) (commentList []*model.Comment, count int64, err error) {

	// 1.查询出id列表
	var commentIdList [] int64
	sqlStr := "select comment_id from comment_rel where question_id = ? and level = ? limit ?,?"
	err = dao.DB.Select(&commentIdList, sqlStr, answerId, 1, offset, limit)
	if err != nil {
		logger.Error("get comment list faild,sqlstr is %s,answer_id is %v", sqlStr, answerId)
		return
	}
	if len(commentIdList) == 0 {
		return
	}

	// 2.通过id列表查询所有的评论信息

	sqlStr = "select comment_id,content,author_id,like_count,comment_count,create_time from comment where comment_id in (?)"

	var commentIdInterfaceList [] interface{}
	for _, commentId := range commentIdList {
		commentIdInterfaceList = append(commentIdInterfaceList, commentId)
	}

	sqlStr, params, err := sqlx.In(sqlStr, commentIdInterfaceList)
	if err != nil {
		logger.Error("sql in faild. sqlstr is %s,err : %#v", sqlStr, err)
		return
	}
	err = dao.DB.Select(&commentList, sqlStr, params...)
	if err != nil {
		logger.Error("sql select faild. sqlstr is %s ,err : %#v", sqlStr, err)
		return
	}

	// 查询评论条数
	sqlStr = "select count(comment_id) from comment_rel where question_id = ? and level = ?"
	err = dao.DB.Get(&count, sqlStr, answerId, 1)
	if err != nil {
		logger.Error("get comment count faild,sqlstr is %s,answer_id is %v", sqlStr, answerId)
		return
	}
	return

}

// 获得二级评论列表
// 1. 通过当前的answer_id 其实也是 question_id 查出当前的评论id列表
// 2. 通过sql.In语句查询出所有的评论信息
func GetRePlyCommentList(answerId int64, offset, limit int64) (commentList []*model.Comment, count int64, err error) {

	// 1.查询出id列表
	var commentIdList [] int64
	sqlStr := "select comment_id from comment_rel where parent_id = ? and level = ? limit ?,?"
	err = dao.DB.Select(&commentIdList, sqlStr, answerId, 2, offset, limit)
	if err != nil {
		logger.Error("get comment list faild,sqlstr is %s,answer_id is %v", sqlStr, answerId)
		return
	}
	if len(commentIdList) == 0 {
		return
	}

	// 2.通过id列表查询所有的评论信息

	sqlStr = "select comment_id,content,author_id,like_count,comment_count,create_time from comment where comment_id in (?)"

	var commentIdInterfaceList [] interface{}
	for _, commentId := range commentIdList {
		commentIdInterfaceList = append(commentIdInterfaceList, commentId)
	}

	sqlStr, params, err := sqlx.In(sqlStr, commentIdInterfaceList)
	if err != nil {
		logger.Error("sql in faild. sqlstr is %s,err : %#v", sqlStr, err)
		return
	}
	err = dao.DB.Select(&commentList, sqlStr, params...)
	if err != nil {
		logger.Error("sql select faild. sqlstr is %s ,err : %#v", sqlStr, err)
		return
	}

	// 查询评论条数
	sqlStr = "select count(comment_id) from comment_rel where question_id = ? and level = ?"
	err = dao.DB.Get(&count, sqlStr, answerId, 2)
	if err != nil {
		logger.Error("get comment count faild,sqlstr is %s,answer_id is %v", sqlStr, answerId)
		return
	}
	return

}


// 记录点赞数亩
func UpdateCommentCountCount(commentId int64) (err error) {

	sqlStr := "update comment set like_count = like_count+1 where comment_id = ?"
	_, err = dao.DB.Exec(sqlStr, commentId)
	if err != nil {
		logger.Error("update answer voteup_count faild, err: %#v", err)
		return
	}
	return
}

