package ms_subarray

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test(t *testing.T) {
	assert.Equal(t, 2, minSubArrayLen(7, []int{2, 3, 1, 2, 4, 3}))
}
