package jet_stream

type Options struct {
	servers  []string
	user     string
	password string
	param    map[string]interface{}
}

type Option func(*Options)

func WithServers(servers []string) Option {
	return func(o *Options) {
		o.servers = servers
	}
}

func WithAuth(user, password string) Option {
	return func(o *Options) {
		o.user = user
		o.password = password
	}
}

func WithParam(param map[string]interface{}) Option {
	return func(o *Options) {
		o.param = param
	}
}