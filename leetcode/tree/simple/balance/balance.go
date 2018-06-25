package balance

/*
110. 平衡二叉树
给定一个二叉树，判断它是否是高度平衡的二叉树。

本题中，一棵高度平衡二叉树定义为：

一个二叉树每个节点 的左右两个子树的高度差的绝对值不超过1。
示例 1:

给定二叉树 [3,9,20,null,null,15,7]

    3
   / \
  9  20
    /  \
   15   7
返回 true 。

示例 2:

给定二叉树 [1,2,2,3,3,null,null,4,4]

       1
      / \
     2   2
    / \
   3   3
  / \
 4   4
返回 false 。
*/
/**
* Definition for a binary tree node.
 */
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}

	if absInt(maxDepth(root.Left) - maxDepth(root.Right)) > 1 {
		return false
	} else {
		return isBalanced(root.Left) && isBalanced(root.Right)
	}
}

func maxDepth(root *TreeNode) int {
	if root == nil {
		return  0
	}

	if root.Left == nil && root.Right == nil {
		return 1
	}

	return 1 + maxInt(maxDepth(root.Left), maxDepth(root.Right))
}

func maxInt(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func absInt(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
