package lsw

import "strings"

/*
3. 无重复字符的最长子串
给定一个字符串，找出不含有重复字符的最长子串的长度。

示例：

给定 "abcabcbb" ，没有重复字符的最长子串是 "abc" ，那么长度就是3。

给定 "bbbbb" ，最长的子串就是 "b" ，长度是1。

给定 "pwwkew" ，最长子串是 "wke" ，长度是3。请注意答案必须是一个子串，"pwke" 是 子序列  而不是子串。
*/
func lengthOfLongestSubstringEffective(s string) int {
	last, left, maxCount := make([]int, 256), 0, 0
	for i := range last {
		last[i] = -1
	}

	for i := 0; i < len(s); i++ {
		if last[s[i]] >= left {
			left = last[s[i]] + 1
		}

		last[s[i]] = i
		//一直记录最大长度
		maxCount = maxInt(maxCount, i-left+1)
	}

	return maxCount
}

//滑动窗口
func lengthOfLongestSubstringWindow(s string) int {
	set := map[byte]struct{}{}
	iB, iE, maxCount := 0, 0, 0
	for iB < len(s) && iE < len(s) {
		if _, ok := set[s[iE]]; !ok {
			set[s[iE]] = struct{}{}
			iE++
			maxCount = maxInt(maxCount, iE-iB)
		} else {
			delete(set, s[iB])
			iB++
		}
	}

	return maxCount
}

func lengthOfLongestSubstring(s string) int {
	if len(s) <= 1 || s == "" {
		return 0
	}

	maxSubCount, curSubCount := 1, 1
	for i := 0; i < len(s); i++ {
		for idx := i + 1; idx < len(s); idx++ {
			curS, curSub := s[i:idx], s[idx:idx+1]
			//fmt.Printf("curS:%s, curSub:%s\n", curS, curSub)
			if strings.Contains(curS, curSub) {
				break
			} else {
				curSubCount++
			}
		}
		maxSubCount = maxInt(maxSubCount, curSubCount)
		curSubCount = 1
	}

	return maxInt(maxSubCount, curSubCount)
}

func maxInt(x, y int) int {
	if x > y {
		return x
	}
	return y
}
