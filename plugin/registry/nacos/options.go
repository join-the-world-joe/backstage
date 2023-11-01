package nacos

import "github.com/nacos-group/nacos-sdk-go/v2/vo"

type Options struct {
	param vo.NacosClientParam
}

type Option func(*Options)

func WithNacosClientParam(param vo.NacosClientParam) Option {
	return func(o *Options) {
		o.param = param
	}
}