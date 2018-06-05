package sumarray

import "math"

/*
给定一个整数数组 nums ，找到一个具有最大和的连续子数组（子数组最少包含一个元素），返回其最大和。

示例:

输入: [-2,1,-3,4,-1,2,1,-5,4],
输出: 6
解释: 连续子数组 [4,-1,2,1] 的和最大，为 6。
进阶:

如果你已经实现复杂度为 O(n) 的解法，尝试使用更为精妙的分治法求解。
 */
// https://leetcode-cn.com/problems/maximum-subarray/description/
func maxSubArray(nums []int) int {
	maxSum, thisSum := math.MinInt64, 0
	for _, num := range nums {
		thisSum += num

		if maxSum < thisSum {
			maxSum = thisSum
		}

		// 对于数组中的某个元素,如果前序列和为正数,那么再加上该元素
		// 就有可能获得更大的数. 如果前序列和为负数,那么肯定不需要
		// 前序列
		if thisSum < 0 {
			thisSum = 0
		}
	}

	return maxSum
}
