package rect

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test(t *testing.T) {
	assert.Equal(t, 4, largestRectangleArea1([]int{2, 4}))
}
