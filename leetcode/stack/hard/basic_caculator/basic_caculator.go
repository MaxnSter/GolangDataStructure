package basic_caculator

import (
	"strconv"
	"unicode"

	"github.com/MaxnSter/GolangDataStructure/leetcode/stack"
)

/*
224. 基本计算器
实现一个基本的计算器来计算一个简单的字符串表达式的值。

字符串表达式可以包含左括号 ( ，右括号 )，加号 + ，减号 -，非负整数和空格  。

示例 1:

输入: "1 + 1"
输出: 2
示例 2:

输入: " 2-1 + 2 "
输出: 3
示例 3:

输入: "(1+(4+5+2)-3)+(6+8)"
输出: 23
*/
func calculate(s string) int {
	token := make([]string, 0)
	for i := 0; i < len(s); {
		if unicode.IsNumber(rune(s[i])) {
			digitBegin := i
			for i < len(s) && unicode.IsNumber(rune(s[i])) {
				i++
			}

			token = append(token, s[digitBegin:i])
			continue
		}

		if unicode.IsSpace(rune(s[i])) {
			i++
			continue
		}

		token = append(token, string(s[i]))
		i++
	}

	convert := convertTo(token)
	return evalRPN(convert)
}

func convertTo(s []string) []string {
	output := make([]string, 0)
	operatorStack := stack.NewStack()

	for _, ss := range s {
		if !isOperator(ss) {
			output = append(output, ss)
			continue
		}

		if operatorStack.Empty() {
			operatorStack.Push(ss)
			continue
		}

		if ss == ")" {
			//弹出元素直到")"
			for operatorStack.Top().(string) != "(" {
				op := operatorStack.Top().(string)
				output = append(output, op)
				operatorStack.Pop()
			}

			//把"("弹出,不记录
			operatorStack.Pop()
			continue
		}

		if ss == "(" {
			//遇到左括号,直接入栈
			operatorStack.Push(ss)
			continue
		}

		//弹出所有元素,直到遇到比自己优先级低的
		//也就意味着,如果栈不为空并且栈顶元素优先级大于等于当前优先级,就一直弹出
		for !operatorStack.Empty() && priority[ss] <= priority[operatorStack.Top().(string)] {
			op := operatorStack.Top().(string)
			output = append(output, op)
			operatorStack.Pop()
		}

		//然后自己入栈
		operatorStack.Push(ss)

	}

	for !operatorStack.Empty() {
		op := operatorStack.Top().(string)
		output = append(output, op)
		operatorStack.Pop()
	}

	return output
}

func isOperator(s string) bool {
	return s == "*" || s == "-" || s == "+" || s == "/" || s == "(" || s == ")"
}

var priority = map[string]int{
	"*": 2,
	"/": 2,
	"+": 1,
	"-": 1,
	"(": 3,
}

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
