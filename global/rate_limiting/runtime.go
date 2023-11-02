package rate_limiting

import (
	"backstage/common/macro/rate_limit"
	"errors"
	"fmt"
	"sync"
	"time"
)

var _rateLimit sync.Map // key: major-minor, value: the period for making request

func Store(major, minor string, period time.Duration) {
	_rateLimit.Store(rate_limit.Key(major, minor), period)
}

func Load(major, minor string) (time.Duration, error) {
	key := rate_limit.Key(major, minor)
	value, ok := _rateLimit.Load(key)
	if ok {
		return value.(time.Duration), nil
	}
	return 0, errors.New(fmt.Sprintf("%s doesn't exist", key))
}
