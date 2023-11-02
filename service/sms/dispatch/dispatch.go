package dispatch

import (
	"backstage/common/payload"
	"backstage/common/protocol/sms"
	"backstage/global/log"
	"backstage/service/sms/business"
	"context"
	"encoding/json"
)

func Dispatch(packet *payload.PacketInternal) {
	minor := packet.GetRequest().GetHeader().GetMinor()
	log.Info("Dispatch.Minor: ", minor)
	switch minor {
	case sms.SendVerificationCodeReq_:
		req := &sms.SendVerificationCodeReq{}
		err := json.Unmarshal(packet.Request.GetBody(), req)
		if err != nil {
			log.Error("Dispatch.json.Unmarshal failure, err: ", err.Error())
			return
		}
		rsp := &sms.SendVerificationCodeRsp{}
		err = business.SendVerificationCode(context.Background(), req, rsp)
		if err != nil {
			log.Error("Dispatch.SendVerificationCode failure, err: ", err.Error())
			return
		}
	default:
		log.ErrorF("unknown minor [%v]", packet.GetRequest().GetHeader().GetMinor())
	}
}
