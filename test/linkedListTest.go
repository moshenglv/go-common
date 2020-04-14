package test

import (
	"fmt"
	linkedList2 "go-common/structure/linkedList"
)

func LinkedListTest()  {
	list := linkedList2.NewLinkedStack()
	list.Append(3)
	list.Display()
	list.Append(5)
	list.Display()
	list.Insert(0,1)
	list.Display()
	list.Insert(1,2)
	list.Display()
	list.Insert(3,4)
	list.Display()
	list.Append(7)
	list.Display()

	list.Delete(0)
	list.Display()
	fmt.Println(list.Size())
	list.Delete(list.Size()-1)
	list.Display()

	list.Delete(2)
	list.Display()

	list.Insert(2,4)
	list.Display()
}
