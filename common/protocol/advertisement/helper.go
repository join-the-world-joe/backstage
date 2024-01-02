package advertisement

import (
	"backstage/common/macro/service"
	"backstage/global"
	"backstage/global/rpc"
	"context"
	"github.com/spf13/cast"
)

// Carousel
func FetchVersionOfADOfCarousel(ctx context.Context, req *FetchVersionOfADOfCarouselReq, rsp *FetchVersionOfADOfCarouselRsp) error {
	srv, err := global.SelectService(service.Advertisement)
	if err != nil {
		return err
	}
	xClient, err := rpc.GetXClient(service.Advertisement, srv.Id, srv.Ip, cast.ToString(srv.Port))
	if err != nil {
		return err
	}
	return xClient.Call(ctx, "FetchVersionOfADOfCarousel", req, rsp)
}

func FetchIdListOfADOfCarousel(ctx context.Context, req *FetchIdListOfADOfCarouselReq, rsp *FetchIdListOfADOfCarouselRsp) error {
	srv, err := global.SelectService(service.Advertisement)
	if err != nil {
		return err
	}
	xClient, err := rpc.GetXClient(service.Advertisement, srv.Id, srv.Ip, cast.ToString(srv.Port))
	if err != nil {
		return err
	}
	return xClient.Call(ctx, "FetchIdListOfADOfCarousel", req, rsp)
}

func FetchRecordsOfADOfCarousel(ctx context.Context, req *FetchRecordsOfADOfCarouselReq, rsp *FetchRecordsOfADOfCarouselRsp) error {
	srv, err := global.SelectService(service.Advertisement)
	if err != nil {
		return err
	}
	xClient, err := rpc.GetXClient(service.Advertisement, srv.Id, srv.Ip, cast.ToString(srv.Port))
	if err != nil {
		return err
	}
	return xClient.Call(ctx, "FetchRecordsOfADOfCarousel", req, rsp)
}

// Deals
func FetchVersionOfADOfDeals(ctx context.Context, req *FetchVersionOfADOfDealsReq, rsp *FetchVersionOfADOfDealsRsp) error {
	srv, err := global.SelectService(service.Advertisement)
	if err != nil {
		return err
	}
	xClient, err := rpc.GetXClient(service.Advertisement, srv.Id, srv.Ip, cast.ToString(srv.Port))
	if err != nil {
		return err
	}
	return xClient.Call(ctx, "FetchVersionOfADOfDeals", req, rsp)
}

func FetchIdListOfADOfDeals(ctx context.Context, req *FetchIdListOfADOfDealsReq, rsp *FetchIdListOfADOfDealsRsp) error {
	srv, err := global.SelectService(service.Advertisement)
	if err != nil {
		return err
	}
	xClient, err := rpc.GetXClient(service.Advertisement, srv.Id, srv.Ip, cast.ToString(srv.Port))
	if err != nil {
		return err
	}
	return xClient.Call(ctx, "FetchIdListOfADOfDeals", req, rsp)
}

func FetchRecordsOfADOfDeals(ctx context.Context, req *FetchRecordsOfADOfDealsReq, rsp *FetchRecordsOfADOfDealsRsp) error {
	srv, err := global.SelectService(service.Advertisement)
	if err != nil {
		return err
	}
	xClient, err := rpc.GetXClient(service.Advertisement, srv.Id, srv.Ip, cast.ToString(srv.Port))
	if err != nil {
		return err
	}
	return xClient.Call(ctx, "FetchRecordsOfADOfDeals", req, rsp)
}

// Hots
func FetchVersionOfADOfCamping(ctx context.Context, req *FetchVersionOfADOfCampingReq, rsp *FetchVersionOfADOfCampingRsp) error {
	srv, err := global.SelectService(service.Advertisement)
	if err != nil {
		return err
	}
	xClient, err := rpc.GetXClient(service.Advertisement, srv.Id, srv.Ip, cast.ToString(srv.Port))
	if err != nil {
		return err
	}
	return xClient.Call(ctx, "FetchVersionOfADOfCamping", req, rsp)
}

func FetchIdListOfADOfCamping(ctx context.Context, req *FetchIdListOfADOfCampingReq, rsp *FetchIdListOfADOfCampingRsp) error {
	srv, err := global.SelectService(service.Advertisement)
	if err != nil {
		return err
	}
	xClient, err := rpc.GetXClient(service.Advertisement, srv.Id, srv.Ip, cast.ToString(srv.Port))
	if err != nil {
		return err
	}
	return xClient.Call(ctx, "FetchIdListOfADOfCamping", req, rsp)
}

func FetchRecordsOfADOfCamping(ctx context.Context, req *FetchRecordsOfADOfCampingReq, rsp *FetchRecordsOfADOfCampingRsp) error {
	srv, err := global.SelectService(service.Advertisement)
	if err != nil {
		return err
	}
	xClient, err := rpc.GetXClient(service.Advertisement, srv.Id, srv.Ip, cast.ToString(srv.Port))
	if err != nil {
		return err
	}
	return xClient.Call(ctx, "FetchRecordsOfADOfCamping", req, rsp)
}

// Barbecue
func FetchVersionOfADOfBarbecue(ctx context.Context, req *FetchVersionOfADOfBarbecueReq, rsp *FetchVersionOfADOfBarbecueRsp) error {
	srv, err := global.SelectService(service.Advertisement)
	if err != nil {
		return err
	}
	xClient, err := rpc.GetXClient(service.Advertisement, srv.Id, srv.Ip, cast.ToString(srv.Port))
	if err != nil {
		return err
	}
	return xClient.Call(ctx, "FetchVersionOfADOfBarbecue", req, rsp)
}

func FetchIdListOfADOfBarbecue(ctx context.Context, req *FetchIdListOfADOfBarbecueReq, rsp *FetchIdListOfADOfBarbecueRsp) error {
	srv, err := global.SelectService(service.Advertisement)
	if err != nil {
		return err
	}
	xClient, err := rpc.GetXClient(service.Advertisement, srv.Id, srv.Ip, cast.ToString(srv.Port))
	if err != nil {
		return err
	}
	return xClient.Call(ctx, "FetchIdListOfADOfBarbecue", req, rsp)
}

func FetchRecordsOfADOfBarbecue(ctx context.Context, req *FetchRecordsOfADOfBarbecueReq, rsp *FetchRecordsOfADOfBarbecueRsp) error {
	srv, err := global.SelectService(service.Advertisement)
	if err != nil {
		return err
	}
	xClient, err := rpc.GetXClient(service.Advertisement, srv.Id, srv.Ip, cast.ToString(srv.Port))
	if err != nil {
		return err
	}
	return xClient.Call(ctx, "FetchRecordsOfADOfBarbecue", req, rsp)
}

// Snacks
func FetchVersionOfADOfSnacks(ctx context.Context, req *FetchVersionOfADOfSnacksReq, rsp *FetchVersionOfADOfSnacksRsp) error {
	srv, err := global.SelectService(service.Advertisement)
	if err != nil {
		return err
	}
	xClient, err := rpc.GetXClient(service.Advertisement, srv.Id, srv.Ip, cast.ToString(srv.Port))
	if err != nil {
		return err
	}
	return xClient.Call(ctx, "FetchVersionOfADOfSnacks", req, rsp)
}

func FetchIdListOfADOfSnacks(ctx context.Context, req *FetchIdListOfADOfSnacksReq, rsp *FetchIdListOfADOfSnacksRsp) error {
	srv, err := global.SelectService(service.Advertisement)
	if err != nil {
		return err
	}
	xClient, err := rpc.GetXClient(service.Advertisement, srv.Id, srv.Ip, cast.ToString(srv.Port))
	if err != nil {
		return err
	}
	return xClient.Call(ctx, "FetchIdListOfADOfSnacks", req, rsp)
}

func FetchRecordsOfADOfSnacks(ctx context.Context, req *FetchRecordsOfADOfSnacksReq, rsp *FetchRecordsOfADOfSnacksRsp) error {
	srv, err := global.SelectService(service.Advertisement)
	if err != nil {
		return err
	}
	xClient, err := rpc.GetXClient(service.Advertisement, srv.Id, srv.Ip, cast.ToString(srv.Port))
	if err != nil {
		return err
	}
	return xClient.Call(ctx, "FetchRecordsOfADOfSnacks", req, rsp)
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
