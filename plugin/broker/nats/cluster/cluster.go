package cluster

import (
	"backstage/abstract/broker"
	broker2 "backstage/common/broker"
	"fmt"
	"github.com/nats-io/nats.go"
	"sync"
)

type _broker struct {
	opts     *Options
	nopts    *nats.Options
	conn     *nats.Conn
	Mutex    sync.Mutex
	subjects sync.Map
	servers  []string
	user     string
	password string
}

func NewBroker(opts ...Option) (broker.Broker, error) {
	servers := []string{}
	user, password := "", ""

	options := Options{}

	for _, o := range opts {
		o(&options)
	}

	if len(options.servers) > 0 {
		servers = options.servers
	}

	if options.user != "" {
		user = options.user
	}

	if options.password != "" {
		password = options.password
	}

	return &_broker{
		user:     user,
		servers:  servers,
		opts:     &options,
		password: password,
		Mutex:    sync.Mutex{},
		nopts:    &nats.Options{Servers: servers, User: user, Password: password},
	}, nil
}

func (p *_broker) Name() string {
	return broker2.NATS_CLUSTER
}

func (p *_broker) Connect() error {
	conn, err := p.nopts.Connect()
	if err != nil {
		return err
	}
	p.conn = conn
	return nil
}

func (p *_broker) Disconnect() error {
	if !p.conn.IsClosed() {
		p.conn.Close()
	}
	return nil
}

func (p *_broker) Publish(topic string, msg []byte) error {
	return p.conn.Publish(topic, msg)
}

func (p *_broker) Subscribe(topic string, handle broker.Handler) error {
	subj, err := p.conn.Subscribe(topic, func(msg *nats.Msg) {
		handle(msg.Subject, msg.Data)
	})
	if err != nil {
		return err
	}

	p.subjects.Store(topic, subj)
	return nil
}

func (p *_broker) UnSubscribe(topic string) error {
	temp, exist := p.subjects.Load(topic)
	if exist {
		p.subjects.Delete(topic)
		return temp.(*nats.Subscription).Unsubscribe()
	}

	return fmt.Errorf(fmt.Sprintf("the subscribtion for %s doesn't exist", topic))
}

func (p *_broker) Debug() {
	p.subjects.Range(
		func(key, value interface{}) bool {
			fmt.Println("Key: ", key, "Val: ", value)
			return true
		},
	)
}
