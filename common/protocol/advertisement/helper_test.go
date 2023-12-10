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
	req := &FetchIdListOfADOfCarouselReq{}
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

//
//func TestFetchADOfCarousel(t *testing.T) {
//	diagnostic.SetupLogger()
//	diagnostic.SetupRegistry()
//	req := &FetchADOfCarouselReq{}
//	rsp := &FetchADOfCarouselRsp{}
//	err := FetchADOfCarousel(context.Background(), req, rsp)
//	if err != nil {
//		t.Fatal(err)
//	}
//	t.Log("code: ", rsp.Code)
//	t.Log("version: ", rsp.Version)
//	t.Log("body: ", string(rsp.Body))
//}
//
//func TestFetchADOfDealsOfToday(t *testing.T) {
//	diagnostic.SetupLogger()
//	diagnostic.SetupRegistry()
//	req := &FetchADOfDealsOfTodayReq{}
//	rsp := &FetchADOfDealsOfTodayRsp{}
//	err := FetchADOfDealsOfToday(context.Background(), req, rsp)
//	if err != nil {
//		t.Fatal(err)
//	}
//	t.Log("code: ", rsp.Code)
//	t.Log("body:", string(rsp.Body))
//	t.Log("version: ", rsp.Version)
//}
//
//func TestFetchVersionOfADOfDealsOfToday(t *testing.T) {
//	diagnostic.SetupLogger()
//	diagnostic.SetupRegistry()
//	req := &FetchVersionOfADOfDealsOfTodayReq{}
//	rsp := &FetchVersionOfADOfDealsOfTodayRsp{}
//	err := FetchVersionOfADOfDealsOfToday(context.Background(), req, rsp)
//	if err != nil {
//		t.Fatal(err)
//	}
//	t.Log("code: ", rsp.Code)
//	t.Log("version: ", rsp.Version)
//}
//
//func TestFetchADOfHotDeals(t *testing.T) {
//	diagnostic.SetupLogger()
//	diagnostic.SetupRegistry()
//	req := &FetchADOfHotDealsReq{}
//	rsp := &FetchADOfHotDealsRsp{}
//	err := FetchADOfHotDeals(context.Background(), req, rsp)
//	if err != nil {
//		t.Fatal(err)
//	}
//	t.Log("code: ", rsp.Code)
//	t.Log("body:", string(rsp.Body))
//	t.Log("version: ", rsp.Version)
//}
//
//func TestFetchVersionOfADOfHotDeals(t *testing.T) {
//	diagnostic.SetupLogger()
//	diagnostic.SetupRegistry()
//	req := &FetchVersionOfADOfHotDealsReq{}
//	rsp := &FetchVersionOfADOfHotDealsRsp{}
//	err := FetchVersionOfADOfHotDeals(context.Background(), req, rsp)
//	if err != nil {
//		t.Fatal(err)
//	}
//	t.Log("code: ", rsp.Code)
//	t.Log("version: ", rsp.Version)
//}
//
//func TestFetchADOfBBQProducts(t *testing.T) {
//	diagnostic.SetupLogger()
//	diagnostic.SetupRegistry()
//	req := &FetchADOfBBQProductsReq{}
//	rsp := &FetchADOfBBQProductsRsp{}
//	err := FetchADOfBBQProducts(context.Background(), req, rsp)
//	if err != nil {
//		t.Fatal(err)
//	}
//	t.Log("code: ", rsp.Code)
//	t.Log("body:", string(rsp.Body))
//	t.Log("version: ", rsp.Version)
//}
//
//func TestFetchVersionOfADOfBBQProducts(t *testing.T) {
//	diagnostic.SetupLogger()
//	diagnostic.SetupRegistry()
//	req := &FetchVersionOfADOfBBQProductsReq{}
//	rsp := &FetchVersionOfADOfBBQProductsRsp{}
//	err := FetchVersionOfADOfBBQProducts(context.Background(), req, rsp)
//	if err != nil {
//		t.Fatal(err)
//	}
//	t.Log("code: ", rsp.Code)
//	t.Log("version: ", rsp.Version)
//}
//
//func TestFetchADOfSnackProducts(t *testing.T) {
//	diagnostic.SetupLogger()
//	diagnostic.SetupRegistry()
//	req := &FetchADOfSnackProductsReq{}
//	rsp := &FetchADOfSnackProductsRsp{}
//	err := FetchADOfSnackProducts(context.Background(), req, rsp)
//	if err != nil {
//		t.Fatal(err)
//	}
//	t.Log("code: ", rsp.Code)
//	t.Log("version: ", rsp.Version)
//	t.Log("body:", string(rsp.Body))
//}
//
//func TestFetchVersionOfADOfSnackProducts(t *testing.T) {
//	diagnostic.SetupLogger()
//	diagnostic.SetupRegistry()
//	req := &FetchVersionOfADOfSnackProductsReq{}
//	rsp := &FetchVersionOfADOfSnackProductsRsp{}
//	err := FetchVersionOfADOfSnackProducts(context.Background(), req, rsp)
//	if err != nil {
//		t.Fatal(err)
//	}
//	t.Log("code: ", rsp.Code)
//	t.Log("version: ", rsp.Version)
//}
