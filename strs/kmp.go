package strs

// KMP实现了kmp算法
// 其中haystack中存在needle，返回下标
func KMP(haystack, needle string) int {
	if len(needle) == 0 {
		return 0
	}

	prefixTable := prefixTable(needle)

	// j为needle的下标，i为haystack的下标
	for j, i := 0, 0; i < len(haystack); {
		if haystack[i] == needle[j] {
			// 当值相等，且j为needle的最后一个元素时，说明匹配成功，则直接返回下标
			if j == len(needle)-1 {
				return i - j
			}
			// 继续匹配haystack和needle的下一个元素
			j++
			i++
		} else {
			// 当不匹配时，且needle不是第一个元素，则返回匹配needle的第prefixTable[j-1]
			if j != 0 {
				j = prefixTable[j-1]
			} else {
				i++
			}
		}
	}

	return -1
}

/*
prefixTable获取到needle的前缀表
何为前缀？
	举例：
		abacdaba 的前缀为3，因为字符串的前三个aba与后三个aba一样
		aabacdaba 的前缀为1，因为第一个a与最后一个a一样

needle的前缀表为每一位的前缀组成的列表
例如：needle = "abcdabca"
第一位的前缀为0
第二位的前缀为0
第三位的前缀为0
第四位的前缀为0
第五位的前缀为1
第六位的前缀为2
第七位的前缀为3
第八位的前缀为1

算法：
j i
a b c d a b c a

初始化两个下标j,i为1
当needle[i] == needle[j]时，prefixTable[i] = prefixTable[i-1]+1
当needle[i] != needle[j]时，
	j = prefixTable[j-1]，如果needle[i] == needle[j]时，prefixTable[i] = prefixTable[j]+1
	否则循环寻找j，知道j=0
*/
func prefixTable(needle string) []int {
	prefixTable := make([]int, len(needle))
	prefixTable[0] = 0 // 第一个元素默认为0

	for j, i := 0, 1; i < len(needle); i++ {
		// 当needle[i] == needle[j]时，直接当前位置的前缀值等于前一个位置的前缀值+1
		if needle[i] == needle[j] {
			prefixTable[i] = prefixTable[i-1] + 1
			j++
		} else {
			// 当needle[i] != needle[j]时，
			// j = prefixTable[j-1]，当needle[j] != needle[i], 继续j = prefixTable[j-1]，直到j=0
			// 否则，prefixTable[i] = prefixTable[j] + 1
			for j > 0 {
				j = prefixTable[j-1]
				if needle[i] == needle[j] {
					prefixTable[i] = prefixTable[j] + 1
					j++
					break
				} else {
					prefixTable[i] = 0
				}
			}
		}
	}

	return prefixTable
}
