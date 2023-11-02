package redis

import (
	"backstage/abstract/locker"
	"backstage/common/cache/string/lock"
	"backstage/common/conf"
	"backstage/global/cache/redis"
	"encoding/json"
	"fmt"
	"github.com/BurntSushi/toml"
	"testing"
	"time"
)

var redis_conf = `
[Redis.test1]
	Name = "Redis Server"
	Servers = ["192.168.130.128:16381"]
	Password = "123456"

[Redis.test2]
	Name = "Redis Server"
	Servers = ["192.168.130.128:16382"]
	Password = "123456"

[Redis.test3]
	Name = "Redis Server"
	Servers = ["192.168.130.128:16383"]
	Password = "123456"
`

func getLocker(config, which string, db int64) locker.MLock {
	cf := &conf.CacheConf{}
	if err := toml.Unmarshal([]byte(config), &cf); err != nil {
		panic(err)
	}

	client, err := redis.GetClient(cf, which, db)
	if err != nil {
		panic(err)
	}

	l, err := NewLocker(
		WithClient(client),
	)
	if err != nil {
		panic(err)
	}
	return l
}

func TestRedis(t *testing.T) {
	db := int64(0)
	which := "test1"
	cf := &conf.CacheConf{}
	if err := toml.Unmarshal([]byte(redis_conf), &cf); err != nil {
		t.Fatal(err)
	}

	client, err := redis.GetClient(cf, which, db)
	if err != nil {
		t.Fatal(err)
	}

	_, err = NewLocker(
		WithClient(client),
	)
	if err != nil {
		t.Fatal(err)
	}
}

func TestLock(t *testing.T) {
	id := 1
	which := template.GetWhich(id)
	db := template.GetDB()
	lll := getLocker(redis_conf, which, db)
	locCtx, err := lll.Lock(template.GetKey(id), template.GetHoldingTime(), template.GetTimeout(), template.GetRetry())
	if err != nil {
		t.Fatal(err)
	}
	temp, err := json.Marshal(locCtx)
	if err != nil {
		t.Fatal(err)
	}
	t.Log("Context: ", string(temp))
}

func TestUnlock(t *testing.T) {
	id := 1
	which := template.GetWhich(id)
	db := template.GetDB()
	lll := getLocker(redis_conf, which, db)
	ctx, err := lll.Lock(template.GetKey(id), template.GetHoldingTime(), template.GetTimeout(), template.GetRetry())
	if err != nil {
		t.Fatal(err)
	}
	temp, err := json.Marshal(ctx)
	if err != nil {
		t.Fatal(err)
	}
	t.Log("Context: ", string(temp))

	lll.Unlock(ctx)
}

func TestRefresh(t *testing.T) {
	addition := time.Second * 3
	nbOfThread := 10
	id := 1
	which := template.GetWhich(id)
	db := template.GetDB()
	lll := getLocker(redis_conf, which, db)
	ctx, err := lll.Lock(template.GetKey(id), template.GetHoldingTime(), template.GetTimeout(), template.GetRetry())
	if err != nil {
		t.Fatal(err)
	}
	temp, err := json.Marshal(ctx)
	if err != nil {
		t.Fatal(err)
	}
	t.Log("Context: ", string(temp))

	for i := 0; i < nbOfThread; i++ {
		go func(index int) {
			for {
				ctx, err := lll.Lock(template.GetKey(id), template.GetHoldingTime(), template.GetTimeout(), template.GetRetry())
				if err != nil {
					//fmt.Println("index = ", index, ", acquire lock fail")
					continue
				} else {
					fmt.Println("index = ", index, ", acquire lock")
				}
				err = lll.Refresh(ctx, addition)
				if err != nil {
					panic(err)
				}
				fmt.Println("index = ", index, ", do something")
				time.Sleep(time.Second * 2)
				fmt.Println("index = ", index, ", release lock")
				err = lll.Unlock(ctx)
				if err != nil {
					fmt.Println("index = ", index, ", release lock fail, err = ", err)
					continue
				} else {
					time.Sleep(time.Millisecond * 300)
					//fmt.Println("index = ", index, ", after release lock")
				}
			}
		}(i)
	}

	select {}
}
