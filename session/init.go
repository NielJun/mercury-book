package session

import (
	"fmt"
)

var (
	sessionMgr SessionManager
)

// 对启动时候Session管理器做一个配置初始化
// "memory"表示内存形式的session 管理器
// "Redis"表示Redis作为存储的管理器

func Init(provider string, addr string, options ...string) (err error) {
	switch provider {
	case "memory":
		sessionMgr = NewMemorySessionManager()
	case "redis":
		sessionMgr = NewRedisSessionManager()
	default:
		err = fmt.Errorf("Not support the type of provider")
		return
	}
	return
}

// 向外暴露两个方法 这是通过SessionId获得session
func Get(SessionId string) (session Session, err error) {
	return sessionMgr.Get(SessionId)
}

// 创建一个Session
func CreateSession() (session Session, err error) {
	return sessionMgr.CreateSession()
}
