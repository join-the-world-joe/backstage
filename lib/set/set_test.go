package set

import "testing"

func TestSet(t *testing.T) {
	m1 := "4_3"
	m2 := "5_3"
	m3 := "6_3"
	s := NewSet()
	s.SAdd(m1)
	s.SAdd(m2)
	s.SAdd(m3)
	ok := s.SAdd(m1)
	if !ok {
		t.Log("!ok")
	}
	t.Log(s.SMembers())
	t.Log(s.SCard())
	t.Log(s.SisMember(m1))
	s.SRem(m1)
	t.Log(s.SCard())
	t.Log(s.SisMember(m1))
	t.Log(s.SMembers())
}
