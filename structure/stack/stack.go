package stack

type Element interface{}

type Interface interface {
	Push(e ...Element)

	Pop() Element

	Peek() Element

	IsEmpty() bool

	Clear()

	Size() int
}
