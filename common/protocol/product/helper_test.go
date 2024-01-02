package product

import (
	"backstage/diagnostic"
	"context"
	"testing"
)

func TestFetchIdListOfProduct(t *testing.T) {
	userId := int64(1)
	behavior := 1
	productName := []byte("product1")
	diagnostic.SetupLogger()
	diagnostic.SetupRegistry()
	req := &FetchIdListOfProductReq{
		UserId:      userId,
		Behavior:    behavior,
		ProductName: productName,
	}
	rsp := &FetchIdListOfProductRsp{}
	err := FetchIdListOfProduct(context.Background(), req, rsp)
	if err != nil {
		t.Fatal(err)
	}
	t.Log("rsp: ", string(rsp.Body))
}

func TestFetchRecordsOfGood(t *testing.T) {
	userId := int64(1)
	productIdList := []int64{1, 2, 3}
	diagnostic.SetupLogger()
	diagnostic.SetupRegistry()
	req := &FetchRecordsOfProductReq{
		UserId:        userId,
		ProductIdList: productIdList,
	}
	rsp := &FetchRecordsOfProductRsp{}
	err := FetchRecordsOfProduct(context.Background(), req, rsp)
	if err != nil {
		t.Fatal(err)
	}
	t.Log("rsp: ", string(rsp.Body))
}
