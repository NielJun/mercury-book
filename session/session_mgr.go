package session

type SessionManager interface {

	Get(SessionId string) (session Session, err error)
	CreateSession() (session Session, err error)
	Init(addr string,options ... string)(err error)
}