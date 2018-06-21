package search_in_rotated

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test(t *testing.T) {
	assert.Equal(t, 1, search([]int{3, 1}, 1))
	//assert.Equal(t, 4, search( []int{4,5,6,7,0,1,2}, 0))
	//assert.Equal(t, -1, search( []int{4,5,6,7,0,1,2}, 3))
}
