package nacos

import (
	"backstage/abstract/registry"
	"errors"
	"github.com/nacos-group/nacos-sdk-go/v2/clients"
	"github.com/nacos-group/nacos-sdk-go/v2/clients/naming_client"
	"github.com/nacos-group/nacos-sdk-go/v2/vo"
	"time"
)

const (
	Name = "Nacos Registry"
)

type _registry struct {
	opts *Options
	nc   naming_client.INamingClient
}

func NewRegistry(opts ...Option) (registry.Registry, error) {
	options := Options{}

	for _, o := range opts {
		o(&options)
	}

	nc, err := clients.NewNamingClient(options.param)
	if err != nil {
		return nil, err
	}

	return &_registry{
		nc:   nc,
		opts: &options,
	}, nil
}

func (p *_registry) Name() string {
	return Name
}

func (p *_registry) Register(service *registry.Service) error {
	b, err := p.nc.RegisterInstance(
		vo.RegisterInstanceParam{
			Ip:          service.Ip,
			Port:        service.Port,
			Weight:      10,
			Enable:      true,
			Healthy:     true,
			ServiceName: service.Name,
			GroupName:   service.Group,
			Metadata:    map[string]string{registry.Id: service.Id, registry.Version: service.Version, registry.LastActive: time.Now().Format(time.RFC3339)},
			Ephemeral:   false,
		},
	)
	if err != nil {
		return err
	}
	if !b {
		return errors.New("register doesn't success")
	}
	return nil
}

func (p *_registry) DeRegister(service *registry.Service) error {
	b, err := p.nc.DeregisterInstance(
		vo.DeregisterInstanceParam{
			Ip:          service.Ip,
			ServiceName: service.Name,
			Port:        service.Port,
			GroupName:   service.Group,
		},
	)
	if err != nil {
		return err
	}
	if !b {
		return errors.New("DeRegister doesn't success")
	}
	return nil
}

func (p *_registry) ListServices(service *registry.Service) ([]*registry.Service, error) {
	list, err := p.nc.SelectAllInstances(
		vo.SelectAllInstancesParam{
			GroupName:   service.Group,
			ServiceName: service.Name,
		},
	)
	if err != nil {
		return nil, err
	}
	NodeList := make([]*registry.Service, 0, len(list))
	for _, v := range list {
		NodeList = append(NodeList, &registry.Service{
			Ip:   v.Ip,
			Port: v.Port,
			Name: v.ServiceName,
			Id: func() string {
				if version, exist := v.Metadata[registry.Id]; exist {
					return version
				}
				return registry.Unknown
			}(),
			Version: func() string {
				if version, exist := v.Metadata[registry.Version]; exist {
					return version
				}
				return registry.Unknown
			}(),
			LastActive: func() string {
				if lastActive, exist := v.Metadata[registry.LastActive]; exist {
					return lastActive
				}
				return registry.Unknown
			}(),
		})
	}
	return NodeList, nil
}

func (p *_registry) Destroy() {
	p.nc.CloseClient()
}
