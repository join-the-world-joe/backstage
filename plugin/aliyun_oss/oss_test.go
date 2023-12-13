package aliyun_oss

import (
	"encoding/json"
	"testing"
)

var id = ""
var secret = ""
var endpoint = ""
var bkt = ""

func TestListObject(t *testing.T) {
	aliyunOss, err := NewOSS(
		WithAccessKeyId(id),
		WithAccessKeySecret(secret),
		WithEndpoint(endpoint),
	)
	if err != nil {
		t.Fatal(err)
	}

	objs, err := aliyunOss.ListObject(bkt)
	if err != nil {
		t.Fatal(err)
	}

	bytes, err := json.Marshal(objs)
	if err != nil {
		t.Fatal(err)
	}
	t.Log("bytes: ", string(bytes))
}

func TestPutObject(t *testing.T) {
	objectFileName := "3/6.jpg"
	nativeFileName := "D:\\Projects\\github\\express\\asset\\image\\6.jpg"
	aliyunOss, err := NewOSS(
		WithAccessKeyId(id),
		WithAccessKeySecret(secret),
		WithEndpoint(endpoint),
	)
	if err != nil {
		t.Fatal(err)
	}

	err = aliyunOss.Put(nativeFileName, bkt, objectFileName)
	if err != nil {
		t.Fatal(err)
	}
}

func TestGetObject(t *testing.T) {
	objectFileName := "3/2.jpg"
	nativeFileName := "D:\\Projects\\github\\express\\asset\\image\\2.jpg"
	aliyunOss, err := NewOSS(
		WithAccessKeyId(id),
		WithAccessKeySecret(secret),
		WithEndpoint(endpoint),
	)
	if err != nil {
		t.Fatal(err)
	}

	err = aliyunOss.Get(nativeFileName, bkt, objectFileName)
	if err != nil {
		t.Fatal(err)
	}
}

func TestDeleteObject(t *testing.T) {
	objectFileName := "3/6.jpg"
	aliyunOss, err := NewOSS(
		WithAccessKeyId(id),
		WithAccessKeySecret(secret),
		WithEndpoint(endpoint),
	)
	if err != nil {
		t.Fatal(err)
	}

	err = aliyunOss.Delete(bkt, objectFileName)
	if err != nil {
		t.Fatal(err)
	}
}
