package maximum_sub_array

/*

给定 n 个整数，找出平均数最大且长度为 k 的连续子数组，并输出该最大平均数。

示例 1:

输入: [1,12,-5,-6,50,3], k = 4
输出: 12.75
解释: 最大平均数 (12-5-6+50)/4 = 51/4 = 12.75


注意:

1 <= k <= n <= 30,000。
所给数据范围 [-10,000，10,000]

https://leetcode-cn.com/problems/maximum-average-subarray-i/description/
*/
func findMaxAverage(nums []int, k int) float64 {
	maxAvg := avg(nums[:k])
	for i := 1; i+k <= len(nums); i++ {
		curAvg := avg(nums[i : i+k])
		if curAvg > maxAvg {
			maxAvg = curAvg
		}
	}
	return maxAvg
}

func avg(nums []int) float64 {
	var sum int
	for _, v := range nums {
		sum += v
	}

	return float64(sum) / float64(len(nums))
}

func sum(nums []int) int {
	var sum int
	for i := 0; i < len(nums); i++ {
		sum += nums[i]
	}
	return sum
}

func maxInt(x, y int) int {
	if x < y {
		return x
	}
	return y
}

// 滑动窗口
func findMaxAverage1(nums []int, k int) float64 {
	var curMaxAvg, maxAvg int
	curMaxAvg = sum(nums[:k])
	maxAvg = curMaxAvg

	for i := k; i < len(nums); i++ {
		curMaxAvg = curMaxAvg + nums[i] - nums[i-k]
		maxAvg = maxInt(curMaxAvg, maxAvg)
	}

	return float64(maxAvg) / float64(k)
}
