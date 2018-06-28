package insert_interval

/*
57. 插入区间
给出一个无重叠的 ，按照区间起始端点排序的区间列表。

在列表中插入一个新的区间，你需要确保列表中的区间仍然有序且不重叠（如果有必要的话，可以合并区间）。

示例 1:

输入: intervals = [[1,3],[6,9]], newInterval = [2,5]
输出: [[1,5],[6,9]]
示例 2:

输入: intervals = [[1,2],[3,5],[6,7],[8,10],[12,16]], newInterval = [4,8]
输出: [[1,2],[3,10],[12,16]]
解释: 这是因为新的区间 [4,8] 与 [3,5],[6,7],[8,10] 重叠。
*/
/**
* Definition for an interval.
*
 */
type Interval struct {
	Start int
	End   int
}

func insert(intervals []Interval, newInterval Interval) []Interval {
	intervals = append(intervals, newInterval)

	var idx int
	for idx = len(intervals) - 1; idx > 0 && intervals[idx-1].Start > newInterval.Start; idx-- {
		intervals[idx] = intervals[idx-1]
	}
	//找到插入点
	intervals[idx] = newInterval

	//开始合并区间
	last, result := intervals[0], make([]Interval, 0)
	for i := 1; i < len(intervals); i++ {
		if intervals[i].Start <= last.End {
			last.End = maxInt(last.End, intervals[i].End)
		} else {
			result = append(result, last)
			last = intervals[i]
		}
	}

	//数组长度至少为1,所以这里安全
	result = append(result, last)
	return result
}

func maxInt(x, y int) int {
	if x > y {
		return x
	}
	return y
}
