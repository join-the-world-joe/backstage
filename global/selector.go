package global

import (
	"backstage/abstract/registry"
	"backstage/abstract/selector"
	service2 "backstage/common/macro/config"
	selector2 "backstage/lib/selector"
	"encoding/json"
	"errors"
	"fmt"
	"sync"
)

var g_service_lock sync.Mutex // for serviceList
var nextList sync.Map
var discoveringService sync.Map
var serviceList = func() map[string][]*registry.Service {
	return make(map[string][]*registry.Service)
}()

func nextService(next selector.Next) (*registry.Service, error) {
	ss, err := next()
	if err != nil {
		return nil, err
	}
	if len(ss) <= 0 {
		return nil, errors.New("no service could be found")
	}
	return toService(ss)
}

func toService(srvInfo string) (*registry.Service, error) {
	srv := new(registry.Service)
	err := json.Unmarshal([]byte(srvInfo), srv)
	if err != nil {
		return nil, err
	}
	return srv, nil
}

func SelectService(name string) (*registry.Service, error) {
	if temp, ok := nextList.Load(name); ok {
		return nextService(temp.(selector.Next))
	} else {
		g_service_lock.Lock()
		defer g_service_lock.Unlock()

		if temp, ok = nextList.Load(name); ok {
			return nextService(temp.(selector.Next))
		}

		if _, ok := discoveringService.Load(name); ok {
			return nil, errors.New(fmt.Sprintf("SelectService: discovering service instances[%s]", name))
		}

		// service discovery
		if svl2, err := Registry().ListServices(
			&registry.Service{
				Group: service2.ServiceGroup,
				Name:  name,
			},
		); err == nil {
			subscribe(name)
			if len(svl2) <= 0 {
				discoveringService.Store(name, nil)
				return nil, errors.New(fmt.Sprintf("SelectService: no service instances[%s] found", name))
			}
			ss, err := setget(name, updateServiceList(name, svl2))
			if err != nil {
				return nil, err
			}
			return toService(ss)
		}
		return nil, errors.New(fmt.Sprintf("Service %s without next()", name))
	}
}

func setget(name string, svl1 []*registry.Service) (string, error) {
	// non-thread-save ...
	m := make(map[string]int)
	for _, ss := range svl1 {
		bytes, err := json.Marshal(&ss)
		if err != nil {
			return "", err
		}
		m[string(bytes)] = 0
	}
	slt := selector2.NewSelector()
	slt.Set(name, m)
	nx, err := slt.Select(name, selector2.WithRoundRobin())
	if err != nil {
		return "", nil
	}
	nextList.Store(name, nx)
	return nx()
}

func updateServiceList(name string, l []*registry.Service) []*registry.Service {
	// non-thread-save ...
	if len(l) <= 0 {
		delete(serviceList, name)
		return l
	}
	serviceList[name] = l
	return l
}
