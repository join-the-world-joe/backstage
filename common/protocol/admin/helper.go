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

func InsertRecordOfGood(ctx context.Context, req *InsertRecordOfGoodReq, rsp *InsertRecordOfGoodRsp) error {
	srv, err := global.SelectService(service.Admin)
	if err != nil {
		return err
	}
	xClient, err := rpc.GetXClient(service.Admin, srv.Id, srv.Ip, cast.ToString(srv.Port))
	if err != nil {
		return err
	}
	return xClient.Call(ctx, "InsertRecordOfGood", req, rsp)
}

func FetchIdListOfGood(ctx context.Context, req *FetchIdListOfGoodReq, rsp *FetchIdListOfGoodRsp) error {
	srv, err := global.SelectService(service.Admin)
	if err != nil {
		return err
	}
	xClient, err := rpc.GetXClient(service.Admin, srv.Id, srv.Ip, cast.ToString(srv.Port))
	if err != nil {
		return err
	}
	return xClient.Call(ctx, "FetchIdListOfGood", req, rsp)
}

func FetchRecordsOfGood(ctx context.Context, req *FetchRecordsOfGoodReq, rsp *FetchRecordsOfGoodRsp) error {
	srv, err := global.SelectService(service.Admin)
	if err != nil {
		return err
	}
	xClient, err := rpc.GetXClient(service.Admin, srv.Id, srv.Ip, cast.ToString(srv.Port))
	if err != nil {
		return err
	}
	return xClient.Call(ctx, "FetchRecordsOfGood", req, rsp)
}

func SoftDeleteRecordOfGood(ctx context.Context, req *SoftDeleteRecordsOfGoodReq, rsp *SoftDeleteRecordsOfGoodRsp) error {
	srv, err := global.SelectService(service.Admin)
	if err != nil {
		return err
	}
	xClient, err := rpc.GetXClient(service.Admin, srv.Id, srv.Ip, cast.ToString(srv.Port))
	if err != nil {
		return err
	}
	return xClient.Call(ctx, "SoftDeleteRecordOfGood", req, rsp)
}

func UpdateRecordOfGood(ctx context.Context, req *UpdateRecordOfGoodReq, rsp *UpdateRecordOfGoodRsp) error {
	srv, err := global.SelectService(service.Admin)
	if err != nil {
		return err
	}
	xClient, err := rpc.GetXClient(service.Admin, srv.Id, srv.Ip, cast.ToString(srv.Port))
	if err != nil {
		return err
	}
	return xClient.Call(ctx, "UpdateRecordOfGood", req, rsp)
}

func FetchIdListOfAdvertisement(ctx context.Context, req *FetchIdListOfAdvertisementReq, rsp *FetchIdListOfAdvertisementRsp) error {
	srv, err := global.SelectService(service.Admin)
	if err != nil {
		return err
	}
	xClient, err := rpc.GetXClient(service.Admin, srv.Id, srv.Ip, cast.ToString(srv.Port))
	if err != nil {
		return err
	}
	return xClient.Call(ctx, "FetchIdListOfAdvertisement", req, rsp)
}

func FetchRecordsOfAdvertisement(ctx context.Context, req *FetchRecordsOfAdvertisementReq, rsp *FetchRecordsOfAdvertisementRsp) error {
	srv, err := global.SelectService(service.Admin)
	if err != nil {
		return err
	}
	xClient, err := rpc.GetXClient(service.Admin, srv.Id, srv.Ip, cast.ToString(srv.Port))
	if err != nil {
		return err
	}
	return xClient.Call(ctx, "FetchRecordsOfAdvertisement", req, rsp)
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

func InsertRecordOfADOfHots(ctx context.Context, req *InsertRecordOfADOfHotsReq, rsp *InsertRecordOfADOfHotsRsp) error {
	srv, err := global.SelectService(service.Admin)
	if err != nil {
		return err
	}
	xClient, err := rpc.GetXClient(service.Admin, srv.Id, srv.Ip, cast.ToString(srv.Port))
	if err != nil {
		return err
	}
	return xClient.Call(ctx, "InsertRecordOfADOfHots", req, rsp)
}

func RemoveOutdatedRecordsOfADOfHots(ctx context.Context, req *RemoveOutdatedRecordsOfADOfHotsReq, rsp *RemoveOutdatedRecordsOfADOfHotsRsp) error {
	srv, err := global.SelectService(service.Admin)
	if err != nil {
		return err
	}
	xClient, err := rpc.GetXClient(service.Admin, srv.Id, srv.Ip, cast.ToString(srv.Port))
	if err != nil {
		return err
	}
	return xClient.Call(ctx, "RemoveOutdatedRecordsOfADOfHots", req, rsp)
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
