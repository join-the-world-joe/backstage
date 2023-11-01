package broker

import (
	"backstage/abstract/broker"
	selector2 "backstage/abstract/selector"
	"backstage/common/conf"
	"backstage/global/log"
	"fmt"
)

func Broker(cf *conf.BrokerConf, strategy selector2.Strategy, topic string, handle broker.Handler) error {
	var err error
	if len(cf.Broker) > 0 {
		clear()                        // first, clear all old brokers
		choice := make(map[string]int) // for selector
		for which, v := range cf.Broker {
			brk, err := connectToBroker(v.Category, v.Servers, v.User, v.Password, v.Param)
			if err != nil {
				return err
			}
			if len(topic) > 0 && handle != nil {
				if err = brk.Subscribe(topic, handle); err != nil {
					return err
				}
				log.Info(fmt.Sprintf("Subscribe %s from Broker %s", topic, which))
			}
			choice[which] = 0
			store(which, brk)
		}

		sel.Set("Broker", choice)
		next, err = sel.Select("Broker", strategy)
		if err != nil {
			return err
		}
	}

	return nil
}

func Select() (string, error) {
	return next()
}

func Publish(which, topic string, msg []byte) error {
	_broker, err := getBroker(which)
	if err != nil {
		return err
	}

	return _broker.Publish(topic, msg)
}

func Subscribe(which, topic string, handle broker.Handler) error {
	_broker, err := getBroker(which)
	if err != nil {
		return err
	}
	return _broker.Subscribe(topic, handle)
}

func UnSubscribe(which, topic string) error {
	_broker, err := getBroker(which)
	if err != nil {
		return err
	}
	return _broker.UnSubscribe(topic)
}

func Disconnect(which string) error {
	_broker, err := getBroker(which)
	if err != nil {
		return err
	}
	return _broker.Disconnect()
}

func Debug(which string) {
	_broker, err := getBroker(which)
	if err != nil {
		fmt.Println("err: ", err)
		return
	}
	_broker.Debug()
}
