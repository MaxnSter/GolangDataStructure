package map_gap

/*
164. 最大间距
给定一个无序的数组，找出数组在排序之后，相邻元素之间最大的差值。

如果数组元素个数小于 2，则返回 0。

示例 1:

输入: [3,6,9,1]
输出: 3
解释: 排序后的数组是 [1,3,6,9], 其中相邻元素 (3,6) 和 (6,9) 之间都存在最大差值 3。
示例 2:

输入: [10]
输出: 0
解释: 数组元素个数小于 2，因此返回 0。
说明:

你可以假设数组中所有元素都是非负整数，且数值在 32 位有符号整数范围内。
请尝试在线性时间复杂度和空间复杂度的条件下解决此问题。
*/

type bk struct {
	min, max int
}

func maximumGap(nums []int) int {
	if len(nums) <= 1 {
		return 0
	}

	max, min, n := nums[0], nums[0], len(nums)
	for _, v := range nums {
		if v < min {
			min = v
		}

		if v > max {
			max = v
		}
	}
	if max == min {
		return 0
	}

	bkSize := float64(max-min) / float64(n-1)
	buckets, emptyBk := make([]bk, n), bk{min: -1, max: -1}
	for i := 0; i < len(buckets); i++ {
		buckets[i] = emptyBk
	}

	for i := 0; i < n; i++ {
		bkId := int(float64(nums[i]-min) / bkSize)
		if buckets[bkId].min == -1 {
			buckets[bkId].min, buckets[bkId].max = nums[i], nums[i]
		} else {
			if nums[i] < buckets[bkId].min {
				buckets[bkId].min = nums[i]
			}

			if nums[i] > buckets[bkId].max {
				buckets[bkId].max = nums[i]
			}
		}
	}

	result, last := -1, 0
	for i := 1; i < len(buckets); i++ {
		if buckets[i].min == -1 {
			continue
		}

		// wrong, 因为直接用i-1,上一个bucket可能是空的
		//if buckets[i].min-buckets[-1].max > result {
		//	result = buckets[i].min - buckets[-1].max
		//}
		if buckets[i].min-buckets[last].max > result {
			result = buckets[i].min - buckets[last].max
		}
		last = i
	}
	return result
}
