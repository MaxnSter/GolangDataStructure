package lsw

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test(t *testing.T) {
	assert.Equal(t, 3, lengthOfLongestSubstringWindow("abcabcbb"))
	assert.Equal(t, 1, lengthOfLongestSubstringWindow("bbbbbbbbbbb"))
	assert.Equal(t, 2, lengthOfLongestSubstringWindow("au"))
	assert.Equal(t, 0, lengthOfLongestSubstringWindow(""))
	assert.Equal(t, 1, lengthOfLongestSubstringWindow("a"))
	assert.Equal(t, 2, lengthOfLongestSubstringWindow("cdd"))
	assert.Equal(t, 2, lengthOfLongestSubstringWindow("abba"))
}
