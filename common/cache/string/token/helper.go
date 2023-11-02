package token

import (
	"fmt"
	"time"
)

func GetWhich() string {
	return "test"
}

func GetKey(unique string) string {
	return fmt.Sprintf(format, unique)
}

func GetDB() int64 {
	return 0
}

func GetExpire() time.Duration {
	return expire
}
