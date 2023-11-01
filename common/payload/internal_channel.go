package payload

import (
	"backstage/abstract/fifo"
	fifo2 "backstage/lib/fifo"
	"fmt"
)

type PacketInternalChannel struct {
	c    chan *PacketInternal
	fifo fifo.FIFO
}

func NewPacketInternalChannel(bufferSize int) *PacketInternalChannel {
	return &PacketInternalChannel{
		c: make(chan *PacketInternal),
		fifo: fifo2.NewFIFO(
			fifo2.WithBufferSize(bufferSize),
		),
	}
}

func (p *PacketInternalChannel) Push(packet *PacketInternal) error {
	return p.fifo.Push(packet)
}

func (p *PacketInternalChannel) Pop() (*PacketInternal, error) {
	any, err := p.fifo.Pop()
	if err != nil {
		return nil, err
	}

	if packet, ok := any.(*PacketInternal); ok {
		return packet, nil
	}

	return nil, fmt.Errorf("packet from pop is not in type of a pointer of payload.payload")
}

func (p *PacketInternalChannel) Wait() <-chan *PacketInternal {
	go func() {
		any := <-p.fifo.Channel()
		if packet, ok := any.(*PacketInternal); ok {
			p.c <- packet
		}
	}()
	return p.c
}

func (p *PacketInternalChannel) Destroy() {
	p.fifo.Destroy()
}
