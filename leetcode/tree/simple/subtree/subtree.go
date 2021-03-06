package subtree

/*
572. 另一个树的子树
给定两个非空二叉树 s 和 t，检验 s 中是否包含和 t 具有相同结构和节点值的子树。
s 的一个子树包括 s 的一个节点和这个节点的所有子孙。s 也可以看做它自身的一棵子树。

示例 1:
给定的树 s:

     3
    / \
   4   5
  / \
 1   2
给定的树 t：

   4
  / \
 1   2
返回 true，因为 t 与 s 的一个子树拥有相同的结构和节点值。

示例 2:
给定的树 s：

     3
    / \
   4   5
  / \
 1   2
    /
   0
给定的树 t：

   4
  / \
 1   2
返回 false。
*/
/**
* Definition for a binary tree node.
 */
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSubtree(s *TreeNode, t *TreeNode) bool {
	if s == nil || t == nil {
		/*
			if s == nil && t == nil {
				return true
			}

			if s == nil && t != nil {
				return false
			}

			if s != nil && t == nil {
				return false
			}

			s != nil, t != nil 做操作...

			以后请简化成现在这样
		*/
		return s == nil && t == nil
	}

	return isEqual(s, t) ||
		isSubtree(s.Left, t) || isSubtree(s.Right, t)
}

func isEqual(s *TreeNode, t *TreeNode) bool {
	if s == nil || t == nil {
		return s == nil && t == nil
	}

	if s.Val != t.Val {
		return false
	}

	return isEqual(s.Left, t.Left) && isEqual(s.Right, t.Right)
}
