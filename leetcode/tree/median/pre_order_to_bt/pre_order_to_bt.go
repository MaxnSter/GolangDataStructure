package pre_order_to_bt

/*
105. 从前序与中序遍历序列构造二叉树
根据一棵树的前序遍历与中序遍历构造二叉树。

注意:
你可以假设树中没有重复的元素。

例如，给出

前序遍历 preorder = [3,9,20,15,7]
中序遍历 inorder = [9,3,15,20,7]
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

func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 || len(inorder) == 0 {
		return nil
	}

	//注意点1:base case2
	if len(preorder) == 1 {
		return &TreeNode{Val: preorder[0]}
	}

	rootVal := preorder[0]
	root := &TreeNode{Val: rootVal}

	pos := idxOfVal(rootVal, inorder)
	inLeft := inorder[:pos]
	inRight := inorder[pos+1:]

	//注意点2:切片是[),所以还要加1
	//注意点3:preOrder的下个坐标,完全可以由pos计算出来
	preLeft := preorder[1 : pos-0+1] //这里的0指的是start idx of preOrder
	preRight := preorder[pos+1:]

	root.Left = buildTree(preLeft, inLeft)
	root.Right = buildTree(preRight, inRight)
	return root
}

func idxOfVal(val int, nums []int) int {
	for i, v := range nums {
		if v == val {
			return i
		}
	}

	panic("val not exist")
}
