package unique

/*
96. 不同的二叉搜索树
给定一个整数 n，求以 1 ... n 为节点组成的二叉搜索树有多少种？

示例:

输入: 3
输出: 5
解释:
给定 n = 3, 一共有 5 种不同结构的二叉搜索树:

   1         3     3      2      1
    \       /     /      / \      \
     3     2     1      1   3      2
    /     /       \                 \
   2     1         2                 3
*/
/*
G(n)=∑ni=1(G(i−1)∗G(n−i))G(n)=∑i=1n(G(i−1)∗G(n−i))
需要对G(0)和G(1)特殊处理，令其为1，即G(0)=G(1)=1

Count[i] = ∑ Count[0...k] * [ k+1....i]     0<=k<i-1
*/
func numTrees(n int) int {
	dp := make([]int, n+1)
	dp[0], dp[1] = 1, 1

	for i := 2; i <= n; i++ {
		for j := 1; j <= i; j++ {
			//左子树共有节点j-1,右子树共有节点i-j
			dp[i] += dp[j-1] * dp[i-j]
		}
	}

	return dp[n]
}
