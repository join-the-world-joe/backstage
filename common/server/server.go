package server

import "github.com/smallnest/rpcx/server"

func NewServer(opts ...Option) error {
	options := Options{}

	for _, o := range opts {
		o(&options)
	}

	srv := server.NewServer()
	err := srv.RegisterName(options.servicePath, options.server, "")
	if err != nil {
		return err
	}
	go srv.Serve("tcp", options.host + ":" + options.port)

	return nil
}
