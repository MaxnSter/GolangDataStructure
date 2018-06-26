package validate_bst

import "math"

/*
98. 验证二叉搜索树
给定一个二叉树，判断其是否是一个有效的二叉搜索树。

一个二叉搜索树具有如下特征：

节点的左子树只包含小于当前节点的数。
节点的右子树只包含大于当前节点的数。
所有左子树和右子树自身必须也是二叉搜索树。
示例 1:

输入:
    2
   / \
  1   3
输出: true
示例 2:

输入:
    5
   / \
  1   4
     / \
    3   6
输出: false
解释: 输入为: [5,1,4,null,null,3,6]。
     根节点的值为 5 ，但是其右子节点值为 4 。
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

func isValidBST(root *TreeNode) bool {
	if root == nil || root.Left == nil && root.Right == nil {
		return true
	}

	return isValid(root, math.MaxInt64, math.MinInt64)
}

func isValid(root *TreeNode, max int, min int) bool {
	if root == nil {
		return true
	}

	//如果是bst,则应该: root.Val < max, root.val > min

	if root.Val <= min || root.Val >= max {
		return false
	}

	//要点:比左子树结点值大的数始终是路径上最小的
	//比右子树节点值小的数始终是路径上最大的
	return isValid(root.Left, MinInt(max, root.Val), min) &&
		isValid(root.Right, max, MaxInt(min, root.Val))
}

func MaxInt(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func MinInt(x, y int) int {
	if x < y {
		return x
	}
	return y
}
