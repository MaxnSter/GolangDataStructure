package mode_bst

/*
501. 二叉搜索树中的众数
给定一个有相同值的二叉搜索树（BST），找出 BST 中的所有众数（出现频率最高的元素）。

假定 BST 有如下定义：

结点左子树中所含结点的值小于等于当前结点的值
结点右子树中所含结点的值大于等于当前结点的值
左子树和右子树都是二叉搜索树
例如：
给定 BST [1,null,2,2],

   1
    \
     2
    /
   2
返回[2].

提示：如果众数超过1个，不需考虑输出顺序

进阶：你可以不使用额外的空间吗？（假设由递归产生的隐式调用栈的开销不被计算在内）
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

func findMode(root *TreeNode) []int {
	m, result, max := make(map[int]int), make([]int, 0), 0
	inorder(root, m, &max)

	for k, count := range m {
		if count == max {
			result = append(result, k)
		}
	}

	return result
}

func inorder(root *TreeNode, m map[int]int, max *int) {
	if root == nil {
		return
	}

	inorder(root.Left, m, max)
	m[root.Val]++
	*max = maxInt(*max, m[root.Val])
	inorder(root.Right, m, max)
}

func maxInt(x, y int) int {
	if x > y {
		return x
	}
	return y
}
