package payload

import "encoding/json"

type Result struct {
	Code    int
	Message string
	Data    interface{}
}

func (p *Result) SetCode(code int) {
	p.Code = code
}

func (p *Result) SetMessage(msg string) {
	p.Message = msg
}

func (p *Result) SetData(data interface{}) {
	p.Data = data
}

func (p *Result) ToBytes() []byte {
	bytes, err := json.Marshal(p)
	if err != nil {
		return []byte("")
	}

	return bytes
}

func (p *Result) ToString() string {
	bytes, err := json.Marshal(p)
	if err != nil {
		return ""
	}
	return string(bytes)
}
