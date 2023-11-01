package broker

import (
	"backstage/abstract/registry"
	"backstage/common/broker"
	"backstage/common/payload"
	"backstage/global"
	"backstage/global/log"
	"backstage/global/routing"
	"encoding/json"
	"errors"
	"fmt"
)

func Forward(packet *payload.PacketInternal) error {
	var err error
	var srv *registry.Service
	if srv, err = routing.Load(payload.GetUpstreamServiceName(packet)); err == nil {
		if srv == nil {
			srv, err = global.SelectService(payload.GetUpstreamServiceName(packet))
			if err != nil {
				return err
			}
		}
	} else {
		srv, err = global.SelectService(payload.GetUpstreamServiceName(packet))
		if err != nil {
			return err
		}
		if srv == nil {
			return errors.New("Forward.global.SelectService, srv == nil")
		}
	}

	which, err := Select()
	if err != nil {
		return err
	}

	bytes, err := json.Marshal(packet)
	if err != nil {
		return err
	}

	log.Debug("Forward To: \n Broker: ", which, "\n", "Service: ", *srv)

	err = Publish(
		which,
		fmt.Sprintf(broker.Forward, payload.GetUpstreamServiceName(packet), srv.Id),
		bytes,
	)
	if err != nil {
		return err
	}
	return nil
}
