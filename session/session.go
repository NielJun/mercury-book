package session

// 服务器端的session令牌
type Session interface {
	Get(key string) (interface{}, error)
	Set(key string, value interface{}) error
	Delete(key string) error
	IsModify() bool
	Save() (err error)

	Id() string
}
