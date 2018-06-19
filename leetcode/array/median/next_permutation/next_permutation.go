package next_permutation

/*
实现获取下一个排列的函数，算法需要将给定数字序列重新排列成字典序中下一个更大的排列。

如果不存在下一个更大的排列，则将数字重新排列成最小的排列（即升序排列）。

必须原地修改，只允许使用额外常数空间。

以下是一些例子，输入位于左侧列，其相应输出位于右侧列。
1,2,3 → 1,3,2
3,2,1 → 1,2,3
1,1,5 → 1,5,1
*/
func nextPermutation(nums []int) {
	if len(nums) < 2 {
		return
	}

	for i, ii := len(nums)-2, len(nums)-1; i >= 0; i-- {
		if nums[i] < nums[ii] {
			// 找到i-1之后的第一个大于nums[i]的数
			var j = len(nums) - 1
			for nums[j] <= nums[i] {
				j--
			}

			nums[j], nums[i] = nums[i], nums[j]
			//注意,是从ii开始reverse
			reverse(nums[ii:])
			return
		}
		ii--
	}
	reverse(nums)
}

func reverse(nums []int) {
	idxB, idxE := 0, len(nums)-1
	for idxB < idxE {
		nums[idxB], nums[idxE] = nums[idxE], nums[idxB]
		idxB++
		idxE--
	}
}
