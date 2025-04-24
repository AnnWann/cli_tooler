package utility

type Stack[T any] struct {
	elements []T
}

func (s *Stack[T]) Push(element T) {
	s.elements = append(s.elements, element)
}

func (s *Stack[T]) Pop() T {
	if len(s.elements) == 0 {
		var zeroValue T
		return zeroValue
	}
	element := s.elements[len(s.elements)-1]
	s.elements = s.elements[:len(s.elements)-1]
	return element
}

func (s *Stack[T]) Peek() T {
	if len(s.elements) == 0 {
		var zeroValue T
		return zeroValue
	}
	return s.elements[len(s.elements)-1]
}

func (s *Stack[T]) IsEmpty() bool {
	return len(s.elements) == 0
}
