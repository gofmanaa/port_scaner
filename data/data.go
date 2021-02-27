package data

import "sync"

type Set struct {
	sync.Mutex
	set map[interface{}]struct{}
}

func NewSet() *Set {
	return &Set{set: make(map[interface{}]struct{})}
}

func (s *Set) Add(data interface{}) {
	s.Lock()
	defer s.Unlock()
	s.set[data] = struct{}{}
}

func (s *Set) Delete(data interface{}) {
	s.Lock()
	defer s.Unlock()
	delete(s.set, data)
}

func (s * Set) GetInt() ([]int) {
	res := []int{}
	for  key := range s.set {
		res = append(res, key.(int))
	}

	return res
}
