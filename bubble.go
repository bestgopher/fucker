package sort

// 冒泡排序
// 从头开始比较每一对相邻元素，如果第1个比第2个大，就交换它们的位置。
// 执行完一轮后，最末尾那个元素就是最大的元素
// 优化：
// 		1.当s本来就是有序或者在执行了一部分冒泡后序列就有序了，可以提前终止冒泡排序
//		  但是这种情况发生的几率很小，并且多个三条指令，有时会增加运行时间
//		func BubbleSort(s []int) {
//			for i := len(s); i > 0; i-- {
//				isSorted := true // 增加一个标志位，标示序列是否已经有序了
//
//				for j := 1; j < i; j++ {
//					if s[j] < s[j-1] {
//						s[j], s[j-1] = s[j-1], s[j]
//						isSorted = false
//					}
//				}
//
//				if isSorted {
//					break
//				}
//			}
//		}
//		2.如果序列尾部已经局部有序，可以记录最后一次交换的位置，减少比较的次数
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
