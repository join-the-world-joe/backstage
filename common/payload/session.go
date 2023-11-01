package payload

import "reflect"

type Session struct {
	// important
	Role               string `json:"role"`
	Token              string `json:"token"`
	Sequence           uint64 `json:"sequence"`
	LoginServerName    string `json:"login_server_name"`
	LoginServerId      string `json:"login_server_id"`
	LoginServerHost    string `json:"login_server_host"`
	LoginServerRPCPort string `json:"login_server_port"`
	// info
	UserId     int64  `json:"user_id"`
	Lang       string `json:"lang"`
	OS         string `json:"os"`
	Mac        string `json:"mac"`
	RemoteAddr string `json:"remote_addr"`
	Version    string `json:"version"`
}

func (p *Session) Reset() *Session {
	v := reflect.ValueOf(p).Elem()
	v.Set(reflect.Zero(v.Type()))
	return p
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

func (p *Session) SetToken(token string) {
	p.Token = token
}

func (p *Session) GetToken() string {
	return p.Token
}

func (p *Session) SetRole(role string) {
	p.Role = role
}

func (p *Session) GetRole() string {
	return p.Role
}
