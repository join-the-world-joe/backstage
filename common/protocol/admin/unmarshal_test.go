package admin

import (
	"backstage/common/payload"
	"encoding/json"
	"testing"
)

func TestUnmarshalInsertRecordOfAdvertisement(t *testing.T) {
	bytes := []byte{123, 34, 104, 101, 97, 100, 101, 114, 34, 58, 123, 34, 109, 97, 106, 111, 114, 34, 58, 34, 53, 34, 44, 34, 109, 105, 110, 111, 114, 34, 58, 34, 51, 55, 34, 125, 44, 34, 98, 111, 100, 121, 34, 58, 123, 34, 110, 97, 109, 101, 34, 58, 91, 53, 53, 44, 53, 53, 44, 53, 53, 93, 44, 34, 116, 105, 116, 108, 101, 34, 58, 91, 53, 53, 44, 53, 53, 93, 44, 34, 115, 101, 108, 108, 105, 110, 103, 95, 112, 111, 105, 110, 116, 34, 58, 34, 91, 91, 50, 51, 48, 44, 49, 52, 55, 44, 49, 52, 49, 44, 50, 51, 48, 44, 49, 51, 48, 44, 49, 51, 50, 44, 50, 51, 48, 44, 49, 53, 53, 44, 49, 55, 56, 93, 93, 34, 44, 34, 115, 101, 108, 108, 105, 110, 103, 95, 112, 114, 105, 99, 101, 34, 58, 48, 44, 34, 100, 101, 115, 99, 114, 105, 112, 116, 105, 111, 110, 34, 58, 91, 93, 44, 34, 115, 116, 97, 116, 117, 115, 34, 58, 49, 44, 34, 117, 114, 108, 34, 58, 91, 93, 44, 34, 112, 114, 111, 100, 117, 99, 116, 95, 105, 100, 34, 58, 55, 55, 55, 44, 34, 115, 116, 111, 99, 107, 34, 58, 48, 44, 34, 112, 108, 97, 99, 101, 95, 111, 102, 95, 111, 114, 105, 103, 105, 110, 34, 58, 91, 93, 125, 125}

	packet := &payload.PacketClient{}
	err := json.Unmarshal(bytes, packet)
	if err != nil {
		t.Fatal(err)
	}
	req := &InsertRecordOfAdvertisementReq{}
	err = json.Unmarshal(packet.Body, req)
	if err != nil {
		t.Fatal(err)
	}
	t.Log("bytes: ", string(bytes))
	t.Log("selling points: ", string(req.SellingPoint))

	var biBytes = [][]byte{}
	err = json.Unmarshal([]byte(req.SellingPoint), &biBytes)
	if err != nil {
		t.Fatal(err)
	}
	for _, v := range biBytes {
		t.Log("v :", string(v))
	}

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
	//bytes := `[[230,163,174]]`
	bytes := ` ["森林","大树","point3"]`
	//var temp [][]byte
	var temp []string
	err := json.Unmarshal([]byte(bytes), &temp)
	if err != nil {
		t.Fatal(err)
	}
	t.Log("temp", temp)
	t.Log("temp[0]", string(temp[0]))
}
