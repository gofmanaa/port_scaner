package data

type Set struct {
	set map[interface{}]struct{}
}

func NewSet() *Set {
	return &Set{set: make(map[interface{}]struct{})}
}

func (s *Set) Add(data interface{}) {
	s.set[data] = struct{}{}
}

func (s *Set) Delete(data interface{}) {
	delete(s.set, data)
}

func (s * Set) GetInt() ([]int) {
	res := []int{}
	for  key := range s.set {
		res = append(res, key.(int))
	}

	return res
}
