package _sum_closest

import "sort"

/*
16. 最接近的三数之和
给定一个包括 n 个整数的数组 nums 和 一个目标值 target。
找出 nums 中的三个整数，使得它们的和与 target 最接近。返回这三个数的和。假定每组输入只存在唯一答案。

例如，给定数组 nums = [-1，2，1，-4], 和 target = 1.

与 target 最接近的三个数的和为 2. (-1 + 2 + 1 = 2).
*/
func threeSumClosest(nums []int, target int) int {
	sort.Ints(nums)
	maxGap, result := sumInt(nums[:3]...)-target, sumInt(nums[:3]...)

	for i := 0; i < len(nums); i++ {
		idxL := i + 1
		idxR := len(nums) - 1

		for idxL < idxR {
			sum := sumInt(nums[i], nums[idxL], nums[idxR])
			if gap := sum - target; absInt(gap) < absInt(maxGap) {
				result = sumInt(nums[i], nums[idxL], nums[idxR])
				maxGap = gap
			}

			if sum > target {
				idxR--
			} else {
				idxL++
			}
		}
	}

	return result
}

func absInt(num int) int {
	if num < 0 {
		return -num
	}
	return num
}

func sumInt(nums ...int) (sum int) {
	for _, num := range nums {
		sum += num
	}
	return
}
