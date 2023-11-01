package payload

import "encoding/json"

type PacketInternal struct {
	Request  *PacketClient
	Response *PacketClient
	Session  *Session
}

func (p *PacketInternal) ToBytes() ([]byte, error) {
	return json.Marshal(p)
}

func ParsePacketInternal(bytes []byte) (*PacketInternal, error) {
	packet := new(PacketInternal)
	return packet, json.Unmarshal(bytes, packet)
}

func (p *PacketInternal) GetRequest() *PacketClient {
	return p.Request
}

func (p *PacketInternal) GetResponse() *PacketClient {
	return p.Response
}

func (p PacketInternal) SetResponse(response *PacketClient) {
	p.Response = response
}

func (p *PacketInternal) GetSession() *Session {
	return p.Session
}
