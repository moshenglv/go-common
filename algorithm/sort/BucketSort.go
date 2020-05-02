package sort

/**
size:桶中数据个数
*/
func BucketSort(arr *[]int, size int) {
	//获取数组最大最小值
	max, min := getMaxMinInArr(*arr)
	//桶数
	count := (max-min)/size + 1
	//二维切片
	buckets := make([][]int, count)

	//数据入桶
	for _, v := range *arr {
		index := (v - min) / size //放入桶的下标
		buckets[index] = append(buckets[index], v)
	}

	//桶内排序,并写回原数组
	tmpPos := 0
	for _, bucket := range buckets {
		sortInBucket(&bucket)
		copy((*arr)[tmpPos:], bucket)
		tmpPos += len(bucket)
	}

}

/*
桶内排序
*/
func sortInBucket(bucket *[]int) { //此处实现插入排序方式，其实可以用任意其他排序方式
	InsertSort(bucket)
}

/*
获取数组最大最小值
*/
func getMaxMinInArr(arr []int) (int, int) {
	max := arr[0]
	min := arr[0]
	for i := 1; i < len(arr); i++ {
		if arr[i] > max {
			max = arr[i]
		}
		if arr[i] < min {
			min = arr[i]
		}
	}
	return max, min
}
