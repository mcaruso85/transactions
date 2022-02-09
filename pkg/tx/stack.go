package tx

type stack []interface{}

func (s *stack) IsEmpty() bool {
	return len(*s) == 0
}

func (s *stack) Push(any interface{}) {
	*s = append(*s, any)
}

func (s *stack) Pop() (interface{}, bool) {
	if s.IsEmpty() {
		return nil, false
	} else {
		index := len(*s) - 1
		element := (*s)[index]
		*s = (*s)[:index]
		return element, true
	}
}
