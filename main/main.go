package main

import (
	"fmt"
	"go-common/test"
)

func main() {
	arr := []int{3, 5, 3, 9, 10, 23, 4, 6, 123, 22, 33}
	//sort.HeapSort(arr)
	fmt.Println(arr)
	test.HeapTest()

}
