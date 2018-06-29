package mult_strings

import (
	"strconv"
	"strings"

	"github.com/MaxnSter/GolangDataStructure/leetcode/stack"
)

/*
43. 字符串相乘
给定两个以字符串形式表示的非负整数 num1 和 num2，返回 num1 和 num2 的乘积，它们的乘积也表示为字符串形式。

示例 1:

输入: num1 = "2", num2 = "3"
输出: "6"
示例 2:

输入: num1 = "123", num2 = "456"
输出: "56088"
说明：

num1 和 num2 的长度小于110。
num1 和 num2 只包含数字 0-9。
num1 和 num2 均不以零开头，除非是数字 0 本身。
不能使用任何标准库的大数类型（比如 BigInteger）或直接将输入转换为整数来处理。
*/
//处理好进位就好,可以简化很多,只是懒得写了
func multiply(num1 string, num2 string) string {
	if num1 == "0" || num2 == "0" {
		return "0"
	}

	curMult, allMults := &strings.Builder{}, make([]string, 0)

	for idx1 := len(num1) - 1; idx1 >= 0; idx1-- {
		result, carry := stack.NewStack(), 0

		for i := 0; i < len(num1)-1-idx1; i++ {
			result.Push(0)
		}

		for idx2 := len(num2) - 1; idx2 >= 0; idx2-- {
			mult := int((num1[idx1]-'0')*(num2[idx2]-'0')) + carry
			result.Push(mult % 10)
			carry = mult / 10
		}

		if carry != 0 {
			result.Push(carry)
		}

		for !result.Empty() {
			node := strconv.Itoa(result.Top().(int))
			result.Pop()

			curMult.WriteString(node)
		}

		allMults = append(allMults, curMult.String())
		curMult.Reset()
	}

	result := "0"
	for _, mult := range allMults {
		result = sum(mult, result)
	}

	return result
}

func sum(num1 string, num2 string) string {
	result, carry := stack.NewStack(), 0
	idx1, idx2 := len(num1)-1, len(num2)-1

	for idx1 >= 0 && idx2 >= 0 {
		sum := int(num1[idx1]-'0'+num2[idx2]-'0') + carry
		carry = sum / 10
		result.Push(sum % 10)

		idx1--
		idx2--
	}

	for idx1 >= 0 {
		sum := int(num1[idx1]-'0') + carry
		carry = sum / 10
		result.Push(sum % 10)

		idx1--
	}

	for idx2 >= 0 {
		sum := int(num2[idx2]-'0') + carry
		carry = sum / 10
		result.Push(sum % 10)

		idx2--
	}

	if carry > 0 {
		result.Push(carry)
	}

	sum := &strings.Builder{}
	for !result.Empty() {
		sum.WriteString(strconv.Itoa(result.Top().(int)))
		result.Pop()
	}

	return sum.String()
}
