package largest_num_atleasr_twice

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test(t *testing.T) {
	assert.Equal(t, 1, dominantIndex([]int{3, 6, 1, 0}))
	assert.Equal(t, -1, dominantIndex([]int{1, 2, 3, 4}))
}
