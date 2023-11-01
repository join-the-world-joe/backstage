package verification_code

import (
	"fmt"
	"time"
)

func GetWhich() string {
	return "test"
}

func GetKey(behavior, countryCode, phoneNumber string) string {
	return fmt.Sprintf(Format, behavior, countryCode, phoneNumber)
}

func GetDB() int64 {
	return 0
}

func GetExpire() time.Duration {
	return Expire
}
