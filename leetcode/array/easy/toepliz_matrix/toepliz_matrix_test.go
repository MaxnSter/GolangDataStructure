package toepliz_matrix

import "testing"

func TestMatrix (t *testing.T) {
	isToeplitzMatrix([][]int {
		{1, 2, 3, 4},
		{5, 1, 2, 3},
		{9, 5, 1, 2},
	})
}
