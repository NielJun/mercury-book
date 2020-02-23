package session

import (
	"encoding/json"
	"github.com/garyburd/redigo/redis"
	"sync"
)

const (
	SESSION_FLAG_NONE = iota
	SESSION_FLAG_MODIFY
	SESSION_FLAG_LOAD
)

// 通过redis春初的session令牌
type RedisSession struct {
	sessionId  string                 // session的ID
	pool       *redis.Pool            // 用于通信的redis迟
	sessionMap map[string]interface{} // 存储用的session map
	flag       int                    // 表示session当前状态的旗标 比如说 已修改 、 未保存、已经加载等
	rwLock     sync.RWMutex           // 读写锁🔒
}

// Redis令牌的构造函数
func NewRedisSession(id string, pool *redis.Pool) *RedisSession {
	s := &RedisSession{
		sessionId:  id,
		pool:       pool,
		sessionMap: make(map[string]interface{}),
		flag:       SESSION_FLAG_LOAD,
		rwLock:     sync.RWMutex{},
	}
	return s
}

// 从数据集离main取出特定key的数据
func (s *RedisSession) Get(key string) (value interface{}, err error) {
	s.rwLock.RLock()
	defer s.rwLock.RUnlock()

	// 实现延迟加载
	if s.flag == SESSION_FLAG_NONE {

		err := s.loadFromRedis()
		if err != nil {
			return nil, err
		}
	}
	value, ok := s.sessionMap[key]
	if !ok {
		err = ErrKeyNotExistInSession
		return
	}
	return
}

// 从redis里面加载数据集
// 加载好的数据集放在s里面
func (s *RedisSession) loadFromRedis() (err error) {

	// 通过pool获得链接
	conn := s.pool.Get()
	reply, err := conn.Do("GET", s.sessionId)
	if err != nil {
		return
	}

	// 从redis里取出数据
	data, err := redis.String(reply, err)
	if err != nil {
		return
	}
	json.Unmarshal([]byte(data), &s.sessionMap)
	if err != nil {
		return
	}
	return
}

// 两种方案
// 方案一 ： 来个请求就立即刷新
// 方案二 ： 等请求处理完成再刷新（ 批量处理 ） 通过Save函数进行刷新
func (s *RedisSession) Set(key string, value interface{}) error {

	// 方案二
	s.rwLock.Lock()
	defer s.rwLock.Unlock()
	s.sessionMap[key] = value
	s.flag = SESSION_FLAG_MODIFY
	return nil
}

// 删除对应key的数据
func (s *RedisSession) Delete(key string) error {

	s.rwLock.RLock()
	defer s.rwLock.Unlock()

	delete(s.sessionMap, key)
	return nil
}

// 把修改后的数据存储在redis里面
// 1. 判断是不是修改状态
// 2. 将将要写入的数据序列化
// 3. 获取链接 写入数据
func (s *RedisSession) Save() (err error) {
	s.rwLock.Lock()
	defer s.rwLock.Unlock()

	if s.flag != SESSION_FLAG_MODIFY {
		return
	}
	// 把修改好的map转化为json
	data, err := json.Marshal(s.sessionMap)
	if err != nil {
		return
	}

	conn := s.pool.Get()
	_, err = conn.Do("SET", s.sessionId, data)
	if err != nil {
		return
	}
	return
}

// 判断是否被修改过
func (r *RedisSession) IsModify() bool {
	if r.flag == SESSION_FLAG_MODIFY {
		return true
	}
	return false
}

// 获得当前的Id
func (r * RedisSession)Id() string  {
	return r.sessionId
}