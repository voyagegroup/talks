package stack

type Stack struct {
	s []interface{}
}

func (s *Stack) Pop() interface{} {
	return s.s[len(s.s)-1]
}

func (s *Stack) Push(i interface{}) {
	s.s = append(s.s, i)
}
