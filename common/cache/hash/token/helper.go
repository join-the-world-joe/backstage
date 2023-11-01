package token

import (
	"fmt"
	"time"
)

func GetWhich() string {
	return "test"
}

func GetKey(uuid string) string {
	return fmt.Sprintf(Format, uuid)
}

func GetDB() int64 {
	return 0
}

func GetExpire() time.Duration {
	return Expire
}
