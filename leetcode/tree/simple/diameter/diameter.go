package diameter

/*
543. 二叉树的直径
给定一棵二叉树，你需要计算它的直径长度。一棵二叉树的直径长度是任意两个结点路径长度中的最大值。这条路径可能穿过根结点。

示例 :
给定二叉树

          1
         / \
        2   3
       / \
      4   5
返回 3, 它的长度是路径 [4,2,1,3] 或者 [5,2,1,3]。

注意：两结点之间的路径长度是以它们之间边的数目表示。
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

func diameterOfBinaryTree(root *TreeNode) int {
	if root == nil || root.Left == nil && root.Right == nil {
		return 0
	}

	maxL, maxR := getHeight(root.Left), getHeight(root.Right)
	//题目没理解好,cnm
	return maxInt(maxL+maxR,
		maxInt(diameterOfBinaryTree(root.Right), diameterOfBinaryTree(root.Left)))
}

func getHeight(root *TreeNode) int {
	if root == nil {
		return 0
	}

	return 1 + maxInt(getHeight(root.Left), getHeight(root.Right))

}

func maxInt(x, y int) int {
	if x > y {
		return x
	}
	return y
}
