package path_sum

/*
112. 路径总和
给定一个二叉树和一个目标和，判断该树中是否存在根节点到叶子节点的路径，这条路径上所有节点值相加等于目标和。

说明: 叶子节点是指没有子节点的节点。

示例:
给定如下二叉树，以及目标和 sum = 22，

              5
             / \
            4   8
           /   / \
          11  13  4
         /  \      \
        7    2      1
返回 true, 因为存在目标和为 22 的根节点到叶子节点的路径 5->4->11->2。
*/
/**
* Definition for a binary tree node.
 */
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func hasPathSum(root *TreeNode, sum int) bool {
	if root == nil {
		return false
	}

	if root.Left == nil && root.Right == nil {
		return root.Val == sum
	}

	return hasPathSum(root.Left, sum-root.Val) ||
		hasPathSum(root.Right, sum-root.Val)

}

func chkPathSum(root *TreeNode, pathSum int, sum int, result *bool) {
	if root == nil {
		return
	}

	//注意1,题目说的是叶子节点
	if root.Left == nil && root.Right == nil {
		//注意2,叶子节点计算path值也要加上自己的val
		if pathSum+root.Val == sum {
			*result = true
		}
		return
	}

	chkPathSum(root.Left, root.Val+pathSum, sum, result)
	chkPathSum(root.Right, root.Val+pathSum, sum, result)
}
