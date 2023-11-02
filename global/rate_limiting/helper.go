package rate_limiting

import (
	"backstage/global/config"
	"backstage/global/log"
	"github.com/spf13/cast"
	"time"
)

func UpdateRateLimit() {
	cf := config.RateLimitingConf()
	if cf == nil {
		log.Error("UpdateRateLimit failure, cf == nil")
		return
	}
	for _, rateLimit := range cf.RateLimiting {
		Store(cast.ToString(rateLimit.Major), cast.ToString(rateLimit.Minor), time.Duration(int32(rateLimit.Period))*time.Millisecond)
	}
}

func GetRateLimitingConfig() ([]string, []int, []int, []int) {
	nameList, majorList, minorList, periodList := []string{}, []int{}, []int{}, []int{}
	cf := config.RateLimitingConf()
	if cf == nil {
		log.Error("GetRateLimitingConfig failure, cf == nil")
		return nameList, majorList, minorList, periodList
	}

	for name, rateLimit := range cf.RateLimiting {
		nameList = append(nameList, name)
		majorList = append(majorList, rateLimit.Major)
		minorList = append(minorList, rateLimit.Minor)
		periodList = append(periodList, rateLimit.Period)
	}

	return nameList, majorList, minorList, periodList
}
