package count_ct

import "github.com/MaxnSter/GolangDataStructure/leetcode/queue"

/*
222. 完全二叉树的节点个数
给出一个完全二叉树，求出该树的节点个数。

说明：

完全二叉树的定义如下：在完全二叉树中，除了最底层节点可能没填满外，其余每层节点数都达到最大值，并且最下面一层的节点都集中在该层最左边的若干位置。若最底层为第 h 层，则该层包含 1~ 2h 个节点。

示例:

输入:
    1
   / \
  2   3
 / \  /
4  5 6

输出: 6
*/
/**
* Definition for a binary tree node.
 */
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func countNodesEffective(root *TreeNode) int {
	if root == nil {
		return 0
	}

	//利用完全二叉树的性质
	depL, depR := depth(root.Left, true), depth(root.Right, false)
	if depL == depR {
		return (1 << uint(depL)) - 1
	} else {
		//此时,可以分解为n多个完全二叉树
		return 1 + countNodes(root.Left) + countNodes(root.Right)
	}
}

func depth(root *TreeNode, isLeft bool) (count int) {
	if root == nil {
		return 0
	}

	for root != nil {
		count++

		if isLeft {
			root = root.Left
		} else {
			root = root.Right
		}
	}
	return
}

func countNodes(root *TreeNode) int {
	if root == nil {
		return 0
	}

	q, count := queue.NewQueue(), 0
	q.PushBack(root)

	for !q.Empty() {
		size := q.Size()

		for i := 0; i < size; i++ {
			node := q.Peek().(*TreeNode)
			q.PopFront()
			count++

			if node.Left != nil {
				q.PushBack(node.Left)
			}

			if node.Right != nil {
				q.PushBack(node.Right)
			}
		}
	}
	return count
}
