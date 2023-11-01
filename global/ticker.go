package global

import "time"

var _ticker *time.Ticker

func SetTicker(t *time.Ticker) {
	_ticker = t
}

func Ticker() *time.Ticker {
	return _ticker
}
