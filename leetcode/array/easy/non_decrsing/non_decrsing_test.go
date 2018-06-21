package non_decrsing

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test(t *testing.T) {
	assert.Equal(t, true, checkPossibility([]int{4, 2, 3}))
	assert.Equal(t, false, checkPossibility([]int{4, 2, 1}))
}
