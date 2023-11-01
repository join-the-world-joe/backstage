package jet_stream

import (
	"backstage/abstract/broker"
	broker2 "backstage/common/broker"
	"encoding/json"
	"fmt"
	"github.com/nats-io/nats.go"
	"sync"
)

type _broker struct {
	opts        *Options
	nopts       *nats.Options
	conn        *nats.Conn
	Mutex       sync.Mutex
	subjectsMap sync.Map
	servers     []string
	user        string
	password    string
	js          nats.JetStreamContext

	// contextual variables
	subOpt     []nats.SubOpt
	streamInfo *nats.StreamInfo

	// extracted from param
	stream   string
	durable  string
	subjects []string
}

func NewBroker(opts ...Option) (broker.Broker, error) {
	servers := []string{}
	subOpt := []nats.SubOpt{}
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

	if options.param == nil {
		return nil, fmt.Errorf("param == nil")
	}

	stream, durable, subjects, err := extractParam(options.param)
	if err != nil {
		return nil, err
	}

	if durable != "" {
		subOpt = append(subOpt, nats.Durable(durable), nats.ManualAck())
	}

	return &_broker{
		user:     user,
		subOpt:   subOpt,
		stream:   stream,
		durable:  durable,
		subjects: subjects,
		servers:  servers,
		opts:     &options,
		password: password,
		Mutex:    sync.Mutex{},
		nopts:    &nats.Options{Servers: servers, User: user, Password: password},
	}, nil
}

func (p *_broker) Name() string {
	return broker2.NATS_JETSTREAM
}

func (p *_broker) Connect() error {
	conn, err := p.nopts.Connect()
	if err != nil {
		return err
	}

	js, err := conn.JetStream()
	if err != nil {
		conn.Close()
		return err
	}

	//err = js.DeleteStream("Message")
	//if err != nil {
	//	return err
	//}

	//ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	//defer cancel()
	//for info := range js.StreamsInfo(nats.Context(ctx)) {
	//	bytes, err := json.Marshal(info)
	//	if err != nil {
	//		panic(err)
	//	}
	//	fmt.Println("sinfo: ", string(bytes))
	//}

	if p.streamInfo, err = js.StreamInfo(p.stream); err != nil {
		p.streamInfo, err = js.AddStream(&nats.StreamConfig{
			Name:     p.stream,
			Subjects: p.subjects,
		})
		if err != nil {
			conn.Close()
			return err
		}
	}
	//else {
	//	p.streamInfo, err = js.UpdateStream(&nats.StreamConfig{
	//		Name:     p.stream,
	//		Subjects: p.subjects,
	//	})
	//	if err != nil {
	//		conn.Close()
	//		return err
	//	}
	//}

	p.js = js
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
	_, err := p.js.Publish(topic, msg)
	if err != nil {
		return err
	}
	return nil
}

func (p *_broker) Subscribe(topic string, handle broker.Handler) error {
	subj, err := p.js.Subscribe(topic, func(msg *nats.Msg) {
		handle(topic, msg.Data)
		msg.Ack()
	}, p.subOpt...)
	if err != nil {
		return err
	}

	p.subjectsMap.Store(topic, subj)
	return nil
}

func (p *_broker) UnSubscribe(topic string) error {
	temp, exist := p.subjectsMap.Load(topic)
	if exist {
		p.subjectsMap.Delete(topic)
		return temp.(*nats.Subscription).Unsubscribe()
	}

	return fmt.Errorf(fmt.Sprintf("the subscribtion for %s doesn't exist", topic))
}

func (p *_broker) Debug() {
	bDumpBrokerInfo := false
	bDumpDurable := true
	bDumpStreamInfo := true
	bDeleteStream := false
	bDeleteDurable := false

	if bDumpBrokerInfo {
		bytes, err := json.Marshal(p)
		if err != nil {
			panic(err)
		}
		fmt.Println("bi: ", string(bytes))
	}

	if bDumpDurable {
		di, err := p.js.ConsumerInfo(p.stream, p.durable)
		if err != nil {
			panic(err)
		}
		bytes, err := json.Marshal(di)
		if err != nil {
			panic(err)
		}
		fmt.Println("di: ", string(bytes))
	}

	if bDumpStreamInfo {
		streamInfo, err := p.js.StreamInfo(p.stream)
		if err != nil {
			panic(err)
		}
		bytes, err := json.Marshal(streamInfo)
		if err != nil {
			panic(err)
		}
		fmt.Println("sinfo: ", string(bytes))
	}

	if bDeleteStream {
		err := p.js.DeleteStream(p.stream)
		if err != nil {
			panic(err)
		}
	}

	if bDeleteDurable {
		err := p.js.DeleteConsumer(p.stream, p.durable)
		if err != nil {
			panic(err)
		}
	}
}

func extractParam(param map[string]interface{}) (string, string, []string, error) {
	var ok bool
	subjects := []string{}
	stream, durable := "", ""

	bytes, err := json.Marshal(&param)
	if err != nil {
		return "", "", nil, err
	}

	fmt.Println("extractParam.Param", string(bytes))

	if temp, exist := param["stream"]; exist {
		if stream, ok = temp.(string); !ok {
			return "", "", nil, fmt.Errorf("param[stream] not in type string")
		}
	}

	if temp, exist := param["durable"]; exist {
		if durable, ok = temp.(string); !ok {
			return "", "", nil, fmt.Errorf("param[durable] not in type string")
		}
	}

	if temp, exist := param["subjects"]; exist {
		if ti, ok := temp.([]interface{}); ok {
			for _, v := range ti {
				if vStr, ok := v.(string); ok {
					subjects = append(subjects, vStr)
				} else {
					return "", "", nil, fmt.Errorf("param[durable] not in type string")
				}
			}
		}
	}

	return stream, durable, subjects, nil
}
