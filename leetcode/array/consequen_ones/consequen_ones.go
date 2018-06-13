package consequen_ones

/*
给定一个二进制数组， 计算其中最大连续1的个数。

示例 1:

输入: [1,1,0,1,1,1]
输出: 3
解释: 开头的两位和最后的三位都是连续1，所以最大连续1的个数是 3.
注意：

输入的数组只包含 0 和1。
输入数组的长度是正整数，且不超过 10,000
https://leetcode-cn.com/problems/max-consecutive-ones/description/
*/
func findMaxConsecutiveOnes(nums []int) int {
	var maxOnes, curOnes int
	for i := 0; i < len(nums); i++ {
		if nums[i] == 0 {
			if curOnes > maxOnes {
				maxOnes = curOnes
			}
			curOnes = 0
		} else {
			curOnes++
		}
	}

	if maxOnes < curOnes {
		maxOnes = curOnes

	}
	return maxOnes
}