package limiter

import "time"

type Limiter interface {
	Name() string
	Take() time.Time
	Allow() bool
	Destroy()
}
