package min_cost_climbing_stairs

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test(t *testing.T) {
	//assert.Equal(t,15, minCostClimbingStairs([]int{10, 15, 20}))
	//assert.Equal(t,6, minCostClimbingStairs([]int{1, 100, 1, 1, 1, 100, 1, 1, 100, 1}))
	//assert.Equal(t,2, minCostClimbingStairs([]int{0, 2, 2, 1}))
	assert.Equal(t,2, minCostClimbingStairs([]int{0, 1, 2, 2}))
}
