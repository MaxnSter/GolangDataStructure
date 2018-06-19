package majority

import "sort"

/*
给定一个大小为 n 的数组，找到其中的众数。众数是指在数组中出现次数大于 ⌊ n/2 ⌋ 的元素。

你可以假设数组是非空的，并且给定的数组总是存在众数。

示例 1:

输入: [3,2,3]
输出: 3
示例 2:

输入: [2,2,1,1,1,2,2]
输出: 2
https://leetcode-cn.com/problems/majority-element/description/
*/

// 正确姿势,众数能抵消其他不是众数的元素
func majorityElement(nums []int) int {

	majorityNum, count := nums[0], 1
	for i := 1; i < len(nums); i++ {
		if nums[i] == majorityNum {
			count++
		} else {
			count--
		}

		if count == 0 {
			majorityNum = nums[i]
			count = 1
		}
	}

	//检查出现次数
	count = 0
	for _, v := range nums {
		if v == majorityNum {
			count++
		}
	}

	if count > len(nums)/2 {
		return majorityNum
	}
	return -1
}

// 先排序,然后找出众数
func majorityElement1(nums []int) int {
	sort.Ints(nums)

	majorityNum, count := nums[0], 1
	for i := 1; i < len(nums); i++ {
		if nums[i] == nums[i-1] {
			count++
			if count > len(nums)/2 {
				return majorityNum
			}
		} else {
			majorityNum, count = nums[i], 1
		}
	}

	return majorityNum
}

// hashMap 一遍过
func majorityElement2(nums []int) int {
	m := map[int]int{}
	for _, v := range nums {
		if _, ok := m[v]; ok {
			m[v]++
		} else {
			m[v] = 1
		}

		if m[v] > len(nums)/2 {
			return v
		}
	}

	return -1
}
