package rect

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test(t *testing.T) {
	assert.Equal(t, 10, LargestRectangleArea([]int{2, 1, 5, 6, 2, 3}))
}
