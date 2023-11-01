package global

import (
	"backstage/abstract/registry"
	"backstage/common/macro/config"
	"backstage/common/macro/service"
	"backstage/global/log"
	"encoding/json"
	"fmt"
	"sync"
	"time"
)

var _registry registry.Registry

func SetRegistry(r registry.Registry) {
	_registry = r
}

func Registry() registry.Registry {
	return _registry
}

var subscriptionCounter int
var _subscription_lock sync.Mutex               // for reg
var subscription = func() map[string]struct{} { // map[service_name]
	return make(map[string]struct{})
}()

var subChan = func() chan string {
	return make(chan string, service.DefaultRegistrySubscribeBufferSize)
}()

func init() {
	go func() {
		for {
			select {
			case name := <-subChan:
				if l, err := Registry().ListServices(
					&registry.Service{
						Group: config.ServiceGroup,
						Name:  name,
					},
				); err == nil {
					log.Debug(fmt.Sprintf("Registry.ServiceList(%s): %v", name, l))
					g_service_lock.Lock()
					updateServiceList(name, l)
					g_service_lock.Unlock()
					if len(l) <= 0 {
						nextList.Delete(name)
						discoveringService.Store(name, nil)
						repost(name)
						break
					}
					setget(name, l)
					discoveringService.Delete(name)
					repost(name)
					break
				} else {
					log.Error(err.Error())
				}
			}
		}
	}()
}

func subscribe(name string) {
	_subscription_lock.Lock()
	defer _subscription_lock.Unlock()

	if _, exist := subscription[name]; !exist {
		subscription[name] = struct{}{}
		subChan <- name
		subscriptionCounter++
	}
}

func repost(name string) {
	go func() {
		<-time.After(registry.DefaultRepostInterval)
		subChan <- name
	}()
}

func DumpRegistry() {
	_subscription_lock.Lock()
	defer _subscription_lock.Unlock()
	bytes, err := json.Marshal(subscription)
	if err != nil {
		log.Error(err.Error())
		return
	}
	log.Warn(fmt.Sprintf("Subscription(%v): %v", subscriptionCounter, string(bytes)))
	g_service_lock.Lock()
	defer g_service_lock.Unlock()
	for k1, v1 := range serviceList {
		log.WarnF("Service %v: ", k1)
		for _, v2 := range v1 {
			if bs, err := json.Marshal(v2); err == nil {
				log.Warn(string(bs))
			}
		}
	}
}
