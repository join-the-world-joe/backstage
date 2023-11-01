package phonedata

import (
	"encoding/json"
	"testing"
)

func TestExtractPhoneInfo(t *testing.T) {
	phone := "18677455433"
	file := ""
	pd, err := NewPhoneData(file)
	if err != nil {
		t.Error(err)
		return
	}

	rec, err := pd.Find(phone)
	if err != nil {
		t.Error(err)
		return
	}

	bytes, err := json.Marshal(rec)
	if err != nil {
		t.Error(err)
		return
	}

	t.Log(string(bytes))
}
