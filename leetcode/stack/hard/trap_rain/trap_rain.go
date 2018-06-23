package trap_rain

/*
42. 接雨水
给定 n 个非负整数表示每个宽度为 1 的柱子的高度图，计算按此排列的柱子，下雨之后能接多少雨水。



上面是由数组 [0,1,0,2,1,0,1,3,2,1,2,1] 表示的高度图，在这种情况下，可以接 6 个单位的雨水（蓝色部分表示雨水）。 感谢 Marcos 贡献此图。

示例:

输入: [0,1,0,2,1,0,1,3,2,1,2,1]
输出: 6
*/
func trap(height []int) int {
	//先找到最高点maxHeight
	maxHeightIdx := 0
	for i, v := range height {
		if v > height[maxHeightIdx] {
			maxHeightIdx = i
		}
	}

	//[0,maxHeight)计算雨水
	vWater, secondHeightIdx := 0, 0
	for i := 0; i < maxHeightIdx; i++ {
		if height[i] > height[secondHeightIdx] {
			secondHeightIdx = i
		} else {
			vWater += height[secondHeightIdx] - height[i]
		}
	}

	//[maxHeight+1, end)计算雨水,倒序
	secondHeightIdx = len(height) - 1
	for i := len(height) - 1; i > maxHeightIdx; i-- {
		if height[i] > height[secondHeightIdx] {
			secondHeightIdx = i
		} else {
			vWater += height[secondHeightIdx] - height[i]
		}
	}

	return vWater
}
