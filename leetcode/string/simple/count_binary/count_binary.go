package count_binary

/*
696. 计数二进制子串
给定一个字符串 s，计算具有相同数量0和1的非空(连续)子字符串的数量，并且这些子字符串中的所有0和所有1都是组合在一起的。

重复出现的子串要计算它们出现的次数。

示例 1 :

输入: "00110011"
输出: 6
解释: 有6个子串具有相同数量的连续1和0：“0011”，“01”，“1100”，“10”，“0011” 和 “01”。

请注意，一些重复出现的子串要计算它们出现的次数。

另外，“00110011”不是有效的子串，因为所有的0（和1）没有组合在一起。
示例 2 :

输入: "10101"
输出: 4
解释: 有4个子串：“10”，“01”，“10”，“01”，它们具有相同数量的连续1和0。
注意：

s.length 在1到50,000之间。
s 只包含“0”或“1”字符。
*/

//思路:
//查出这个串里相邻的0，1的个数，如：00110011这个的个数就是[2,2,2,2]，这样每相邻两个数取到最小值的和就是最后的结果。
//如:00011001查出个数为[3,2,2,1], res = min(3,2)+min(2,2)+min(2,1)=5
func countBinarySubstrings(s string) int {
	if len(s) <= 1 {
		return 0
	}

	curCount, counts := 1, make([]int, 0)
	for i := 1; i < len(s); i++ {
		if s[i] == s[i-1] {
			curCount++
		} else {
			counts = append(counts, curCount)
			curCount = 1
		}
	}

	counts = append(counts, curCount)

	if len(counts) <= 1 {
		return 0
	}

	result := 0
	for i := 1; i < len(counts); i++ {
		result += minInt(counts[i], counts[i-1])
	}
	return result
}

//简化版本
//巧用pre=0来避免首次遍历相加,巧用联机算法记录状态
func countBinarySubstringsLess(s string) int {
	if len(s) <= 1 {
		return 0
	}

	pre, cur, result := 0,1,0
	for i := 1; i < len(s); i++ {
		if s[i] == s[i-1] {
			cur++
		} else {
			result += minInt(pre, cur)
			pre, cur = cur, 1
		}
	}

	return result + minInt(pre, cur)
}

func minInt(x, y int) int {
	if x < y {
		return x
	}
	return y
}
