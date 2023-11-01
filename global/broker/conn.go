package broker

import (
	"backstage/abstract/broker"
	broker2 "backstage/common/broker"
	"backstage/plugin/broker/nats/cluster"
	"backstage/plugin/broker/nats/jet_stream"
	server2 "backstage/plugin/broker/nats/server"
	"fmt"
)

func getBroker(which string) (broker.Broker, error) {
	_broker, err := load(which)
	if err != nil {
		return nil, err
	}
	return _broker, nil
}

func connectToBroker(category string, servers []string, user, password string, param map[string]interface{}) (broker.Broker, error) {
	var err error
	var _broker broker.Broker
	switch category {
	case broker2.NATS_SERVER:
		_broker, err = server2.NewBroker(
			server2.WithServer(servers[0]),
			server2.WithAuth(user, password),
		)
		if err != nil {
			return nil, err
		}
		return _broker, _broker.Connect()
	case broker2.NATS_CLUSTER:
		_broker, err = cluster.NewBroker(
			cluster.WithServers(servers),
			cluster.WithAuth(user, password),
		)
		if err != nil {
			return nil, err
		}
		return _broker, _broker.Connect()
	case broker2.NATS_JETSTREAM:
		_broker, err = jet_stream.NewBroker(
			jet_stream.WithServers(servers),
			jet_stream.WithAuth(user, password),
			jet_stream.WithParam(param),
		)
		if err != nil {
			return nil, err
		}
		return _broker, _broker.Connect()
	default:
		return nil, fmt.Errorf("unknow category: %s", category)
	}
}
