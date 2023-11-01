package template

import "time"

const (
	mod         = 3
	Format      = "string.lock.template.%v" // value as uuid
	ttl         = time.Hour * 24
	holdingTime = time.Second * 10
	retry       = time.Millisecond * 200
	timeout     = time.Second * 3
)
