package session

import (
	uuid "github.com/satori/go.uuid"
	"sync"
)

type MemorySessionManager struct {
	sessionMap map[string]Session
	rwLock     sync.RWMutex
}

// MemorySessionManager的构造函数
func NewMemorySessionManager() *MemorySessionManager {
	m := &MemorySessionManager{
		sessionMap: make(map[string]Session, 1024),
		rwLock:     sync.RWMutex{},
	}
	return m
}

func (sessionMgr *MemorySessionManager) Init(addr string, options ...string) (err error) {
	return
}

func (ssessionMgr *MemorySessionManager) Get(SessionId string) (s Session, err error) {

	ssessionMgr.rwLock.RLock()
	defer ssessionMgr.rwLock.RUnlock()

	s, ok := ssessionMgr.sessionMap[SessionId]
	if !ok {
		err = ErrSessionKeyNotExist
	}
	return
}

// 创建一个session
func (sessionMgr *MemorySessionManager) CreateSession() (session Session, err error) {

	sessionMgr.rwLock.Lock()
	defer sessionMgr.rwLock.Unlock()

	id := uuid.NewV4()
	sessionId := id.String()
	session = NewMemorySession(sessionId)

	sessionMgr.sessionMap[sessionId] = session
	return
}
