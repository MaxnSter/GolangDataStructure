package _sum_closest

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test(t *testing.T) {
	//assert.Equal(t, 2, threeSumClosest([]int{-1, 2, 1, -4}, 1))
	assert.Equal(t, 82, threeSumClosest([]int{1, 2, 4, 8, 16, 32, 64, 128}, 82))
}
