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

func GetExpire() time.Duration {
	return Expire
}
