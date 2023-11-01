package breaker

import "backstage/abstract/breaker"

const (
	Name = "Breaker"
)

type _breaker struct {
	breakChan    chan interface{}
	continueChan chan interface{}
	opts         *Options
}

func NewBreaker(opts ...Option) (breaker.Breaker, error) {
	options := Options{}

	for _, o := range opts {
		o(&options)
	}

	return &_breaker{
		opts:         &options,
		breakChan:    make(chan interface{}, 1),
		continueChan: make(chan interface{}, 1),
	}, nil
}

func (p *_breaker) Name() string {
	return Name
}

func (p *_breaker) Break() <-chan interface{} {
	p.continueChan <- ""
	return p.breakChan
}

func (p *_breaker) Continue() <-chan interface{} {
	return p.continueChan
}

func (p *_breaker) Resume() {
	p.breakChan <- ""
}
