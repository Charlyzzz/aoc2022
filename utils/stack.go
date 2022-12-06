package utils

type Stack[T any] struct {
	elems []T
}

func (s *Stack[T]) Append(elem T) {
	s.elems = append(s.elems, elem)
}

func (s *Stack[T]) Pop() T {
	elem := s.elems[0]
	s.elems = s.elems[1:]
	return elem
}

func (s *Stack[T]) PopN(n int) []T {
	elems := s.elems[:n]
	s.elems = s.elems[n:]
	return elems
}

func (s *Stack[T]) Peek() T {
	return s.elems[0]
}

func (s *Stack[T]) Push(elem T) {
	s.elems = append([]T{elem}, s.elems...)
}
