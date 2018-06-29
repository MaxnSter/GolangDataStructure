package prefix

import "bytes"

/*
14. 最长公共前缀
编写一个函数来查找字符串数组中的最长公共前缀。

如果不存在公共前缀，返回空字符串 ""。

示例 1:

输入: ["flower","flow","flight"]
输出: "fl"
示例 2:

输入: ["dog","racecar","car"]
输出: ""
解释: 输入不存在公共前缀。
说明:

所有输入只包含小写字母 a-z 。
*/
func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}

	curPrefix := &bytes.Buffer{}
	for i := 0; i < len(strs[0]); i++ {
		if allContainsAt(strs, i, strs[0][i]) {
			curPrefix.WriteByte(strs[0][i])
		} else {
			break
		}
	}
	return curPrefix.String()
}

func allContainsAt(strs []string, idx int, sub byte) bool {
	for _, str := range strs {
		if len(str) <= idx || str[idx] != sub {
			return false
		}
	}
	return true
}
