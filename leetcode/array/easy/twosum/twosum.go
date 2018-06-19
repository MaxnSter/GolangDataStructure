// https://leetcode-cn.com/problems/two-sum/description/
package twosum

// 时间复杂度 0(n^2)
// 空间复杂度 0(1)
func TwoSumSlow(nums []int, target int) []int {
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				return []int{i, j}
			}
		}
	}

	return nil
}

// 时间复杂度 0(n)
// 空间复杂度 0(n)
func TwoSumFast(nums []int, target int) []int {
	m := make(map[int]int)

	for i, v := range nums {
		complement := target - v

		if idx, ok := m[complement]; ok {
			return []int{idx, i}
		}
		m[v] = i
	}

	return nil
}
