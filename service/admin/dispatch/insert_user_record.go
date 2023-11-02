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

func insertUserRecord(packet *payload.PacketInternal) {
	req := &admin.InsertUserRecordReq{}
	rsp := &admin.InsertUserRecordRsp{}

	err := json.Unmarshal(packet.GetRequest().GetBody(), req)
	if err != nil {
		log.Error("Dispatch.insertUserRecord.json.Unmarshal failure, err: ", err.Error())
		return
	}

	req.Id = packet.GetSession().GetUserId()

	err = business.InsertUserRecord(context.Background(), req, rsp)
	if err != nil {
		log.Error("Dispatch.insertUserRecord.business.InsertUserRecord failure, err: ", err.Error())
		return
	}

	bytes, err := json.Marshal(rsp)
	if err != nil {
		log.Error("Dispatch.insertUserRecord.json.Marshal failure, err: ", err.Error())
		return
	}

	packet.Response = &payload.PacketClient{
		Header: &payload.Header{
			Major: major.Admin,
			Minor: admin.InsertUserRecordRsp_,
		},
		Body: bytes,
	}

	err = route.Downstream(
		config.DownstreamProtocol(),
		packet,
	)
	if err != nil {
		log.Error("Dispatch.insertUserRecord.route.Downstream failure, err: ", err.Error())
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
			Minor:      admin.InsertUserRecordReq_,
			Request:    convert.Bytes2StringArray(packet.GetRequest().GetBody()),
			Permission: permission.InsertUserRecord,
			Response:   convert.Bytes2StringArray(bytes),
			Timestamp:  time.Now().Unix(),
		},
	); err != nil {
		log.ErrorF("insertUserRecord failure, err :", err.Error())
	}
}
