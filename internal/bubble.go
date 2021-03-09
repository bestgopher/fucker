package internal

/*
冒泡排序

基本步骤:
	从头开始比较序列的每一对相邻元素，如果第1个比第2个大，就交换它们的位置。执行完一轮后，最末尾那个元素就是最大的元素。如 BubbleSort1
不足:
	但是此代码有很大优化的地方，比如当传入的序列本身就是有序的或者执行完一些步骤后排序就完成了，此时没有必要进行后面无效的排序步骤
	例如：1, 2, 3, 4, 5, 6

优化1:
	加个标志位，标示着没有进入交换两个数的逻辑，这时说明序列有序，如 BubbleSort2
优化1缺点:
	对于大多数情况优化1的执行情况还要比未优化之前差，因为优化1出现的时机本来就比较少，而对于常见情况，它却多执行一些代码。
优化2:
	但是有种情况是常见的，就是一部分有序。在冒泡中，如果序列的尾部子序列已有序，我们就不需要冒泡到最后了。
	例如: 7, 5, 2, 3, 4, 10, 11, 12
	在这个序列中，第一个数7冒泡到10前面就不会向上冒泡。并且10， 11， 12也已经有序了。因此第二次冒泡的时候，只需要冒泡到7的位置即可
最终代码 BubbleSort
*/
func BubbleSort(s []int) {
	for i := len(s) - 1; i > 0; i-- {
		// sortedIndex的初始值在数组完全有序的时候也要有用
		sortedIndex := 0
		for j := 1; j <= i; j++ {
			if s[j] < s[j-1] {
				s[j], s[j-1] = s[j-1], s[j]
				sortedIndex = j
			}
		}

		i = sortedIndex
	}
}

// Deprecated
func BubbleSort1(s []int) {
	for i := len(s); i > 0; i-- {
		for j := 1; j < i; j++ {
			if s[j] < s[j-1] {
				s[j], s[j-1] = s[j-1], s[j]
			}
		}
	}
}

// Deprecated
func BubbleSort2(s []int) {
	for i := len(s); i > 0; i-- {
		flag := true
		for j := 1; j < i; j++ {
			if s[j] < s[j-1] {
				s[j], s[j-1] = s[j-1], s[j]
				flag = false
			}
		}

		if flag {
			break
		}
	}
}
