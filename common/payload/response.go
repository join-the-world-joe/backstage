package payload

import (
	"backstage/common/code"
	"encoding/json"
	"fmt"
)

type Response struct {
	Code int `json:"code"`
	//Body    []byte `json:"body"`
	Body json.RawMessage `json:"body"`
	body interface{}
}

func (p *Response) SetCode(c int) *Response {
	p.Code = c
	return p
}

func (p *Response) SetBody(body interface{}) *Response {
	p.body = body
	return p
}

func (p *Response) SetRaw(body []byte) *Response {
	p.Body = body
	return p
}

func (p *Response) Bytes() []byte {
	var err error
	if p.body != nil {
		p.Body, err = json.Marshal(p.body)
		if err != nil {
			return []byte(fmt.Sprintf(` {"code":%v, "msg":%s}`, code.InternalError, code.Message(code.InternalError)))
		}
	}

	bytes, err := json.Marshal(p)
	if err != nil {
		return []byte(fmt.Sprintf(` {"code":%v, "msg":%s}`, code.InternalError, code.Message(code.InternalError)))
	}
	return bytes
}
