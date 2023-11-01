package global

import "backstage/abstract/notifier"

var _notifier notifier.Notifier

func SetNotifier(n notifier.Notifier) {
	_notifier = n
}

func Notifier() notifier.Notifier {
	return _notifier
}
