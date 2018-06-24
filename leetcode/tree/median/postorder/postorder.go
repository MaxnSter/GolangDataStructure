package postorder

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
