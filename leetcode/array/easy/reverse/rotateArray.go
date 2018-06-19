package reverse

/*
给定一个数组，将数组中的元素向右移动 k 个位置，其中 k 是非负数。

示例 1:

输入: [1,2,3,4,5,6,7] 和 k = 3
输出: [5,6,7,1,2,3,4]
解释:
向右旋转 1 步: [7,1,2,3,4,5,6]
向右旋转 2 步: [6,7,1,2,3,4,5]
向右旋转 3 步: [5,6,7,1,2,3,4]
示例 2:

输入: [-1,-100,3,99] 和 k = 2
输出: [3,99,-1,-100]
解释:
向右旋转 1 步: [99,-1,-100,3]
向右旋转 2 步: [3,99,-1,-100]
说明:

尽可能想出更多的解决方案，至少有三种不同的方法可以解决这个问题。
要求使用空间复杂度为 O(1) 的原地算法。

https://leetcode-cn.com/problems/rotate-array/description/
*/

// 方法1:旋转3次
func rotate(nums []int, k int) {

	if len(nums) <= 1 {
		return
	}

	// 当k超过数组长度时
	k = k % len(nums)
	if k == 0 {
		return
	}

	// 性能最好
	reverseArray(nums, 0, len(nums)-1)
	reverseArray(nums, 0, k-1)
	reverseArray(nums, k, len(nums)-1)
}

func reverseArray(nums []int, begin, end int) {
	for begin < end {
		nums[begin], nums[end] = nums[end], nums[begin]
		begin++
		end--
	}
}

// 方法2:数组向右移动k次
func rotate1(nums []int, k int) {
	if len(nums) <= 1 {
		return
	}

	k = k % len(nums)
	if k == 0 {
		return
	}

	// 当k很大时,时间复杂度o(n^2)
	for i := 0; i < k; i++ {
		moveArray(nums)
	}
}

func moveArray(nums []int) {
	if len(nums) <= 1 {
		return
	}

	l := len(nums)
	last := nums[l-1]

	for i := l - 2; i >= 0; i-- {
		nums[i+1] = nums[i]
	}

	nums[0] = last
}

// 方法3:暂时保存k个最后的数,只要注意k大于数组长度的情况就好了
func rotate2(nums []int, k int) {
	k = k % len(nums)
	if k == 0 {
		return
	}

	l := len(nums)
	tmp := make([]int, k)
	lastK := nums[l-k:]
	beforeK := nums[:l-k]

	// 暂存后k个元素[len-k:]
	for i := 0; i < len(lastK); i++ {
		tmp[i] = lastK[i]
	}

	// 前[:len-k]个元素向右移动
	for i := len(beforeK) - 1; i >= 0; i-- {
		nums[i+k] = beforeK[i]
	}

	// 前[:k]个元素赋值为tmp存放的值
	for i := 0; i < k; i++ {
		nums[i] = tmp[i]
	}

}
