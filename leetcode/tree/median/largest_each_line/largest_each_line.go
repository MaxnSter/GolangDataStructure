package largest_each_line

import (
	"math"

	"github.com/MaxnSter/GolangDataStructure/leetcode/queue"
)

/*
515. 在每个树行中找最大值
您需要在二叉树的每一行中找到最大的值。

示例：

输入:

          1
         / \
        3   2
       / \   \
      5   3   9

输出: [1, 3, 9]
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

func largestValues(root *TreeNode) []int {
	if root == nil {
		return nil
	}

	result := make([]int, 0)
	q := queue.NewQueue()
	q.PushBack(root)

	for !q.Empty() {
		size, max := q.Size(), math.MinInt64

		for i := 0; i < size; i++ {
			node := q.Peek().(*TreeNode)
			q.PopFront()

			if node.Val > max {
				max = node.Val
			}

			if node.Left != nil {
				q.PushBack(node.Left)
			}
			if node.Right != nil {
				q.PushBack(node.Right)
			}
		}

		result = append(result, max)
	}
	return result
}
