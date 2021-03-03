package sort

// 冒泡排序
// 从头往后挨个比较，交换元素较小的使其在前，往后遍历。
func BubbleSort(s []int) {

	for i := len(s); i > 0; i-- {
		for j := 1; j < i; j++ {
			if s[j] < s[j-1] {
				s[j], s[j-1] = s[j-1], s[j]
			}
		}
	}
}
