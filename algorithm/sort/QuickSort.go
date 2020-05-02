package sort

func QuickSort(arr *[]interface{}) {

}

func QuickSort1(list []int, left, right int) {
	if right < left {
		return
	}
	flag := list[left]
	start := left
	end := right
	for {
		if start == end {
			break
		}
		for list[end] >= flag && end > start {
			end--
		}
		for list[start] <= flag && end > start {
			start++
		}
		if end > start {
			SwapGo(list, start, end)
		}
	}
	SwapGo(list, left, start)
	QuickSort1(list, left, start-1)
	QuickSort1(list, start+1, right)
}

func SwapGo(list []int, i, j int) {
	list[i], list[j] = list[j], list[i]
}
