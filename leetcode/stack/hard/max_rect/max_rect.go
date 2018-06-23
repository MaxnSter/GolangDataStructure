package max_rect

/*
85. 最大矩形
给定一个仅包含 0 和 1 的二维二进制矩阵，找出只包含 1 的最大矩形，并返回其面积。

示例:

输入:
[
  ["1","0","1","0","0"],
  ["1","0","1","1","1"],
  ["1","1","1","1","1"],
  ["1","0","0","1","0"]
]
输出: 6
*/

//转换为直方图的题目
func maximalRectangle(matrix [][]byte) int {
	if len(matrix) == 0 {
		return 0
	}

	row, col := len(matrix), len(matrix[0])

	heights := make([][]int, row)
	for i := 0; i < len(heights); i++ {
		heights[i] = make([]int, col)
	}

	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			if matrix[i][j] == '0' {
				heights[i][j] = 0
			} else {
				if i == 0 {
					heights[i][j] = 1
				} else {
					heights[i][j] += heights[i-1][j] + 1
				}
			}
		}
	}

	max := 0
	for i := 0; i < len(heights); i++ {
		max = maxInt(largestRectangleArea(heights[i]), max)
	}

	return max
}

func largestRectangleArea(height []int) int {
	maxArea, idxs := 0, newStack()
	for i := 0; i <= len(height); i++ {
		var curVal int
		if i == len(height) {
			curVal = -1
		} else {
			curVal = height[i]
		}

		//这个我感觉是有意精简版了,要理解思路,还是要看原本的版本
		for !idxs.empty() && curVal < height[idxs.top()] {
			var h, w int

			h = height[idxs.top()]
			idxs.pop()

			if idxs.empty() {
				w = i
			} else {
				w = i - idxs.top() - 1
			}

			maxArea = maxInt(maxArea, h*w)
		}

		idxs.push(i)
	}

	return maxArea
}

func maxInt(x, y int) int {
	if x > y {
		return x
	}
	return y
}

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
