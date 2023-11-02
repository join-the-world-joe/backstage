package runtime

import "sync/atomic"

const (
	_begin = uint64(0)
	_end   = uint64(1000) // todo: set end of the sequence
)

var _sequence uint64

func init() {
	_sequence = _begin
}

func PopSequence() uint64 {
	// todo: once _sequence reaches the end, set it to _begin
	return atomic.AddUint64(&_sequence, 1)
}
