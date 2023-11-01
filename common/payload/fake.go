package payload

var _FakeRsp FakeRsp

type FakeReq struct {

}

type FakeRsp struct {

}

func GetFakeRsp() *FakeRsp {
	return &_FakeRsp
}
