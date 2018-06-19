package three_sum

import "sort"

/*
15. 三数之和
给定一个包含 n 个整数的数组 nums，判断 nums 中是否存在三个元素 a，b，c ，使得 a + b + c = 0 ？找出所有满足条件且不重复的三元组。

注意：答案中不可以包含重复的三元组。

例如, 给定数组 nums = [-1, 0, 1, 2, -1, -4]，

满足要求的三元组集合为：
[
  [-1, 0, 1],
  [-1, -1, 2]
]

*/
func threeSum(nums []int) [][]int {
	if len(nums) < 3 {
		return nil
	}

	res := make([][]int, 0)
	sort.Ints(nums)
	if nums[0] > 0 {
		return nil
	}

	for i := 0; i < len(nums); i++ {
		//想想为什么idxL是i+i,而不是你认为的从i的最左边0开始
		//因为i从0开始遍历,遍历到j时,[0,j)的数肯定都验证过了,不需要再重复遍历
		idxL := i + 1
		idxR := len(nums) - 1

		for idxL < idxR {
			sum := nums[i] + nums[idxL] + nums[idxR]
			if sum == 0 {
				sums := []int{nums[i], nums[idxL], nums[idxR]}
				if !contain3(res, sums) {
					res = append(res, sums)
				}

				idxL++
				idxR--
			} else if sum < 0 {
				idxL++
			} else {
				idxR--
			}
		}
	}

	return res
}

func contain3(numss [][]int, nums []int) bool {
	for i := 0; i < len(numss); i++ {
		if numss[i][0] == nums[0] &&
			numss[i][1] == nums[1] &&
			numss[i][2] == nums[2] {
			return true
		}
	}
	return false
}
