package sum

import (
	"strconv"
	"strings"

	"github.com/MaxnSter/GolangDataStructure/leetcode/stack"
)

/*
415. 字符串相加
给定两个字符串形式的非负整数 num1 和num2 ，计算它们的和。

注意：

num1 和num2 的长度都小于 5100.
num1 和num2 都只包含数字 0-9.
num1 和num2 都不包含任何前导零。
你不能使用任何內建 BigInteger 库， 也不能直接将输入的字符串转换为整数形式。
*/
//注意进位就好
func addStrings(num1 string, num2 string) string {
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
