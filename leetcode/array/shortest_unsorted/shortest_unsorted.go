package shortest_unsorted

import "sort"

/*
给定一个整数数组，你需要寻找一个连续的子数组，如果对这个子数组进行升序排序，那么整个数组都会变为升序排序。

你找到的子数组应是最短的，请输出它的长度。

示例 1:

输入: [2, 6, 4, 8, 10, 9, 15]
输出: 5
解释: 你只需要对 [6, 4, 8, 10, 9] 进行升序排序，那么整个表都会变为升序排序。
说明 :

输入的数组长度范围在 [1, 10,000]。
输入的数组可能包含重复元素 ，所以升序的意思是<=。
*/
func findUnsortedSubarray(nums []int) int {
	numsOrder := make([]int, len(nums))
	copy(numsOrder, nums)
	sort.Ints(numsOrder)

	idxB, idxE := 0, -1
	for i := range numsOrder {
		if nums[i] != numsOrder[i] {
			idxB = i
			break
		}
	}

	for i := len(numsOrder) - 1; i >= 0; i-- {
		if nums[i] != numsOrder[i] {
			idxE = i
			break
		}
	}

	return idxE - idxB + 1
}

// O(N)时间, O(1)空间
func findUnsortedSubarrayEffective(nums []int) int {
	// end is -2 is because it works if the array is already in ascending order
	start, end := -1, -2
	min, max := nums[len(nums)-1], nums[0]
	n := len(nums)

	for i := 0; i < len(nums); i++ {
		max = maxInt(max, nums[i])
		min = minInt(min, nums[n-1-i])

		//如果是有序,那么应该是nums[i]>=max
		if nums[i] < max {
			end = i
		}

		//如果是有序,那么应该是nums[n-1-i]<=min
		if nums[n-1-i] > min {
			start = n - 1 - i
		}
	}
	return end - start + 1
}

func minInt(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func maxInt(x, y int) int {
	if x < y {
		return y
	}
	return x
}
