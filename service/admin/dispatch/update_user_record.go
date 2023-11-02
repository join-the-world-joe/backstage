package dispatch

import (
	"backstage/common/db/mgo/backend/track"
	"backstage/common/macro/permission"
	"backstage/common/major"
	"backstage/common/payload"
	"backstage/common/protocol/admin"
	"backstage/common/route"
	"backstage/global/config"
	"backstage/global/log"
	"backstage/global/mgo"
	"backstage/service/admin/business"
	"backstage/utils/convert"
	"context"
	"encoding/json"
	"time"
)

func updateUserRecord(packet *payload.PacketInternal) {
	req := &admin.UpdateUserRecordReq{}
	rsp := &admin.UpdateUserRecordRsp{}

	err := json.Unmarshal(packet.GetRequest().GetBody(), req)
	if err != nil {
		log.Error("Dispatch.updateUserRecord.json.Unmarshal failure, err: ", err.Error())
		return
	}

	req.Id = packet.GetSession().GetUserId()

	err = business.UpdateUserRecord(context.Background(), req, rsp)
	if err != nil {
		log.Error("Dispatch.updateUserRecord.business.InsertUserRecord failure, err: ", err.Error())
		return
	}

	bytes, err := json.Marshal(rsp)
	if err != nil {
		log.Error("Dispatch.updateUserRecord.json.Marshal failure, err: ", err.Error())
		return
	}

	packet.Response = &payload.PacketClient{
		Header: &payload.Header{
			Major: major.Admin,
			Minor: admin.UpdateUserRecordRsp_,
		},
		Body: bytes,
	}

	err = route.Downstream(
		config.DownstreamProtocol(),
		packet,
	)
	if err != nil {
		log.Error("Dispatch.updateUserRecord.route.Downstream failure, err: ", err.Error())
		return
	}

	if _, err = mgo.InsertDoc(
		config.MongoConf(),
		context.Background(),
		track.GetWhich(),
		track.GetDBName(),
		track.GetTableName(),
		&track.Model{
			Operator:   packet.GetSession().GetName(),
			Major:      major.Admin,
			Minor:      admin.UpdateUserRecordReq_,
			Request:    convert.Bytes2StringArray(packet.GetRequest().GetBody()),
			Permission: permission.UpdateUserRecord,
			Response:   convert.Bytes2StringArray(bytes),
			Timestamp:  time.Now().Unix(),
		},
	); err != nil {
		log.ErrorF("insertUserRecord failure, err :", err.Error())
	}
}
