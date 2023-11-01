package rate

import (
	"testing"
	"time"
)

func TestTake(t *testing.T) {
	limiter, err := NewLimiter(
		WithCapacity(10),
		WithFrequency(2),
	)
	if err != nil {
		t.Error(err)
		return
	}

	begin := time.Now()

	count := 0
	for {
		limiter.Take()
		count++
		//t.Log(count)
		if time.Now().Sub(begin) > time.Second*3 {
			break
		}
	}
	t.Log("spent: ", time.Now().Sub(begin), ", count: ", count)
	limiter.Destroy()
}

func TestAllow(t *testing.T) {
	limiter, err := NewLimiter(
		WithCapacity(10),
		WithFrequency(2),
	)
	if err != nil {
		t.Error(err)
		return
	}

	count := 0
	last := time.Now()

	for {
		b := limiter.Allow()
		if b {
			count++
			t.Log("count: ", count, ", spent: ", time.Now().Sub(last))
			last = time.Now()
		}
		if count > 16 {
			break
		}
	}

	limiter.Destroy()
}
