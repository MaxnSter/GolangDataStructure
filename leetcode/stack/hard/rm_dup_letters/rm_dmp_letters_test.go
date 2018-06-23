package rm_dup_letters

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test(t *testing.T) {
	assert.Equal(t, "abc", removeDuplicateLetters("bcabc"))
	assert.Equal(t, "acdb", removeDuplicateLetters("cbacdcbc"))
}