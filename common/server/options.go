package server


type Options struct {
	servicePath string
	host string
	port string
	server interface{}
}

type Option func(*Options)

func WithServicePath(servicePath string) Option {
	return func(o *Options) {
		o.servicePath = servicePath
	}
}

func WithHost(host string) Option {
	return func(o *Options) {
		o.host = host
	}
}

func WithPort(port string) Option {
	return func(o *Options) {
		o.port = port
	}
}

func WithServer(server interface{}) Option {
	return func(o *Options) {
		o.server = server
	}
}