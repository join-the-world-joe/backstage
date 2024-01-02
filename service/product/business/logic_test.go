package business

import (
	"encoding/json"
	"testing"
)

func TestUnmarshalList(t *testing.T) {
	js := `[1, 2, 3]`
	ls := []int{}
	err := json.Unmarshal([]byte(js), &ls)
	if err != nil {
		t.Error(err)
	}
	t.Log(ls[0], ls[1], ls[2])
}
