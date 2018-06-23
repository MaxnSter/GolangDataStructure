package queue_to_stack

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test(t *testing.T) {
	s := Constructor()
	s.Push(1)
	s.Push(2)
	s.Pop()
	assert.Equal(t, 1, s.Top())
}
