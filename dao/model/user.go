package model

// 登录注册相关的数据结构
type UserInfo struct {
	UserId   uint64  `json:"user_id" db:"user_id"`
	NickName string `json:"nickname" db:"nickname"`
	Sex      int    `json:"sex" db:"sex"`
	UserName string `json:"user" db:"username"`
	Email    string `json:"email" db:"email"`
	Password string `json:"password" db:"password"`
}

const (

	Boy = 1
	Girl = 2

)