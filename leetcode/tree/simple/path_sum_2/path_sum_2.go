package path_sum_2

/*
113. 路径总和 II
给定一个二叉树和一个目标和，找到所有从根节点到叶子节点路径总和等于给定目标和的路径。

说明: 叶子节点是指没有子节点的节点。

示例:
给定如下二叉树，以及目标和 sum = 22，

              5
             / \
            4   8
           /   / \
          11  13  4
         /  \    / \
        7    2  5   1
返回:

[
   [5,4,11,2],
   [5,8,4,5]
]
*/
/**
* Definition for a binary tree node.
 */
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func pathSum(root *TreeNode, sum int) [][]int {
	if root == nil {
		return nil
	}

	path, result := make([]int, 0), make([][]int, 0)
	calPathSum(root, &path, sum, &result)
	return result
}

func calPathSum(root *TreeNode, pathSum *[]int, sum int, results *[][]int) {
	if root == nil {
		return
	}

	//技巧,用减法代替加法,这样就不用每次遍历都计算了
	sum -= root.Val
	if root.Left == nil && root.Right == nil {
		*pathSum = append(*pathSum, root.Val)

		if sum == 0{
			result := make([]int, len(*pathSum))
			copy(result, *pathSum)
			*results = append(*results, result)
		}
		*pathSum = (*pathSum)[:len(*pathSum)-1]
		return
	}

	*pathSum = append(*pathSum, root.Val)
	calPathSum(root.Left, pathSum, sum, results)
	calPathSum(root.Right, pathSum, sum, results)

	//遍历完一整条路径后,往上回溯时,删除对应的路径
	*pathSum = (*pathSum)[:len(*pathSum)-1]
}

func sumInts(nums ...int) (sum int) {
	for _, num := range nums {
		sum += num
	}
	return sum
}
