package notify

import (
	"backstage/common/notifier"
	"backstage/global"
	"backstage/global/broker"
)

func Common(argument []string) {
	n := len(argument)
	if n > 0 {
		switch argument[0] {
		case notifier.Dump:
			if n == 2 {
				if argument[1] == notifier.Broker { // Dump Broker
					broker.DumpBroker()
				} else if argument[1] == notifier.Registry { // Dump Registry
					global.DumpRegistry()
				}
			} else if n == 3 {

			}
			break
		}
	}
}
