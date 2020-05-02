package stack

/**
使用切片实现栈
*/
type ArrayStack []Element

func NewArrayStack() *ArrayStack {
	return &ArrayStack{}
}

func (s *ArrayStack) Push(e ...Element) {
	*s = append(*s, e...)
}

func (s *ArrayStack) Pop() Element {
	if s.IsEmpty() {
		return nil
	}
	val := (*s)[s.Size()-1]
	*s = (*s)[:s.Size()-1]
	return val
}

func (s *ArrayStack) Peek() Element {
	if s.IsEmpty() {
		return nil
	}
	return (*s)[s.Size()-1]
}

func (s *ArrayStack) Clear() {
	*s = ArrayStack{}
}

func (s *ArrayStack) IsEmpty() bool {
	return len(*s) == 0
}

func (s *ArrayStack) Size() int {
	return len(*s)
}
