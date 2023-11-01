package breaker

type Breaker interface {
	Name() string
	Break() <-chan interface{}
	Continue() <-chan interface{}
	Resume()
}