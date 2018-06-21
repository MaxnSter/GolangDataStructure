package flowers

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test(t *testing.T) {
	assert.Equal(t, true, canPlaceFlowers([]int{1, 0, 0, 0, 1}, 1))
	assert.Equal(t, false, canPlaceFlowers([]int{1, 0, 0, 0, 1}, 2))
	assert.Equal(t, false, canPlaceFlowers([]int{1, 0, 0, 0, 0, 1}, 2))
}
