package min_cost_climbing_stairs

/*
数组的每个索引做为一个阶梯，第 i个阶梯对应着一个非负数的体力花费值 cost[i](索引从0开始)。

每当你爬上一个阶梯你都要花费对应的体力花费值，然后你可以选择继续爬一个阶梯或者爬两个阶梯。

您需要找到达到楼层顶部的最低花费。在开始时，你可以选择从索引为 0 或 1 的元素作为初始阶梯。

示例 1:

输入: cost = [10, 15, 20]
输出: 15
解释: 最低花费是从cost[1]开始，然后走两步即可到阶梯顶，一共花费15。
 示例 2:

输入: cost = [1, 100, 1, 1, 1, 100, 1, 1, 100, 1]
输出: 6
解释: 最低花费方式是从cost[0]开始，逐个经过那些1，跳过cost[3]，一共花费6。
注意：

cost 的长度将会在 [2, 1000]。
每一个 cost[i] 将会是一个Integer类型，范围为 [0, 999]。
*/
func minCostClimbingStairs(cost []int) int {
	cost1 := costClimbingStairs(cost)
	cost2 := costClimbingStairs(cost[1:])

	if cost1 < cost2 {
		return cost1
	} else {
		return cost2
	}
}

func costClimbingStairs(cost []int) int {
	//注意点1:起点固定
	stepCost, stepB, stepE := cost[0], 1, 2
	curStep := 0

	for stepE < len(cost) {
		if cost[stepE] <= cost[stepB] {
			//相等时候优先选择后面的那一台阶
			curStep = stepE
		} else {
			curStep = stepB
		}

		stepCost += cost[curStep]
		stepB, stepE = curStep+1, curStep+2
	}

	return stepCost
}
