package service

type Service interface {
	Name() string
	Init() error
	Start() error
	Stop()
	Run()
	Destroy()
}
