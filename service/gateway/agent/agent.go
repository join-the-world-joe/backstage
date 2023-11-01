package agent

import (
	"backstage/abstract/notifier"
	"backstage/common/payload"
	"backstage/common/route"
	"backstage/global/config"
	"backstage/global/crypto"
	"backstage/global/log"
	notifier2 "backstage/lib/notifier"
	"backstage/service/gateway/runtime"
	"backstage/utils/convert"
	"fmt"
	"github.com/gorilla/websocket"
	"time"
)

type Agent struct {
	conn          *websocket.Conn
	channel       *payload.PacketClientChannel
	session       *payload.Session
	closeNotifier notifier.Notifier
}

func NewAgent(conn *websocket.Conn, session *payload.Session, c *payload.PacketClientChannel) (*Agent, error) {
	notify, err := notifier2.NewNotifier(
		notifier2.WithBufferSize(2),
		notifier2.WithEmitTimeout(time.Microsecond*100),
	)
	if err != nil {
		return nil, err
	}
	return &Agent{
		conn:          conn,
		channel:       c,
		session:       session,
		closeNotifier: notify,
	}, nil
}

func (p *Agent) Serve() error {
	go receiver(p)

	for {
		p.conn.SetReadDeadline(time.Now().Add(time.Second * time.Duration(runtime.WebsocketReadDeadline())))

		// send whatever received from websocket to broker
		_, message, err := p.conn.ReadMessage()
		if err != nil {
			log.Error(fmt.Sprintf("Agent[%v] Sender: error: %v", p.session.GetUserId(), err.Error()))
			p.closeNotifier.Emit("close")
			return err
		}

		log.Debug("Message: ", convert.Bytes2StringArray(message))

		plainText := message
		if runtime.EncryptionEnable() {
			plainText, err = crypto.AESDecrypt(message)
			if err != nil {
				log.Error(err.Error())
				continue
			}
		}

		log.Debug("PlainText(Bytes): ", convert.Bytes2StringArray(plainText))
		log.Debug("PlainText(String): ", string(plainText))

		packet, err := payload.ParsePacketClient(plainText)
		if err == nil {
			// todo: check if packet if valid
			log.Debug("*******************************************************************************")
			log.Debug("Major: ", packet.Header.Major, ", Minor: ", packet.Header.Minor)
			log.Debug("UserId: ", p.session.GetUserId())
			log.Debug("Sequence: ", p.session.GetSequence())
			log.Debug("Body: ", string(packet.Body))
			log.Debug("*******************************************************************************")
			if err = route.Upstream(config.DownstreamProtocol(), &payload.PacketInternal{Session: p.session, Request: packet}); err != nil {
				log.Error("agent.Upstream fail: ", err.Error())
				continue
			}
			log.Debug(fmt.Sprintf("Agent[%v] Upstream: [%s]", p.session.GetUserId(), string(plainText)))
		} else {
			log.Error(fmt.Sprintf("Agent[%v] Parse Packet[%v] Error %v", p.session.GetUserId(), convert.Bytes2StringArray(plainText), err.Error()))
			continue
		}
	}
}

func receiver(p *Agent) {
	defer log.WarnF("Agent[%v].Receiver Closed", p.session.GetUserId())
loop:
	for {
		select {
		case <-p.closeNotifier.Wait():
			return
		case packet := <-p.channel.Wait():
			bytes, err := packet.ToBytes()
			if err == nil {
				log.Debug(fmt.Sprintf("Agent[%v] Downstream: [%s]", p.session.GetUserId(), string(bytes)))
				//fmt.Println("Encryption: ", runtime.GetEncryption())
				if runtime.EncryptionEnable() {
					bytes, err = crypto.AESEncrypt(bytes)
					if err != nil {
						log.Error(err.Error())
						continue
					}
				}
				log.Debug(fmt.Sprintf("Agent[%v] Downstream.Content: [%s]", p.session.GetUserId(), convert.Bytes2StringArray(bytes)))
				err = p.conn.WriteMessage(websocket.BinaryMessage, bytes)
				if err != nil {
					break loop
				}
			} else {
				// fail to marshal packet to bytes
				log.Error(err.Error())
			}
		}
	}
}
