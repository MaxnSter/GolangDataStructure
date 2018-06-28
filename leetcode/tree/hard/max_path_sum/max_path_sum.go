package max_path_sum

import "math"

/*
124. 二叉树中的最大路径和
给定一个非空二叉树，返回其最大路径和。

本题中，路径被定义为一条从树中任意节点出发，达到任意节点的序列。该路径至少包含一个节点，且不一定经过根节点。

示例 1:

输入: [1,2,3]

       1
      / \
     2   3

输出: 6
示例 2:

输入: [-10,9,20,null,null,15,7]

   -10
   / \
  9  20
    /  \
   15   7

输出: 42
*/

/**
 * Definition for a binary tree node.
 */
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxPathSum(root *TreeNode) int {
	max := math.MinInt64
	helper(root, &max)
	return max
}

func helper(root *TreeNode, max *int) int {
	if root == nil {
		return 0
	}

	//自底向上收集
	l := maxInt(helper(root.Left, max), 0)
	r := maxInt(helper(root.Right, max), 0)

	sum := maxInt(l+root.Val, r+root.Val)
	*max = maxInt(*max, l+r+root.Val)

	//当前结点路径下的最大值
	return sum
}

func maxInt(x, y int) int {
	if x > y {
		return x
	}
	return y
}
