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
