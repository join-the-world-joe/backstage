package business

import (
	"backstage/common/code"
	"backstage/common/protocol/gateway"
	"backstage/global/log"
	"backstage/global/rate_limiting"
	"backstage/utils/json"
	"context"
)

func FetchRateLimitingConfig(ctx context.Context, req *gateway.FetchRateLimitingConfigReq, rsp *gateway.FetchRateLimitingConfigRsp) error {
	nameList, majorList, minorList, periodList := rate_limiting.GetRateLimitingConfig()
	js := json.New()
	for k, v := range nameList {
		js.SetPath([]string{"rate_limit_list", v, "major"}, majorList[k])
		js.SetPath([]string{"rate_limit_list", v, "minor"}, minorList[k])
		js.SetPath([]string{"rate_limit_list", v, "period"}, periodList[k])
	}
	bytes, err := js.Encode()
	if err != nil {
		log.ErrorF("FetchRateLimitingConfig failure, err: ", err.Error())
		rsp.Code = code.InternalError
		return nil
	}
	rsp.Body = bytes
	rsp.Code = code.Success
	return nil
}
