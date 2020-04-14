package stack

/**
使用链表实现栈
*/

type node struct {
	Value Element
	Next  *node
}

type LinkedStack struct {
	Head *node
	size int
}

func NewLinkedStack() *LinkedStack {
	return &LinkedStack{
		Head: nil,
		size:0,
	}
}

func (s *LinkedStack)Push(e...Element)  {
	for _,v := range e{
		node := node{
			Value: v,
			Next: s.Head,
		}
		s.Head = &node
		s.size++
	}
}

func (s *LinkedStack)Pop() Element {
	if s.IsEmpty(){
		return nil
	}
	top := s.Head
	s.Head = top.Next
	s.size--
	return top.Value
}

func (s *LinkedStack)Peek() Element {
	if s.IsEmpty(){
		return nil
	}
	return s.Head.Value
}

func (s *LinkedStack)Clear(){
	*s = LinkedStack{
		Head:&node{},
		size:0,
	}
}

func (s *LinkedStack)IsEmpty() bool {
	return s.size == 0
}

func (s *LinkedStack)Size() int {
	return s.size
}