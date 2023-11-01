package rpc

import (
	"errors"
	"fmt"
	"github.com/smallnest/rpcx/client"
	"sync"
)

var g_xclient_map sync.Map // it holds all connected xclient

func store(unique string, client client.XClient) {
	g_xclient_map.Store(unique, client)
}

func load(unique string) (client.XClient, error) {
	value, ok := g_xclient_map.Load(unique)
	if ok {
		return value.(client.XClient), nil
	}
	return nil, errors.New(fmt.Sprintf("%s doesn't exist", unique))
}

func Dump() {
	a := g_xclient_map
	fmt.Println("a = ", a)
}
