package business

import (
	"backstage/common/major"
	"backstage/common/payload"
	"backstage/common/protocol/inform"
	"backstage/global/log"
	"backstage/service/backend_gateway/runtime"
	"encoding/json"
)

func ForceOffline(userId int64) {
	session, err := runtime.LoadSession(userId)
	if err != nil {
		log.ErrorF("ForceOffline.runtime.LoadSession failure, err: ", err.Error())
		return
	}

	bytes, err := json.Marshal(&inform.Notification{
		Event:   inform.EventForceOffline,
		Message: inform.MessageForceOffline,
	})
	if err != nil {
		log.ErrorF("ForceOffline failure, err: ", err.Error())
		return
	}

	packet := &payload.PacketClient{
		Header: &payload.Header{
			Major: major.Inform,
			Minor: inform.Notification_,
		},
		Body: bytes,
	}

	channel, err := runtime.LoadChannel(session.GetSequence())
	if err != nil {
		return
	}
	if err := channel.Push(packet); err != nil {
		return
	}

	log.DebugF("ForceOffline.session.IP: %v", session.GetClientIP())

	session.SetForceOffline(true)
	session.SetPacketClient(packet)
}
