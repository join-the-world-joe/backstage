package backend_gateway

import (
	"backstage/abstract/config"
	"backstage/abstract/notifier"
	"backstage/abstract/registry"
	"backstage/abstract/selector"
	"backstage/abstract/service"
	broker3 "backstage/common/broker"
	service2 "backstage/common/macro/service"
	"backstage/common/protocol/gateway"
	"backstage/global"
	broker2 "backstage/global/broker"
	config2 "backstage/global/config"
	"backstage/global/log"
	"backstage/global/rate_limiting"
	notifier2 "backstage/lib/notifier"
	"backstage/plugin/config/nacos"
	"backstage/plugin/logger/zap"
	nacos2 "backstage/plugin/registry/nacos"
	"backstage/service/backend_gateway/conf"
	"backstage/service/backend_gateway/notify"
	"backstage/service/backend_gateway/rpc"
	"backstage/service/backend_gateway/runtime"
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/BurntSushi/toml"
	"github.com/gin-gonic/gin"
	"github.com/nacos-group/nacos-sdk-go/v2/common/constant"
	"github.com/nacos-group/nacos-sdk-go/v2/vo"
	"os"
	"os/signal"
	"strings"
	"time"
)

type _service struct {
	opts *Options

	conf conf.ServiceConf

	exitChan chan os.Signal

	ctx    context.Context
	cancel context.CancelFunc

	isStop     bool
	lastActive time.Time
}

func NewService(opts ...Option) (service.Service, error) {
	options := Options{
		config: make(map[string][]string),
	}

	for _, o := range opts {
		o(&options)
	}

	return &_service{
		opts: &options,
	}, nil
}

func (p *_service) Name() string {
	return service2.BackendGateway
}

func (p *_service) Init() error {
	var err error

	// for logger
	lg, err := zap.NewLogger(
		zap.WithLevel(-1),
		zap.WithFilePath(p.opts.logFilePath),
		zap.WithFileName(p.opts.logFileName),
		zap.WithCallerSkip(2),
	)
	if err != nil {
		return err
	}
	log.SetLogger(lg)

	// config
	handle, err := nacos.NewConfig(
		nacos.WithNacosClientParam(
			vo.NacosClientParam{
				ClientConfig: &constant.ClientConfig{
					LogDir:               p.opts.nacosLogDir,
					CacheDir:             p.opts.nacosCacheDir,
					LogLevel:             p.opts.nacosLogLevel,
					TimeoutMs:            config.DefaultTimeout,
					NamespaceId:          p.opts.nacosNamespaceId,
					NotLoadCacheAtStart:  true,
					UpdateCacheWhenEmpty: true,
				},
				ServerConfigs: []constant.ServerConfig{
					{
						Scheme:   p.opts.nacosScheme,
						IpAddr:   p.opts.nacosHost,
						Port:     p.opts.nacosHttpPort,
						GrpcPort: p.opts.nacosGrpcPort,
					},
				},
			},
		),
	)
	if err != nil {
		return err
	}

	buffer := &bytes.Buffer{}
	for k, v := range p.opts.config {
		//log.DebugF("Load Config %v-%v from nacos %v", k, v, fmt.Sprintf("%v:%v", p.opts.nacosHost, p.opts.nacosGrpcPort))
		for _, v2 := range v {
			bs, err := handle.Load(
				config.Parameter{
					Group:  k,
					DataId: v2,
				},
			)
			if err != nil {
				return err
			}
			_, err = buffer.Write(bs)
			if err != nil {
				return err
			}
			//log.DebugF("Content: \n %v", string(bs))
			buffer.WriteString("\n")
		}
	}

	log.Debug("bytes: \n", buffer.String())

	if err = toml.Unmarshal(buffer.Bytes(), &p.conf); err != nil {
		return err
	}

	if bs, err := json.Marshal(p.conf); err == nil {
		log.Info("Service Config: \n ", string(bs))
	} else {
		return err
	}

	if p.Name() != p.conf.Servant.Name {
		return errors.New(fmt.Sprintf("service.name(%s) is not equal to the Servant.Name(%s)", p.Name(), p.conf.Servant.Name))
	}

	// registry
	r, err := nacos2.NewRegistry(
		nacos2.WithNacosClientParam(
			vo.NacosClientParam{
				ClientConfig: &constant.ClientConfig{
					LogDir:               p.opts.nacosLogDir,
					CacheDir:             p.opts.nacosCacheDir,
					LogLevel:             p.opts.nacosLogLevel,
					TimeoutMs:            p.opts.registerTTL * 1000, // Healthy timeout
					NamespaceId:          p.opts.nacosNamespaceId,
					NotLoadCacheAtStart:  true,
					UpdateCacheWhenEmpty: true,
				},
				ServerConfigs: []constant.ServerConfig{
					{
						Scheme:   p.opts.nacosScheme,
						IpAddr:   p.opts.nacosHost,
						Port:     p.opts.nacosHttpPort,
						GrpcPort: p.opts.nacosGrpcPort,
					},
				},
			},
		),
	)
	if err != nil {
		return err
	}

	// context
	p.ctx, p.cancel = context.WithCancel(context.Background())

	// for notifier
	n, err := notifier2.NewNotifier(
		notifier2.WithBufferSize(p.opts.notifyBufferSize),
		notifier2.WithEmitTimeout(time.Microsecond*100),
	)
	if err != nil {
		return err
	}

	p.exitChan = make(chan os.Signal, 1)

	config2.SetMongoConf(&p.conf.MongoConf)
	config2.SetRateLimitingCallback(rate_limiting.UpdateRateLimit)
	config2.SetRateLimiting(&p.conf.RateLimitingConf)
	config2.SetRouteConf(&p.conf.RouteConf)
	config2.SetCacheConf(&p.conf.CacheConf)
	config2.SetBrokerConf(&p.conf.BrokerConf)
	config2.SetRPCServerConf(&p.conf.RPCServerConf)
	config2.SetRBACConf(&p.conf.RBACConf)
	config2.SetGracefulShutdownConf(&p.conf.GracefulShutdownConf)

	global.SetRegistry(r)
	global.SetNotifier(n)
	global.SetHost(p.opts.host)
	global.SetServiceId(p.opts.id)
	global.SetConfigHandler(handle)
	runtime.SetServiceConf(&p.conf)
	global.SetRouter(gin.Default())
	global.SetRPCPort(p.opts.rpcPort)
	global.SetVersion(p.opts.version)
	global.SetHTTPPort(p.opts.httpPort)
	log.SetLevel(p.conf.Servant.LogLevel)
	global.SetNacosGroup(p.opts.nacosGroup)
	global.SetServiceName(p.conf.Servant.Name)
	global.SetServicePath(p.conf.Servant.Name)
	return nil
}

func (p *_service) Start() error {
	log.Info(fmt.Sprintf("Start [%v.%v]", global.ServiceName(), global.ServiceId()))
	if config2.BrokerEnable() {
		err := broker2.Broker(
			config2.BrokerConf(),
			selector.RoundRobin,
			fmt.Sprintf(broker3.P2P, global.ServiceName(), global.ServiceId()),
			func(topic string, msg []byte) {
				req := new(gateway.P2PReq)
				err := json.Unmarshal(msg, req)
				if err != nil {
					log.Error(err.Error())
					return
				}
				channel, err := runtime.LoadChannel(req.Sequence)
				if err != nil {
					log.Error(err.Error())
					return
				}
				err = channel.Push(req.Packet)
				if err != nil {
					log.Error(err.Error())
					return
				}
			},
		)
		return err
	}

	if config2.RPCEnable(p.Name()) {
		if err := rpc.Setup(); err != nil {
			return err
		}
		log.Info(fmt.Sprintf("Setup RPC Server: Host: %v, Port: %v, ServicePath: %v", global.Host(), global.RPCPort(), global.ServicePath()))
	}

	global.Router().GET(runtime.WebsocketEndpoint(), handler)

	global.SetService(
		&registry.Service{
			Id:      global.ServiceId(),
			Ip:      global.Host(),
			Port:    global.RPCPort(),
			Name:    global.ServiceName(),
			Group:   global.NacosGroup(),
			Version: global.Version(),
		},
	)
	if err := global.Registry().Register(global.Service()); err != nil {
		return err
	}

	signal.Notify(p.exitChan, os.Interrupt)

	return nil
}

func (p *_service) Stop() {
	p.isStop = true
	global.Ticker().Stop()
	p.lastActive = time.Now()
	global.Registry().DeRegister(global.Service())
}

func (p *_service) Run() {

	runtime.Listener()(global.ConfigHandler())

	go global.Router().Run(fmt.Sprintf(":%v", global.HTTPPort()))

	log.Info(fmt.Sprintf("Run [%v.%v] successfully.", global.ServiceName(), global.ServiceId()))

	global.SetTicker(time.NewTicker(time.Duration(int(p.opts.registerInterval)) * time.Second))

	for {
		select {
		case <-global.Ticker().C:
			if !p.isStop {
				if err := global.Registry().Register(global.Service()); err == nil {
					log.Info(fmt.Sprintf("Register [%v.%v] successfully", global.ServiceName(), global.ServiceId()))
				} else {
					log.Error(fmt.Sprintf("Register [%v.%v] fail, %s", global.ServiceName(), global.ServiceId(), err.Error()))
				}
			}
		case <-time.After(time.Duration(config2.GracefulShutdownCheckInterval()) * time.Second):
			if p.isStop {
				if time.Now().Sub(p.lastActive) > time.Duration(config2.GracefulShutdownTimeout())*time.Second {
					log.Warn(fmt.Sprintf("Shutdown [%v.%v] right now!", global.ServiceName(), global.ServiceId()))
					p.cancel()
				}
			}
		case <-p.exitChan:
			p.Destroy()
		case <-p.ctx.Done():
			log.Warn(fmt.Sprintf("Shutdown [%v.%v]  successfully.", global.ServiceName(), global.ServiceId()))
			return
		case cmd := <-global.Notifier().Wait():
			arguments := strings.Split(cmd, " ")
			log.Warn("CMD: ", cmd, ", Arguments: ", arguments)
			if strings.Compare(cmd, notifier.Destroy) == 0 {
				p.Destroy()
			} else if strings.Compare(cmd, notifier.Stop) == 0 {
				p.Stop()
			} else {
				notify.Handler(arguments)
			}
		}
	}
}

func (p *_service) Destroy() {
	global.Registry().DeRegister(global.Service())
	p.cancel()
}
