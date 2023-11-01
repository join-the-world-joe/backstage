package token

import "time"

const (
	Mod    = 1
	Format = "hash:token:%v" // key as uuid
	Expire = time.Hour * 24
)
