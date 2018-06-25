package path_longest_same

/*
687. 最长同值路径
给定一个二叉树，找到最长的路径，这个路径中的每个节点具有相同值。 这条路径可以经过也可以不经过根节点。

注意：两个节点之间的路径长度由它们之间的边数表示。

示例 1:

输入:

              5
             / \
            4   5
           / \   \
          1   1   5
输出:

2
示例 2:

输入:

              1
             / \
            4   5
           / \   \
          4   4   5
输出:

2
注意: 给定的二叉树不超过10000个结点。 树的高度不超过1000。
*/
/**
* Definition for a binary tree node.
*
 */
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func longestUnivaluePath(root *TreeNode) int {
	var max int
	helper(root, &max)
	return max
}

func helper(root *TreeNode, max *int) int {
	if root == nil {
		return 0
	}

	l := helper(root.Left, max)
	r := helper(root.Right, max)

	left, right := 0, 0
	if root.Left != nil && root.Val == root.Left.Val {
		left = l + 1
	}

	if root.Right != nil && root.Val == root.Right.Val {
		right = r + 1
	}

	*max = maxInt(*max, left+right)
	return maxInt(left, right)
}

func maxInt(x, y int) int {
	if x > y {
		return x
	}
	return y
}
