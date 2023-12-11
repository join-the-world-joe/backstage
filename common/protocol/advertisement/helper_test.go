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
	req := &FetchIdListOfADOfDealsReq{}
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

func TestFetchVersionOfADOfHots(t *testing.T) {
	diagnostic.SetupLogger()
	diagnostic.SetupRegistry()
	req := &FetchVersionOfADOfHotsReq{}
	rsp := &FetchVersionOfADOfHotsRsp{}
	err := FetchVersionOfADOfHots(context.Background(), req, rsp)
	if err != nil {
		t.Fatal(err)
	}
	t.Log("code: ", rsp.Code)
	t.Log("body: ", string(rsp.Body))
}

func TestFetchIdListOfADOfHots(t *testing.T) {
	diagnostic.SetupLogger()
	diagnostic.SetupRegistry()
	req := &FetchIdListOfADOfHotsReq{}
	rsp := &FetchIdListOfADOfHotsRsp{}
	err := FetchIdListOfADOfHots(context.Background(), req, rsp)
	if err != nil {
		t.Fatal(err)
	}
	t.Log("code: ", rsp.Code)
	t.Log("body: ", string(rsp.Body))
}

func TestFetchRecordsOfADOfHots(t *testing.T) {
	diagnostic.SetupLogger()
	diagnostic.SetupRegistry()
	advertisementIdList := []int64{1, 2, 3}
	req := &FetchRecordsOfADOfHotsReq{
		AdvertisementIdList: advertisementIdList,
	}
	rsp := &FetchRecordsOfADOfHotsRsp{}
	err := FetchRecordsOfADOfHots(context.Background(), req, rsp)
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
