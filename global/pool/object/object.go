package object

import (
	"sync"
)

var g_object_pool *sync.Pool

type Object struct {
	Id int32
}

func init() {
	g_object_pool = &sync.Pool{
		New: func() interface{} {
			return &Object{}
		},
	}
}

func Push(p *Object) {
	g_object_pool.Put(p)
}

func Pop() *Object {
	return g_object_pool.Get().(*Object)
}
