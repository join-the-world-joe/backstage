package advertisement

import (
	"backstage/common/macro/service"
	"backstage/global"
	"backstage/global/rpc"
	"context"
	"github.com/spf13/cast"
)

func FetchADOfCarousel(ctx context.Context, req *FetchADOfCarouselReq, rsp *FetchADOfCarouselRsp) error {
	srv, err := global.SelectService(service.Advertisement)
	if err != nil {
		return err
	}
	xClient, err := rpc.GetXClient(service.Advertisement, srv.Id, srv.Ip, cast.ToString(srv.Port))
	if err != nil {
		return err
	}
	return xClient.Call(ctx, "FetchADOfCarousel", req, rsp)
}

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

func FetchADOfDealsOfToday(ctx context.Context, req *FetchADOfDealsOfTodayReq, rsp *FetchADOfDealsOfTodayRsp) error {
	srv, err := global.SelectService(service.Advertisement)
	if err != nil {
		return err
	}
	xClient, err := rpc.GetXClient(service.Advertisement, srv.Id, srv.Ip, cast.ToString(srv.Port))
	if err != nil {
		return err
	}
	return xClient.Call(ctx, "FetchADOfDealsOfToday", req, rsp)
}

func FetchVersionOfADOfDealsOfToday(ctx context.Context, req *FetchVersionOfADOfDealsOfTodayReq, rsp *FetchVersionOfADOfDealsOfTodayRsp) error {
	srv, err := global.SelectService(service.Advertisement)
	if err != nil {
		return err
	}
	xClient, err := rpc.GetXClient(service.Advertisement, srv.Id, srv.Ip, cast.ToString(srv.Port))
	if err != nil {
		return err
	}
	return xClient.Call(ctx, "FetchVersionOfADOfDealsOfToday", req, rsp)
}

func FetchADOfHotDeals(ctx context.Context, req *FetchADOfHotDealsReq, rsp *FetchADOfHotDealsRsp) error {
	srv, err := global.SelectService(service.Advertisement)
	if err != nil {
		return err
	}
	xClient, err := rpc.GetXClient(service.Advertisement, srv.Id, srv.Ip, cast.ToString(srv.Port))
	if err != nil {
		return err
	}
	return xClient.Call(ctx, "FetchADOfHotDeals", req, rsp)
}

func FetchVersionOfADOfHotDeals(ctx context.Context, req *FetchVersionOfADOfHotDealsReq, rsp *FetchVersionOfADOfHotDealsRsp) error {
	srv, err := global.SelectService(service.Advertisement)
	if err != nil {
		return err
	}
	xClient, err := rpc.GetXClient(service.Advertisement, srv.Id, srv.Ip, cast.ToString(srv.Port))
	if err != nil {
		return err
	}
	return xClient.Call(ctx, "FetchVersionOfADOfHotDeals", req, rsp)
}

func FetchADOfBBQProducts(ctx context.Context, req *FetchADOfBBQProductsReq, rsp *FetchADOfBBQProductsRsp) error {
	srv, err := global.SelectService(service.Advertisement)
	if err != nil {
		return err
	}
	xClient, err := rpc.GetXClient(service.Advertisement, srv.Id, srv.Ip, cast.ToString(srv.Port))
	if err != nil {
		return err
	}
	return xClient.Call(ctx, "FetchADOfBBQProducts", req, rsp)
}

func FetchVersionOfADOfBBQProducts(ctx context.Context, req *FetchVersionOfADOfBBQProductsReq, rsp *FetchVersionOfADOfBBQProductsRsp) error {
	srv, err := global.SelectService(service.Advertisement)
	if err != nil {
		return err
	}
	xClient, err := rpc.GetXClient(service.Advertisement, srv.Id, srv.Ip, cast.ToString(srv.Port))
	if err != nil {
		return err
	}
	return xClient.Call(ctx, "FetchVersionOfADOfBBQProducts", req, rsp)
}

func FetchADOfSnackProducts(ctx context.Context, req *FetchADOfSnackProductsReq, rsp *FetchADOfSnackProductsRsp) error {
	srv, err := global.SelectService(service.Advertisement)
	if err != nil {
		return err
	}
	xClient, err := rpc.GetXClient(service.Advertisement, srv.Id, srv.Ip, cast.ToString(srv.Port))
	if err != nil {
		return err
	}
	return xClient.Call(ctx, "FetchADOfSnackProducts", req, rsp)
}

func FetchVersionOfADOfSnackProducts(ctx context.Context, req *FetchVersionOfADOfSnackProductsReq, rsp *FetchVersionOfADOfSnackProductsRsp) error {
	srv, err := global.SelectService(service.Advertisement)
	if err != nil {
		return err
	}
	xClient, err := rpc.GetXClient(service.Advertisement, srv.Id, srv.Ip, cast.ToString(srv.Port))
	if err != nil {
		return err
	}
	return xClient.Call(ctx, "FetchVersionOfADOfSnackProducts", req, rsp)
}
