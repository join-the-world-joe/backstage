package locker

import "time"

type Context struct {
	Id           string
	UUID         string
	Prefix       string
	From         string
	FromInMS     int64
	To           string
	ToInMS       int64
	ExtendTo     string
	ExtendToInMS int64
	Signature    string
}

type MLock interface {
	TryLock(id string, sec time.Duration) (*Context, error)
	Lock(id string, holdingTime time.Duration, lockTimeout time.Duration, retry time.Duration) (*Context, error)
	Refresh(ctx *Context, addition time.Duration) error
	Unlock(ctx *Context) error
}
