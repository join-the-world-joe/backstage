package server

type Options struct {
	server  string
	user     string
	password string
}

type Option func(*Options)

func WithServer(server string) Option {
	return func(o *Options) {
		o.server = server
	}
}

func WithAuth(user, password string) Option {
	return func(o *Options) {
		o.user = user
		o.password = password
	}
}