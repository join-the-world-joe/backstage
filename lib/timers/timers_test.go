package timers

import (
	"testing"
	"time"
)

func TestAddTimer(t *testing.T) {
	timerMng := NewTimers().Start()

	timerMng.AddTimer(
		NewTimer(
			WithId("event"),
			WithLoop(20),
			WithInterval(time.Millisecond*200),
		),
	)

	timerMng.AddTimer(
		NewTimer(
			WithId("notifier_1"),
			WithLoop(2),
			WithInterval(time.Second),
		),
	)

	timerMng.AddTimer(
		NewTimer(
			WithId("destroy"),
			WithLoop(1),
			WithInterval(time.Second*5),
		),
	)

	count := 0

	for {
		select {
		case timer := <-timerMng.Wait():
			t.Log(timer.Id)
			count++
			if count > 10 {
				timerMng.RemoveTimer(timer.Id)
			}
			if timer.Id == "destroy" {
				timerMng.Destroy()
				return
			}
		}
	}
}
