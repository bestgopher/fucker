package internal

import (
	"gopkg.in/bestgopher/fucker"
)

// 快速排序
//
//	基于分治法，但是是把所有的复杂操作在递归调用之前做完。
//	主要思想是应用分治法把序列s分解为子序列，递归地排序每个子序列，然后通过简单串联的方式合并这些已排序的子序列。
//
// 步骤：
//
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
//
// 此排序为非原地，有大量的容器创建和删除
// Deprecated
func QuickSort1(s []interface{}, f fucker.CompareFunc) {
	n := len(s)
	if n < 2 {
		return
	}

	p := s[0]
	var (
		L []interface{}
		E []interface{}
		G []interface{}
	)

	for _, v := range s {
		switch f(v, p) {
		case fucker.Less:
			L = append(L, v)
		case fucker.Equal:
			E = append(E, v)
		case fucker.Greater:
			G = append(G, v)
		}
	}

	QuickSort1(L, f)
	QuickSort1(G, f)

	copy(s[:len(L)], L)
	copy(s[len(L):len(L)+len(E)], E)
	copy(s[len(L)+len(E):len(L)+len(E)+len(G)], G)
}

// 原地排序的快速排序实现
// 步骤：
//
//	1.指定一个基准值p(p为s[begin]), begin为0，end为len(s)-1
//	2.首先从后往前遍历，如果遍历的元素大于基准值，则继续往前遍历, end = end-1
//	3.如果从后往前遍历时的值小于基准值，begin的值赋值为s[end], begin = begin + 1
//	4.然后从前往后遍历，遍历值小于或者等于基准值时，继续向后遍历，begin = begin + 1
//	5.当遍历值大于或者等于基准值时，s[end] = s[begin]
//	6.再次重复2-5步，直到begin == end为止
//	7.最后s[begin] = p
//	8.递归快速排序以begin为分界线的两部分序列
func QuickSort(s []interface{}, f fucker.CompareFunc) {
	n := len(s)
	if n < 2 {
		return
	}

	// p为基准值
	p, begin, end := s[0], 0, n-1

	for begin < end {
		for begin < end {
			if f(p, s[end]) == fucker.Less {
				end--
			} else {
				s[begin] = s[end]
				begin++
				break
			}
		}

		for begin < end {
			if f(p, s[begin]) == fucker.Greater {
				begin++
			} else {
				s[end] = s[begin]
				end--
				break
			}
		}
	}

	s[begin] = p
	QuickSort(s[:begin], f)
	QuickSort(s[begin+1:], f)
}
