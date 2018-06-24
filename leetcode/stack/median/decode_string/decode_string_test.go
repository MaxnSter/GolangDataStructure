package decode_string

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test(t *testing.T) {
	assert.Equal(t, "aaabcbc", decodeString2("3[a]2[bc]"))
	assert.Equal(t, "accaccacc", decodeString2("3[a2[c]]"))
	assert.Equal(t, "abcabccdcdcdef", decodeString2("2[abc]3[cd]ef"))
}
