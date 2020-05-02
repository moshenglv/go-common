package main

import (
	"fmt"
	"go-common/algorithm/sort"
)

func main() {
	arr := []int{3, 5, 3, 7, 9, 10, 23, 4, 6, 123, 22, 33}
	sort.RadixSort(&arr)
	fmt.Println(arr)
}
