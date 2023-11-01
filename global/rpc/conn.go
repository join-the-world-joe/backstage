package rpc

import (
	"github.com/smallnest/rpcx/client"
	"sync"
)

var g_lock sync.Mutex

func GetXClient(name, id, host, port string) (client.XClient, error) {
	fullName := name
	if len(id) > 0 {
		fullName = fullName + "-" + id
	}

	_client, err := load(fullName)
	if err == nil { // created
		return _client, nil
	}

	g_lock.Lock()
	defer g_lock.Unlock()

	_client, err = load(fullName)
	if err == nil { // created
		return _client, nil
	}

	_client, err = connectToRPCServer(
		host,
		port,
		name,
	)
	if err != nil {
		return nil, err
	}
	store(fullName, _client)
	return _client, nil
}

func connectToRPCServer(host, port, servicePath string) (client.XClient, error) {
	d, err := client.NewPeer2PeerDiscovery("tcp@"+host+":"+port, "")
	if err != nil {
		return nil, err
	}
	return client.NewXClient(servicePath, client.Failfast, client.RandomSelect, d, client.DefaultOption), nil
}
