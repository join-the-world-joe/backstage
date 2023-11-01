package selector

import (
	"backstage/abstract/selector"
	"math/rand"
	"sync"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func Random(values map[string]int) selector.Next {

	elements := []string{}
	for k, _ := range values {
		elements = append(elements, k)
	}

	return func() (string, error) {
		i := rand.Int() % len(elements)
		return elements[i], nil
	}
}

func RoundRobin(values map[string]int) selector.Next {
	var i = 0
	var mtx sync.Mutex

	elements := []string{}
	for k, _ := range values {
		elements = append(elements, k)
	}

	return func() (string, error) {
		mtx.Lock()
		v := elements[i%len(elements)]
		i++
		mtx.Unlock()

		return v, nil
	}
}

type mark struct {
	Key string
	Min int
	Max int
}

func Weight(values map[string]int) selector.Next {
	i := 1
	count := 0
	items := []*mark{}

	for k, v := range values {
		count += v
		item := &mark{Key: k}
		item.Min = i
		i += v
		item.Max = i
		items = append(items, item)
	}

	return func() (string, error) {
		out := ""
		ind := rand.Intn(count) + 1
		for _, v := range items {
			if ind >= v.Min && ind < v.Max {
				out = v.Key
				break
			}
		}
		return out, nil
	}
}
