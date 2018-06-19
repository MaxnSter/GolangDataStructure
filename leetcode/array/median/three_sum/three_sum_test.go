package three_sum

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test(t *testing.T) {
	assert.Equal(t, [][]int{
		{0,0,0},
	}, threeSum([]int{0, 0, 0,0}))
}
