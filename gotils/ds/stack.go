package ds

type Stack[T any] []T

func (s Stack[T]) Push(v T) Stack[T] {
	return append(s, v)
}

func (s Stack[T]) Pop() (Stack[T], T, bool) {
	if len(s) < 1 {
        var zv T
		return s, zv, true
	}

    res := s[len(s)-1]
	s = s[:len(s)-1]

	return s, res, false
}

func (s Stack[T]) Len() int {
	if s == nil {
		return 0
	}

	return len(s)
}
