package gateway

import "encoding/json"

type FetchRateLimitingConfigReq struct {
}

type FetchRateLimitingConfigRsp struct {
	Code int             `json:"code"`
	Body json.RawMessage `json:"body"`
}
