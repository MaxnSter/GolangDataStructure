package last_word

/*
58. 最后一个单词的长度
给定一个仅包含大小写字母和空格 ' ' 的字符串，返回其最后一个单词的长度。

如果不存在最后一个单词，请返回 0 。

说明：一个单词是指由字母组成，但不包含任何空格的字符串。

示例:

输入: "Hello World"
输出: 5
*/

func lengthOfLastWord(s string) int {
	lIdx, length := len(s)-1, 0
	for ; lIdx >= 0 && s[lIdx] == ' '; lIdx-- {
	}

	if lIdx == 0 && s[lIdx] == ' ' {
		return 0
	}

	for ; lIdx >= 0 && s[lIdx] != ' '; lIdx-- {
		length++
	}

	return length
}
