package rpc

import (
	"backstage/common/payload"
	"backstage/global"
	"backstage/global/log"
	breaker2 "backstage/lib/breaker"
	"context"
	"time"
)

type Sync struct {
}

var count = 1

func (p *Sync) Break(ctx context.Context, req *payload.FakeReq, rsp *payload.FakeRsp) error {
	breaker, err := breaker2.NewBreaker()
	if err != nil {
		return err
	}

	global.Breaker() <- breaker
	//defer breaker.Resume()
	defer func() {
		if r := recover(); r != nil {
			breaker.Resume()
			log.Error("recover break, err = ", r)
			return
		}
		breaker.Resume()
		return
	}()
	<-breaker.Continue()

	if count%5 == 0 {
		count++
		panic("any panic")
	}
	count++

	log.Debug("Do something")
	time.Sleep(time.Second * 2)
	return nil
}
