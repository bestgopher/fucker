package internal

import (
	fucker "gopkg.in/bestgopher/fucker.v1"
)

/*
选择排序

基本步骤:

	1.从序列中找出最大的元素，然后与最末尾的元素交换位置。
	2.忽略步骤一找到的最大元素，重复执行步骤一

优点:

	性能可能优于冒泡排序，因为不用每次比较了就交换元素位置，而是扫描完一遍后再交换元素位置。

缺点:

	最优时间复杂度和最差时间复杂度都是O(n^2)，因为选择完一圈不能确定序列是否有序。
*/
func SelectionSort(s []interface{}, f fucker.CompareFunc) {

	for i := len(s) - 1; i > 0; i-- {
		maxIndex := 0

		for j := 1; j <= i; j++ {
			// 注:这里只有加上等于增加算法的稳定(但是算法本身不是稳定的)
			if f(s[j], s[maxIndex]) != fucker.Less {
				maxIndex = j
			}
		}

		s[maxIndex], s[i] = s[i], s[maxIndex]
	}
}
