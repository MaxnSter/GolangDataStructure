package search_for_range

import "sort"

/*
34. 搜索范围
给定一个按照升序排列的整数数组 nums，和一个目标值 target。找出给定目标值在数组中的开始位置和结束位置。

你的算法时间复杂度必须是 O(log n) 级别。

如果数组中不存在目标值，返回 [-1, -1]。

示例 1:

输入: nums = [5,7,7,8,8,10], target = 8
输出: [3,4]
示例 2:

输入: nums = [5,7,7,8,8,10], target = 6
输出: [-1,-1]
*/
func searchRange(nums []int, target int) []int {
	if len(nums) == 0 {
		return nil
	}
	//找到这个数第一次出现的位置
	idx := sort.Search(len(nums), func(i int) bool {
		return nums[i] >= target
	})
	if  idx == len(nums) || nums[idx] != target{
		return []int{-1, -1}
	}

	idxSame := idx
	for i := idxSame; i < len(nums); i++ {
		if nums[i] != nums[idx] {
			break
		}
		idxSame = i
	}

	return []int{idx, idxSame}
}
