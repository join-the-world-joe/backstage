package nacos

import (
	"backstage/abstract/config"
	"errors"
	"github.com/nacos-group/nacos-sdk-go/v2/clients"
	"github.com/nacos-group/nacos-sdk-go/v2/clients/config_client"
	"github.com/nacos-group/nacos-sdk-go/v2/vo"
)

const (
	Name = "Nacos"
)

type _config struct {
	opts *Options
	cc   config_client.IConfigClient
}

func NewConfig(opts ...Option) (config.Config, error) {
	options := Options{}

	for _, o := range opts {
		o(&options)
	}

	cc, err := clients.NewConfigClient(options.param)
	if err != nil {
		return nil, err
	}

	return &_config{
		cc:   cc,
		opts: &options,
	}, nil
}

func (p *_config) Name() string {
	return Name
}

func (p *_config) Publish(param config.Parameter) error {
	b, err := p.cc.PublishConfig(vo.ConfigParam{
		Group:   param.Group,
		DataId:  param.DataId,
		Content: param.Content,
	})
	if err != nil {
		return err
	}
	if !b {
		return errors.New("PublishConfig doesn't success")
	}
	return nil
}

func (p *_config) Subscribe(param config.Parameter) error {
	return p.cc.ListenConfig(vo.ConfigParam{
		Group:    param.Group,
		DataId:   param.DataId,
		OnChange: param.OnChange,
	})
}

func (p *_config) UnSubscribe(param config.Parameter) error {
	return p.cc.CancelListenConfig(vo.ConfigParam{
		Group:  param.Group,
		DataId: param.DataId,
	})
}

func (p *_config) Load(param config.Parameter) ([]byte, error) {
	temp, err := p.cc.GetConfig(vo.ConfigParam{
		Group:  param.Group,
		DataId: param.DataId,
	})
	if err != nil {
		return nil, err
	}
	return []byte(temp), nil
}

func (p *_config) Destroy() {
	p.cc.CloseClient()
}
