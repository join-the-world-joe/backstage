package set

import (
	"go-micro-framework/abstract/set"
	"sync"
)

const (
	Name = "Built-in Set"
)

type _set struct {
	sync.Mutex

	opts  *Options
	count int64
	m     map[string]interface{}
}

func NewSet(opts ...Option) set.Set {
	options := Options{}

	for _, o := range opts {
		o(&options)
	}

	return &_set{
		opts: &options,
		m:    make(map[string]interface{}),
	}
}

func (p *_set) Name() string {
	return Name
}

func (p *_set) SAdd(member string) bool {
	p.Lock()
	defer p.Unlock()
	if _, exist := p.m[member]; exist {
		return false
	}
	p.m[member] = nil
	p.count++
	return true
}

func (p *_set) SCard() int64 {
	return p.count
}

func (p *_set) SisMember(member string) bool {
	p.Lock()
	defer p.Unlock()
	_, exist := p.m[member]
	return exist
}

func (p *_set) SMembers() []string {
	p.Lock()
	defer p.Unlock()
	list := make([]string, 0, p.count)
	for member, _ := range p.m {
		list = append(list, member)
	}
	return list
}

func (p *_set) SRem(member string) {
	p.Lock()
	defer p.Unlock()
	if _, exist := p.m[member]; exist {
		delete(p.m, member)
		p.count--
	}
}
