package broker

import (
	"backstage/abstract/broker"
	"backstage/abstract/selector"
	"backstage/global/log"
	selector2 "backstage/lib/selector"
	"errors"
	"fmt"
	"sync"
)

var g_lock sync.Mutex

// it holds the context of all connected brokers
var next selector.Next
var sel = selector2.NewSelector()

var _g_broker_map = func() map[string]broker.Broker {
	return make(map[string]broker.Broker)
}()

func store(unique string, bkr broker.Broker) {
	g_lock.Lock()
	defer g_lock.Unlock()
	_g_broker_map[unique] = bkr
}

func load(unique string) (broker.Broker, error) {
	g_lock.Lock()
	defer g_lock.Unlock()
	if brk, exist := _g_broker_map[unique]; exist {
		return brk, nil
	}
	return nil, errors.New(fmt.Sprintf("%s doesn't exist", unique))
}

func remove(unique string) {
	g_lock.Lock()
	defer g_lock.Unlock()
	delete(_g_broker_map, unique)
}

func clear() {
	g_lock.Lock()
	defer g_lock.Unlock()
	for k, _ := range _g_broker_map {
		delete(_g_broker_map, k)
	}
}

func DumpBroker() {
	g_lock.Lock()
	defer g_lock.Unlock()
	for _, v := range _g_broker_map {
		log.WarnF("Broker: %v Connected", v.Name())
	}
}
