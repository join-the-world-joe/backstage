package fifo

import (
	"testing"
	"time"
)

func TestPushPop(t *testing.T) {
	begin := 1
	end := 10
	c := NewFIFO(
		WithBufferSize(end - begin + 1),
	)
	for i := begin; i <= end; i++ {
		if err := c.Push(i); err != nil {
			t.Fatal(err)
		}
	}
	for i := begin; i <= end; i++ {
		num, err := c.Pop()
		if err != nil {
			t.Fatal(err)
		}
		value, ok := num.(int)
		if !ok {
			t.Fatal("!ok")
		}
		if value != i {
			t.Fatal("value != i")
		}
	}
	t.Log("done!")
}

func TestPushTimeout(t *testing.T) {
	pushTimeout := time.Second
	c := NewFIFO(
		WithBufferSize(1),
		WithPushTimeout(pushTimeout),
	)
	c.Push(1)
	begin := time.Now()
	if err := c.Push(2); err != nil {
		t.Log(err, time.Now().Sub(begin))
	}
}

func TestPopTimeout(t *testing.T) {
	popTimeout := time.Second
	c := NewFIFO(
		WithBufferSize(1),
		WithPopTimeout(popTimeout),
	)
	begin := time.Now()
	if _, err := c.Pop(); err != nil {
		t.Log(err, time.Now().Sub(begin))
	}
}

func TestWaitForChannel(t *testing.T) {
	begin := 1
	end := 10
	total := end - begin + 1
	count := 0
	ticker := time.NewTicker(time.Millisecond * 300)

	c := NewFIFO(
		WithBufferSize(total),
	)

OUTSIDE:
	for {
		select {
		case <-ticker.C:
			c.Push(count + 1)
		case value := <-c.Channel():
			t.Log("value = ", value)
			count++
		default:
			if count == total {
				break OUTSIDE
			}
		}
	}

	t.Log("done")
}
