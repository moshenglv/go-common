package sort

func ShellSort(arr *[]int) {
	if len(*arr) <= 1 {
		return
	}

	gap := len(*arr) / 2
	for gap > 0 {
		for i := gap; i < len(*arr); i++ {
			temp := (*arr)[i]
			j := i - gap
			for ; j >= 0 && (*arr)[j] > (*arr)[i]; j -= gap { //与插入排序一致，只是比较跨度为gap
				(*arr)[i] = (*arr)[j]
			}
			(*arr)[j+gap] = temp
		}
		gap = gap / 2
	}
}
