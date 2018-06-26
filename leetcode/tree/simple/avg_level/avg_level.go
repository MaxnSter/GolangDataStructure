package avg_level

import "github.com/MaxnSter/GolangDataStructure/leetcode/queue"

/*
637. 二叉树的层平均值
给定一个非空二叉树, 返回一个由每层节点平均值组成的数组.

示例 1:

输入:
    3
   / \
  9  20
    /  \
   15   7
输出: [3, 14.5, 11]
解释:
第0层的平均值是 3,  第1层是 14.5, 第2层是 11. 因此返回 [3, 14.5, 11].
注意：

节点值的范围在32位有符号整数范围内。
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

func averageOfLevels(root *TreeNode) []float64 {
	if root == nil {
		return nil
	}

	result := make([]float64, 0)
	q := queue.NewQueue()
	q.PushBack(root)

	for !q.Empty() {
		sum, size := 0, q.Size()
		for i := 0; i < size; i++ {
			node := q.Peek().(*TreeNode)
			q.PopFront()

			sum += node.Val

			if node.Left != nil {
				q.PushBack(node.Left)
			}
			if node.Right != nil {
				q.PushBack(node.Right)
			}
		}
		result = append(result, float64(sum)/float64(size))
	}

	return result
}
