package oss

import (
	"backstage/diagnostic"
	"context"
	"testing"
)

func TestFetchHeaderListOfObjectFileListOfAdvertisement(t *testing.T) {
	diagnostic.SetupLogger()
	diagnostic.SetupRegistry()

	userId := int64(1)
	advertisementId := int64(1)
	nameListOfFile := []string{
		"0.webp",
		"1.jpg",
		"2.png",
		"3.jpeg",
		"4.bmp",
		"5.svg",
		"6.svg",
		"7.gif",
	}
	req := &FetchHeaderListOfObjectFileListOfAdvertisementReq{
		UserId:          userId,
		AdvertisementId: advertisementId,
		NameListOfFile:  nameListOfFile,
	}
	rsp := &FetchHeaderListOfObjectFileListOfAdvertisementRsp{}

	err := FetchHeaderListOfObjectFileListOfAdvertisement(context.Background(), req, rsp)
	if err != nil {
		t.Fatal(err)
	}

	t.Log("code: ", rsp.Code)
	t.Log("body: ", string(rsp.Body))
}

func TestVerifyObjectFileListOfAdvertisement(t *testing.T) {
	diagnostic.SetupLogger()
	diagnostic.SetupRegistry()

	userId := int64(1)
	advertisementId := int64(1)
	nameListOfObjectFile := []string{"3.jpg", "2.jpg"}
	req := &VerifyObjectFileListOfAdvertisementReq{
		UserId:               userId,
		AdvertisementId:      advertisementId,
		NameListOfObjectFile: nameListOfObjectFile,
	}
	rsp := &VerifyObjectFileListOfAdvertisementRsp{}

	err := VerifyObjectFileListOfAdvertisement(context.Background(), req, rsp)
	if err != nil {
		t.Fatal(err)
	}
	t.Log("code: ", rsp.Code)
	t.Log("body: ", string(rsp.Body))
}

func TestRemoveListOfObjectFile(t *testing.T) {
	diagnostic.SetupLogger()
	diagnostic.SetupRegistry()

	userId := int64(1)
	listOfObjectFile := []string{
		"48/0.png",
		"48/3.png",
	}
	req := &RemoveListOfObjectFileReq{
		UserId:           userId,
		ListOfObjectFile: listOfObjectFile,
	}
	rsp := &RemoveListOfObjectFileRsp{}
	err := RemoveListOfObjectFile(context.Background(), req, rsp)
	if err != nil {
		t.Fatal(err)
	}
	t.Log("code: ", rsp.Code)
	t.Log("body: ", string(rsp.Body))
}
