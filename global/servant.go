package global

var _id string
var _host string
var _port uint64
var _version string
var _rpc_port uint64
var _service_name string
var _service_path string

func SetHost(host string) {
	_host = host
}

func Host() string {
	return _host
}

func SetHTTPPort(port uint64) {
	_port = port
}

func HTTPPort() uint64 {
	return _port
}

func SetVersion(verion string) {
	_version = verion
}

func Version() string {
	return _version
}

func SetServiceId(id string) {
	_id = id
}

func ServiceId() string {
	return _id
}

func SetServiceName(name string) {
	_service_name = name
}

func ServiceName() string {
	return _service_name
}

func SetServicePath(servicePath string) {
	_service_path = servicePath
}

func ServicePath() string {
	return _service_path
}

func SetRPCPort(port uint64) {
	_rpc_port = port
}

func RPCPort() uint64 {
	return _rpc_port
}
