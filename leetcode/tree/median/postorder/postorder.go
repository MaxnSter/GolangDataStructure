package postorder

import "github.com/MaxnSter/GolangDataStructure/leetcode/stack"

/*
145. 二叉树的后序遍历
给定一个二叉树，返回它的 后序 遍历。

示例:

输入: [1,null,2,3]
   1
    \
     2
    /
   3

输出: [3,2,1]
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

func postorderTraversal(root *TreeNode) []int {
	result := make([]int, 0)
	traversal(root, &result)
	return result
}

func traversal(root *TreeNode, result *[]int) {
	if root == nil {
		return
	}

	traversal(root.Left, result)
	traversal(root.Right, result)
	*result = append(*result, root.Val)
}

//非递归,用两个栈实现
func postorderTraversalNoRecursive(root *TreeNode) []int {
	if root == nil {
		return nil
	}

	s1, s2, result := stack.NewStack(), stack.NewStack(), make([]int, 0)
	s1.Push(root)

	for !s1.Empty() {
		node := s1.Top().(*TreeNode)
		s1.Pop()

		s2.Push(node)

		if node.Left != nil {
			s1.Push(node.Left)
		}
		if node.Right != nil {
			s1.Push(node.Right)
		}
	}

	for !s2.Empty() {
		result = append(result, s2.Top().(int))
		s2.Pop()
	}
	return result
}
