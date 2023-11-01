package routing

import (
	"backstage/abstract/registry"
	"errors"
	"fmt"
	"sync"
)

var _routing sync.Map

func Store(service string, bind *registry.Service) {
	_routing.Store(service, bind)
}

func Load(service string) (*registry.Service, error) {
	value, ok := _routing.Load(service)
	if ok {
		return value.(*registry.Service), nil
	}
	return nil, errors.New(fmt.Sprintf("%s doesn't exist", service))
}
