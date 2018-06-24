package zigzag_level

import "github.com/MaxnSter/GolangDataStructure/leetcode/queue"

/*
103. 二叉树的锯齿形层次遍历

给定一个二叉树，返回其节点值的锯齿形层次遍历。（即先从左往右，再从右往左进行下一层遍历，以此类推，层与层之间交替进行）。

例如：
给定二叉树 [3,9,20,null,null,15,7],

    3
   / \
  9  20
    /  \
   15   7
返回锯齿形层次遍历如下：

[
  [3],
  [20,9],
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

func zigzagLevelOrder(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}

	q, result := queue.NewQueue(), make([][]int, 0)
	q.PushBack(root)
	count := 1

	for !q.Empty() {
		level := make([]int, 0)
		s := q.Size()

		for i := 0; i < s; i++ {
			node := q.Peek().(*TreeNode)
			q.PopFront()

			level = append(level, node.Val)
			if node.Left != nil {
				q.PushBack(node.Left)
			}
			if node.Right != nil {
				q.PushBack(node.Right)
			}
		}

		count++
		if count%2 == 1 {
			reverse(level)
		}
		result = append(result, level)
	}
	return result
}

func reverse(nums []int) {
	b, e := 0, len(nums)-1
	for b < e {
		nums[b], nums[e] = nums[e], nums[b]
		b++
		e--
	}
}
