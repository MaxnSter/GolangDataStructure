package median_two_sorted

import (
	"sort"

	"github.com/MaxnSter/GolangDataStructure/kth"
)

/*
两个排序数组的中位数
给定两个大小为 m 和 n 的有序数组 nums1 和 nums2 。

请找出这两个有序数组的中位数。要求算法的时间复杂度为 O(log (m+n)) 。

示例 1:

nums1 = [1, 3]
nums2 = [2]

中位数是 2.0
示例 2:

nums1 = [1, 2]
nums2 = [3, 4]

中位数是 (2 + 3)/2 = 2.5
*/
//使用find kth算法,哈哈哈
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	n1, n2 := len(nums1), len(nums2)
	var min, max int = findMinAndMax(nums1, nums2)

	k := (n1 + n2+1) / 2
	searchFunc := func(guess int) (smaller, same int) {
		smaller1, same1 := search(nums1, guess)
		smaller2, same2 := search(nums2, guess)
		smaller, same = smaller1+smaller2, same1+same2
		return
	}

	median, _ := kth.FindKth(searchFunc, k, n1+n2, min, max)
	if (n1+n2)%2 == 1 {
		return float64(median)
	} else {
		//偶数
		median2, _ := kth.FindKth(searchFunc, k+1, n1+n2, min, max)
		return (float64(median) + float64(median2)) / 2
	}
}

func findMinAndMax(nums1[]int, nums2[]int) (min, max int) {
	n1, n2 := len(nums1), len(nums2)
	if n1 == 0 {
		return nums2[0], nums2[n2-1]
	}
	if n2 == 0 {
		return nums1[0], nums1[n1-1]
	}

	if nums1[0] < nums2[0] {
		min = nums1[0]
	} else {
		min = nums2[0]
	}
	if nums1[n1-1] > nums2[n2-1] {
		max = nums1[n1-1]
	} else {
		max = nums2[n2-1]
	}

	return
}

func search(nums []int, guess int) (smaller, same int) {
	if len(nums) == 0 {
		return 0, 0
	}

	idx := sort.Search(len(nums), func(i int) bool {
		return nums[i] >= guess
	})

	if idx == len(nums) {
		return len(nums), 0
	} else {
		smaller = idx
		for i := idx; i < len(nums); i++ {
			if nums[i] != guess {
				break
			} else {
				same++
			}
		}
		return
	}
}
