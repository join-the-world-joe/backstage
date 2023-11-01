package notifier

import (
	"testing"
	"time"
)

func TestChannelNotifier(t *testing.T) {
	notify, err := NewNotifier(
		WithBufferSize(3),
		WithEmitTimeout(time.Microsecond*100),
	)
	if err != nil {
		t.Error(err)
		return
	}

	//notify := new(channel.ChanNotifier)

	go func() {
		//time.Sleep(time.Second*2)
		for {
			select {
			case sig := <-notify.Wait():
				t.Log("sig = ", sig)
				if sig == "" {
					t.Log("return")
					return
				}
			}
		}
	}()

	count := 0
	for {
		//time.Sleep(time.Second*2)
		err = notify.Emit("test")
		if err != nil {
			t.Error(err)
			return
		}
		count++
		if count > 5 {
			break
		}
		time.Sleep(time.Second)
	}

	notify.Destroy()

	time.Sleep(time.Second * 3)
}
