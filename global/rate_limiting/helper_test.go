package rate_limiting

import (
	"backstage/common/macro/rate_limit"
	"backstage/diagnostic"
	"backstage/utils/json"
	"testing"
	"time"
)

func TestGetRateLimitingConfig(t *testing.T) {
	diagnostic.SetupLogger()
	diagnostic.SetupRateLimiting()
	nameList, majorList, minorList, periodList := GetRateLimitingConfig()
	js := json.New()
	for k, v := range nameList {
		js.SetPath([]string{"rate_limit_list", v, "major"}, majorList[k])
		js.SetPath([]string{"rate_limit_list", v, "minor"}, minorList[k])
		js.SetPath([]string{"rate_limit_list", v, "period"}, periodList[k])
	}
	bytes, err := js.Encode()
	if err != nil {
		t.Fatal(err)
	}
	t.Log(string(bytes))
}

func TestDuration(t *testing.T) {
	millisecond := 1000
	t.Log(rate_limit.DefaultPeriod)
	t.Log(time.Duration(int32(millisecond)) * time.Millisecond)

}

func TestStoreLoad(t *testing.T) {
	major, minor := "1", "2"
	Store(major, minor, time.Duration(int32(1000))*time.Millisecond)
	period, err := Load(major, minor)
	if err != nil {
		t.Fatal(err)
	}
	t.Log("period: ", period)
}
