package payload

import "reflect"

type Session struct {
	// important
	Id                 string `json:"id"` // session id, unique
	Name               string `json:"name"`
	MemberId           string `json:"member_id"`
	Sequence           uint64 `json:"sequence"`
	LoginServerName    string `json:"login_server_name"`
	LoginServerId      string `json:"login_server_id"`
	LoginServerHost    string `json:"login_server_host"`
	LoginServerRPCPort string `json:"login_server_port"`
	ClientIP           string `json:"-"`

	// info
	UserId     int64  `json:"user_id"`
	Lang       string `json:"lang"`
	OS         string `json:"os"`
	Mac        string `json:"mac"`
	RemoteAddr string `json:"remote_addr"`
	Version    string `json:"version"`

	// inform
	ForceOffline bool          `json:"force_offline"`
	Packet       *PacketClient `json:"-"`
}

func (p *Session) Reset() *Session {
	v := reflect.ValueOf(p).Elem()
	v.Set(reflect.Zero(v.Type()))
	return p
}

func (p *Session) SetId(id string) {
	p.Id = id
}

func (p *Session) GetId() string {
	return p.Id
}

func (p *Session) SetClientIP(ip string) {
	p.ClientIP = ip
}

func (p *Session) GetClientIP() string {
	return p.ClientIP
}

func (p *Session) SetPacketClient(Packet *PacketClient) {
	p.Packet = Packet
}

func (p *Session) GetPacketClient() *PacketClient {
	return p.Packet
}

func (p *Session) SetForceOffline(b bool) {
	p.ForceOffline = b
}

func (p *Session) GetForceOffline() bool {
	return p.ForceOffline
}

func (p *Session) SetLoginServerName(srvName string) {
	p.LoginServerName = srvName
}

func (p *Session) GetLoginServerName() string {
	return p.LoginServerName
}

func (p *Session) SetLoginServerId(srvId string) {
	p.LoginServerId = srvId
}

func (p *Session) GetLoginServerId() string {
	return p.LoginServerId
}

func (p *Session) SetLoginServerHost(host string) {
	p.LoginServerHost = host
}

func (p *Session) GetLoginServerHost() string {
	return p.LoginServerHost
}

func (p *Session) SetLoginServerRPCPort(port string) {
	p.LoginServerRPCPort = port
}

func (p *Session) GetLoginServerRPCPort() string {
	return p.LoginServerRPCPort
}

func (p *Session) SetSequence(sequence uint64) {
	p.Sequence = sequence
}

func (p *Session) GetSequence() uint64 {
	return p.Sequence
}

func (p *Session) SetUserId(userId int64) {
	p.UserId = userId
}

func (p *Session) GetUserId() int64 {
	return p.UserId
}

func (p *Session) SetName(name string) {
	p.Name = name
}

func (p *Session) GetName() string {
	return p.Name
}

func (p *Session) SetMemberId(memberId string) {
	p.MemberId = memberId
}

func (p *Session) GetMemberId() string {
	return p.MemberId
}
