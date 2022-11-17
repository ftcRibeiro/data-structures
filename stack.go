package datastructure

func NewStack[T any]() *Stack[T] {
	return &Stack[T]{
		Root: nil,
	}
}

func (s *Stack[T]) Push(val T) {
	s.Root = &Node[T]{
		Next: s.Root,
		Val:  val,
	}
}

func (s *Stack[T]) Pop() (item *T) {
	if s.Root == nil {
		return
	}
	item = &s.Root.Val
	s.Root = s.Root.Next
	return
}

func (s *Stack[T]) Empty() bool {
	return s.Root == nil
}
