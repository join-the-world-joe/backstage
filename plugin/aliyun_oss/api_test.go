package aliyun_oss

import (
	oss2 "backstage/common/macro/oss"
	"backstage/diagnostic"
	"backstage/global/config"
	"crypto/hmac"
	"crypto/sha1"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"mime"
	"net/http"
	"os"
	"testing"
	"time"
)

func TestPutImage(t *testing.T) {
	diagnostic.SetupOSS()
	id := config.OSSConf().OSS[oss2.AliYun].ID
	secret := []byte(config.OSSConf().OSS[oss2.AliYun].Secret)
	bucket := oss2.AdvertisementImageBucket
	endpoint := config.OSSConf().OSS[oss2.AliYun].Endpoint
	objectFile := "5.jpg"
	contentType := "image/png"
	date := time.Now().UTC().Format(http.TimeFormat)
	//plainText := "PUT\n\n" + contentType + "\n" + date + "\n" + "/" + bucket + "/" + objectFile
	plainText := "PUT\n\n" + contentType + "\n" + date + "\n" + "x-oss-date:" + date + "\n" + "/" + bucket + "/" + objectFile

	key := []byte(secret)
	mac := hmac.New(sha1.New, key)
	mac.Write([]byte(plainText))
	// Base64编码。
	signature := base64.StdEncoding.EncodeToString(mac.Sum(nil))

	url := "https://" + bucket + "." + endpoint + "/" + objectFile

	file, err := os.Open("D:\\Projects\\github\\mini_program\\asset\\image\\5.jpg")
	if err != nil {
		t.Fatal(err)
	}
	//payload := strings.NewReader("{go:test}")

	req, _ := http.NewRequest("PUT", url, file)

	req.Header.Add("Authorization", "OSS "+id+":"+signature)
	req.Header.Add("Content-Type", contentType)
	req.Header.Add("Date", date)
	req.Header.Add("x-oss-date", date)

	t.Log(req.Header)

	bytes, err := json.Marshal(&req.Header)
	if err != nil {
		t.Fatal(err)
	}

	t.Log("headers: ", string(bytes))

	t.Log(url)

	res, _ := http.DefaultClient.Do(req)
	//
	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	fmt.Println(res)
	fmt.Println(string(body))
}

func TestPutVideo(t *testing.T) {
	diagnostic.SetupOSS()
	id := config.OSSConf().OSS[oss2.AliYun].ID
	secret := []byte(config.OSSConf().OSS[oss2.AliYun].Secret)
	bucket := oss2.AdvertisementImageBucket
	endpoint := config.OSSConf().OSS[oss2.AliYun].Endpoint
	objectFile := "1.mp4"
	contentType := mime.TypeByExtension(objectFile)
	date := time.Now().UTC().Format(http.TimeFormat)
	//plainText := "PUT\n\n" + contentType + "\n" + date + "\n" + "/" + bucket + "/" + objectFile
	plainText := "PUT\n\n" + contentType + "\n" + date + "\n" + "x-oss-date:" + date + "\n" + "/" + bucket + "/" + objectFile

	key := []byte(secret)
	mac := hmac.New(sha1.New, key)
	mac.Write([]byte(plainText))
	// Base64编码。
	signature := base64.StdEncoding.EncodeToString(mac.Sum(nil))

	url := "https://" + bucket + "." + endpoint + "/" + objectFile

	file, err := os.Open("C:\\Users\\Joe\\Desktop\\oss\\1.mp4")
	if err != nil {
		t.Fatal(err)
	}
	//payload := strings.NewReader("{go:test}")

	req, _ := http.NewRequest("PUT", url, file)

	req.Header.Add("Authorization", "OSS "+id+":"+signature)
	req.Header.Add("Content-Type", contentType)
	req.Header.Add("Date", date)
	req.Header.Add("x-oss-date", date)

	t.Log(req.Header)

	bytes, err := json.Marshal(&req.Header)
	if err != nil {
		t.Fatal(err)
	}

	t.Log("headers: ", string(bytes))

	t.Log(url)

	//res, _ := http.DefaultClient.Do(req)
	//
	//defer res.Body.Close()
	//body, _ := ioutil.ReadAll(res.Body)
	//
	//fmt.Println(res)
	//fmt.Println(string(body))
}
