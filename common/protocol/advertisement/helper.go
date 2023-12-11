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
func FetchVersionOfADOfHots(ctx context.Context, req *FetchVersionOfADOfHotsReq, rsp *FetchVersionOfADOfHotsRsp) error {
	srv, err := global.SelectService(service.Advertisement)
	if err != nil {
		return err
	}
	xClient, err := rpc.GetXClient(service.Advertisement, srv.Id, srv.Ip, cast.ToString(srv.Port))
	if err != nil {
		return err
	}
	return xClient.Call(ctx, "FetchVersionOfADOfHots", req, rsp)
}

func FetchIdListOfADOfHots(ctx context.Context, req *FetchIdListOfADOfHotsReq, rsp *FetchIdListOfADOfHotsRsp) error {
	srv, err := global.SelectService(service.Advertisement)
	if err != nil {
		return err
	}
	xClient, err := rpc.GetXClient(service.Advertisement, srv.Id, srv.Ip, cast.ToString(srv.Port))
	if err != nil {
		return err
	}
	return xClient.Call(ctx, "FetchIdListOfADOfHots", req, rsp)
}

func FetchRecordsOfADOfHots(ctx context.Context, req *FetchRecordsOfADOfHotsReq, rsp *FetchRecordsOfADOfHotsRsp) error {
	srv, err := global.SelectService(service.Advertisement)
	if err != nil {
		return err
	}
	xClient, err := rpc.GetXClient(service.Advertisement, srv.Id, srv.Ip, cast.ToString(srv.Port))
	if err != nil {
		return err
	}
	return xClient.Call(ctx, "FetchRecordsOfADOfHots", req, rsp)
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

//func FetchADOfCarousel(ctx context.Context, req *FetchADOfCarouselReq, rsp *FetchADOfCarouselRsp) error {
//	srv, err := global.SelectService(service.Advertisement)
//	if err != nil {
//		return err
//	}
//	xClient, err := rpc.GetXClient(service.Advertisement, srv.Id, srv.Ip, cast.ToString(srv.Port))
//	if err != nil {
//		return err
//	}
//	return xClient.Call(ctx, "FetchADOfCarousel", req, rsp)
//}
//

//
//func FetchADOfDealsOfToday(ctx context.Context, req *FetchADOfDealsOfTodayReq, rsp *FetchADOfDealsOfTodayRsp) error {
//	srv, err := global.SelectService(service.Advertisement)
//	if err != nil {
//		return err
//	}
//	xClient, err := rpc.GetXClient(service.Advertisement, srv.Id, srv.Ip, cast.ToString(srv.Port))
//	if err != nil {
//		return err
//	}
//	return xClient.Call(ctx, "FetchADOfDealsOfToday", req, rsp)
//}
//
//
//func FetchADOfHotDeals(ctx context.Context, req *FetchADOfHotDealsReq, rsp *FetchADOfHotDealsRsp) error {
//	srv, err := global.SelectService(service.Advertisement)
//	if err != nil {
//		return err
//	}
//	xClient, err := rpc.GetXClient(service.Advertisement, srv.Id, srv.Ip, cast.ToString(srv.Port))
//	if err != nil {
//		return err
//	}
//	return xClient.Call(ctx, "FetchADOfHotDeals", req, rsp)
//}
//
//func FetchVersionOfADOfHotDeals(ctx context.Context, req *FetchVersionOfADOfHotDealsReq, rsp *FetchVersionOfADOfHotDealsRsp) error {
//	srv, err := global.SelectService(service.Advertisement)
//	if err != nil {
//		return err
//	}
//	xClient, err := rpc.GetXClient(service.Advertisement, srv.Id, srv.Ip, cast.ToString(srv.Port))
//	if err != nil {
//		return err
//	}
//	return xClient.Call(ctx, "FetchVersionOfADOfHotDeals", req, rsp)
//}
//
//func FetchADOfBBQProducts(ctx context.Context, req *FetchADOfBBQProductsReq, rsp *FetchADOfBBQProductsRsp) error {
//	srv, err := global.SelectService(service.Advertisement)
//	if err != nil {
//		return err
//	}
//	xClient, err := rpc.GetXClient(service.Advertisement, srv.Id, srv.Ip, cast.ToString(srv.Port))
//	if err != nil {
//		return err
//	}
//	return xClient.Call(ctx, "FetchADOfBBQProducts", req, rsp)
//}
//
//func FetchVersionOfADOfBBQProducts(ctx context.Context, req *FetchVersionOfADOfBBQProductsReq, rsp *FetchVersionOfADOfBBQProductsRsp) error {
//	srv, err := global.SelectService(service.Advertisement)
//	if err != nil {
//		return err
//	}
//	xClient, err := rpc.GetXClient(service.Advertisement, srv.Id, srv.Ip, cast.ToString(srv.Port))
//	if err != nil {
//		return err
//	}
//	return xClient.Call(ctx, "FetchVersionOfADOfBBQProducts", req, rsp)
//}
//
//func FetchADOfSnackProducts(ctx context.Context, req *FetchADOfSnackProductsReq, rsp *FetchADOfSnackProductsRsp) error {
//	srv, err := global.SelectService(service.Advertisement)
//	if err != nil {
//		return err
//	}
//	xClient, err := rpc.GetXClient(service.Advertisement, srv.Id, srv.Ip, cast.ToString(srv.Port))
//	if err != nil {
//		return err
//	}
//	return xClient.Call(ctx, "FetchADOfSnackProducts", req, rsp)
//}
//
//func FetchVersionOfADOfSnackProducts(ctx context.Context, req *FetchVersionOfADOfSnackProductsReq, rsp *FetchVersionOfADOfSnackProductsRsp) error {
//	srv, err := global.SelectService(service.Advertisement)
//	if err != nil {
//		return err
//	}
//	xClient, err := rpc.GetXClient(service.Advertisement, srv.Id, srv.Ip, cast.ToString(srv.Port))
//	if err != nil {
//		return err
//	}
//	return xClient.Call(ctx, "FetchVersionOfADOfSnackProducts", req, rsp)
//}
//
