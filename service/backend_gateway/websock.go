package backend_gateway

import (
	"backstage/common/payload"
	"backstage/global"
	"backstage/global/log"
	"backstage/service/backend_gateway/agent"
	"backstage/service/backend_gateway/business"
	"backstage/service/backend_gateway/conf"
	"backstage/service/backend_gateway/runtime"
	"backstage/utils/proxy"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/gorilla/websocket"
	"github.com/spf13/cast"
	"net/http"
	"strings"
)

func handler(c *gin.Context) {

	log.Debug("Client: ", c.Request.RemoteAddr, " Connected")
	defer log.Debug("Client: ", c.Request.RemoteAddr, " Disconnected")

	conn, err := (&websocket.Upgrader{
		ReadBufferSize:    runtime.WebsocketReadBufferSize(),
		WriteBufferSize:   runtime.WebsocketWriteBufferSize(),
		EnableCompression: true,
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
		Subprotocols: c.Request.Header["Sec-Websocket-Protocol"],
	}).Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		log.Error("backend_gateway.handler.websocket.Upgrader.Upgrade failure, err: ", err.Error())
		return
	}

	defer conn.Close()

	conn.SetReadLimit(runtime.WebsocketReadLimit())

	session, err := business.Logon(c, conn)
	if err != nil {
		log.Error("backend_gateway.handler.business.Logon failure, err: ", err.Error())
		return
	}

	if previousSession, err := runtime.LoadSession(session.GetUserId()); err == nil && previousSession != nil {
		// force previous session to logout
		if strings.Compare(previousSession.GetClientIP(), proxy.ClientIP(c.Request)) != 0 {
			log.DebugF("previous session.IP %v, current session.IP %v", previousSession.GetClientIP(), proxy.ClientIP(c.Request))
			business.ForceOffline(session.GetUserId())
		}
	}

	sequence := runtime.PopSequence()

	runtime.StoreSession(session.GetUserId(), session)
	defer runtime.RemoveSession(session.GetUserId())

	channel := runtime.StoreChannel(sequence, payload.NewPacketClientChannel(sequence, conf.DefaultP2PChannelSize))
	defer channel.Destroy()

	session.SetSequence(sequence)
	session.SetClientIP(proxy.ClientIP(c.Request))
	session.SetId(uuid.New().String())
	session.SetLoginServerName(global.ServiceName())
	session.SetLoginServerId(global.ServiceId())
	session.SetLoginServerHost(global.Host())
	session.SetLoginServerRPCPort(cast.ToString(global.RPCPort()))

	log.Debug(fmt.Sprintf("User: %v Verified, Sequence %v, ClientIP %v", session.GetUserId(), sequence, session.GetClientIP()))

	agt, err := agent.NewAgent(
		conn,
		session,
		channel,
	)
	if err != nil {
		log.Error("backend_gateway.handler.agent.NewAgent failure, err: ", err.Error())
		return
	}
	defer runtime.RemoveChannel(sequence)

	agt.Serve()
}
