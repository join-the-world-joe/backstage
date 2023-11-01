package registry

import "time"

const (
	Id         = "id"
	Version    = "version"
	LastActive = "last_active"
	ServerPath = "server_path"
	Unknown    = "unknown"

	DefaultTTL      = 60 // 60 seconds
	DefaultInterval = 55 // 55 seconds

	DefaultRepostInterval = time.Minute
)

type Service struct {
	Id         string
	Ip         string
	Port       uint64
	Name       string
	Group      string
	Version    string
	LastActive string
}

type Registry interface {
	Name() string
	Register(*Service) error
	DeRegister(*Service) error
	ListServices(*Service) ([]*Service, error)
	Destroy()
}
