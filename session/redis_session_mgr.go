package session

import (
	"github.com/garyburd/redigo/redis"
	uuid "github.com/satori/go.uuid"
	"sync"
	"time"
)

type RedisSessionManager struct {
	addr       string             // 链接redis的地址
	pwd        string             // 链接redis的密码
	pool       *redis.Pool        // redis的链接池
	sessionMap map[string]Session // 存储session的map
	rwLock     sync.RWMutex       // 读写锁🔒
}

// RedisSessionManager的构造函数
func NewRedisSessionManager(/*addr string,pwd string,*/) *RedisSessionManager {
	//r := &RedisSessionManager{
	//	addr:       "",
	//	pwd:        "",
	//	pool:       nil,
	//	sessionMap: nil,
	//	rwLock:     sync.RWMutex{},
	//}
	return nil
}

// 通过指定的sessionId取得对应的Session对象
func (r *RedisSessionManager) Get(sessionId string) (s Session, err error) {
	r.rwLock.RLock()
	defer r.rwLock.RUnlock()

	s, ok := r.sessionMap[sessionId]
	if !ok  {
		err = ErrSessionKeyNotExist
		return
	}
	return

}

// 创建一个session对象
func (r *RedisSessionManager) CreateSession() (s Session, err error) {
	r.rwLock.Lock()
	defer r.rwLock.Unlock()
	id := uuid.NewV4()

	sessionId := id.String()
	s = NewRedisSession(sessionId, r.pool)
	r.sessionMap[sessionId] = s
	return

}

// 在redis催出链接时候进行初始化
func (r *RedisSessionManager) Init(addr string, options ...string) (err error) {
	if len(options) > 0 {
		r.pwd = options[0]
	}

	r.pool = NewRedisPool(addr, r.pwd)
	r.addr = addr
	return
}

// 新生成一个RedisPool
func NewRedisPool(serverAddr, pwd string) *redis.Pool {
	return &redis.Pool{
		Dial: func() (conn redis.Conn, e error) {
			connect, err := redis.Dial("tcp", serverAddr)
			if err != nil {
				return nil, err;
			}
			return connect, err;
		},
		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			if time.Since(t) < time.Minute {
				return nil
			}
			_, err := c.Do("PING")
			return err
		},
		MaxIdle:     64,
		MaxActive:   1024,
		IdleTimeout: 240 * time.Second,
		Wait:        false,
	}
}

