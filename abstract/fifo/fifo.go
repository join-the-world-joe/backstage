package fifo

type FIFO interface {
	Name() string
	Push(interface{}) error
	Pop() (interface{}, error)
	Channel() <-chan interface{}
	Destroy()
}
