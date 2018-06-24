package basic_caculator

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test(t *testing.T) {
	assert.Equal(t, 27, calculate("100000000/1/2/3/4/5/6/7/8/9/10"))
}
