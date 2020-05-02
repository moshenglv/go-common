package test

import (
	"fmt"
	"go-common/structure/heap"
)

func HeapTest() {
	//arr := []int{8, 7, 6, 5, 4, 3, 2}

	//heap := heap2.NewHeap(arr, 15)
	//fmt.Println(heap)
	//heap.Insert(9)
	//fmt.Println(heap)
	//heap.RemoveMax()
	//fmt.Println(heap)

	//构建最大堆法1：自底向上堆化
	arr1 := []int{2, 3, 5, 6, 8, 7, 4}
	heap1 := heap.NewHeap(arr1, 15)
	fmt.Println(*heap1)
	heap1.BuildMapHeap()
	fmt.Println(*heap1)

	//构建最大堆法1：自底向上堆化
	arr2 := []int{2, 3, 5, 6, 8, 7, 4}
	heap2 := heap.NewEmptyHeap(15)
	for _, v := range arr2 {
		heap2.Insert(v)
	}
	fmt.Println(*heap2)

	//排序
	fmt.Println("排序")
	arr3 := []int{2, 3, 5, 6, 8, 7, 4}
	heap3 := heap.NewHeap(arr3, 15)
	heap3.Sort()
	fmt.Println(*heap3)

}
