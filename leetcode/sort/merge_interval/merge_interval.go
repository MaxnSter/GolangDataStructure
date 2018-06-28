package merge_interval

import "sort"

/*
56. 合并区间
给出一个区间的集合，请合并所有重叠的区间。

示例 1:

输入: [[1,3],[2,6],[8,10],[15,18]]
输出: [[1,6],[8,10],[15,18]]
解释: 区间 [1,3] 和 [2,6] 重叠, 将它们合并为 [1,6].
示例 2:

输入: [[1,4],[4,5]]
输出: [[1,5]]
解释: 区间 [1,4] 和 [4,5] 可被视为重叠区间。
*/
/**
* Definition for an interval.
*
 */
type Interval struct {
	Start int
	End   int
}

type intervs []Interval

func (is intervs) Less(i, j int) bool { return is[i].Start < is[j].Start }
func (is intervs) Len() int           { return len(is) }
func (is intervs) Swap(i, j int)      { is[i], is[j] = is[j], is[i] }

func merge(intervals []Interval) []Interval {
	if len(intervals) <= 1 {
		return intervals
	}

	sort.Sort(intervs(intervals))
	result := make([]Interval, 0)

	last := intervals[0]
	for i := 1; i < len(intervals); i++ {
		if intervals[i].Start <= last.End {
			//记录一个连续区间
			last.End = maxInt(last.End, intervals[i].End)
		} else {
			result = append(result, last)
			last = intervals[i]
		}
	}

	result = append(result, last)
	return result
}

func maxInt(x, y int) int {
	if x > y {
		return x
	}
	return y
}
