package intersection

import "sort"

/*
349. 两个数组的交集
给定两个数组，写一个函数来计算它们的交集。

例子:

 给定 num1= [1, 2, 2, 1], nums2 = [2, 2], 返回 [2].

提示:

每个在结果中的元素必定是唯一的。
我们可以不考虑输出结果的顺序。
*/
func intersection(nums1 []int, nums2 []int) []int {
	idx1, idx2, result := 0, 0, make([]int, 0)

	sort.Ints(nums1)
	sort.Ints(nums2)

	for idx1 < len(nums1) && idx2 < len(nums2) {
		if nums1[idx1] < nums2[idx2] {
			idx1++
		} else if nums1[idx1] > nums2[idx2] {
			idx2++
		} else {
			same := nums1[idx1]
			result = append(result, same)

			for idx1 < len(nums1) && nums1[idx1] == same {
				idx1++
			}

			for idx2 < len(nums2) && nums2[idx2] == same {
				idx2++
			}
		}
	}

	return result
}
