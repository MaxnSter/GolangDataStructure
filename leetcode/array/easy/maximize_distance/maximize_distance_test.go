package maximize_distance

import (
	"testing"

	"github.com/magiconair/properties/assert"
)

func Test(t *testing.T) {
	assert.Equal(t, maxDistToClosest([]int{1,0,0,0,1,0,1}), 2)
	assert.Equal(t, maxDistToClosest([]int{1,0,0,0}), 3)
	assert.Equal(t, maxDistToClosest([]int{0,0,1,0,1,1}), 2)
	assert.Equal(t, maxDistToClosest([]int{0,1,1,1,0,0,1,0,0}), 2)
	assert.Equal(t, maxDistToClosest([]int{1,1,1,0,1,0,1,1,0,0,1}), 1)
}