package path_sum_3

/*
437. 路径总和 III
给定一个二叉树，它的每个结点都存放着一个整数值。

找出路径和等于给定数值的路径总数。

路径不需要从根节点开始，也不需要在叶子节点结束，但是路径方向必须是向下的（只能从父节点到子节点）。

二叉树不超过1000个节点，且节点数值范围是 [-1000000,1000000] 的整数。

示例：

root = [10,5,-3,3,2,null,11,3,-2,null,1], sum = 8

      10
     /  \
    5   -3
   / \    \
  3   2   11
 / \   \
3  -2   1

返回 3。和等于 8 的路径有:

1.  5 -> 3
2.  5 -> 2 -> 1
3.  -3 -> 11
*/
/**
* Definition for a binary tree node.
 */
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func pathSum(root *TreeNode, sum int) int {
	var count int
	dfs1(root, sum, &count)
	return count
}

func dfs1(root *TreeNode, sum int, count *int) {
	if root == nil {
		return
	}

	dfs(root, sum, count)
	dfs1(root.Left, sum, count)
	dfs1(root.Right, sum, count)
}

func dfs(root *TreeNode, sum int, count *int) {
	if root == nil {
		return
	}

	sum -= root.Val
	if sum == 0 {
		*count += 1
	}

	dfs(root.Left, sum, count)
	dfs(root.Right, sum, count)
}