package sort

func RadixSort(arr *[]int) {
	max, _ := getMaxMinInArr(*arr)
	for bit := 1; max/bit > 0; bit *= 10 {
		n := len(*arr)
		bitData := make([][]int, 10)
		for i := 0; i < n; i++ {
			num := ((*arr)[i] / bit) % 10 //获取数字个位数...
			bitData[num] = append(bitData[num], (*arr)[i])
		}
		pos := 0
		for j := 0; j < 10; j++ { //遍历十个桶取出数据回填数组
			for k := 0; k < len(bitData[j]); k++ {
				(*arr)[pos] = bitData[j][k]
				pos++
			}
		}
	}
}
