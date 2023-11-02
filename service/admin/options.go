package admin

type Options struct {
	// service
	id       string
	host     string
	httpPort uint64 // http port
	rpcPort  uint64 // rpc port
	version  string
	logLevel string

	// logger
	logFilePath string // for logger
	logFileName string // for logger

	// nacos
	nacosGroup       string
	config           map[string][]string // for nacos, key as group, value as data id
	nacosLogDir      string
	nacosCacheDir    string
	nacosLogLevel    string
	nacosNamespaceId string
	nacosScheme      string
	nacosHost        string
	nacosHttpPort    uint64
	nacosGrpcPort    uint64

	// registry
	registerTTL      uint64 // the duration for key expiration
	registerInterval uint64 // for ticker

	forwardChannelBufferSize int // for forward
	breakerBufferSize        int // for breaker
	notifyBufferSize         int // for cmd

}

type Option func(*Options)

func WithId(id string) Option {
	return func(o *Options) {
		o.id = id
	}
}

func WithHost(host string) Option {
	return func(o *Options) {
		o.host = host
	}
}

func WithHTTPPort(port uint64) Option {
	return func(o *Options) {
		o.httpPort = port
	}
}

func WithRPCPort(port uint64) Option {
	return func(o *Options) {
		o.rpcPort = port
	}
}

func WithVersion(version string) Option {
	return func(o *Options) {
		o.version = version
	}
}

func WithLogLevel(level string) Option {
	return func(o *Options) {
		o.logLevel = level
	}
}

func WithRegisterTTL(ttl uint64) Option {
	return func(o *Options) {
		o.registerTTL = ttl
	}
}

func WithRegisterInterval(interval uint64) Option {
	return func(o *Options) {
		o.registerInterval = interval
	}
}

func WithNacosGroup(group string) Option {
	return func(o *Options) {
		o.nacosGroup = group
	}
}

func WithNacosLogDir(dir string) Option {
	return func(o *Options) {
		o.nacosLogDir = dir
	}
}

func WithNacosCacheDir(dir string) Option {
	return func(o *Options) {
		o.nacosCacheDir = dir
	}
}

func WithNacosLogLevel(level string) Option {
	return func(o *Options) {
		o.nacosLogLevel = level
	}
}

func WithNacosNamespaceId(id string) Option {
	return func(o *Options) {
		o.nacosNamespaceId = id
	}
}

func WithNacosScheme(scheme string) Option {
	return func(o *Options) {
		o.nacosScheme = scheme
	}
}

func WithNacosHost(host string) Option {
	return func(o *Options) {
		o.nacosHost = host
	}
}

func WithNacosHttpPort(port uint64) Option {
	return func(o *Options) {
		o.nacosHttpPort = port
	}
}

func WithNacosGrpcPort(port uint64) Option {
	return func(o *Options) {
		o.nacosGrpcPort = port
	}
}

func WithConfig(group, dataId string) Option {
	return func(o *Options) {
		o.config[group] = append(o.config[group], dataId)
	}
}

func WithNotifyBufferSize(bufferSize int) Option {
	return func(o *Options) {
		o.notifyBufferSize = bufferSize
	}
}

func WithLogFilePath(filePath string) Option {
	return func(o *Options) {
		o.logFilePath = filePath
	}
}

func WithLogFileName(file string) Option {
	return func(o *Options) {
		o.logFileName = file
	}
}

func WithForwardChannelBufferSize(size int) Option {
	return func(o *Options) {
		o.forwardChannelBufferSize = size
	}
}
