package lps

/*
5. 最长回文子串
给定一个字符串 s，找到 s 中最长的回文子串。你可以假设 s 的最大长度为1000。

示例 1：

输入: "babad"
输出: "bab"
注意: "aba"也是一个有效答案。
示例 2：

输入: "cbbd"
输出: "bb"
*/

//马拉车先放弃了...
func longestPalindrome3(s string) string {

}

//中心节点扩散
func longestPalindrome2(s string) string {
	start, end := 0, 0
	for i := 0; i < len(s); i++ {
		l1 := palindrome(s, i, i)
		l2 := palindrome(s, i, i+1) //回文数为偶数
		l := maxInt(l1, l2)

		if l > end-start {
			start = i - (l-1)/2 //奇数,偶数情况
			end = i + l/2
		}
	}

	return s[start : end+1]
}

func maxInt(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func palindrome(s string, left, r int) int {
	for left >= 0 && r < len(s) && s[left] == s[r] {
		left--
		r++
	}

	return r - left - 1
}

//暴力法
func longestPalindrome(s string) string {
	if len(s) <= 1 {
		return s
	}

	maxStr, maxCount := "", 0
	for i := 0; i < len(s); i++ {
		for j := i; j < len(s); j++ {
			count := palidromeLen(s[i : j+1])
			if count != -1 && count > maxCount {
				maxCount = count
				maxStr = s[i : j+1]
			}
		}
	}

	return maxStr
}

func palidromeLen(s string) int {
	iB, iE := 0, len(s)-1
	for iB < iE {
		if s[iB] != s[iE] {
			return -1
		}
		iB++
		iE--
	}

	return len(s)
}
