package reverse

import "math"

// https://leetcode-cn.com/problems/reverse-integer/description/
func reverse(x int) int {
	var result int
	for x != 0 {
		// golang中负数取模得到一个负数
		result = result*10 + x%10
		x /= 10
	}

	if result > math.MaxInt32 || result < math.MinInt32 {
		return 0
	}
	return result
}
