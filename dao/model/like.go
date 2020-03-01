package model

const (
	LikeTypeAnswer  = 1 //回答的点赞
	LikeTypeComment = 2 //评论的点赞

)

type Like struct {
	Id       int64 `json:"id"`
	LikeType int   `json:"like_type"`
}
