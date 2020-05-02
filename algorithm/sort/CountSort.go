package sort

func CountSort(arr *[]int) {
	max, _ := getMaxMinInArr(*arr)
	temp := make([]int, max+1)
	for _, v := range *arr {
		temp[v]++
	}
	index := 0
	for i, v := range temp {
		for ; v > 0; v-- {
			(*arr)[index] = i
			index++
		}
	}
}
