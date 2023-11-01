package selector

type Strategy int8

const (
	RoundRobin Strategy = iota
	Random
	Weight
)

type Selector interface {
	Name() string
	Set(string, map[string]int) // map[choice]weight
	Select(string, Strategy) (Next, error)
	Destroy()
}

type Next func() (string, error)
