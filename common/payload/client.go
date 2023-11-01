package payload

import "encoding/json"

type PacketClient struct {
	Header *Header         `json:"header"`
	Body   json.RawMessage `json:"body"`
	//Body []byte `json:"body"`
}

func NewPacketClient(major, minor string, body []byte) *PacketClient {
	return &PacketClient{
		Header: &Header{
			Major: major,
			Minor: minor,
		},
		Body: body,
	}
}

func (p *PacketClient) ToBytes() ([]byte, error) {
	return json.Marshal(p)
}

func ParsePacketClient(bytes []byte) (*PacketClient, error) {
	packet := new(PacketClient)
	return packet, json.Unmarshal(bytes, packet)
}

func (p *PacketClient) GetHeader() *Header {
	return p.Header
}

func (p *PacketClient) GetBody() []byte {
	return p.Body
}
