package oss

import (
	"backstage/diagnostic"
	"context"
	"testing"
)

func TestFetchHeaderListOfObjectFileList(t *testing.T) {
	diagnostic.SetupLogger()
	diagnostic.SetupRegistry()

	userId := int64(1)
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
	req := &FetchHeaderListOfObjectFileListReq{
		UserId:         userId,
		NameListOfFile: nameListOfFile,
	}
	rsp := &FetchHeaderListOfObjectFileListRsp{}

	err := FetchHeaderListOfObjectFileList(context.Background(), req, rsp)
	if err != nil {
		t.Fatal(err)
	}

	t.Log("code: ", rsp.Code)
	t.Log("body: ", string(rsp.Body))
}

func TestVerifyObjectFileList(t *testing.T) {
	diagnostic.SetupLogger()
	diagnostic.SetupRegistry()

	userId := int64(1)
	ossFolder := ""
	nameListOfObjectFile := []string{"3.jpg", "2.jpg"}
	req := &VerifyObjectFileListReq{
		UserId:               userId,
		OSSFolder:            ossFolder,
		NameListOfObjectFile: nameListOfObjectFile,
	}
	rsp := &VerifyObjectFileListRsp{}

	err := VerifyObjectFileList(context.Background(), req, rsp)
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
