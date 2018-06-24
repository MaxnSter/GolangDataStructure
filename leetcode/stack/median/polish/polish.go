package polish

import (
	"strconv"

	"github.com/MaxnSter/GolangDataStructure/leetcode/stack"
)

/*
150. 逆波兰表达式求值
根据逆波兰表示法，求表达式的值。

有效的运算符包括 +, -, *, / 。每个运算对象可以是整数，也可以是另一个逆波兰表达式。

说明：

整数除法只保留整数部分。
给定逆波兰表达式总是有效的。换句话说，表达式总会得出有效数值且不存在除数为 0 的情况。
示例 1：

输入: ["2", "1", "+", "3", "*"]
输出: 9
解释: ((2 + 1) * 3) = 9
示例 2：

输入: ["4", "13", "5", "/", "+"]
输出: 6
解释: (4 + (13 / 5)) = 6
示例 3：

输入: ["10", "6", "9", "3", "+", "-11", "*", "/", "*", "17", "+", "5", "+"]
输出: 22
解释:
  ((10 * (6 / ((9 + 3) * -11))) + 17) + 5
= ((10 * (6 / (12 * -11))) + 17) + 5
= ((10 * (6 / -132)) + 17) + 5
= ((10 * 0) + 17) + 5
= (0 + 17) + 5
= 17 + 5
= 22
*/

var funcs = map[string]func(int, int) int{
	"+": func(i int, i2 int) int {
		return i + i2
	},

	"-": func(i int, i2 int) int {
		return i - i2
	},

	"*": func(i int, i2 int) int {
		return i * i2
	},

	"/": func(i int, i2 int) int {
		return i / i2
	},
}

func isOperator(token string) bool {
	_, ok := funcs[token]
	return ok
}

//这个还是很简单的,只是要注意除法和减法的被操作数
//栈顶是操作数,栈底是被操作数
func evalRPN(tokens []string) int {
	s := stack.NewStack()
	for _, token := range tokens {
		if isOperator(token) {
			op2 := s.Top().(int)
			s.Pop()

			op1 := s.Top().(int)
			s.Pop()

			s.Push(funcs[token](op1, op2))
		} else {
			num, err := strconv.Atoi(token)
			if err != nil {
				panic(err)
			}
			s.Push(num)
		}
	}

	return s.Top().(int)
}
