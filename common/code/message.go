package code

import (
	"fmt"
)

func Message(code int) string {
	if msg, exist := en[code]; exist {
		return msg
	}
	return fmt.Sprintf(`code[%v], %s`, code, en[ServiceError])
}
