package internal

import (
	"github.com/bestgopher/fucker"
)

/*
插入排序
	非常类似扑克牌的排序
步骤
	1.在执行过程中，插入排序会将序列分为2部分：头部是已经排序的，尾部是待排序的
	2.从头开始扫描每一个元素，每当扫到一个元素，就将它插入到头部合适的位置，使得头部数据依然保持有序

逆序对
	数组[2, 3, 8, 6, 1]的逆序对为(2, 1), (3, 1), (8, 1), (6, 1), (8, 6)
	插入排序的时间复杂度与逆序对的数量成正比，逆序对的数量越多，插入排序的时间复杂度越高

优化1
	将交换改为挪动，找到插入的下标，然后把此下标及其后面的元素向后挪动一个位置，在插入元素到这个下标
优化2
	通过二分搜索找到插入的位置(因为前面部分是有序的，所以这里可以是用二分搜索)
	返回的位置上的元素要大于待插入元素，且此位置前面的元素小于等于待插入元素
*/
func InsertionSort(s []interface{}, f fucker.CompareFunc) {
	for i := 1; i < len(s); i++ {

		if f(s[i], s[i-1]) == fucker.Greater {
			continue
		}

		// 二分查找
		start, end, middle := 0, i, 0
		for {
			middle = (start + end) >> 1
			if end-start <= 1 || middle == 0 || middle == i ||
				f(s[middle], s[i]) == fucker.Greater && f(s[middle-1], s[i]) != fucker.Greater {
				break
			} else if f(s[middle], s[i]) == fucker.Greater {
				end = middle
			} else if f(s[middle], s[i]) != fucker.Greater {
				start = middle + 1
			}
		}

		v := s[i]
		for x := i; x > middle; x-- {
			s[x] = s[x-1]
		}
		s[middle] = v
	}
}

// Deprecated
func InsertionSort1(s []int) {
	for i := 1; i < len(s); i++ {
		v := s[i]
		for j := i; j > 0; j-- {
			if s[j-1] > v {
				s[j] = s[j-1]
			} else {
				s[j] = v
				break
			}
		}
	}
}
