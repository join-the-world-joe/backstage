package rpc

import (
	"backstage/common/payload"
	"backstage/common/protocol/admin"
	"backstage/global"
	"backstage/service/admin/business"
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

func (p *Async) SignIn(ctx context.Context, req *admin.SignInReq, rsp *admin.SignInRsp) error {
	return business.SignIn(ctx, req, rsp)
}

func (p *Async) FetchMenuListOfCondition(ctx context.Context, req *admin.FetchMenuListOfConditionReq, rsp *admin.FetchMenuListOfConditionRsp) error {
	return business.FetchMenuListOfCondition(ctx, req, rsp)
}

func (p *Async) InsertRecordOfGood(ctx context.Context, req *admin.InsertRecordOfGoodReq, rsp *admin.InsertRecordOfGoodRsp) error {
	return business.InsertRecordOfGood(ctx, req, rsp)
}

func (p *Async) FetchIdListOfGood(ctx context.Context, req *admin.FetchIdListOfGoodReq, rsp *admin.FetchIdListOfGoodRsp) error {
	return business.FetchIdListOfGood(ctx, req, rsp)
}

func (p *Async) FetchRecordsOfGood(ctx context.Context, req *admin.FetchRecordsOfGoodReq, rsp *admin.FetchRecordsOfGoodRsp) error {
	return business.FetchRecordsOfGood(ctx, req, rsp)
}

func (p *Async) SoftDeleteRecordOfGood(ctx context.Context, req *admin.SoftDeleteRecordsOfGoodReq, rsp *admin.SoftDeleteRecordsOfGoodRsp) error {
	return business.SoftDeleteRecordsOfGood(ctx, req, rsp)
}

func (p *Async) UpdateRecordOfGood(ctx context.Context, req *admin.UpdateRecordOfGoodReq, rsp *admin.UpdateRecordOfGoodRsp) error {
	return business.UpdateRecordOfGood(ctx, req, rsp)
}

func (p *Async) FetchIdListOfAdvertisement(ctx context.Context, req *admin.FetchIdListOfAdvertisementReq, rsp *admin.FetchIdListOfAdvertisementRsp) error {
	return business.FetchIdListOfAdvertisement(ctx, req, rsp)
}

func (p *Async) FetchRecordsOfAdvertisement(ctx context.Context, req *admin.FetchRecordsOfAdvertisementReq, rsp *admin.FetchRecordsOfAdvertisementRsp) error {
	return business.FetchRecordsOfAdvertisement(ctx, req, rsp)
}

func (p *Async) InsertRecordOfAdvertisement(ctx context.Context, req *admin.InsertRecordOfAdvertisementReq, rsp *admin.InsertRecordOfAdvertisementRsp) error {
	return business.InsertRecordOfAdvertisement(ctx, req, rsp)
}

func (p *Async) SoftDeleteRecordsOfAdvertisement(ctx context.Context, req *admin.SoftDeleteRecordsOfAdvertisementReq, rsp *admin.SoftDeleteRecordsOfAdvertisementRsp) error {
	return business.SoftDeleteRecordsOfAdvertisement(ctx, req, rsp)
}

func (p *Async) UpdateRecordOfAdvertisement(ctx context.Context, req *admin.UpdateRecordOfAdvertisementReq, rsp *admin.UpdateRecordOfAdvertisementRsp) error {
	return business.UpdateRecordOfAdvertisement(ctx, req, rsp)
}

func (p *Async) InsertRecordOfADOfCarousel(ctx context.Context, req *admin.InsertRecordOfADOfCarouselReq, rsp *admin.InsertRecordOfADOfCarouselRsp) error {
	return business.InsertRecordOfADOfCarousel(ctx, req, rsp)
}

func (p *Async) RemoveOutdatedRecordsOfADOfCarousel(ctx context.Context, req *admin.RemoveOutdatedRecordsOfADOfCarouselReq, rsp *admin.RemoveOutdatedRecordsOfADOfCarouselRsp) error {
	return business.RemoveOutdatedRecordsOfADOfCarousel(ctx, req, rsp)
}

func (p *Async) InsertRecordOfADOfBarbecue(ctx context.Context, req *admin.InsertRecordOfADOfBarbecueReq, rsp *admin.InsertRecordOfADOfBarbecueRsp) error {
	return business.InsertRecordOfADOfBarbecue(ctx, req, rsp)
}

func (p *Async) RemoveOutdatedRecordsOfADOfBarbecue(ctx context.Context, req *admin.RemoveOutdatedRecordsOfADOfBarbecueReq, rsp *admin.RemoveOutdatedRecordsOfADOfBarbecueRsp) error {
	return business.RemoveOutdatedRecordsOfADOfBarbecue(ctx, req, rsp)
}

func (p *Async) InsertRecordOfADOfDeals(ctx context.Context, req *admin.InsertRecordOfADOfDealsReq, rsp *admin.InsertRecordOfADOfDealsRsp) error {
	return business.InsertRecordOfADOfDeals(ctx, req, rsp)
}

func (p *Async) RemoveOutdatedRecordsOfADOfDeals(ctx context.Context, req *admin.RemoveOutdatedRecordsOfADOfDealsReq, rsp *admin.RemoveOutdatedRecordsOfADOfDealsRsp) error {
	return business.RemoveOutdatedRecordsOfADOfDeals(ctx, req, rsp)
}

func (p *Async) InsertRecordOfADOfCamping(ctx context.Context, req *admin.InsertRecordOfADOfCampingReq, rsp *admin.InsertRecordOfADOfCampingRsp) error {
	return business.InsertRecordOfADOfCamping(ctx, req, rsp)
}

func (p *Async) RemoveOutdatedRecordsOfADOfCamping(ctx context.Context, req *admin.RemoveOutdatedRecordsOfADOfCampingReq, rsp *admin.RemoveOutdatedRecordsOfADOfCampingRsp) error {
	return business.RemoveOutdatedRecordsOfADOfCamping(ctx, req, rsp)
}

func (p *Async) InsertRecordOfADOfSnacks(ctx context.Context, req *admin.InsertRecordOfADOfSnacksReq, rsp *admin.InsertRecordOfADOfSnacksRsp) error {
	return business.InsertRecordOfADOfSnacks(ctx, req, rsp)
}

func (p *Async) RemoveOutdatedRecordsOfADOfSnacks(ctx context.Context, req *admin.RemoveOutdatedRecordsOfADOfSnacksReq, rsp *admin.RemoveOutdatedRecordsOfADOfSnacksRsp) error {
	return business.RemoveOutdatedRecordsOfADOfSnacks(ctx, req, rsp)
}
