package parenthese

import "github.com/MaxnSter/GolangDataStructure/leetcode/stack"

/*
20. 有效的括号
给定一个只包括 '('，')'，'{'，'}'，'['，']' 的字符串，判断字符串是否有效。

有效字符串需满足：

左括号必须用相同类型的右括号闭合。
左括号必须以正确的顺序闭合。
注意空字符串可被认为是有效字符串。

示例 1:

输入: "()"
输出: true
示例 2:

输入: "()[]{}"
输出: true
示例 3:

输入: "(]"
输出: false
示例 4:

输入: "([)]"
输出: false
示例 5:

输入: "{[]}"
输出: true
*/
//思路:如果是左括号,直接入栈,如果是右括号,判断栈顶元素是否和自己对应
func isValid(s string) bool {
	if s == "" {
		return true
	}

	stk := stack.NewStack()
	for i := 0; i < len(s); i++ {
		if s[i] == '(' || s[i] == '[' || s[i] == '{' {
			stk.Push(s[i])
			continue
		}

		if v, ok := expected[s[i]]; !ok || stk.Empty() || stk.Top().(byte) != v {
			return false
		}

		stk.Pop()
	}

	if stk.Empty() {
		return true
	} else {
		return false
	}
}

var expected = map[byte]byte{
	')': '(',
	']': '[',
	'}': '{',
}
