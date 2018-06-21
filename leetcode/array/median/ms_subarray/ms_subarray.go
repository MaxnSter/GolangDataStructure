package ms_subarray

/*
209. 长度最小的子数组
给定一个含有 n 个正整数的数组和一个正整数 s ，找出该数组中满足其和 ≥ s 的长度最小的子数组。如果不存在符合条件的子数组，返回 0。

示例:

输入: s = 7, nums = [2,3,1,2,4,3]
输出: 2
解释: 子数组 [4,3] 是该条件下的长度最小的子数组。
进阶:

如果你已经完成了O(n) 时间复杂度的解法, 请尝试 O(n log n) 时间复杂度的解法。
*/
func minSubArrayLen(s int, nums []int) int {
	if len(nums) < 1 {
		return 0
	}

	if len(nums) == 1 {
		if nums[0] >= s {
			return 1
		} else {
			return 0
		}
	}

	//滑动窗口
	l, r, sum := 0, -1, 0
	gap := len(nums) + 1

	//思考为什么是l<len(nums)
	for l < len(nums) {
		if sum < s && r+1 < len(nums) {
			r++
			sum += nums[r]
		} else {
			sum -= nums[l]
			l++
		}

		if sum >= s {
			gap = minInt(gap, r-l+1)
		}
	}

	if gap == len(nums)+1 {
		return 0
	}
	return gap
}

func minInt(x, y int) int {
	if x < y {
		return x
	}
	return y
}
