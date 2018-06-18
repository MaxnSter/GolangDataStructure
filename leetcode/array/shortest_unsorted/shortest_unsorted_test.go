package shortest_unsorted

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test(t *testing.T) {
	assert.Equal(t,5, findUnsortedSubarray([]int{2, 6, 4, 8, 10, 9, 15}))
	assert.Equal(t,4, findUnsortedSubarray([]int{1,3,2,2,2}))
	assert.Equal(t,4, findUnsortedSubarray([]int{1,3,5,4,2}))
}

func Test1(t *testing.T) {
	assert.Equal(t,5, findUnsortedSubarrayEffective([]int{2, 6, 4, 8, 10, 9, 15}))
	assert.Equal(t,4, findUnsortedSubarrayEffective([]int{1,3,2,2,2}))
	assert.Equal(t,4, findUnsortedSubarrayEffective([]int{1,3,5,4,2}))
}
