package business

import (
	"backstage/common/code"
	oss2 "backstage/common/macro/oss"
	"backstage/common/major"
	"backstage/common/protocol/oss"
	"backstage/global/config"
	"backstage/global/log"
	"backstage/plugin/aliyun_oss"
	"context"
	"encoding/json"
	"github.com/spf13/cast"
)

type OutputOfListOfObjectFile struct {
	ListOfObjectFile []string `json:"list_of_object_file"`
}

func RemoveListOfObjectFile(ctx context.Context, req *oss.RemoveListOfObjectFileReq, rsp *oss.RemoveListOfObjectFileRsp) error {
	if !hasPermission(
		cast.ToInt(major.OSS),
		cast.ToInt(oss.RemoveListOfObjectFileReq_),
		req.UserId,
	) {
		rsp.Code = code.AccessDenied
		return nil
	}

	output := &OutputOfListOfObjectFile{
		ListOfObjectFile: req.ListOfObjectFile,
	}

	aliyunOss, err := aliyun_oss.NewOSS(
		aliyun_oss.WithAccessKeyId(config.OSSConf().OSS[oss2.AliYun].ID),
		aliyun_oss.WithAccessKeySecret(config.OSSConf().OSS[oss2.AliYun].Secret),
		aliyun_oss.WithEndpoint(config.OSSConf().OSS[oss2.AliYun].Endpoint),
	)
	if err != nil {
		log.Error("RemoveListOfObjectFile.aliyun_oss.NewOSS failure, err: ", err)
		rsp.Code = code.InternalError
		return nil
	}

	for _, v := range req.ListOfObjectFile {
		err = aliyunOss.Delete(oss2.AdvertisementImageBucket, v)
		if err != nil {
			log.Error("RemoveListOfObjectFile.aliyunOss.Delete failure, err: ", err)
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
