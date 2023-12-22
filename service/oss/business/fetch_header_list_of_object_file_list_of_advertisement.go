package business

import (
	"backstage/common/code"
	oss2 "backstage/common/macro/oss"
	"backstage/common/major"
	"backstage/common/protocol/oss"
	"backstage/global/config"
	"backstage/global/log"
	"context"
	"crypto/hmac"
	"crypto/sha1"
	"encoding/base64"
	"encoding/json"
	"github.com/spf13/cast"
	"mime"
	"net/http"
	"path"
	"time"
)

type ObjectFileRequestHeader struct {
	Authorization string `json:"authorization"`
	ContentType   string `json:"content_type"`
	Date          string `json:"date"`
	XOSSDate      string `json:"x_oss_date"`
}

type OutputOfHeaderListOfObjectFileListOfAdvertisement struct {
	AdvertisementId int64                               `json:"advertisement_id"`
	RequestHeader   map[string]*ObjectFileRequestHeader `json:"request_header"`
	CommonPath      string                              `json:"common_path"`
	Host            string                              `json:"host"`
}

func FetchHeaderListOfObjectFileListOfAdvertisement(ctx context.Context, req *oss.FetchHeaderListOfObjectFileListOfAdvertisementReq, rsp *oss.FetchHeaderListOfObjectFileListOfAdvertisementRsp) error {
	if !hasPermission(
		cast.ToInt(major.OSS),
		cast.ToInt(oss.FetchHeaderListOfObjectFileListOfAdvertisementReq_),
		req.UserId,
	) {
		rsp.Code = code.AccessDenied
		return nil
	}

	if req.AdvertisementId <= 0 || len(req.NameListOfFile) <= 0 {
		rsp.Code = code.InvalidData
		return nil
	}

	id := config.OSSConf().OSS[oss2.AliYun].ID
	bucket := oss2.AdvertisementImageBucket
	secret := []byte(config.OSSConf().OSS[oss2.AliYun].Secret)
	endpoint := config.OSSConf().OSS[oss2.AliYun].Endpoint
	output := &OutputOfHeaderListOfObjectFileListOfAdvertisement{
		AdvertisementId: req.AdvertisementId,
		RequestHeader:   map[string]*ObjectFileRequestHeader{},
		Host:            bucket + "." + endpoint,
		CommonPath:      "https://" + bucket + "." + endpoint + "/", // add object file to tail
	}

	for _, v := range req.NameListOfFile {
		objectFile := v
		contentType := mime.TypeByExtension(path.Ext(v))
		date := time.Now().UTC().Format(http.TimeFormat)
		//plainText := "PUT\n\n" + contentType + "\n" + date + "\n" + "/" + bucket + "/" + objectFile
		plainText := "PUT\n\n" + contentType + "\n" + date + "\n" + "x-oss-date:" + date + "\n" + "/" + bucket + "/" + objectFile

		mac := hmac.New(sha1.New, secret)
		mac.Write([]byte(plainText))
		signature := base64.StdEncoding.EncodeToString(mac.Sum(nil))
		output.RequestHeader[v] = &ObjectFileRequestHeader{
			Date:          date,
			Authorization: "OSS " + id + ":" + signature,
			ContentType:   contentType,
			XOSSDate:      date,
		}
	}

	bytes, err := json.Marshal(output)
	if err != nil {
		log.Error("json.Marshal failure, err: ", err)
		rsp.Code = code.InvalidData
		return nil
	}

	rsp.Body = bytes
	rsp.Code = code.Success
	return nil
}
