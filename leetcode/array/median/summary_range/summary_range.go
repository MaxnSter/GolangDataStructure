package summary_range

import (
	"fmt"
	"strconv"
)

/*
228. 汇总区间
给定一个无重复元素的有序整数数组，返回数组区间范围的汇总。

示例 1:

输入: [0,1,2,4,5,7]
输出: ["0->2","4->5","7"]
解释: 0,1,2 可组成一个连续的区间; 4,5 可组成一个连续的区间。
示例 2:

输入: [0,2,3,4,6,8,9]
输出: ["0","2->4","6","8->9"]
解释: 2,3,4 可组成一个连续的区间; 8,9 可组成一个连续的区间。
*/
func summaryRanges(nums []int) []string {
	if len(nums) == 0 {
		return nil
	}

	if len(nums) == 1 {
		return []string{fmt.Sprintf("%d", nums[0])}
	}

	idxB, idxE := 0, 0
	res := make([]string, 0)
	for i := 1; i < len(nums); i++ {
		if nums[i]-nums[i-1] == 1 {
			//连续
			idxE = idxE + 1
		} else {
			if idxB == idxE {
				res = append(res, strconv.Itoa(nums[idxB]))
			} else {
				res = append(res, fmt.Sprintf("%d->%d", nums[idxB], nums[idxE]))
			}

			//重置下标
			idxB, idxE = i, i
		}
	}

	if idxB == idxE {
		res = append(res, strconv.Itoa(nums[idxB]))
	} else {
		res = append(res, fmt.Sprintf("%d->%d", nums[idxB], nums[idxE]))
	}
	return res
}
