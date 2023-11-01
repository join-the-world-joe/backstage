package verification_code

import "time"

const (
	Mod    = 1
	Format = "string:verification_code:%v:%v-%v" // behavior:country_code-phone_number
	Expire = time.Hour * 24
)

// behavior
const (
	Register = "Register"
	Login    = "Login"  // for server
	SignIn   = "SignIn" // for backend
)
