package session

import "errors"

var (

	ErrSessionNotExist = errors.New("session Not Exit")
	ErrSessionKeyNotExist = errors.New("session Key Not Exit")
	ErrKeyNotExistInSession = errors.New(" Key Not In The session ")
)
