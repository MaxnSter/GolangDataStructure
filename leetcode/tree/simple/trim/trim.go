package trim

/*
669. 修剪二叉搜索树
给定一个二叉搜索树，同时给定最小边界L 和最大边界 R。通过修剪二叉搜索树，使得所有节点的值在[L, R]中 (R>=L)
你可能需要改变树的根节点，所以结果应当返回修剪好的二叉搜索树的新的根节点。

示例 1:

输入:
    1
   / \
  0   2

  L = 1
  R = 2

输出:
    1
      \
       2
示例 2:

输入:
    3
   / \
  0   4
   \
    2
   /
  1

  L = 1
  R = 3

输出:
      3
     /
   2
  /
 1
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

func trimBST(root *TreeNode, L int, R int) *TreeNode {
	if root == nil {
		return root
	}

	//要利用好搜索二叉树的性质啊!
	if root.Val < L {
		//根节点包括其左子树全部不符合要求了,全部裁掉
		return trimBST(root.Right, L, R)
	} else if root.Val > R {
		//同上
		return trimBST(root.Left, L, R)
	} else {
		root.Left = trimBST(root.Left, L, R)
		root.Right = trimBST(root.Right, L, R)
		return root
	}

}
