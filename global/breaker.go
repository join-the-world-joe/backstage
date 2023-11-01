package global

import "backstage/abstract/breaker"

var _breaker chan breaker.Breaker

func SetBreaker(brk chan breaker.Breaker) {
	_breaker = brk
}

func Breaker() chan breaker.Breaker {
	return _breaker
}
