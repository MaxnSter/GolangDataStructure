package leaf_nums

import "fmt"

/*
129. 求根到叶子节点数字之和
给定一个二叉树，它的每个结点都存放一个 0-9 的数字，每条从根到叶子节点的路径都代表一个数字。

例如，从根到叶子节点路径 1->2->3 代表数字 123。

计算从根到叶子节点生成的所有数字之和。

说明: 叶子节点是指没有子节点的节点。

示例 1:

输入: [1,2,3]
    1
   / \
  2   3
输出: 25
解释:
从根到叶子节点路径 1->2 代表数字 12.
从根到叶子节点路径 1->3 代表数字 13.
因此，数字总和 = 12 + 13 = 25.
示例 2:

输入: [4,9,0,5,1]
    4
   / \
  9   0
 / \
5   1
输出: 1026
解释:
从根到叶子节点路径 4->9->5 代表数字 495.
从根到叶子节点路径 4->9->1 代表数字 491.
从根到叶子节点路径 4->0 代表数字 40.
因此，数字总和 = 495 + 491 + 40 = 1026.
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

func sumNumbers(root *TreeNode) int {
	curSum, numbers := 0, make([]int, 0)
	pathTree(root, curSum, &numbers)

	curSum = 0
	for _, v := range numbers {
		curSum += v
	}
	return curSum
}

func pathTree(root *TreeNode, curSum int, numbers *[]int) {
	if root == nil {
		return
	}

	curSum = curSum*10 + root.Val

	if root.Left == nil && root.Right == nil {
		//我的base case, 也就是只有一个节点时,直接append
		//因此,我在递归过程中,也应该在root.Left == nil && root.Right == nil时append,而不是
		//在root==nil时append
		*numbers = append(*numbers, curSum)
	}
	pathTree(root.Left, curSum, numbers)
	pathTree(root.Right, curSum, numbers)
}
