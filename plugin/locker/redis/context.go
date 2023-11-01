package redis

import (
	"fmt"
	"github.com/google/uuid"
	"go-micro-framework/abstract/locker"
	"time"
)

func genContext(id string, sec time.Duration) *locker.Context {
	var ctx = new(locker.Context)
	ctx.Id = id
	ctx.UUID = uuid.New().String()
	ctx.Prefix = ctx.UUID + ":"
	ctx.FromInMS = time.Now().UnixNano() / 1e6
	ctx.From = fmt.Sprintf("%v", ctx.FromInMS)
	ctx.ToInMS = ctx.FromInMS + sec.Milliseconds()
	ctx.To = fmt.Sprintf("%v", ctx.ToInMS)
	ctx.Signature = ctx.Prefix + ctx.From + ":" + ctx.To
	return ctx
}

func updateTimeInfo(ctx *locker.Context, sec time.Duration) {
	ctx.ExtendToInMS = (time.Now().UnixNano() / 1e6) + sec.Milliseconds()
	ctx.ExtendTo = fmt.Sprintf("%v", ctx.ExtendToInMS)
	ctx.Signature = ctx.Prefix + ctx.From + ":" + ctx.ExtendTo
}
