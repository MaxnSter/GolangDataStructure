package twosum

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func findTarget(root *TreeNode, k int) bool {
	m := make(map[int]struct{})
	var result bool

	traverseTree(root, &m, k, &result)
	return result
}

func traverseTree(root *TreeNode, m *map[int]struct{}, k int, finded *bool) {
	if root == nil {
		return
	}

	complement := k - root.Val
	if _, ok := (*m)[complement]; ok {
		*finded = true
	}
	(*m)[root.Val] = struct{}{}
	traverseTree(root.Left, m, k, finded)
	traverseTree(root.Right, m, k, finded)
}
