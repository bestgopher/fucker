package internal

// 快速排序
//		基于分治法，但是是把所有的复杂操作在递归调用之前做完。
// 		主要思想是应用分治法把序列s分解为子序列，递归地排序每个子序列，然后通过简单串联的方式合并这些已排序的子序列。
// 步骤：
//	1.分解：
//		如果s有至少2个元素(如果s只有1个或者0个元素，什么都不用做)，从s中选择一个特定元素x，称之为基准值。
//		一般情况下，如果s中最后一个元素作为基准值x。从s中移除所有的元素，并把它们放在3个序列中：
//			L：存储s中小于x的值
//			E：存储s中等于x的值
//			G：存储s中大于x的值
//		如果s中的元素是互异的，那么E将只含有一个元素 -- 基准值自己。
//	2.解决子问题：
//		递归地排序序列L和G
//	3.合并
//		把s中的元素按照先插入L中的元素、然后插入E中的元素、最后插入G中的元素的顺序放回。
func QuickSort(s []int) {
	n := len(s)
	if n < 2 {
		return
	}

	p := s[0]
	var (
		L []int
		E []int
		G []int
	)

	for _, v := range s {
		switch {
		case v < p:
			L = append(L, v)
		case v == p:
			E = append(E, v)
		case v > p:
			G = append(G, v)
		}
	}

	QuickSort(L)
	QuickSort(G)

	copy(s[:len(L)], L)
	copy(s[len(L):len(L)+len(E)], E)
	copy(s[len(L)+len(E):len(L)+len(E)+len(G)], G)
}
