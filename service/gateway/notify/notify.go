package notify

import (
	"backstage/common/notifier"
	"backstage/common/notify"
	"backstage/service/gateway/runtime"
	"github.com/spf13/cast"
)

func Handler(argument []string) {
	notify.Common(argument)
	n := len(argument)
	if n > 0 {
		switch argument[0] {
		case notifier.Dump:
			if n == 2 {
				if argument[1] == notifier.Config { // Dump Config
					runtime.DumpConfig()
				}
			} else if n == 3 {
				if argument[1] == notifier.Session { // Dump Session UserId
					if userId := cast.ToInt64(argument[2]); userId > 0 {
						runtime.DumpSession(userId)
					}
				}
			}
			break
		}
	}
}
