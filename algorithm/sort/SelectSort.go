package sort

func SelectSort(arr *[]int) {
	if len(*arr) <= 1 {
		return
	}
	for i := 0; i < len(*arr); i++ {
		value := (*arr)[i]
		index := i
		for j := i + 1; j < len(*arr); j++ {
			if (*arr)[j] < value {
				value = (*arr)[j]
				index = j
			}
		}
		(*arr)[index], (*arr)[i] = (*arr)[i], value
	}
}
