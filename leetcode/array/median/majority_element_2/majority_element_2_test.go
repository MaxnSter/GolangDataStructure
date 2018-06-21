package majority_element_2

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test(t *testing.T) {
	assert.Equal(t, []int{3}, majorityElement([]int{3, 2, 3}))
	assert.Equal(t, []int{1, 2}, majorityElement([]int{1, 1, 1, 3, 3, 2, 2, 2}))
	assert.Equal(t, []int{-1}, majorityElement([]int{0, -1, 2, -1}))
}
