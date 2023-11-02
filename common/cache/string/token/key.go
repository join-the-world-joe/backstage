package token

import "time"

const (
	mod    = 1
	format = "string:token:%v" //
	expire = time.Hour * 10    // ten hours
)
