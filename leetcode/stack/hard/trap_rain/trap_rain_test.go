package trap_rain

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test(t *testing.T) {
	assert.Equal(t, 6,trap([]int{0,1,0,2,1,0,1,3,2,1,2,1}) )
}
