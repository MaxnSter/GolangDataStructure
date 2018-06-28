package second_min

import "math"

/*
671. 二叉树中第二小的节点
给定一个非空特殊的二叉树，每个节点都是正数，并且每个节点的子节点数量只能为 2 或 0。
如果一个节点有两个子节点的话，那么这个节点的值不大于它的子节点的值。

给出这样的一个二叉树，你需要输出所有节点中的第二小的值。如果第二小的值不存在的话，输出 -1 。

示例 1:

输入:
    2
   / \
  2   5
     / \
    5   7

输出: 5
说明: 最小的值是 2 ，第二小的值是 5 。
示例 2:

输入:
    2
   / \
  2   2

输出: -1
说明: 最小的值是 2, 但是不存在第二小的值。
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

func findSecondMinimumValue(root *TreeNode) int {
	if root == nil || root.Left == nil && root.Right == nil {
		return -1
	}

	min := math.MaxInt64
	traverse(root, math.MinInt64, &min)

	secondMin := math.MaxInt64
	traverse(root, min, &secondMin)

	if secondMin == math.MaxInt64 {
		return -1
	} else {
		return secondMin
	}
}

func traverse(root *TreeNode, min int, secondMin *int) {
	if root == nil {
		return
	}

	if root.Val > min && root.Val < *secondMin {
		*secondMin = root.Val
	}
	traverse(root.Left, min, secondMin)
	traverse(root.Right, min, secondMin)
}
