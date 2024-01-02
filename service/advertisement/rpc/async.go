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
func (p *Async) FetchVersionOfADOfCamping(ctx context.Context, req *advertisement.FetchVersionOfADOfCampingReq, rsp *advertisement.FetchVersionOfADOfCampingRsp) error {
	return business.FetchVersionOfADOfCamping(ctx, req, rsp)
}

func (p *Async) FetchIdListOfADOfCamping(ctx context.Context, req *advertisement.FetchIdListOfADOfCampingReq, rsp *advertisement.FetchIdListOfADOfCampingRsp) error {
	return business.FetchIdListOfADOfCamping(ctx, req, rsp)
}

func (p *Async) FetchRecordsOfADOfCamping(ctx context.Context, req *advertisement.FetchRecordsOfADOfCampingReq, rsp *advertisement.FetchRecordsOfADOfCampingRsp) error {
	return business.FetchRecordsOfADOfCamping(ctx, req, rsp)
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

// advertisement
func (p *Async) FetchIdListOfAdvertisement(ctx context.Context, req *advertisement.FetchIdListOfAdvertisementReq, rsp *advertisement.FetchIdListOfAdvertisementRsp) error {
	return business.FetchIdListOfAdvertisement(ctx, req, rsp)
}

func (p *Async) FetchRecordsOfAdvertisement(ctx context.Context, req *advertisement.FetchRecordsOfAdvertisementReq, rsp *advertisement.FetchRecordsOfAdvertisementRsp) error {
	return business.FetchRecordsOfAdvertisement(ctx, req, rsp)
}
