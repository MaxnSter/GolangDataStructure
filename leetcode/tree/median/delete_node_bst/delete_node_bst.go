package delete_node_bst

/*
450. 删除二叉搜索树中的节点
给定一个二叉搜索树的根节点 root 和一个值 key，删除二叉搜索树中的 key 对应的节点，并保证二叉搜索树的性质不变。返回二叉搜索树（有可能被更新）的根节点的引用。

一般来说，删除节点可分为两个步骤：

首先找到需要删除的节点；
如果找到了，删除它。
说明： 要求算法时间复杂度为 O(h)，h 为树的高度。

示例:

root = [5,3,6,2,4,null,7]
key = 3

    5
   / \
  3   6
 / \   \
2   4   7

给定需要删除的节点值是 3，所以我们首先找到 3 这个节点，然后删除它。

一个正确的答案是 [5,4,6,2,null,null,7], 如下图所示。

    5
   / \
  4   6
 /     \
2       7

另一个正确答案是 [5,2,6,null,4,null,7]。

    5
   / \
  2   6
   \   \
    4   7
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

func deleteNode(root *TreeNode, key int) *TreeNode {
	if root == nil {
		return nil
	}

	if root.Val > key {
		//值在左侧
		root.Left = deleteNode(root.Left, key)
	} else if root.Val < key {
		//值在右侧
		root.Right = deleteNode(root.Right, key)
	} else if root.Left != nil && root.Right != nil {
		//两个子结点,找到右边的最小值,或左边的最大值
		//minInR := findMin(root.Right)
		//root.Val = minInR.Val
		//root.Right = deleteNode(root.Right, root.Val)

		maxInL := findMax(root.Left)
		root.Val = maxInL.Val
		root.Left = deleteNode(root.Left, root.Val)
	} else {
		if root.Left == nil { // also handles 0 child
			root = root.Right
		} else if root.Right == nil {
			root = root.Left
		}
	}

	return root
}

func findMax(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	for root.Right != nil {
		root = root.Right
	}
	return root
}

func findMin(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	for root.Left != nil {
		root = root.Left
	}
	return root
}
