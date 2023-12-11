package rpc

import (
	"backstage/common/payload"
	"backstage/common/protocol/advertisement"
	"backstage/global"
	"backstage/service/advertisement/business"
	"context"
)

type Async struct {
}

func (p *Async) Forward(ctx context.Context, packet *payload.PacketInternal, rsp *interface{}) error {
	if err := global.Forward().Push(packet); err != nil {
		return err
	}
	return nil
}

// Carousel
func (p *Async) FetchVersionOfADOfCarousel(ctx context.Context, req *advertisement.FetchVersionOfADOfCarouselReq, rsp *advertisement.FetchVersionOfADOfCarouselRsp) error {
	return business.FetchVersionOfADOfCarousel(ctx, req, rsp)
}

func (p *Async) FetchIdListOfADOfCarousel(ctx context.Context, req *advertisement.FetchIdListOfADOfCarouselReq, rsp *advertisement.FetchIdListOfADOfCarouselRsp) error {
	return business.FetchIdListOfADOfCarousel(ctx, req, rsp)
}

func (p *Async) FetchRecordsOfADOfCarousel(ctx context.Context, req *advertisement.FetchRecordsOfADOfCarouselReq, rsp *advertisement.FetchRecordsOfADOfCarouselRsp) error {
	return business.FetchRecordsOfADOfCarousel(ctx, req, rsp)
}

// Deals
func (p *Async) FetchVersionOfADOfDeals(ctx context.Context, req *advertisement.FetchVersionOfADOfDealsReq, rsp *advertisement.FetchVersionOfADOfDealsRsp) error {
	return business.FetchVersionOfADOfDeals(ctx, req, rsp)
}

func (p *Async) FetchIdListOfADOfDeals(ctx context.Context, req *advertisement.FetchIdListOfADOfDealsReq, rsp *advertisement.FetchIdListOfADOfDealsRsp) error {
	return business.FetchIdListOfADOfDeals(ctx, req, rsp)
}

func (p *Async) FetchRecordsOfADOfDeals(ctx context.Context, req *advertisement.FetchRecordsOfADOfDealsReq, rsp *advertisement.FetchRecordsOfADOfDealsRsp) error {
	return business.FetchRecordsOfADOfDeals(ctx, req, rsp)
}

// Deals
func (p *Async) FetchVersionOfADOfHots(ctx context.Context, req *advertisement.FetchVersionOfADOfHotsReq, rsp *advertisement.FetchVersionOfADOfHotsRsp) error {
	return business.FetchVersionOfADOfHots(ctx, req, rsp)
}

func (p *Async) FetchIdListOfADOfHots(ctx context.Context, req *advertisement.FetchIdListOfADOfHotsReq, rsp *advertisement.FetchIdListOfADOfHotsRsp) error {
	return business.FetchIdListOfADOfHots(ctx, req, rsp)
}

func (p *Async) FetchRecordsOfADOfHots(ctx context.Context, req *advertisement.FetchRecordsOfADOfHotsReq, rsp *advertisement.FetchRecordsOfADOfHotsRsp) error {
	return business.FetchRecordsOfADOfHots(ctx, req, rsp)
}

// Barbecue
func (p *Async) FetchVersionOfADOfBarbecue(ctx context.Context, req *advertisement.FetchVersionOfADOfBarbecueReq, rsp *advertisement.FetchVersionOfADOfBarbecueRsp) error {
	return business.FetchVersionOfADOfBarbecue(ctx, req, rsp)
}

func (p *Async) FetchIdListOfADOfBarbecue(ctx context.Context, req *advertisement.FetchIdListOfADOfBarbecueReq, rsp *advertisement.FetchIdListOfADOfBarbecueRsp) error {
	return business.FetchIdListOfADOfBarbecue(ctx, req, rsp)
}

func (p *Async) FetchRecordsOfADOfBarbecue(ctx context.Context, req *advertisement.FetchRecordsOfADOfBarbecueReq, rsp *advertisement.FetchRecordsOfADOfBarbecueRsp) error {
	return business.FetchRecordsOfADOfBarbecue(ctx, req, rsp)
}

// Snacks
func (p *Async) FetchVersionOfADOfSnacks(ctx context.Context, req *advertisement.FetchVersionOfADOfSnacksReq, rsp *advertisement.FetchVersionOfADOfSnacksRsp) error {
	return business.FetchVersionOfADOfSnacks(ctx, req, rsp)
}

func (p *Async) FetchIdListOfADOfSnacks(ctx context.Context, req *advertisement.FetchIdListOfADOfSnacksReq, rsp *advertisement.FetchIdListOfADOfSnacksRsp) error {
	return business.FetchIdListOfADOfSnacks(ctx, req, rsp)
}

func (p *Async) FetchRecordsOfADOfSnacks(ctx context.Context, req *advertisement.FetchRecordsOfADOfSnacksReq, rsp *advertisement.FetchRecordsOfADOfSnacksRsp) error {
	return business.FetchRecordsOfADOfSnacks(ctx, req, rsp)
}

//func (p *Async) FetchADOfCarousel(ctx context.Context, req *advertisement.FetchADOfCarouselReq, rsp *advertisement.FetchADOfCarouselRsp) error {
//	return business.FetchADOfCarousel(ctx, req, rsp)
//}
//

//
//func (p *Async) FetchADOfDealsOfToday(ctx context.Context, req *advertisement.FetchADOfDealsOfTodayReq, rsp *advertisement.FetchADOfDealsOfTodayRsp) error {
//	return business.FetchADOfDealsOfToday(ctx, req, rsp)
//}
//
//func (p *Async) FetchVersionOfADOfDealsOfToday(ctx context.Context, req *advertisement.FetchVersionOfADOfDealsOfTodayReq, rsp *advertisement.FetchVersionOfADOfDealsOfTodayRsp) error {
//	return business.FetchVersionOfADOfDealsOfToday(ctx, req, rsp)
//}
//
//func (p *Async) FetchADOfHotDeals(ctx context.Context, req *advertisement.FetchADOfHotDealsReq, rsp *advertisement.FetchADOfHotDealsRsp) error {
//	return business.FetchADOfHotDeals(ctx, req, rsp)
//}
//
//func (p *Async) FetchVersionOfADOfHotDeals(ctx context.Context, req *advertisement.FetchVersionOfADOfHotDealsReq, rsp *advertisement.FetchVersionOfADOfHotDealsRsp) error {
//	return business.FetchVersionOfADOfHotDeals(ctx, req, rsp)
//}
//
//func (p *Async) FetchADOfBBQProducts(ctx context.Context, req *advertisement.FetchADOfBBQProductsReq, rsp *advertisement.FetchADOfBBQProductsRsp) error {
//	return business.FetchADOfBBQProducts(ctx, req, rsp)
//}
//
//func (p *Async) FetchVersionOfADOfBBQProducts(ctx context.Context, req *advertisement.FetchVersionOfADOfBBQProductsReq, rsp *advertisement.FetchVersionOfADOfBBQProductsRsp) error {
//	return business.FetchVersionOfADOfBBQProducts(ctx, req, rsp)
//}
//
//func (p *Async) FetchADOfSnackProducts(ctx context.Context, req *advertisement.FetchADOfSnackProductsReq, rsp *advertisement.FetchADOfSnackProductsRsp) error {
//	return business.FetchADOfSnackProducts(ctx, req, rsp)
//}
//
//func (p *Async) FetchVersionOfADOfSnackProducts(ctx context.Context, req *advertisement.FetchVersionOfADOfSnackProductsReq, rsp *advertisement.FetchVersionOfADOfSnackProductsRsp) error {
//	return business.FetchVersionOfADOfSnackProducts(ctx, req, rsp)
//}
