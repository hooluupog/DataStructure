package stack

type Stack struct {
	top  *Element
	size int
}
type Element struct {
	value interface{}
	next  *Element
}

func (s *Stack) Length() int {
	return s.size
}

func (s *Stack) Empty() bool {
	return s.size <= 0
}
func (s *Stack) Push(value interface{}) {
	//top point to new Element address and next pointer point to last top pointer postion
	s.top = &Element{value, s.top}
	s.size++
}
func (s *Stack) Pop() (value interface{}) {
	if s.size > 0 {
		value, s.top = s.top.value, s.top.next
		s.size--
		return
	}
	return nil
}

func (s *Stack) Top() (value interface{}) {
	value = s.top.value
	return
}
