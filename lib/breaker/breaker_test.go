package breaker

import (
	breaker2 "go-micro-framework/abstract/breaker"
	"testing"
	"time"
)

func TestBreaker(t *testing.T) {

	breaker, err := NewBreaker()
	if err != nil {
		t.Error(err)
		return
	}

	breakerChan := make(chan breaker2.Breaker, 10)

	go func() {
		count := 0
		ticker := time.NewTicker(time.Second)
		for {
			select {
			case <-ticker.C:
				count++
				t.Log("ticker: count: ", count)
			case brk := <-breakerChan:
				<-brk.Break()
				t.Log("Resume()")
			}
		}
	}()

	go func() {
		defer breaker.Resume()
		time.Sleep(time.Second * 5)
		breakerChan <- breaker
		t.Log("Break()")
		t.Log("do something")
		time.Sleep(time.Second * 5)
		<-breaker.Continue()
		t.Log("Continue()")
		return
	}()

	select {}
}
