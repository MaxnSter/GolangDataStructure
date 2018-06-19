package pivot_index

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test(t *testing.T) {
	assert.Equal(t, 3, pivotIndex([]int{1, 7, 3, 6, 5, 6}))
	assert.Equal(t, -1, pivotIndex([]int{1, 2, 3}))
}

func TestEffective(t *testing.T) {
	assert.Equal(t, 3, pivotIndexEffective([]int{1, 7, 3, 6, 5, 6}))
	assert.Equal(t, -1, pivotIndexEffective([]int{1, 2, 3}))
}
