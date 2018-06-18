package longest_ci_sub

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test(t *testing.T) {
	assert.Equal(t, 3, findLengthOfLCIS([]int{1,3,5,4,7}))
	assert.Equal(t, 1, findLengthOfLCIS([]int{2,2,2,2,2}))
}
