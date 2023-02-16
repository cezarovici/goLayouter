package stack

type Stack []any

func (s *Stack) Push(item any) {
	*s = append(*s, item)
}

func (s Stack) IsEmpty() bool {
	return len(s) == 0
}

func (s *Stack) Pop() {
	if !s.IsEmpty() {
		*s = (*s)[:len(*s)-1]
	}
}

func (s Stack) Peek() any {
	if !s.IsEmpty() {
		return s[len(s)-1]
	}

	return nil
}
