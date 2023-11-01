package selector

import (
	"backstage/abstract/selector"
	"fmt"
	"sync"
)

const (
	Name = "Built-in Selector"
)

type _selector struct {
	opts *Options
	kv   sync.Map
}

func WithRandom() selector.Strategy {
	return selector.Random
}

func WithRoundRobin() selector.Strategy {
	return selector.RoundRobin
}

func WithWeight() selector.Strategy {
	return selector.Weight
}

func NewSelector(opts ...Option) selector.Selector {
	options := Options{}

	for _, o := range opts {
		o(&options)
	}

	return &_selector{
		opts: &options,
	}
}

func (p *_selector) Name() string {
	return Name
}

func (p *_selector) Set(key string, values map[string]int) {
	if values == nil || len(values) == 0 {
		if _, exist := p.kv.Load(key); exist {
			p.kv.Delete(key)
		}
		return
	}
	p.kv.Store(key, values)
}

func (p *_selector) Select(key string, strategy selector.Strategy) (selector.Next, error) {
	temp, exist := p.kv.Load(key)
	if exist {
		values, ok := temp.(map[string]int)
		if ok {
			if strategy == selector.RoundRobin {
				return RoundRobin(values), nil
			} else if strategy == selector.Random {
				return Random(values), nil
			} else if strategy == selector.Weight {
				return Weight(values), nil
			} else {
				return nil, fmt.Errorf("strategy[%v] is illegal", strategy)
			}
		}
		return nil, fmt.Errorf("value is not in type []string")
	}

	return nil, fmt.Errorf("no such key[%v]", key)
}

func (p *_selector) Destroy() {
	p.kv.Range(
		func(k, v interface{}) bool {
			p.kv.Delete(k)
			return true
		},
	)
}
