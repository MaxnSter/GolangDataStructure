package four_sum

import "sort"

/*
18. 四数之和
给定一个包含 n 个整数的数组 nums 和一个目标值 target，
判断 nums 中是否存在四个元素 a，b，c 和 d ，使得 a + b + c + d 的值与 target 相等？
找出所有满足条件且不重复的四元组。

注意：

答案中不可以包含重复的四元组。
示例：

给定数组 nums = [1, 0, -1, 0, -2, 2]，和 target = 0。

满足要求的四元组集合为：
[
  [-1,  0, 0, 1],
  [-2, -1, 1, 2],
  [-2,  0, 0, 2]
]
*/
func fourSum(nums []int, target int) [][]int {
	if len(nums) < 4 {
		return nil
	}

	res := make([][]int, 0)
	sort.Ints(nums)
	if nums[0] > 0 {
		return nil
	}

	for i := 0; i < len(nums)-3; i++ {
		for j := i + 1; j < len(nums)-2; j++ {
			idxL := j + 1
			idxR := len(nums) - 1

			for idxL < idxR {
				sum := nums[i] + nums[idxL] + nums[idxR] + nums[j]
				if sum == target {
					sums := []int{nums[i], nums[j], nums[idxL], nums[idxR]}
					if !contain4(res, sums) {
						res = append(res, sums)
					}

					idxL++
					idxR--
				} else if sum < target {
					idxL++
				} else {
					idxR--
				}
			}
		}
	}

	return res
}

func contain4(numss [][]int, nums []int) bool {
	for i := 0; i < len(numss); i++ {
		if numss[i][0] == nums[0] &&
			numss[i][1] == nums[1] &&
			numss[i][2] == nums[2] &&
			numss[i][3] == nums[3] {
			return true
		}
	}
	return false
}
