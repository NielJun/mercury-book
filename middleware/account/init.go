package account

import "github.com/daniel/AnserBlock/session"

func Init(provider string, addr string, options ...string) (err error) {
	return session.Init(provider, addr, options...)
}
