package inorder

import "github.com/MaxnSter/GolangDataStructure/leetcode/stack"

/*
94. 二叉树的中序遍历
给定一个二叉树，返回它的中序 遍历。

示例:

输入: [1,null,2,3]
   1
    \
     2
    /
   3

输出: [1,3,2]
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

func inorderTraversal(root *TreeNode) []int {
	result := make([]int, 0)
	traversal(root, &result)
	return result
}

func traversal(root *TreeNode, result *[]int) {
	if root == nil {
		return
	}

	traversal(root.Left, result)
	*result = append(*result, root.Val)
	traversal(root.Right, result)
}

func inorderTraversalNoRecursive(root *TreeNode) []int {
	if root == nil {
		return nil
	}

	result, s := make([]int, 0), stack.NewStack()
	node := root
	for node != nil || !s.Empty() {
		for node != nil {
			s.Push(node)
			node = node.Left
		}

		if !s.Empty() {
			node = s.Top().(*TreeNode)
			s.Pop()
			result = append(result, node.Val)
			node = node.Right
		}
	}

	return result
}
