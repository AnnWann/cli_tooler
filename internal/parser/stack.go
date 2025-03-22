package parser

type stack struct {
	elements []BinaryTreeNode
}

func (s *stack) push(element BinaryTreeNode) {
	s.elements = append(s.elements, element)
}

func (s *stack) pop() BinaryTreeNode {
	if len(s.elements) == 0 {
		return BinaryTreeNode{}
	}
	element := s.elements[len(s.elements)-1]
	s.elements = s.elements[:len(s.elements)-1]
	return element
}

func (s *stack) peek() BinaryTreeNode {
	if len(s.elements) == 0 {
		return BinaryTreeNode{}
	}
	return s.elements[len(s.elements)-1]
}

func (s *stack) isEmpty() bool {
	return len(s.elements) == 0
}
