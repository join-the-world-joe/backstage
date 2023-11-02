package generic

import (
	"backstage/abstract/breaker"
	"backstage/abstract/config"
	"backstage/abstract/notifier"
	"backstage/abstract/registry"
	"backstage/abstract/selector"
	"backstage/abstract/service"
	"backstage/common/broker"
	service2 "backstage/common/macro/service"
	"backstage/common/payload"
	"backstage/global"
	broker2 "backstage/global/broker"
	config2 "backstage/global/config"
	"backstage/global/log"
	notifier2 "backstage/lib/notifier"
	"backstage/plugin/config/nacos"
	"backstage/plugin/logger/zap"
	nacos2 "backstage/plugin/registry/nacos"
	"backstage/service/generic/api"
	"backstage/service/generic/conf"
	"backstage/service/generic/dispatch"
	"backstage/service/generic/notify"
	"backstage/service/generic/rpc"
	"backstage/service/generic/runtime"
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/BurntSushi/toml"
	"github.com/gin-gonic/gin"
	"github.com/nacos-group/nacos-sdk-go/v2/common/constant"
	"github.com/nacos-group/nacos-sdk-go/v2/vo"
	"github.com/spf13/cast"
	"os"
	"os/signal"
	"strings"
	"time"
)

type _service struct {
	opts *Options

	exitChan chan os.Signal

	conf conf.ServiceConf

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
	return service2.Generic
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
		//fmt.Println("k: ", k, ", v: ", v)
		for _, v2 := range v {
			bs, err := handle.Load(config.Parameter{
				Group:  k,
				DataId: v2,
			})
			if err != nil {
				return err
			}
			_, err = buffer.Write(bs)
			if err != nil {
				return err
			}

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
					LogDir:              p.opts.nacosLogDir,
					CacheDir:            p.opts.nacosCacheDir,
					LogLevel:            p.opts.nacosLogLevel,
					TimeoutMs:           p.opts.registerTTL * 1000, // Healthy timeout
					NamespaceId:         p.opts.nacosNamespaceId,
					NotLoadCacheAtStart: true,
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

	if global.HTTPPort() > 0 {
		global.SetRouter(gin.Default())
	}

	config2.SetRouteConf(&p.conf.RouteConf)
	config2.SetCacheConf(&p.conf.CacheConf)
	config2.SetMySQLConf(&p.conf.MySQLConf)
	config2.SetRBACConf(&p.conf.RBACConf)
	config2.SetBrokerConf(&p.conf.BrokerConf)
	config2.SetMongoConf(&p.conf.MongoConf)
	config2.SetRPCServerConf(&p.conf.RPCServerConf)
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
	global.SetBreaker(make(chan breaker.Breaker, p.opts.breakerBufferSize))
	global.SetForward(payload.NewPacketInternalChannel(p.opts.forwardChannelBufferSize))
	return nil
}

func (p *_service) Start() error {
	log.Info(fmt.Sprintf("Start [%v.%v]", global.ServiceName(), global.ServiceId()))
	if config2.BrokerEnable() {
		err := broker2.Broker(
			config2.BrokerConf(),
			selector.RoundRobin,
			fmt.Sprintf(broker.Forward, global.ServiceName(), global.ServiceId()),
			func(topic string, msg []byte) {
				packet := new(payload.PacketInternal)
				err := json.Unmarshal(msg, packet)
				if err != nil {
					log.Error(err.Error())
					return
				}
				err = global.Forward().Push(packet)
				if err != nil {
					log.Error(err.Error())
					return
				}
			},
		)
		if err != nil {
			return err
		}
	}

	if config2.RPCEnable(p.Name()) {
		if err := rpc.Setup(); err != nil {
			return err
		}
		log.Info(fmt.Sprintf("Setup RPC Server: Host: %v, Port: %v, ServicePath: %v", global.Host(), global.RPCPort(), global.ServicePath()))
	}

	if err := api.Setup(); err != nil {
		return err
	}
	log.Info("Setup API Router successfully.")

	global.SetService(
		&registry.Service{
			Id:      global.ServiceId(),
			Ip:      global.Host(),
			Port:    cast.ToUint64(global.RPCPort()),
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

	if global.HTTPPort() > 0 {
		go global.Router().Run(fmt.Sprintf(":%v", global.HTTPPort()))
	}

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
		case brk := <-global.Breaker():
			log.Warn(fmt.Sprintf("Break [%v.%v] successfully.", global.ServiceName(), global.ServiceId()))
			<-brk.Break()
			log.Warn(fmt.Sprintf("Resume [%v.%v] successfully.", global.ServiceName(), global.ServiceId()))
		case packet := <-global.Forward().Wait():
			dispatch.Dispatch(packet)
		}
	}
}

func (p *_service) Destroy() {
	global.Registry().DeRegister(global.Service())
	p.cancel()
}
