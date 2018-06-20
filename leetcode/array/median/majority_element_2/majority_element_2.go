package majority_element_2

/*
229. 求众数 II
给定一个大小为 n 的数组，找出其中所有出现超过 ⌊ n/3 ⌋ 次的元素。

说明: 要求算法的时间复杂度为 O(n)，空间复杂度为 O(1)。

示例 1:

输入: [3,2,3]
输出: [3]
示例 2:

输入: [1,1,1,3,3,2,2,2]
输出: [1,2]
*/
func majorityElement(nums []int) []int {
	//1.最多有两个
	//2.摩尔投票法
	m, n := 0, 0   //候选者
	cm, cn := 0, 0 //候选个数
	for i := 0; i < len(nums); i++ {
		if nums[i] == nums[m] {
			cm++
		} else if nums[i] == nums[n] {
			cn++
		} else if cm == 0 {
			m = i
			cm = 1
		} else if cn == 0 {
			n = i
			cn = 1
		} else {
			cm--
			cn--
		}
	}

	res := make([]int, 0)
	cm, cn = 0, 0
	for _, v := range nums {
		if v == nums[m] {
			cm++
		} else if v == nums[n] {
			cn++
		}
	}

	if cm > len(nums)/3 {
		res = append(res, nums[m])
	}
	if cn > len(nums)/3 {
		res = append(res, nums[n])
	}

	return res
}
