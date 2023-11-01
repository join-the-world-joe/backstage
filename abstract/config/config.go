package config

const (
	DefaultTimeout = 5000 // 5 seconds
)

type Parameter struct {
	// for file
	FileName string
	Reload   bool

	// for Nacos
	DataId   string
	Group    string
	Content  string
	OnChange func(namespace, group, dataId, data string)
}

type Config interface {
	Name() string
	Publish(Parameter) error
	Subscribe(Parameter) error
	UnSubscribe(Parameter) error
	Load(Parameter) ([]byte, error)
	Destroy()
}