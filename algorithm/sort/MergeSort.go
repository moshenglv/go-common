package sort

func MergeSort(arr *[]int) {
	if len(*arr) <= 1 {
		return
	}
	sort(arr, 0, len(*arr)-1)
}

func sort(arr *[]int, start, end int) {
	if start >= end {
		return
	}
	mid := (start + end) / 2
	sort(arr, start, mid)       //排左边
	sort(arr, mid+1, end)       //排右边
	merge(arr, start, end, mid) //合并
}

func merge(arr *[]int, start, end, mid int) {
	len := end - start + 1         //左右两子数组长度之和
	lIndex, rIndex := start, mid+1 //左右两部分的开始下标
	result := make([]int, 0, len)
	for k := 0; k < len; k++ {
		if lIndex > mid { //left已经遍历完
			result = append(result, (*arr)[rIndex])
			rIndex++
		} else if rIndex > end { //right已遍历完
			result = append(result, (*arr)[lIndex])
			lIndex++
		} else if (*arr)[lIndex] <= (*arr)[rIndex] {
			result = append(result, (*arr)[lIndex])
			lIndex++
		} else {
			result = append(result, (*arr)[rIndex])
			rIndex++
		}
	}
	//拷贝回原数组
	for i, v := range result {
		(*arr)[i+start] = v
	}
}
