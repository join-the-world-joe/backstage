package aliyun_oss

type Options struct {
	accessKeyId     string
	accessKeySecret string
	endpoint        string
}

type Option func(*Options)

func WithAccessKeyId(id string) Option {
	return func(o *Options) {
		o.accessKeyId = id
	}
}

func WithAccessKeySecret(secret string) Option {
	return func(o *Options) {
		o.accessKeySecret = secret
	}
}

func WithEndpoint(endpoint string) Option {
	return func(o *Options) {
		o.endpoint = endpoint
	}
}
