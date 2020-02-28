package dal

import (
	"fmt"
	"github.com/daniel/AnserBlock/dao"
	"github.com/daniel/AnserBlock/dao/model"
	"testing"
)

func init() {
	dns := "root:12345678@tcp(localhost:3306)/mercury?parseTime=true"
	err := dao.Init(dns)
	if err != nil {
		fmt.Printf("Init data base Err ! %v", err)
		return
	}

}

func TestCreatePostComment(t *testing.T) {

	comment := &model.Comment{}
	comment.Content = "这是评论内容"
	comment.CommentId = 12138
	comment.AuthorId = 12145
	comment.QuestionId = 290520334266793985
	comment.ParentId = 290520334266793985

	err := CreatePostComment(comment)

	if err != nil {
		t.Errorf("create pos comment faild")
		return
	}
}

func TestGetAuthorId(t *testing.T) {
	var commentId int64 = 12138

	  authorId,err := GetAuthorId(commentId)
	if err != nil {
		t.Logf("作者的ID 是 : %v ,err： %#v",authorId,err)

	}

	t.Logf("作者的ID 是 : %v ",authorId)

}

func TestCreatePostReplyComment(t *testing.T) {
	comment := &model.Comment{}
	comment.Content = "回复评论，二级的子评论"
	comment.CommentId = 290861371380203521
	err := CreatePostReplyComment(comment)

	if err != nil {
		t.Errorf("create pos comment faild")
		return
	}
}