package max_three_factor

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test(t *testing.T) {
	assert.Equal(t, 24, maximumProductEffetive([]int{1, 2, 3, 4}))
}
