package decode_string

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test(t *testing.T) {
	assert.Equal(t, "aaabcbc", decodeString("3[a]2[bc]"))
	assert.Equal(t, "accaccacc", decodeString("3[a2[c]]"))
	assert.Equal(t, "abcabccdcdcdef", decodeString("2[abc]3[cd]ef"))
}
