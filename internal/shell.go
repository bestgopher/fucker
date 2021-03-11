package internal

import (
	"math"
)

/*
希尔排序
	把序列看作是一个矩阵，分成m列，逐列进行排序
	m从某个整数逐渐减为1
	当m为1时，整个序列将完全有序
也称为递减增量排序
	矩阵的列数取决于步长序列(step sequence)
	比如，如果步长序列为{1, 5, 19, 41, 109...}，就代表依次分为109列，41列，19列，5列，1列
	不同的步长序列，执行效率也不同

希尔本人给出的步长序列是n/2^k，比如n为16时，步长序列时{1, 2, 4, 8}(k依次取1，2，3，4)

比如序列：16 15 14 13 12 11 10 9 8 7 6 5 4 3 2 1
首先分为8列：
	16 15 14 13 12 11 10 9
	 8  7  6  5  4  3  2 1
排序后：
	 8  7  6  5  4  3  2 1
	16 15 14 13 12 11 10 9
序列变成：8 7 6 5 4 3 2 1 16 15 14 13 12 11 10 9

然后分成4列：
	 8 	7  6  5
	 4  3  2  1
	16 15 14 13
	12 11 10  9
排序后：
	4   3  2  1
	8 	7  6  5
	12 11 10  9
	16 15 14 13
序列变成：4  3  2  1  8 7  6  5  12 11 10  9 16 15 14 13

再分成2列
再分成1列
*/
func Shell(s []int) {
	// 先生成步长序列，按照希尔本人的算法来算 n / 2^k
	steps := stepSequence(s)
	// 分成多少列来排序

	for _, step := range steps {
		// col: 列，column
		// 对第col列的元素进行插入排序
		for col := 0; col < step; col++ {
			for i := col + step; i < len(s); i += step {
				for j := i; j > col; j -= step {
					if s[j] < s[j-step] {
						s[j], s[j-step] = s[j-step], s[j]
					} else {
						break
					}
				}
			}
		}
	}
}

// n / 2^k
// 获取步长的方法，希尔本人提供的公式
// 此公式得到步长希尔排序最差复杂度为O(n^2)
// Deprecated
func stepSequence1(s []int) []int {
	steps := make([]int, 0)
	for i := len(s) >> 1; i >= 1; i >>= 1 {
		steps = append(steps, i)
	}
	return steps
}

/*
	9 * (2^k - 2^(k/2)) + 1, k is even
	8 * 2^k - 6 * 2^((k+1)/2) + 1, k is odd
目前已知最好的步长序列，此公式最坏时间复杂度为O(n^(4/3))
*/
func stepSequence(s []int) []int {
	steps := make([]int, 0)

	for k := 0; ; k++ {
		step := 0
		if k%2 == 0 {
			step = 1 + int(9*(math.Pow(2, float64(k))-math.Pow(2, float64(k>>1))))
		} else {
			step = int(8*math.Pow(2, float64(k))-6*math.Pow(2, float64((k+1)>>1))) + 1
		}
		if step > len(s) {
			break
		}

		// 这里要步长从大到小排序
		_steps := make([]int, len(steps)+1)
		_steps[0] = step
		copy(_steps[1:], steps)
		steps = _steps
	}
	return steps
}
