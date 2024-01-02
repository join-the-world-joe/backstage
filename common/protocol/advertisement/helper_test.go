package advertisement

import (
	"backstage/diagnostic"
	"context"
	"testing"
)

func TestFetchVersionOfADOfCarousel(t *testing.T) {
	diagnostic.SetupLogger()
	diagnostic.SetupRegistry()
	req := &FetchVersionOfADOfCarouselReq{}
	rsp := &FetchVersionOfADOfCarouselRsp{}
	err := FetchVersionOfADOfCarousel(context.Background(), req, rsp)
	if err != nil {
		t.Fatal(err)
	}
	t.Log("code: ", rsp.Code)
	t.Log("body: ", string(rsp.Body))
}

func TestFetchIdListOfADOfCarousel(t *testing.T) {
	diagnostic.SetupLogger()
	diagnostic.SetupRegistry()
	userId := int64(1)
	req := &FetchIdListOfADOfCarouselReq{
		UserId: userId,
	}
	rsp := &FetchIdListOfADOfCarouselRsp{}
	err := FetchIdListOfADOfCarousel(context.Background(), req, rsp)
	if err != nil {
		t.Fatal(err)
	}
	t.Log("code: ", rsp.Code)
	t.Log("body: ", string(rsp.Body))
}

func TestFetchRecordsOfADOfCarousel(t *testing.T) {
	diagnostic.SetupLogger()
	diagnostic.SetupRegistry()
	advertisementIdList := []int64{1, 2, 3}
	req := &FetchRecordsOfADOfCarouselReq{
		AdvertisementIdList: advertisementIdList,
	}
	rsp := &FetchRecordsOfADOfCarouselRsp{}
	err := FetchRecordsOfADOfCarousel(context.Background(), req, rsp)
	if err != nil {
		t.Fatal(err)
	}
	t.Log("code: ", rsp.Code)
	t.Log("body: ", string(rsp.Body))
}

func TestFetchVersionOfADOfDeals(t *testing.T) {
	diagnostic.SetupLogger()
	diagnostic.SetupRegistry()
	req := &FetchVersionOfADOfDealsReq{}
	rsp := &FetchVersionOfADOfDealsRsp{}
	err := FetchVersionOfADOfDeals(context.Background(), req, rsp)
	if err != nil {
		t.Fatal(err)
	}
	t.Log("code: ", rsp.Code)
	t.Log("body: ", string(rsp.Body))
}

func TestFetchIdListOfADOfDeals(t *testing.T) {
	diagnostic.SetupLogger()
	diagnostic.SetupRegistry()
	req := &FetchIdListOfADOfDealsReq{
		Behavior: 0,
	}
	rsp := &FetchIdListOfADOfDealsRsp{}
	err := FetchIdListOfADOfDeals(context.Background(), req, rsp)
	if err != nil {
		t.Fatal(err)
	}
	t.Log("code: ", rsp.Code)
	t.Log("body: ", string(rsp.Body))
}

func TestFetchRecordsOfADOfDeals(t *testing.T) {
	diagnostic.SetupLogger()
	diagnostic.SetupRegistry()
	advertisementIdList := []int64{1, 2, 3}
	req := &FetchRecordsOfADOfDealsReq{
		AdvertisementIdList: advertisementIdList,
	}
	rsp := &FetchRecordsOfADOfDealsRsp{}
	err := FetchRecordsOfADOfDeals(context.Background(), req, rsp)
	if err != nil {
		t.Fatal(err)
	}
	t.Log("code: ", rsp.Code)
	t.Log("body: ", string(rsp.Body))
}

func TestFetchVersionOfADOfCamping(t *testing.T) {
	diagnostic.SetupLogger()
	diagnostic.SetupRegistry()
	req := &FetchVersionOfADOfCampingReq{}
	rsp := &FetchVersionOfADOfCampingRsp{}
	err := FetchVersionOfADOfCamping(context.Background(), req, rsp)
	if err != nil {
		t.Fatal(err)
	}
	t.Log("code: ", rsp.Code)
	t.Log("body: ", string(rsp.Body))
}

func TestFetchIdListOfADOfCamping(t *testing.T) {
	diagnostic.SetupLogger()
	diagnostic.SetupRegistry()
	req := &FetchIdListOfADOfCampingReq{}
	rsp := &FetchIdListOfADOfCampingRsp{}
	err := FetchIdListOfADOfCamping(context.Background(), req, rsp)
	if err != nil {
		t.Fatal(err)
	}
	t.Log("code: ", rsp.Code)
	t.Log("body: ", string(rsp.Body))
}

func TestFetchRecordsOfADOfCamping(t *testing.T) {
	diagnostic.SetupLogger()
	diagnostic.SetupRegistry()
	advertisementIdList := []int64{1, 2, 3}
	req := &FetchRecordsOfADOfCampingReq{
		AdvertisementIdList: advertisementIdList,
	}
	rsp := &FetchRecordsOfADOfCampingRsp{}
	err := FetchRecordsOfADOfCamping(context.Background(), req, rsp)
	if err != nil {
		t.Fatal(err)
	}
	t.Log("code: ", rsp.Code)
	t.Log("body: ", string(rsp.Body))
}

func TestFetchVersionOfADOfBarbecue(t *testing.T) {
	diagnostic.SetupLogger()
	diagnostic.SetupRegistry()
	req := &FetchVersionOfADOfBarbecueReq{}
	rsp := &FetchVersionOfADOfBarbecueRsp{}
	err := FetchVersionOfADOfBarbecue(context.Background(), req, rsp)
	if err != nil {
		t.Fatal(err)
	}
	t.Log("code: ", rsp.Code)
	t.Log("body: ", string(rsp.Body))
}

func TestFetchIdListOfADOfBarbecue(t *testing.T) {
	diagnostic.SetupLogger()
	diagnostic.SetupRegistry()
	req := &FetchIdListOfADOfBarbecueReq{}
	rsp := &FetchIdListOfADOfBarbecueRsp{}
	err := FetchIdListOfADOfBarbecue(context.Background(), req, rsp)
	if err != nil {
		t.Fatal(err)
	}
	t.Log("code: ", rsp.Code)
	t.Log("body: ", string(rsp.Body))
}

func TestFetchRecordsOfADOfBarbecue(t *testing.T) {
	diagnostic.SetupLogger()
	diagnostic.SetupRegistry()
	advertisementIdList := []int64{1, 2, 3}
	req := &FetchRecordsOfADOfBarbecueReq{
		AdvertisementIdList: advertisementIdList,
	}
	rsp := &FetchRecordsOfADOfBarbecueRsp{}
	err := FetchRecordsOfADOfBarbecue(context.Background(), req, rsp)
	if err != nil {
		t.Fatal(err)
	}
	t.Log("code: ", rsp.Code)
	t.Log("body: ", string(rsp.Body))
}

func TestFetchVersionOfADOfSnacks(t *testing.T) {
	diagnostic.SetupLogger()
	diagnostic.SetupRegistry()
	req := &FetchVersionOfADOfSnacksReq{}
	rsp := &FetchVersionOfADOfSnacksRsp{}
	err := FetchVersionOfADOfSnacks(context.Background(), req, rsp)
	if err != nil {
		t.Fatal(err)
	}
	t.Log("code: ", rsp.Code)
	t.Log("body: ", string(rsp.Body))
}

func TestFetchIdListOfADOfSnacks(t *testing.T) {
	diagnostic.SetupLogger()
	diagnostic.SetupRegistry()
	req := &FetchIdListOfADOfSnacksReq{}
	rsp := &FetchIdListOfADOfSnacksRsp{}
	err := FetchIdListOfADOfSnacks(context.Background(), req, rsp)
	if err != nil {
		t.Fatal(err)
	}
	t.Log("code: ", rsp.Code)
	t.Log("body: ", string(rsp.Body))
}

func TestFetchRecordsOfADOfSnacks(t *testing.T) {
	diagnostic.SetupLogger()
	diagnostic.SetupRegistry()
	advertisementIdList := []int64{1, 2, 3}
	req := &FetchRecordsOfADOfSnacksReq{
		AdvertisementIdList: advertisementIdList,
	}
	rsp := &FetchRecordsOfADOfSnacksRsp{}
	err := FetchRecordsOfADOfSnacks(context.Background(), req, rsp)
	if err != nil {
		t.Fatal(err)
	}
	t.Log("code: ", rsp.Code)
	t.Log("body: ", string(rsp.Body))
}

func TestFetchIdListOfAdvertisement(t *testing.T) {
	diagnostic.SetupLogger()
	diagnostic.SetupRegistry()
	userId := int64(1)
	name := []byte("title1")
	behavior := 1
	req := &FetchIdListOfAdvertisementReq{
		Behavior:          behavior,
		UserId:            userId,
		AdvertisementName: name,
	}
	rsp := &FetchIdListOfAdvertisementRsp{}
	err := FetchIdListOfAdvertisement(context.Background(), req, rsp)
	if err != nil {
		t.Fatal(err)
	}
	t.Log("bytes: ", string(rsp.Body))
}

func TestFetchRecordsOfAdvertisement(t *testing.T) {
	userId := int64(1)
	idList := []int64{4}
	diagnostic.SetupLogger()
	diagnostic.SetupRegistry()
	req := &FetchRecordsOfAdvertisementReq{
		UserId:              userId,
		AdvertisementIdList: idList,
	}
	rsp := &FetchRecordsOfAdvertisementRsp{}
	err := FetchRecordsOfAdvertisement(context.Background(), req, rsp)
	if err != nil {
		t.Fatal(err)
	}
	t.Log("rsp: ", string(rsp.Body))
}
