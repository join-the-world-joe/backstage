package template

import (
	"time"
)

func GetWhich() string {
	return "test"
}

func GetKey() string {
	return Key
}

func GetDB() int64 {
	return 0
}

func GetExpire() time.Duration {
	return Expire
}
