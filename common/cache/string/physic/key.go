package template

import "time"

const (
	Mod    = 3
	Format = "string.template.%v" // value as uuid
	Expire = time.Hour * 24
)
