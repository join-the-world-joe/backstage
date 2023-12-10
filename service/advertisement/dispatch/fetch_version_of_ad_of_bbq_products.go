package dispatch

//
//import (
//	"backstage/common/major"
//	"backstage/common/payload"
//	"backstage/common/protocol/advertisement"
//	"backstage/common/route"
//	"backstage/global/config"
//	"backstage/global/log"
//	"backstage/service/advertisement/business"
//	"context"
//	"encoding/json"
//)
//
//func fetchVersionOfADOfBBQProducts(packet *payload.PacketInternal) {
//	req := &advertisement.FetchVersionOfADOfBBQProductsReq{}
//	rsp := &advertisement.FetchVersionOfADOfBBQProductsRsp{}
//
//	err := json.Unmarshal(packet.GetRequest().GetBody(), req)
//	if err != nil {
//		log.Error("Dispatch.fetchVersionOfADOfBBQProducts.json.Unmarshal failure, err: ", err.Error())
//		return
//	}
//
//	req.UserId = packet.GetSession().GetUserId()
//
//	err = business.FetchVersionOfADOfBBQProducts(context.Background(), req, rsp)
//	if err != nil {
//		log.Error("Dispatch.fetchVersionOfADOfBBQProducts.business.FetchVersionOfADOfBBQProducts failure, err: ", err.Error())
//		return
//	}
//
//	bytes, err := json.Marshal(rsp)
//	if err != nil {
//		log.Error("Dispatch.fetchVersionOfADOfBBQProducts.json.Marshal failure, err: ", err.Error())
//		return
//	}
//
//	packet.Response = &payload.PacketClient{
//		Header: &payload.Header{
//			Major: major.Advertisement,
//			Minor: advertisement.FetchVersionOfADOfBBQProductsRsp_,
//		},
//		Body: bytes,
//	}
//
//	err = route.Downstream(
//		config.DownstreamProtocol(),
//		packet,
//	)
//	if err != nil {
//		log.Error("Dispatch.fetchVersionOfADOfBBQProducts.route.Downstream failure, err: ", err.Error())
//		return
//	}
//}
