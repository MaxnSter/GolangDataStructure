package a_to_bst

/*
108. 将有序数组转换为二叉搜索树
将一个按照升序排列的有序数组，转换为一棵高度平衡二叉搜索树。

本题中，一个高度平衡二叉树是指一个二叉树每个节点 的左右两个子树的高度差的绝对值不超过 1。

示例:

给定有序数组: [-10,-3,0,5,9],

一个可能的答案是：[0,-3,9,-10,null,5]，它可以表示下面这个高度平衡二叉搜索树：

      0
     / \
   -3   9
   /   /
 -10  5
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

func sortedArrayToBST(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}

	if len(nums) == 1 {
		return &TreeNode{Val: nums[0]}
	}

	mid := len(nums) / 2
	t := &TreeNode{Val: nums[mid]}
	t.Left = sortedArrayToBST(nums[:mid])
	t.Right = sortedArrayToBST(nums[mid+1:])
	return t
	//return arrayToBST(0, len(nums)-1, nums)
}

func arrayToBST(l, r int, nums []int) *TreeNode {
	if l > r {
		return nil
	}

	//递归还应该可以这样用
	mid := l + (r-l)/2
	t := &TreeNode{Val: nums[mid]}
	t.Left = arrayToBST(l, mid-1, nums)
	t.Right = arrayToBST(mid+1, r, nums)
	return t
}
