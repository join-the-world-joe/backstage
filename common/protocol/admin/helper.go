package admin

import (
	"backstage/common/macro/service"
	"backstage/global"
	"backstage/global/rpc"
	"context"
	"github.com/spf13/cast"
)

func SignIn(ctx context.Context, req *SignInReq, rsp *SignInRsp) error {
	srv, err := global.SelectService(service.Admin)
	if err != nil {
		return err
	}
	xClient, err := rpc.GetXClient(service.Admin, srv.Id, srv.Ip, cast.ToString(srv.Port))
	if err != nil {
		return err
	}
	return xClient.Call(ctx, "SignIn", req, rsp)
}

func InsertRecordOfProduct(ctx context.Context, req *InsertRecordOfProductReq, rsp *InsertRecordOfProductRsp) error {
	srv, err := global.SelectService(service.Admin)
	if err != nil {
		return err
	}
	xClient, err := rpc.GetXClient(service.Admin, srv.Id, srv.Ip, cast.ToString(srv.Port))
	if err != nil {
		return err
	}
	return xClient.Call(ctx, "InsertRecordOfProduct", req, rsp)
}

func SoftDeleteRecordOfProduct(ctx context.Context, req *SoftDeleteRecordsOfProductReq, rsp *SoftDeleteRecordsOfProductRsp) error {
	srv, err := global.SelectService(service.Admin)
	if err != nil {
		return err
	}
	xClient, err := rpc.GetXClient(service.Admin, srv.Id, srv.Ip, cast.ToString(srv.Port))
	if err != nil {
		return err
	}
	return xClient.Call(ctx, "SoftDeleteRecordOfProduct", req, rsp)
}

func UpdateRecordOfProduct(ctx context.Context, req *UpdateRecordOfProductReq, rsp *UpdateRecordOfProductRsp) error {
	srv, err := global.SelectService(service.Admin)
	if err != nil {
		return err
	}
	xClient, err := rpc.GetXClient(service.Admin, srv.Id, srv.Ip, cast.ToString(srv.Port))
	if err != nil {
		return err
	}
	return xClient.Call(ctx, "UpdateRecordOfProduct", req, rsp)
}

func InsertRecordOfAdvertisement(ctx context.Context, req *InsertRecordOfAdvertisementReq, rsp *InsertRecordOfAdvertisementRsp) error {
	srv, err := global.SelectService(service.Admin)
	if err != nil {
		return err
	}
	xClient, err := rpc.GetXClient(service.Admin, srv.Id, srv.Ip, cast.ToString(srv.Port))
	if err != nil {
		return err
	}
	return xClient.Call(ctx, "InsertRecordOfAdvertisement", req, rsp)
}

func UpdateRecordOfAdvertisement(ctx context.Context, req *UpdateRecordOfAdvertisementReq, rsp *UpdateRecordOfAdvertisementRsp) error {
	srv, err := global.SelectService(service.Admin)
	if err != nil {
		return err
	}
	xClient, err := rpc.GetXClient(service.Admin, srv.Id, srv.Ip, cast.ToString(srv.Port))
	if err != nil {
		return err
	}
	return xClient.Call(ctx, "UpdateRecordOfAdvertisement", req, rsp)
}

func SoftDeleteRecordsOfAdvertisement(ctx context.Context, req *SoftDeleteRecordsOfAdvertisementReq, rsp *SoftDeleteRecordsOfAdvertisementRsp) error {
	srv, err := global.SelectService(service.Admin)
	if err != nil {
		return err
	}
	xClient, err := rpc.GetXClient(service.Admin, srv.Id, srv.Ip, cast.ToString(srv.Port))
	if err != nil {
		return err
	}
	return xClient.Call(ctx, "SoftDeleteRecordsOfAdvertisement", req, rsp)
}

func InsertRecordOfADOfCarousel(ctx context.Context, req *InsertRecordOfADOfCarouselReq, rsp *InsertRecordOfADOfCarouselRsp) error {
	srv, err := global.SelectService(service.Admin)
	if err != nil {
		return err
	}
	xClient, err := rpc.GetXClient(service.Admin, srv.Id, srv.Ip, cast.ToString(srv.Port))
	if err != nil {
		return err
	}
	return xClient.Call(ctx, "InsertRecordOfADOfCarousel", req, rsp)
}

func RemoveOutdatedRecordsOfADOfCarousel(ctx context.Context, req *RemoveOutdatedRecordsOfADOfCarouselReq, rsp *RemoveOutdatedRecordsOfADOfCarouselRsp) error {
	srv, err := global.SelectService(service.Admin)
	if err != nil {
		return err
	}
	xClient, err := rpc.GetXClient(service.Admin, srv.Id, srv.Ip, cast.ToString(srv.Port))
	if err != nil {
		return err
	}
	return xClient.Call(ctx, "RemoveOutdatedRecordsOfADOfCarousel", req, rsp)
}

func InsertRecordOfADOfDeals(ctx context.Context, req *InsertRecordOfADOfDealsReq, rsp *InsertRecordOfADOfDealsRsp) error {
	srv, err := global.SelectService(service.Admin)
	if err != nil {
		return err
	}
	xClient, err := rpc.GetXClient(service.Admin, srv.Id, srv.Ip, cast.ToString(srv.Port))
	if err != nil {
		return err
	}
	return xClient.Call(ctx, "InsertRecordOfADOfDeals", req, rsp)
}

func RemoveOutdatedRecordsOfADOfDeals(ctx context.Context, req *RemoveOutdatedRecordsOfADOfDealsReq, rsp *RemoveOutdatedRecordsOfADOfDealsRsp) error {
	srv, err := global.SelectService(service.Admin)
	if err != nil {
		return err
	}
	xClient, err := rpc.GetXClient(service.Admin, srv.Id, srv.Ip, cast.ToString(srv.Port))
	if err != nil {
		return err
	}
	return xClient.Call(ctx, "RemoveOutdatedRecordsOfADOfDeals", req, rsp)
}

func InsertRecordOfADOfCamping(ctx context.Context, req *InsertRecordOfADOfCampingReq, rsp *InsertRecordOfADOfCampingRsp) error {
	srv, err := global.SelectService(service.Admin)
	if err != nil {
		return err
	}
	xClient, err := rpc.GetXClient(service.Admin, srv.Id, srv.Ip, cast.ToString(srv.Port))
	if err != nil {
		return err
	}
	return xClient.Call(ctx, "InsertRecordOfADOfCamping", req, rsp)
}

func RemoveOutdatedRecordsOfADOfCamping(ctx context.Context, req *RemoveOutdatedRecordsOfADOfCampingReq, rsp *RemoveOutdatedRecordsOfADOfCampingRsp) error {
	srv, err := global.SelectService(service.Admin)
	if err != nil {
		return err
	}
	xClient, err := rpc.GetXClient(service.Admin, srv.Id, srv.Ip, cast.ToString(srv.Port))
	if err != nil {
		return err
	}
	return xClient.Call(ctx, "RemoveOutdatedRecordsOfADOfCamping", req, rsp)
}

func InsertRecordOfADOfBarbecue(ctx context.Context, req *InsertRecordOfADOfBarbecueReq, rsp *InsertRecordOfADOfBarbecueRsp) error {
	srv, err := global.SelectService(service.Admin)
	if err != nil {
		return err
	}
	xClient, err := rpc.GetXClient(service.Admin, srv.Id, srv.Ip, cast.ToString(srv.Port))
	if err != nil {
		return err
	}
	return xClient.Call(ctx, "InsertRecordOfADOfBarbecue", req, rsp)
}

func RemoveOutdatedRecordsOfADOfBarbecue(ctx context.Context, req *RemoveOutdatedRecordsOfADOfBarbecueReq, rsp *RemoveOutdatedRecordsOfADOfBarbecueRsp) error {
	srv, err := global.SelectService(service.Admin)
	if err != nil {
		return err
	}
	xClient, err := rpc.GetXClient(service.Admin, srv.Id, srv.Ip, cast.ToString(srv.Port))
	if err != nil {
		return err
	}
	return xClient.Call(ctx, "RemoveOutdatedRecordsOfADOfBarbecue", req, rsp)
}

func InsertRecordOfADOfSnacks(ctx context.Context, req *InsertRecordOfADOfSnacksReq, rsp *InsertRecordOfADOfSnacksRsp) error {
	srv, err := global.SelectService(service.Admin)
	if err != nil {
		return err
	}
	xClient, err := rpc.GetXClient(service.Admin, srv.Id, srv.Ip, cast.ToString(srv.Port))
	if err != nil {
		return err
	}
	return xClient.Call(ctx, "InsertRecordOfADOfSnacks", req, rsp)
}

func RemoveOutdatedRecordsOfADOfSnacks(ctx context.Context, req *RemoveOutdatedRecordsOfADOfSnacksReq, rsp *RemoveOutdatedRecordsOfADOfSnacksRsp) error {
	srv, err := global.SelectService(service.Admin)
	if err != nil {
		return err
	}
	xClient, err := rpc.GetXClient(service.Admin, srv.Id, srv.Ip, cast.ToString(srv.Port))
	if err != nil {
		return err
	}
	return xClient.Call(ctx, "RemoveOutdatedRecordsOfADOfSnacks", req, rsp)
}
