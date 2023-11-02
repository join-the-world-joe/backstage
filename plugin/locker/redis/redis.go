package redis

import (
	"backstage/abstract/locker"
	"context"
	"errors"
	"github.com/go-redis/redis/v8"
	"reflect"
	"time"
)

const (
	Name           = "Redis"
	DefaultTimeout = 10 * time.Second
)

type _locker struct {
	opts                 *Options
	acquireLockScriptSha string
	refreshLockScript    string
	releaseLockScript    string
}

func NewLocker(opts ...Option) (locker.MLock, error) {
	var client *redis.Client
	var acquireLockScript, refreshLockScript, releaseLockScript string

	options := Options{timeout: DefaultTimeout}

	for _, o := range opts {
		o(&options)
	}

	if client = options.client; client == nil {
		return nil, errors.New("options.client == nil")
	}

	ctx, cancel := context.WithTimeout(context.Background(), options.timeout)
	defer cancel()

	if err := client.ClientID(ctx).Err(); err != nil {
		return nil, err
	}

	if sha, err := client.ScriptLoad(context.Background(), AcquireLockScript).Result(); err != nil {
		return nil, err
	} else {
		acquireLockScript = sha
	}

	if sha, err := client.ScriptLoad(context.Background(), RefreshLockScript).Result(); err != nil {
		return nil, err
	} else {
		refreshLockScript = sha
	}

	if sha, err := client.ScriptLoad(context.Background(), ReleaseLockScript).Result(); err != nil {
		return nil, err
	} else {
		releaseLockScript = sha
	}

	return &_locker{
		opts:                 &options,
		acquireLockScriptSha: acquireLockScript,
		refreshLockScript:    refreshLockScript,
		releaseLockScript:    releaseLockScript,
	}, nil
}

func (p *_locker) TryLock(id string, sec time.Duration) (*locker.Context, error) {
	return nil, nil
}

func (p *_locker) Lock(id string, holdingTime time.Duration, lockTimeout time.Duration, retry time.Duration) (*locker.Context, error) {
	for end := time.Now().Add(lockTimeout); time.Now().Before(end); {
		ctx := genContext(id, holdingTime)
		ret, err := p.opts.client.EvalSha(context.Background(), p.acquireLockScriptSha, []string{id}, ctx.Signature).Result()
		if err != nil {
			return nil, err
		}

		s := reflect.ValueOf(ret)
		if s.Kind() != reflect.Slice {
			return nil, errors.New("unknown error 1")
		}

		v1, ok1 := s.Index(0).Interface().(string)
		if !ok1 {
			return nil, errors.New("unknown error 3")
		}

		if v1 == "1" {
			return ctx, nil
		} else {
			v2, ok2 := s.Index(1).Interface().(string)
			if !ok2 {
				return nil, errors.New("unknown error 4")
			}
			_ = v2
			//fmt.Println("lock fail, msg = ", v2)
			time.Sleep(retry)
			continue
		}
	}
	return nil, errors.New("timeout")
}

func (p *_locker) Refresh(ctx *locker.Context, addition time.Duration) error {
	updateTimeInfo(ctx, addition)
	ret, err := p.opts.client.EvalSha(context.Background(), p.refreshLockScript, []string{ctx.Id}, ctx.Signature).Result()
	if err != nil {
		return err
	}

	s := reflect.ValueOf(ret)

	v1, ok1 := s.Index(0).Interface().(string)
	if !ok1 {
		return errors.New("unknown error3")
	}

	if v1 == "1" {
		return nil
	} else {
		v2, ok2 := s.Index(1).Interface().(string)
		if !ok2 {
			return errors.New("unknown error 4")
		}
		return errors.New(v2)
	}
}

func (p *_locker) Unlock(ctx *locker.Context) error {
	if ctx == nil {
		return errors.New("ctx == nil")
	}

	ret, err := p.opts.client.EvalSha(context.Background(), p.releaseLockScript, []string{ctx.Id}, ctx.Signature).Result()
	if err != nil {
		return err
	}

	s := reflect.ValueOf(ret)

	v1, ok1 := s.Index(0).Interface().(string)
	if !ok1 {
		return errors.New("unknown error3")
	}

	if v1 == "1" {
		return nil
	} else {
		v2, ok2 := s.Index(1).Interface().(string)
		if !ok2 {
			return errors.New("unknown error 4")
		}
		return errors.New(v2)
	}
}
