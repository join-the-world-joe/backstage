package timestamp

import (
	"testing"
	"time"
)

func TestGetBeginOfADay(t *testing.T) {
	year := 2023
	month := time.Month(11)
	day := 17
	location := "Asia/Kolkata"
	ti := GetBeginOfADay(year, month, day, location)
	t.Log(ti)
	t.Log("timestamp: ", ti)
}

func TestGetEndOfADay(t *testing.T) {
	year := 2023
	month := time.Month(11)
	day := 17
	location := "Asia/Kolkata"
	ti := GetEndOfADay(year, month, day, location)
	t.Log(ti)
	t.Log("timestamp: ", ti)
}
