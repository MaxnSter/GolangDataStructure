package level2

import "github.com/MaxnSter/GolangDataStructure/leetcode/queue"

/*
107. 二叉树的层次遍历 II
给定一个二叉树，返回其节点值自底向上的层次遍历。 （即按从叶子节点所在层到根节点所在的层，逐层从左向右遍历）

例如：
给定二叉树 [3,9,20,null,null,15,7],

    3
   / \
  9  20
    /  \
   15   7
返回其自底向上的层次遍历为：

[
  [15,7],
  [9,20],
  [3]
]
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

func levelOrderBottom(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}

	result := make([][]int, 0)
	q := queue.NewQueue()
	q.PushBack(root)

	for !q.Empty() {
		nums, size := make([]int, 0), q.Size()
		for i := 0; i < size; i++ {
			node := q.Peek().(*TreeNode)
			q.PopFront()

			nums = append(nums, node.Val)
			if node.Left != nil {
				q.PushBack(node.Left)
			}
			if node.Right != nil {
				q.PushBack(node.Right)
			}
		}
		result = append(result, nums)
	}

	iB, iE := 0, len(result)-1
	for iB < iE {
		result[iB], result[iE] = result[iE], result[iB]
		iB++
		iE--
	}
	return result
}
