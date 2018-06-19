package next_permutation

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test(t *testing.T) {
	nums1 := []int{1, 2, 3}
	nextPermutation(nums1)
	assert.Equal(t, []int{1, 3, 2}, nums1)

	nums2 := []int{3, 2, 1}
	nextPermutation(nums2)
	assert.Equal(t, []int{1, 2, 3}, nums2)

	nums3 := []int{1, 5, 1}
	nextPermutation(nums3)
	assert.Equal(t, []int{5, 1, 1}, nums3)

	nums4 := []int{1, 3, 2}
	nextPermutation(nums4)
	assert.Equal(t, []int{2, 1, 3}, nums4)
}
