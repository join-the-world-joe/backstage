package business

import (
	"encoding/json"
	"testing"
)

func TestMarshalOutputOfRecordsOfGood(t *testing.T) {
	output := &OutputOfRecordsOfGood{
		RecordsOfGood: map[int64]*RecordOfGood{
			1: &RecordOfGood{
				Id:   1,
				Name: "name1",
			},
		},
	}

	bytes, err := json.Marshal(output)
	if err != nil {
		t.Fatal(err)
	}
	t.Log("bytes: ", string(bytes))
}

func TestMarshalOutputOfIdListOfGood(t *testing.T) {
	idList := []int64{1, 2, 3}
	output := OutputOfIdListOfGood{IdListOfGood: idList}
	bytes, err := json.Marshal(output)
	if err != nil {
		t.Fatal(err)
	}
	t.Log("output: ", string(bytes))
}
