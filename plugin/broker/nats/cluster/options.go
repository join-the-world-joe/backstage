package cluster

type Options struct {
	servers  []string
	user     string
	password string
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
