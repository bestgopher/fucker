package internal

import (
	fucker "gopkg.in/bestgopher/fucker.v1"
)

// 简单实现，
//
//	步骤：
//		1.找出最大值max
//		2.创建一个max长度的数组
//		3.遍历数据，得到数据值i。然后把数组中下标为i的值加1
//		4.遍历数组，则数组的下标为值，数组的值为个数，依次插入原数据中
//
//	缺点
//		1.浪费空间
//		2.无法对负数排序
//
// Deprecated
func CountingSort1(s []interface{}, f fucker.CompareFunc) {
	max := 0

	for _, v := range s {
		if v.(int) > max {
			max = v.(int)
		}
	}

	array := make([]int, max+1)

	for _, v := range s {
		array[v.(int)]++
	}

	index := 0
	for i, v := range array {
		for j := v; j > 0; j-- {
			s[index] = i
			index++
		}
	}
}

// 对CountingSort1进行改进
//
//	改进：
//		1.找出最大值max和最小值min，
//		2.创建一个max-min+1长度的数组
//		3.遍历数据，得到数据值i。然后把数组中下标为i-min的值加1
//		4.再次遍历数组，将下标array[i] += array[i-1]
//		5.遍历数组，则数组的下标+min为值，数组的值为个数，依次插入原数据中
//
// 优点:
//
//	1.节省空间
//	2.可以对负数排序
func CountingSort(s []interface{}, f fucker.CompareFunc) {
	max, min := 0, 0
	for _, v := range s {
		d := v.(int)
		if d > max {
			max = d
		}
		// 最大最小可能为同一个树，所以这里要分开判断
		if d < min {
			min = d
		}
	}

	array := make([]int, max-min+1)

	for _, v := range s {
		array[v.(int)-min]++
	}

	for i := 1; i < len(array); i++ {
		array[i] += array[i-1]
	}

	newArray := make([]interface{}, len(s))

	for i := len(s) - 1; i >= 0; i-- {
		array[s[i].(int)]--
		newArray[array[s[i].(int)]] = s[i]
	}

	copy(s, newArray)
}
