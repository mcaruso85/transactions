package tx

type stack []interface{}

func (s *stack) len() int {
	return len(*s)
}

func (s *stack) isEmpty() bool {
	return len(*s) == 0
}

func (s *stack) push(any interface{}) {
	*s = append(*s, any)
}

func (s *stack) pop() (interface{}, bool) {
	if s.isEmpty() {
		return nil, false
	} else {
		index := len(*s) - 1
		element := (*s)[index]
		*s = (*s)[:index]
		return element, true
	}
}
