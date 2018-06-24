package level

import "github.com/MaxnSter/GolangDataStructure/leetcode/queue"

/*
102. 二叉树的层次遍历
给定一个二叉树，返回其按层次遍历的节点值。 （即逐层地，从左到右访问所有节点）。

例如:
给定二叉树: [3,9,20,null,null,15,7],

    3
   / \
  9  20
    /  \
   15   7
返回其层次遍历结果：

[
  [3],
  [9,20],
  [15,7]
]

*/
/**
 * Definition for a binary tree node.
 */
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}

	q, result := queue.NewQueue(), make([][]int,0)
	q.PushBack(root)

	for !q.Empty() {
		level := make([]int, 0)
		s := q.Size()

		for i := 0; i < s; i++ {
			node := q.Peek().(*TreeNode)
			q.PopFront()

			level = append(level, node.Val)
			if node.Left != nil {q.PushBack(node.Left)}
			if node.Right != nil {q.PushBack(node.Right)}
		}

		result = append(result, level)
	}

	return result
}
