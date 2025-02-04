package internal

import (
	fucker "gopkg.in/bestgopher/fucker.v1"
)

// 归并排序
// 分支法：
//
//	1.分解：
//		如果输入值的规格小于确定的阈值(如果一个或者两个元素)，我们就通过使用直接了当的方法来解决这些问题并返回所获的的答案。
//		否则，我们把输入值分解为两个或者更多的互斥子集。
//	2.解决子问题：
//		递归地解决这些与子集相关的子问题
//	3.合并：
//		整理这些子问题的解，然后把它们合并成一个整体用以解决最开始的问题。
//
// 归并排序执行过程：
//
//	 1.分解：
//			若S只有0个或者1个元素，直接返回s；此时它已经完成排序了。否则(若S有至少2个元素)，从s中移除所有的元素，然后将它们放在
//			s1、s2两个子序列中，每个序列包含S中一般的元素。这就是说，s1包含S前一半的元素，s2包含s后一半的元素。
//		2.解决子问题：
//			递归地对S1和S2进行排序。
//		3.合并：
//			把这些分别在S1和S2中排好序的元素拿出来并按照顺序合并到S序列中
//
// 图示：
//
//	divide
//	85 24 63 45 17 31 96 50
//
// -->  85 24 63 45		17 31 96 50
// --> 	85 24	63 45	17 31	96 50
// -->	85	24	63	45	17	31	96	50
//
//	merge
//
// -->  24 85	45 63	17 31	50 96
// -->  24 45 85 63	17 31 50 96
// -->	17 24 31 45 50 63 85 96
//
// 时间复杂度为O(nlogn)
func MergeSort(s []interface{}, f fucker.CompareFunc) {
	n := len(s)
	if n < 2 {
		return
	}

	mid := n >> 1
	s1 := s[0:mid]
	s2 := s[mid:n]
	MergeSort(s1, f)
	MergeSort(s2, f)
	merge(s1, s2, s, f)
}

// 合并2个已排序的序列s1和s2到s中.
func merge(s1, s2, s []interface{}, f fucker.CompareFunc) {
	s3 := make([]interface{}, len(s1))
	copy(s3, s1) // 这里只需要一个数组的备份
	// i: s3(s1的copy)的索引，j: s2的索引
	for i, j := 0, 0; i+j < len(s); {
		if j == len(s2) || (i < len(s3) && f(s3[i], s2[j]) != fucker.Greater) {
			s[i+j] = s3[i]
			i++
		} else {
			s[i+j] = s2[j]
			j++
		}

		// 优化：当s3(s1的copy)已经遍历完毕后，其实s2可以不用再继续遍历了
		if len(s3) == i {
			break
		}

		// 优化: 当s2遍历完毕后，直接复制s3(s1的copy)剩余部分就完事了
		if len(s2) == j {
			copy(s[i+j:], s3[i:])
			break
		}
	}
}
