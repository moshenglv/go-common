package sort

func BubbleSort(arr *[]int) {
	if len(*arr) <= 1 {
		return
	}
	flag := false
	for i, _ := range *arr {
		for j := 0; j < (len(*arr) - i - 1); j++ {
			if (*arr)[j] > (*arr)[j+1] {
				Swap(&(*arr)[j], &(*arr)[j+1])
				flag = true
			}
		}
		if !flag {
			break
		}
	}
}
