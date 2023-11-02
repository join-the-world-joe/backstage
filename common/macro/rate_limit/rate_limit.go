package rate_limit

import (
	"fmt"
	"time"
)

const (
	Format        = "%v-%v"
	DefaultPeriod = (time.Duration(1000) * time.Millisecond)
)

func Key(major, minor string) string {
	return fmt.Sprintf(Format, major, minor)
}
