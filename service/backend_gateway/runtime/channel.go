package runtime

import (
	"backstage/common/payload"
	"errors"
	"fmt"
	"sync"
)

var _channelMap sync.Map // it holds all channels

func StoreChannel(sequence uint64, p *payload.PacketClientChannel) *payload.PacketClientChannel {
	_channelMap.Store(sequence, p)
	return p
}

func LoadChannel(sequence uint64) (*payload.PacketClientChannel, error) {
	value, ok := _channelMap.Load(sequence)
	if ok {
		return value.(*payload.PacketClientChannel), nil
	}
	return nil, errors.New(fmt.Sprintf("%d doesn't exist", sequence))
}

func RemoveChannel(sequence uint64) {
	_channelMap.Delete(sequence)
}
