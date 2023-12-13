package aliyun_oss

import (
	"backstage/common/macro/oss"
	"backstage/diagnostic"
	"backstage/global/config"
	"encoding/json"
	"testing"
)

var bkt = ""

func TestListObject(t *testing.T) {
	diagnostic.SetupOSS()
	aliyunOss, err := NewOSS(
		WithAccessKeyId(config.OSSConf().OSS["Aliyun"].ID),
		WithAccessKeySecret(config.OSSConf().OSS["Aliyun"].Secret),
		WithEndpoint(config.OSSConf().OSS["Aliyun"].Endpoint),
	)
	if err != nil {
		t.Fatal(err)
	}

	objs, err := aliyunOss.ListObject(oss.AdvertisementImageBucket)
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
	objectFileName := "6.jpg"
	nativeFileName := "D:\\Projects\\github\\express\\asset\\image\\6.jpg"
	diagnostic.SetupOSS()
	aliyunOss, err := NewOSS(
		WithAccessKeyId(config.OSSConf().OSS["Aliyun"].ID),
		WithAccessKeySecret(config.OSSConf().OSS["Aliyun"].Secret),
		WithEndpoint(config.OSSConf().OSS["Aliyun"].Endpoint),
	)
	if err != nil {
		t.Fatal(err)
	}

	err = aliyunOss.Put(nativeFileName, oss.AdvertisementImageBucket, objectFileName)
	if err != nil {
		t.Fatal(err)
	}
}

func TestGetObject(t *testing.T) {
	objectFileName := "6.jpg"
	nativeFileName := "D:\\Projects\\github\\express\\asset\\image\\6.bak.1.jpg"
	diagnostic.SetupOSS()
	aliyunOss, err := NewOSS(
		WithAccessKeyId(config.OSSConf().OSS["Aliyun"].ID),
		WithAccessKeySecret(config.OSSConf().OSS["Aliyun"].Secret),
		WithEndpoint(config.OSSConf().OSS["Aliyun"].Endpoint),
	)
	if err != nil {
		t.Fatal(err)
	}

	err = aliyunOss.Get(nativeFileName, oss.AdvertisementImageBucket, objectFileName)
	if err != nil {
		t.Fatal(err)
	}
}

func TestDeleteObject(t *testing.T) {
	objectFileName := "3.jpg"
	diagnostic.SetupOSS()
	aliyunOss, err := NewOSS(
		WithAccessKeyId(config.OSSConf().OSS["Aliyun"].ID),
		WithAccessKeySecret(config.OSSConf().OSS["Aliyun"].Secret),
		WithEndpoint(config.OSSConf().OSS["Aliyun"].Endpoint),
	)
	if err != nil {
		t.Fatal(err)
	}

	err = aliyunOss.Delete(oss.AdvertisementImageBucket, objectFileName)
	if err != nil {
		t.Fatal(err)
	}
}
