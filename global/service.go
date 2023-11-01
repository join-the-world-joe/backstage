package global

import "backstage/abstract/registry"

var _service *registry.Service

func SetService(s *registry.Service) {
	_service = s
}

func Service() *registry.Service {
	return _service
}
