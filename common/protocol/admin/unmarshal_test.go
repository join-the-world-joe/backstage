package admin

import (
	"backstage/common/payload"
	"encoding/json"
	"testing"
)

func TestSignInRsp(t *testing.T) {
	//var temp = &SignInRsp{}
}

func TestUnmarshalInsertRecordOfAdvertisement(t *testing.T) {
	//bytes := []byte{123, 34, 104, 101, 97, 100, 101, 114, 34, 58, 123, 34, 109, 97, 106, 111, 114, 34, 58, 34, 53, 34, 44, 34, 109, 105, 110, 111, 114, 34, 58, 34, 51, 55, 34, 125, 44, 34, 98, 111, 100, 121, 34, 58, 123, 34, 110, 97, 109, 101, 34, 58, 91, 53, 53, 44, 53, 53, 44, 53, 53, 93, 44, 34, 116, 105, 116, 108, 101, 34, 58, 91, 53, 53, 44, 53, 53, 93, 44, 34, 115, 101, 108, 108, 105, 110, 103, 95, 112, 111, 105, 110, 116, 34, 58, 34, 91, 91, 50, 51, 48, 44, 49, 52, 55, 44, 49, 52, 49, 44, 50, 51, 48, 44, 49, 51, 48, 44, 49, 51, 50, 44, 50, 51, 48, 44, 49, 53, 53, 44, 49, 55, 56, 93, 93, 34, 44, 34, 115, 101, 108, 108, 105, 110, 103, 95, 112, 114, 105, 99, 101, 34, 58, 48, 44, 34, 100, 101, 115, 99, 114, 105, 112, 116, 105, 111, 110, 34, 58, 91, 93, 44, 34, 115, 116, 97, 116, 117, 115, 34, 58, 49, 44, 34, 117, 114, 108, 34, 58, 91, 93, 44, 34, 112, 114, 111, 100, 117, 99, 116, 95, 105, 100, 34, 58, 55, 55, 55, 44, 34, 115, 116, 111, 99, 107, 34, 58, 48, 44, 34, 112, 108, 97, 99, 101, 95, 111, 102, 95, 111, 114, 105, 103, 105, 110, 34, 58, 91, 93, 125, 125}
	//
	//packet := &payload.PacketClient{}
	//err := json.Unmarshal(bytes, packet)
	//if err != nil {
	//	t.Fatal(err)
	//}
	//req := &InsertRecordOfAdvertisementReq{}
	//err = json.Unmarshal(packet.Body, req)
	//if err != nil {
	//	t.Fatal(err)
	//}
	//t.Log("bytes: ", string(bytes))
	//t.Log("selling points: ", string(req.SellingPoints))
	//
	//var biBytes = [][]byte{}
	//err = json.Unmarshal([]byte(req.SellingPoints), &biBytes)
	//if err != nil {
	//	t.Fatal(err)
	//}
	//for _, v := range biBytes {
	//	t.Log("v :", string(v))
	//}

	//t.Log("selling points.len: ", len(req.SellingPoint))
}

func TestUnmarshalInsertRecordOfAdvertisement1(t *testing.T) {
	bytes := []byte{123, 34, 104, 101, 97, 100, 101, 114, 34, 58, 123, 34, 109, 97, 106, 111, 114, 34, 58, 34, 53, 34, 44, 34, 109, 105, 110, 111, 114, 34, 58, 34, 51, 55, 34, 125, 44, 34, 98, 111, 100, 121, 34, 58, 123, 34, 110, 97, 109, 101, 34, 58, 91, 52, 57, 44, 52, 57, 44, 52, 57, 93, 44, 34, 116, 105, 116, 108, 101, 34, 58, 91, 93, 44, 34, 115, 101, 108, 108, 105, 110, 103, 95, 112, 111, 105, 110, 116, 34, 58, 34, 91, 91, 50, 51, 48, 44, 49, 54, 51, 44, 49, 55, 52, 93, 93, 34, 44, 34, 115, 101, 108, 108, 105, 110, 103, 95, 112, 114, 105, 99, 101, 34, 58, 48, 44, 34, 100, 101, 115, 99, 114, 105, 112, 116, 105, 111, 110, 34, 58, 91, 93, 44, 34, 115, 116, 97, 116, 117, 115, 34, 58, 49, 44, 34, 117, 114, 108, 34, 58, 91, 93, 44, 34, 112, 114, 111, 100, 117, 99, 116, 95, 105, 100, 34, 58, 49, 49, 49, 44, 34, 115, 116, 111, 99, 107, 34, 58, 48, 44, 34, 112, 108, 97, 99, 101, 95, 111, 102, 95, 111, 114, 105, 103, 105, 110, 34, 58, 91, 93, 125, 125}
	temp := map[string]interface{}{}
	err := json.Unmarshal(bytes, &temp)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(temp["selling_point"])
	//var biArray [][]byte
	//err := json.Unmarshal([]byte(bytes), &biArray)
	//if err != nil {
	//	t.Fatal(err)
	//}
	//t.Log("temp: ", temp)
}

func TestUnMarshal(t *testing.T) {
	bytes := `[[229, 164, 167, 229, 174, 182, 229, 165, 189]]`
	//bytes := ` ["森林","大树","point3"]`
	var temp [][]byte
	//var temp []string
	err := json.Unmarshal([]byte(bytes), &temp)
	if err != nil {
		t.Fatal(err)
	}
	t.Log("temp", temp)
	for _, v := range temp {
		t.Log("v: ", string(v))
	}
}

func TestUnmarshalUpdateRecordOfAdvertisement(t *testing.T) {
	bytes := []byte{171, 205, 28, 130, 28, 35, 176, 121, 112, 169, 215, 212, 45, 230, 171, 164, 194, 241, 77, 28, 59, 89, 117, 20, 19, 135, 44, 52, 217, 92, 112, 21, 155, 218, 145, 254, 80, 7, 52, 160, 20, 64, 206, 213, 129, 255, 16, 204, 30, 50, 17, 233, 34, 68, 127, 107, 116, 220, 215, 182, 74, 140, 175, 49, 165, 251, 54, 215, 71, 232, 25, 205, 127, 240, 153, 17, 218, 93, 105, 168, 214, 182, 20, 97, 109, 128, 51, 18, 202, 179, 32, 223, 238, 248, 54, 56}
	packet := &payload.PacketClient{}
	err := json.Unmarshal(bytes, packet)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(packet)
}

func TestUnmarshalThumnail(t *testing.T) {
	//bytes := []byte{102, 100, 115, 97, 102, 49, 49}
	bytes := []byte{102, 100, 115, 97, 102}
	thumbnail := []byte{}

	err := json.Unmarshal(bytes, &thumbnail)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(string(thumbnail))
}
