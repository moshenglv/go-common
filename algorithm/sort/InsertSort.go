package sort

func InsertSort(arr *[]int) {
	if len(*arr) <= 1 {
		return
	}

	for i := 1; i < len(*arr); i++ {
		value := (*arr)[i]
		j := i - 1
		for ; j >= 0; j-- {
			if (*arr)[j] > value {
				(*arr)[j+1] = (*arr)[j]
			} else {
				break
			}
		}
		(*arr)[j+1] = value
	}
}
