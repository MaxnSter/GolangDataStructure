package sub_contain

import "strings"

/*
30. 与所有单词相关联的字串
给定一个字符串 s 和一些长度相同的单词 words。在 s 中找出可以恰好串联 words 中所有单词的子串的起始位置。

注意子串要与 words 中的单词完全匹配，中间不能有其他字符，但不需要考虑 words 中单词串联的顺序。

示例 1:

输入:
  s = "barfoothefoobarman",
  words = ["foo","bar"]
输出: [0,9]
解释: 从索引 0 和 9 开始的子串分别是 "barfoor" 和 "foobar" 。
输出的顺序不重要, [9,0] 也是有效答案。
示例 2:

输入:
  s = "wordgoodstudentgoodword",
  words = ["word","student"]
输出: []
*/
func findSubstring(s string, words []string) []int {
	if s == "" || len(words) <= 0{
		return nil
	}

	result := make([]int, 0)
	l := wordsLen(words)

	for i := 0; i+l <= len(s); i++ {
		if containAll(s[i:i+l], words) && sameAlphas(s[i:i+l], words) {
			result = append(result, i)
		}
	}
	return result
}

func wordsLen(words []string) (l int) {
	for _, word := range words {
		l += len(word)
	}
	return
}

func sameAlphas(str string, words []string) bool {
	m := map[byte]int{}
	for _, word := range words {
		for i := 0; i < len(word); i++ {
			m[word[i]]++
		}
	}

	for i := 0; i < len(str);i++ {
		if _, ok := m[str[i]]; !ok {
			return false
		} else {
			m[str[i]]--
		}
	}

	for _, count := range m {
		if count != 0 {
			return false
		}
	}
	return true
}

func containAll(str string, words []string) bool {
	for _, word := range words {
		if !strings.Contains(str, word) {
			return false
		}
	}

	return true
}
