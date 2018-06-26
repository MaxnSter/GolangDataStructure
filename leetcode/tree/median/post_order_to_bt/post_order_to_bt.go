package post_order_to_bt

/*
106. 从中序与后序遍历序列构造二叉树
根据一棵树的中序遍历与后序遍历构造二叉树。

注意:
你可以假设树中没有重复的元素。

例如，给出

中序遍历 inorder = [9,3,15,20,7]
后序遍历 postorder = [9,15,7,20,3]
返回如下的二叉树：

    3
   / \
  9  20
    /  \
   15   7
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

func buildTree(inorder []int, postorder []int) *TreeNode {
	if len(postorder) == 0 {
		return nil
	}

	if len(postorder) == 1 {
		return &TreeNode{Val: postorder[0]}
	}

	rootVal := postorder[len(postorder)-1]
	pos := idxOf(rootVal, inorder)

	inLeft := inorder[:pos]
	inRight := inorder[pos+1:]

	postLeft := postorder[:pos]
	postRight := postorder[pos : len(postorder)-1]

	root := &TreeNode{Val: rootVal}
	root.Left = buildTree(inLeft, postLeft)
	root.Right = buildTree(inRight, postRight)

	return root
}

func idxOf(val int, nums []int) int {
	for i, v := range nums {
		if v == val {
			return i
		}
	}

	panic("val not exits in nums")
}
