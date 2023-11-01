package broker

type Handler func(topic string, msg []byte)

type Broker interface {
	Name() string
	Connect() error
	Disconnect() error
	Publish(topic string, msg []byte) error
	Subscribe(topic string, handle Handler) error
	UnSubscribe(topic string) error
	Debug()
}