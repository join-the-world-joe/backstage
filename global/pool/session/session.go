package session

import (
	"go-micro-framework/common/payload"
	"sync"
)

var g_session_pool *sync.Pool

func init() {
	g_session_pool = &sync.Pool{
		New: func() interface{} {
			return &payload.Session{}
		},
	}
}

func Push(p *payload.Session) {
	g_session_pool.Put(p)
}

func Pop() *payload.Session {
	return g_session_pool.Get().(*payload.Session)
}
