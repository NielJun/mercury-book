package session

import (
	"sync"
)

// 用于内存存储的Session
type MemorySession struct {
	data   map[string]interface{}
	id     string
	flag   int          // 表示session当前状态的旗标 比如说 已修改 、 未保存、已经加载等
	rwLock sync.RWMutex // 读写所
}

// 构造函数
func NewMemorySession(id string) *MemorySession {
	return &MemorySession{
		data:   make(map[string]interface{}, 8),
		id:     id,
		rwLock: sync.RWMutex{},
	}
}

func (m *MemorySession) Get(key string) (value interface{}, err error) {
	m.rwLock.RLock()
	defer m.rwLock.RUnlock()
	value, ok := m.data[key]
	if !ok {
		err = ErrSessionKeyNotExist
		return
	}
	return
}

func (m *MemorySession) Set(key string, value interface{}) (err error) {
	m.rwLock.Lock()
	defer m.rwLock.Unlock()
	m.data[key] = value
	m.flag = SESSION_FLAG_MODIFY
	return
}

func (m *MemorySession) Delete(key string) (err error) {
	m.rwLock.Lock()
	defer m.rwLock.Unlock()
	m.flag = SESSION_FLAG_MODIFY
	delete(m.data, key)
	return
}

func (m *MemorySession) Save() (err error) {
	return
}

// 判断是否被修改过
func (m *MemorySession) IsModify() bool {
	if m.flag == SESSION_FLAG_MODIFY {
		return true
	}
	return false
}

// 获得当前的Id
func (m *MemorySession) Id() string {
	return m.id
}
