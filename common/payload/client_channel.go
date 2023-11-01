package payload

import (
	"backstage/abstract/fifo"
	fifo2 "backstage/lib/fifo"
	"fmt"
	"sync/atomic"
	"time"
)

type PacketClientChannel struct {
	sequence   uint64
	sendPacket uint64
	establish  time.Time
	lastSend   time.Time
	fifo       fifo.FIFO
	c          chan *PacketClient
}

func NewPacketClientChannel(seq uint64, bufferSize int) *PacketClientChannel {
	return &PacketClientChannel{
		c:          make(chan *PacketClient),
		sequence:   seq,
		sendPacket: 0,
		establish:  time.Now(),
		fifo: fifo2.NewFIFO(
			fifo2.WithBufferSize(bufferSize),
		),
	}
}

func (p *PacketClientChannel) GetSequence() uint64 {
	return p.sequence
}

func (p *PacketClientChannel) GetSendPacket() uint64 {
	return p.sendPacket
}

func (p *PacketClientChannel) Push(packet *PacketClient) error {
	atomic.AddUint64(&p.sendPacket, 1)
	return p.fifo.Push(packet)
}

func (p *PacketClientChannel) Pop() (*PacketClient, error) {
	temp, err := p.fifo.Pop()
	if err != nil {
		return nil, err
	}

	if packet, ok := temp.(*PacketClient); ok {
		return packet, nil
	}

	return nil, fmt.Errorf("packet from pop is not in type of a pointer of payload.payload")
}

func (p *PacketClientChannel) Wait() <-chan *PacketClient {
	go func() {
		temp := <-p.fifo.Channel()
		if packet, ok := temp.(*PacketClient); ok {
			p.c <- packet
		}
	}()
	return p.c
}

func (p *PacketClientChannel) Destroy() {
	//log.Debug("PacketClientChannel.Destroy: ", p.sequence)
	p.fifo.Destroy()
}
