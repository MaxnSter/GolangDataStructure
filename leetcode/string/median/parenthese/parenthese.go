package parenthese

/*
22. 括号生成
给出 n 代表生成括号的对数，请你写出一个函数，使其能够生成所有可能的并且有效的括号组合。

例如，给出 n = 3，生成结果为：

[
  "((()))",
  "(()())",
  "(())()",
  "()(())",
  "()()()"
]
*/
func generateParenthesis(n int) []string {
	if n < 1 {
		return nil
	}

	result := make([]string, 0)
	gen(n, 0, 0, "", &result)
	return result
}

// n:需要多少个括号
// Left:左括号有多少个, right:同理
func gen(n, left, right int, path string, result *[]string) {
	//以n为1举例,考虑这个base case
	if left == n && right == n {
		*result = append(*result, path)
		return
	}

	//以n为2举例,考虑怎么生成组合
	if left < n {
		gen(n, left+1, right, path+"(", result)
	}

	//right < n wrong,因为这样生成的括号包含无效的
	if right < left {
		gen(n, left, right+1, path+")", result)
	}
}
