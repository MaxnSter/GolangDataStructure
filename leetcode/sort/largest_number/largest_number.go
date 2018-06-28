package largest_number

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

/*
179. 最大数
给定一组非负整数，重新排列它们的顺序使之组成一个最大的整数。

示例 1:

输入: [10,2]
输出: 210
示例 2:

输入: [3,30,34,5,9]
输出: 9534330
说明: 输出结果可能非常大，所以你需要返回一个字符串而不是整数。
*/

type mInts []int

func (ints mInts) Less(i, j int) bool {
	// ints[i] = 3, ints[j] = 30
	// 330 > 303
	strIJ := fmt.Sprintf("%d%d", ints[i], ints[j])
	strJI := fmt.Sprintf("%d%d", ints[j], ints[i])
	//想让在数组中,I排在J的前面,那么在这里,strIJ应该大于strJI
	return strIJ > strJI
}

func (ints mInts) Swap(i, j int) {
	ints[i], ints[j] = ints[j], ints[i]
}

func (ints mInts) Len() int {
	return len(ints)
}

func largestNumber(nums []int) string {
	sort.Sort(mInts(nums))

	result := &strings.Builder{}
	for i := 0; i < len(nums); i++ {
		if result.Len() == 0 && nums[i] == 0 {
			continue
		}
		result.WriteString(strconv.Itoa(nums[i]))
	}

	if result.Len() == 0 {
		result.WriteString("0")
	}
	return result.String()
}
