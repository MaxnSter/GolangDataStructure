package four_sum

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test(t *testing.T) {
	assert.Equal(t, [][]int{
		{-1, 0, 0, 1},
		{-2, -1, 1, 2},
		{-2, 0, 0, 2},
	}, fourSum([]int{1, 0, -1, 0, -2, 2}, 0))
}
