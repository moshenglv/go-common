package test

import (
	"fmt"
	"go-common/structure/queue"
)

func QueueTest() {
	q := queue.New()
	fmt.Println("length:", q.Size())
	q.Push("a")
	q.Push(1)
	fmt.Println("length:", q.Size())
	fmt.Println(*q)
	tmp := q.Pop()
	fmt.Println("出队：", *tmp)
	tmp = q.Pop()
	fmt.Println("出队：", *tmp)
	tmp = q.Pop()
	fmt.Println("出队：", tmp)
}
