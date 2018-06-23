package rm_dup_letters

import "github.com/MaxnSter/GolangDataStructure/leetcode/stack"

/*
316. 去除重复字母
给定一个仅包含小写字母的字符串，去除字符串中重复的字母，使得每个字母只出现一次。需保证返回结果的字典序最小（要求不能打乱其他字符的相对位置）。

示例 1:

输入: "bcabc"
输出: "abc"
示例 2:

输入: "cbacdcbc"
输出: "acdb"
*/
func removeDuplicateLetters(s string) string {
	count := make([]int, 26)
	visited := make([]bool, 26)
	for i := 0; i < len(s); i++ {
		count[s[i]-'a']++
	}

	ss := stack.NewStack()
	for i := 0; i < len(s); i++ {
		count[s[i]-'a']--
		if visited[s[i]-'a'] {
			continue
		}

		//后面还有,可以放心丢弃这个元素!
		for !ss.Empty() &&
			s[i] < ss.Top().(byte) &&
			count[ss.Top().(byte)-'a'] > 0 {

			t := ss.Top().(byte)
			ss.Pop()
			visited[t-'a'] = false
		}

		ss.Push(s[i])
		visited[s[i]-'a'] = true
	}

	ss.Reverse()
	result := make([]byte, ss.Size())
	for !ss.Empty() {
		result = append(result, ss.Top().(byte))
		ss.Pop()
	}

	return string(result)
}
