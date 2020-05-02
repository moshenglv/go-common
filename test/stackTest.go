package test

import (
	"fmt"
	"go-common/structure/stack"
)

func Hello() {
	fmt.Println("hello")
}

func StackTest() {
	//测试栈
	//stack := stack.NewLinkedStack()
	stack := stack.NewArrayStack()
	stack.Push(3, 2, 5, 9)
	fmt.Println("入栈顺序:3 2 5 9")
	fmt.Println("size:", stack.Size())
	fmt.Println("pop:", stack.Pop())
	fmt.Println("size:", stack.Size())
	fmt.Println("peek:", stack.Peek())
	fmt.Println("size:", stack.Size())
	fmt.Println("pop:", stack.Pop())
	fmt.Println("size:", stack.Size())
	fmt.Println("pop:", stack.Pop())
	fmt.Println("size:", stack.Size())

	fmt.Println("是否为空：", stack.IsEmpty())
	fmt.Println("清空")
	stack.Clear()
	fmt.Println("是否为空：", stack.IsEmpty())
	fmt.Println("size:", stack.Size())
	fmt.Println("pop:", stack.Pop())
	fmt.Println("peek:", stack.Peek())
}
