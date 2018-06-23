package rect

/*
84. 柱状图中最大的矩形
给定 n 个非负整数，用来表示柱状图中各个柱子的高度。每个柱子彼此相邻，且宽度为 1 。

求在该柱状图中，能够勾勒出来的矩形的最大面积。



以上是柱状图的示例，其中每个柱子的宽度为 1，给定的高度为 [2,1,5,6,2,3]。



图中阴影部分为所能勾勒出的最大矩形面积，其面积为 10 个单位。



示例:

输入: [2,1,5,6,2,3]
输出: 10
*/

type stack struct {
	data []int
}

func newStack() *stack {
	return &stack{data: make([]int, 0)}
}

func (s *stack) empty() bool {
	return s.size() <= 0
}

func (s *stack) size() int {
	return len(s.data)
}

func (s *stack) push(x int) {
	s.data = append(s.data, x)
}

func (s *stack) pop() {
	if s.empty() {
		return
	}

	s.data = s.data[:s.size()-1]
}

func (s *stack) top() int {
	if s.empty() {
		panic("top empty stack")
	}
	return s.data[s.size()-1]
}

func maxInt(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func LargestRectangleArea(heights []int) int {
	n, idxs, maxArea := len(heights), newStack(), 0

	for i := 0; i < len(heights); i++ {
		if idxs.empty() || heights[i] >= heights[idxs.top()]  {
			//大于当前栈顶对应的元素值
			idxs.push(i)
		} else {
			//不断弹出栈顶元素并计算,知道当前元素值大于栈顶
			//递增栈的维护,想想递增栈,就理解了
			for !idxs.empty() && heights[i] < heights[idxs.top()] {
				var h, w int

				h = heights[idxs.top()]
				idxs.pop()

				if idxs.empty() {
					w = i
				} else {
					w = i - idxs.top() - 1
				}

				maxArea = maxInt(maxArea, h * w)
			}

			//循环结束后,i对应元素大于栈顶元素
			idxs.push(i)
		}
	}

	//遍历到最后一个元素,同理,只是这个时候,我们的i变成了n
	for !idxs.empty() {
		var h, w int

		h = heights[idxs.top()]
		idxs.pop()

		if idxs.empty() {
			w = n
		} else {
			w = n - idxs.top() - 1
		}

		maxArea = maxInt(maxArea, h * w)
	}

	return maxArea
}
