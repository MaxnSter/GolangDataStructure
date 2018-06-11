package consistenthash

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHashing(t *testing.T) {
	hash := New(3, func(key []byte) uint32 {
		i, err := strconv.Atoi(string(key))
		assert.Nil(t, err)
		return uint32(i)
	})

	//哈希节点
	//2, 4, 6, 12, 14, 16, 22, 24, 26
	hash.Add("6", "4", "2")

	testCases := map[string]string{
		"2":  "2",
		"11": "2",
		"23": "4",
		"25": "6",
		"27": "2",
	}

	for k, v := range testCases {
		assert.Equal(t, v, hash.Get(k))
	}

	//8, 18, 28
	hash.Add("8")
	testCases["27"] = "8"

	for k, v := range testCases {
		assert.Equal(t, v, hash.Get(k))
	}
}
