package kth_small

/*
230. 二叉搜索树中第K小的元素
给定一个二叉搜索树，编写一个函数 kthSmallest 来查找其中第 k 个最小的元素。

说明：
你可以假设 k 总是有效的，1 ≤ k ≤ 二叉搜索树元素个数。

示例 1:

输入: root = [3,1,4,null,2], k = 1
   3
  / \
 1   4
  \
   2
输出: 1
示例 2:

输入: root = [5,3,6,2,4,null,null,1], k = 3
       5
      / \
     3   6
    / \
   2   4
  /
 1
输出: 3
进阶：
如果二叉搜索树经常被修改（插入/删除操作）并且你需要频繁地查找第 k 小的值，你将如何优化 kthSmallest 函数？
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

func kthSmallest(root *TreeNode, k int) int {
	count, result := 0, 0
	//请利用二叉数的性质!!!!!!
	traverse(root, &count, k, &result)
	return result
}

func traverse(root *TreeNode, count *int, k int, result *int) {
	if root == nil {
		return
	}

	traverse(root.Left, count, k, result)
	*count++
	if *count == k {
		*result = root.Val
	}
	traverse(root.Right, count, k, result)
}
