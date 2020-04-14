package linkedList

import (
	"fmt"
	"sync"
)

type Element interface {}

type node struct {
	Value Element
	Next *node
	Prev *node
}

type LinkedList struct {
	mutex *sync.RWMutex
	Head *node
	Tail *node
	size int
}

func NewLinkedStack() *LinkedList {
	return &LinkedList{
		mutex:new(sync.RWMutex),
		Head: nil,
		Tail: nil,
		size:0,
	}
}

func (l *LinkedList)Get(index int) *node {
	if l.size ==0 || index > l.size - 1 {
		return  nil
	}
	if index == 0 {
		return l.Head
	}
	node := l.Head
	for i:= 1;i<=index;i++ {
		node = node.Next
	}
	return node
}

func (l *LinkedList)Append(e Element) bool {
	if e == nil {
		return false
	}

	node := &node{
		Value: e,
		Prev:nil,
		Next:nil,
	}

	l.mutex.Lock()
	defer l.mutex.Unlock()

	if l.size == 0 {
		l.Head = node
		l.Tail = node
		node.Prev = nil
		node.Next = nil
	}else{
		l.Tail.Next = node
		node.Prev = l.Tail
		node.Next = nil
		l.Tail = node
	}
	l.size++
	return true
}

func (l * LinkedList)Insert(index int,e Element) bool {
	if index > l.size || e == nil {
		return false
	}

	node := &node{
		Value: e,
		Prev:nil,
		Next:nil,
	}

	l.mutex.Lock()
	defer l.mutex.Unlock()
	if index == 0{
		node.Next = l.Head
		node.Prev = nil
		l.Head.Prev = node
		l.Head = node
	}else if index == l.size{
		node.Prev = l.Tail
		l.Tail.Next = node
		l.Tail = node
	}else {
		n := l.Get(index)
		node.Next = n
		node.Prev = n.Prev
		n.Prev.Next = node
		n.Prev = node
	}
	l.size ++
	return true
}

func (l *LinkedList)Delete(index int) bool {
	if index < 0 || index >= l.size {
		return false
	}

	l.mutex.Lock()
	defer l.mutex.Unlock()
	if index ==  0 {
		if l.size == 1{
			l.Head = nil
			l.Tail = nil
		}else{
			l.Head = l.Head.Next
			l.Head.Prev = nil
		}
	}else if index == l.size-1{
		l.Tail = l.Tail.Prev
		l.Tail.Next = nil
	}else{
		n := l.Get(index)
		n.Prev.Next = n.Next
		n.Next.Prev = n.Prev
	}

	l.size --
	return true
}

func (l *LinkedList)Display()  {
	if l == nil || l.size == 0{
		fmt.Println("list is empty")
	}
	if l.size == 1{
		fmt.Printf("[%v]\n", l.Head.Value)
		return
	}

	node := l.Head
	fmt.Print("[")
	for node.Next != nil {
		fmt.Printf("%v,", node.Value)
		node = node.Next
	}
	fmt.Printf("%v]\n", l.Tail.Value)
}

func (l *LinkedList)Size() int {
	return l.size
}