package max_three_factor

import (
	"math"
	"sort"
)

/*
给定一个整型数组，在数组中找出由三个数组成的最大乘积，并输出这个乘积。

示例 1:

输入: [1,2,3]
输出: 6
示例 2:

输入: [1,2,3,4]
输出: 24
注意:

给定的整型数组长度范围是[3,104]，数组中所有的元素范围是[-1000, 1000]。
输入的数组中任意三个数的乘积不会超出32位有符号整数的范围。

这道题给了数组的范围，至少三个，
那么如果是三个的话，就无所谓了，直接相乘返回即可，
但是如果超过了3个，而且有负数存在的话，情况就可能不一样，我们来考虑几种情况，
如果全是负数，三个负数相乘还是负数，为了让负数最大，那么其绝对值就该最小，而负数排序后绝对值小的都在末尾，所以是末尾三个数字相乘，
这个跟全是正数的情况一样。
那么重点在于前半段是负数，后半段是正数，那么最好的情况肯定是两个最小的负数相乘得到一个正数，然后跟一个最大的正数相乘，
这样得到的肯定是最大的数，所以我们让前两个数相乘，再和数组的最后一个数字相乘，就可以得到这种情况下的最大的乘积。
实际上我们并不用分情况讨论数组的正负，只要把这两种情况的乘积都算出来，比较二者取较大值，
*/
func maximumProduct(nums []int) int {
	if len(nums) == 3 {
		return nums[0] * nums[1] * nums[2]
	}

	sort.Ints(nums)
	n := len(nums)
	return maxInt(nums[0]*nums[1]*nums[n-1],
		nums[n-1]*nums[n-2]*nums[n-3])
}

func maxInt(x, y int) int {
	if x > y {
		return x
	}
	return y
}

//找出3个最大的数 || 找出一个最大的和两个最小的，相乘对比也能得到结果，而且是O(n)的时间复杂度
func maximumProductEffetive(nums []int) int {
	if len(nums) == 3 {
		return nums[0] * nums[1] * nums[2]
	}

	max1, max2, max3 := math.MinInt64, math.MinInt64, math.MinInt64
	min1, min2 := math.MaxInt64, math.MaxInt64
	for _, v := range nums {
		if v > max1 {
			max3, max2, max1 = max2, max1, v
		} else if v > max2 {
			max3, max2 = max2, v
		} else if v > max3 {
			max3 = v
		}

		if v < min1 {
			min2, min1 = min1, v
		} else if v < min2 {
			min2 = v
		}
	}

	return maxInt(max1*max2*max3, min1*min2*max1)
}
