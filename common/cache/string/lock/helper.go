package template

import (
	"fmt"
	"strconv"
	"time"
)

func GetWhich(id int) string {
	return fmt.Sprintf("%v%v", "test", strconv.Itoa(id))
}

func GetKey(id int) string {
	return fmt.Sprintf(Format, strconv.Itoa(id))
}

func GetDB() int64 {
	return 0
}

func GetHoldingTime() time.Duration {
	return holdingTime
}

func GetTimeout() time.Duration {
	return timeout
}

func GetRetry() time.Duration {
	return retry
}

func GetMod() int {
	return mod
}
