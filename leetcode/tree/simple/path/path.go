package path

import (
	"fmt"
	"strconv"
)

/*
257. 二叉树的所有路径
给定一个二叉树，返回所有从根节点到叶子节点的路径。

说明: 叶子节点是指没有子节点的节点。

示例:

输入:

   1
 /   \
2     3
 \
  5

输出: ["1->2->5", "1->3"]

解释: 所有根节点到叶子节点的路径为: 1->2->5, 1->3
*/
/**
* Definition for a binary tree node.
 */
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func binaryTreePaths(root *TreeNode) []string {
	result := make([]string, 0)
	path := ""
	pathTree(root, path, &result)
	return result
}

//传入的string,用来记录递归的状态,这是关键
func pathTree(root *TreeNode, path string, result *[]string) {
	if root == nil {
		return
	}

	if path == "" {
		path = strconv.Itoa(root.Val)
	} else {
		path += fmt.Sprintf("->%d", root.Val)
	}

	if root.Left == nil && root.Right == nil {
		//想想base case, 到底应该什么时候append
		*result = append(*result, path)
		return
	}

	pathTree(root.Left, path, result)
	pathTree(root.Right, path, result)
}
