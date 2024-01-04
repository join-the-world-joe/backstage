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
	"fmt"
	"github.com/spf13/cast"
)

type OutputOfObjectFileList struct {
	OSSFolder            string   `json:"oss_folder"`
	NameListOfObjectFile []string `json:"name_list_of_object_file"`
}

func VerifyObjectFileList(ctx context.Context, req *oss.VerifyObjectFileListReq, rsp *oss.VerifyObjectFileListRsp) error {
	if !hasPermission(
		cast.ToInt(major.OSS),
		cast.ToInt(oss.VerifyObjectFileListReq_),
		req.UserId,
	) {
		rsp.Code = code.AccessDenied
		return nil
	}

	fmt.Println(req)

	if len(req.OSSFolder) <= 0 || len(req.NameListOfObjectFile) <= 0 {
		rsp.Code = code.InvalidData
		return nil
	}

	aliyun, err := aliyun_oss.NewOSS(
		aliyun_oss.WithAccessKeyId(config.OSSConf().OSS[oss2.AliYun].ID),
		aliyun_oss.WithAccessKeySecret(config.OSSConf().OSS[oss2.AliYun].Secret),
		aliyun_oss.WithEndpoint(config.OSSConf().OSS[oss2.AliYun].Endpoint),
	)
	if err != nil {
		rsp.Code = code.InternalError
		return nil
	}

	for _, v := range req.NameListOfObjectFile {
		path := fmt.Sprintf(oss2.FormatOfFullPathOfObjectFileOfAdvertisement, req.OSSFolder, v)
		fmt.Println("path: ", path)
		b, err := aliyun.IsObjectExist(oss2.AdvertisementImageBucket, path)
		if err != nil {
			rsp.Code = code.UnexpectedNetworkError
			return nil
		}
		if !b {
			rsp.Code = code.EntryNotFound
			return nil
		}
	}

	output := &OutputOfObjectFileList{
		OSSFolder:            req.OSSFolder,
		NameListOfObjectFile: req.NameListOfObjectFile,
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
