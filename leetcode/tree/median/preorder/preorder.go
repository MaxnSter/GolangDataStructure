package preorder

import (
	"github.com/Arafatk/Dataviz/trees"
	"github.com/MaxnSter/GolangDataStructure/leetcode/stack"
)

/*
144. 二叉树的前序遍历
给定一个二叉树，返回它的 前序 遍历。

 示例:

输入: [1,null,2,3]
   1
    \
     2
    /
   3

输出: [1,2,3]
进阶: 递归算法很简单，你可以通过迭代算法完成吗？
*/

/**
* Definition for a binary tree node.
 */
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func preorderTraversal(root *TreeNode) []int {
	result := make([]int, 0)
	traversal(root, &result)
	return result
}

func traversal(root *TreeNode, result *[]int) {
	if root == nil {
		return
	}

	*result = append(*result, root.Val)
	traversal(root.Left, result)
	traversal(root.Right, result)
}

func inorderTraversalNoRecursive(root *TreeNode) []int {
	if root == nil {
		return nil
	}

	node := root
	result := make([]int, 0)
	s := stack.NewStack()

	for node != nil || !s.Empty() {
		for node != nil {
			result = append(result, node.Val)
			s.Push(node)
			node = node.Left
		}

		if !s.Empty() {
			node = s.Top().(*TreeNode)
			node = node.Right
			s.Pop()
		}
	}

	return result
}
