package decode_string

import (
	"strings"

	"github.com/MaxnSter/GolangDataStructure/leetcode/stack"
)

/*
394. 字符串解码
给定一个经过编码的字符串，返回它解码后的字符串。

编码规则为: k[encoded_string]，表示其中方括号内部的 encoded_string 正好重复 k 次。注意 k 保证为正整数。

你可以认为输入字符串总是有效的；输入字符串中没有额外的空格，且输入的方括号总是符合格式要求的。

此外，你可以认为原始数据不包含数字，所有的数字只表示重复的次数 k ，例如不会出现像 3a 或 2[4] 的输入。

示例:

s = "3[a]2[bc]", 返回 "aaabcbc".
s = "3[a2[c]]", 返回 "accaccacc".
s = "2[abc]3[cd]ef", 返回 "abcabccdcdcdef".
*/

//栈图
//       numStack        strs         result   digits
//3  -> |       |       |   |          ""         3
//
//       numStack       strs          result   digits
//[ ->  |   3   |      | "" |          ""         3
//
//       numStack       strs          result   digits
//a ->  |   3   |      | "" |          "a"        0
//
//       numStack       strs          result   digits
//2 ->  |   3   |      | ""|           "a"        2
//
//
//       numStack       strs          result   digits
//[ ->  |   3   |      | "" |           ""        0
//      |   2   |      | "a"|
//
//       numStack       strs          result   digits
//c ->  |   3   |      | "" |           "c"       0
//      |   2   |      | "a"|
//
//
//       numStack       strs          result   digits
//] ->  |   3   |      | "" |           "acc"       0
//
//
//       numStack       strs          result            digits
//] ->  |   0   |      | "" |         "accaccacc"       0

func decodeString(s string) string {
	result := &strings.Builder{}
	nums := stack.NewStack()
	strs := stack.NewStack()
	digits := 0

	for i := 0; i < len(s); i++ {
		if isNumber(s[i]) {
			digits = digits*10 + int(s[i]-'0')
			continue
		}

		if s[i] == '[' {
			nums.Push(digits)
			digits = 0

			strs.Push(result)
			result = &strings.Builder{}
			continue
		}

		if s[i] == ']' {
			num := nums.Top().(int)
			nums.Pop()

			for i := 0; i < num; i++ {
				strs.Top().(*strings.Builder).WriteString(result.String())
			}
			result = strs.Top().(*strings.Builder)
			strs.Pop()
		}

		if isAlpha(s[i]) {
			result.WriteByte(s[i])
			continue
		}
	}

	return result.String()
}

func isNumber(b byte) bool {
	return b >= '0' && b <= '9'
}

func isAlpha(b byte) bool {
	return (b >= 'a' && b <= 'z') ||
		(b >= 'A' && b <= 'Z')
}

func decodeString2(s string) string {
	numStk, strStk := stack.NewStack(), stack.NewStack()
	result := &strings.Builder{}
	digit := 0

	for i := 0; i < len(s); i++ {
		if isNumber(s[i]) {
			digit = digit*10 + int(s[i]-'0')
			continue
		}

		if isAlpha(s[i]) {
			result.WriteByte(s[i])
			continue
		}

		if s[i] == '[' {
			numStk.Push(digit)
			digit = 0

			strStk.Push(result)
			result = &strings.Builder{}
			continue
		}

		if s[i] == ']' {
			curDigit := numStk.Top().(int)
			numStk.Pop()

			for i := 0; i < curDigit; i++ {
				strStk.Top().(*strings.Builder).WriteString(result.String())
			}

			result = strStk.Top().(*strings.Builder)
			strStk.Pop()
		}
	}

	return result.String()
}
