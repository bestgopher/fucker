package internal

/*
堆排序
	可以认为是选择排序的一种优化
步骤:
	1.对序列进行原地建堆(heapify)，构建出大顶堆
	2.重复执行以下操作，直到堆的元素数量为1
		1.交换堆顶元素和尾元素，此时序列的最大元素就在尾部了
		2.把尾部元素排除再外，也就是堆的元素减1
		3.堆0位置进行1次siftDown操作。
时间复杂度:
	步骤2的执行n-1次，2.3执行logn，因此时间复杂度是O(nlogn)
*/
func HeapSort(s []int) {
	// 构建大顶堆
	heapify(s)
	heapLen := len(s) - 1

	for ; heapLen >= 1; heapLen-- {
		s[0], s[heapLen] = s[heapLen], s[0]
		downHeap(s[:heapLen], 0)
	}

}

// heapify函数是用与构建一个堆
func heapify(s []int) {
	parent := (len(s) - 1) / 2
	for i := parent; i >= 0; i-- {
		downHeap(s, i)
	}
}

// downHeap是自底向上使得元素冒泡
func downHeap(s []int, index int) {
	if index*2+1 < len(s) {
		max := index*2 + 1
		if index*2+2 < len(s) && s[index*2+1] < s[index*2+2] {
			max = index*2 + 2
		}
		if s[max] > s[index] {
			s[max], s[index] = s[index], s[max]
			downHeap(s, max)
		}
	}
}
