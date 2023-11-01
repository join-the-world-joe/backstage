package gateway

import (
	"backstage/common/payload"
	"backstage/global"
	"backstage/global/log"
	"backstage/service/gateway/agent"
	"backstage/service/gateway/business"
	"backstage/service/gateway/conf"
	"backstage/service/gateway/runtime"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"github.com/spf13/cast"
	"net/http"
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
		log.Error("gateway.handler.websocket.Upgrader.Upgrade failure, err: ", err.Error())
		return
	}

	defer conn.Close()

	conn.SetReadLimit(runtime.WebsocketReadLimit())

	session, err := business.Logon(c, conn)
	if err != nil {
		log.Error("gateway.handler.business.Logon failure, err: ", err.Error())
		return
	}

	sequence := runtime.PopSequence()

	runtime.StoreSession(session.GetUserId(), session)
	defer runtime.RemoveSession(session.GetUserId())

	channel := runtime.StoreChannel(sequence, payload.NewPacketClientChannel(sequence, conf.DefaultP2PChannelSize))
	defer channel.Destroy()

	log.Debug(fmt.Sprintf("User: %v Verified, Sequence %v", session.GetUserId(), sequence))

	session.SetSequence(sequence)
	session.SetLoginServerName(global.ServiceName())
	session.SetLoginServerId(global.ServiceId())
	session.SetLoginServerHost(global.Host())
	session.SetLoginServerRPCPort(cast.ToString(global.RPCPort()))

	agt, err := agent.NewAgent(
		conn,
		session,
		channel,
	)
	if err != nil {
		log.Error("gateway.handler.agent.NewAgent failure, err: ", err.Error())
		return
	}
	defer runtime.RemoveChannel(sequence)

	agt.Serve()
}
