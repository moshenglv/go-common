package heap

type Heap struct {
	Arr      []int
	Count    int
	capacity int
}

func NewHeap(arr []int, cap int) *Heap {
	return &Heap{
		Arr:      arr,
		Count:    len(arr),
		capacity: cap,
	}
}

func NewEmptyHeap(cap int) *Heap {
	return &Heap{
		Arr:      make([]int, 0, cap),
		Count:    0,
		capacity: cap,
	}
}

func (heap *Heap) Insert(data int) bool {
	if heap.Count >= heap.capacity {
		return false
	}
	heap.Count++
	heap.Arr = append(heap.Arr, data)
	i := heap.Count - 1
	for (i-1)/2 >= 0 && heap.Arr[(i-1)/2] < heap.Arr[i] {
		heap.Arr[(i-1)/2], heap.Arr[i] = heap.Arr[i], heap.Arr[(i-1)/2]
		i = (i - 1) / 2
	}

	return true
}

func (heap *Heap) RemoveMax() bool {
	if heap.Count == 0 {
		return false
	}
	heap.Arr[0] = heap.Arr[heap.Count-1]
	heap.Count--
	heap.Arr = heap.Arr[:heap.Count]
	heapFromUp(heap.Arr, heap.Count, 0)
	return true
}

func (heap *Heap) BuildMapHeap() {
	for i := heap.Count / 2; i >= 0; i-- {
		heapFromUp(heap.Arr, heap.Count, i)
	}
}

/**
O(nlog(n))
1、取出堆顶元素（必定最大）
2、将最后的元素放到堆顶
3、重新从上至下建堆
4、重复以上过程直到堆中元素为0
*/
func (heap *Heap) Sort() {
	//首先建最大堆O(n)
	heap.BuildMapHeap()
	//排序O(logn)
	k := heap.Count - 1
	for k > 0 {
		heap.Arr[0], heap.Arr[k] = heap.Arr[k], heap.Arr[0]
		//从顶向下重新建堆,下标为k以及以后的不用堆化
		heapFromUp(heap.Arr, k, 0)
		k--
	}
}

/**
n:堆化的最大位置
i:开始堆化的下标
*/
func heapFromUp(a []int, n, i int) { // 自上往下堆化
	for {
		maxPos := i
		if (i*2+1) < n && a[i] < a[i*2+1] {
			maxPos = i*2 + 1
		}
		if (i*2+2) < n && a[maxPos] < a[i*2+2] {
			maxPos = i*2 + 2
		}
		if maxPos == i {
			break
		}
		a[i], a[maxPos] = a[maxPos], a[i]
		i = maxPos
	}
}
