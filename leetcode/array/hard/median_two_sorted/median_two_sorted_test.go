package median_two_sorted

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test(t *testing.T) {
	assert.Equal(t, float64(2), findMedianSortedArrays([]int{1, 3}, []int{2}))
	assert.Equal(t, float64(3.5), findMedianSortedArrays([]int{}, []int{3, 4}))
}
