package selector

import (
	"testing"
	"time"
)

func TestRandom(t *testing.T) {
	begin := 1
	end := 10
	sel := NewSelector()

	sel.Set("gateway", map[string]int{"node_1": 0, "node_2": 0, "node_3": 0, "node_4": 0, "node_5": 0})

	next, err := sel.Select("gateway", WithRandom())
	if err != nil {
		t.Error(err)
		return
	}

	for count := begin; count <= end; count++ {
		node, err := next()
		if err != nil {
			t.Error(err)
			return
		}
		t.Log(node)
		time.Sleep(time.Millisecond * 200)
	}
}

func TestRoundRobin(t *testing.T) {
	begin := 1
	end := 10

	sel := NewSelector()

	sel.Set("gateway", map[string]int{"node_1": 0, "node_2": 0, "node_3": 0, "node_4": 0, "node_5": 0})

	next, err := sel.Select("gateway", WithRoundRobin())
	if err != nil {
		t.Error(err)
		return
	}

	for count := begin; count <= end; count++ {
		node, err := next()
		if err != nil {
			t.Error(err)
			return
		}
		t.Log(node)
		time.Sleep(time.Millisecond * 200)
	}
}

func TestWeight(t *testing.T) {
	begin := 1
	end := 10

	sel := NewSelector()

	sel.Set("gateway", map[string]int{"node_1": 1, "node_2": 0, "node_3": 0, "node_4": 0, "node_5": 1})

	next, err := sel.Select("gateway", WithWeight())
	if err != nil {
		t.Error(err)
		return
	}

	for count := begin; count <= end; count++ {
		node, err := next()
		if err != nil {
			t.Error(err)
			return
		}
		t.Log(node)
		time.Sleep(time.Millisecond * 200)
	}
}
