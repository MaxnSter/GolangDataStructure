package lru

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type simpleStruct struct {
	int
	string
}

type complexStruct struct {
	int
	simpleStruct
}

var getTests = []struct {
	name       string
	keyToAdd   interface{}
	keyToGet   interface{}
	expectedOk bool
}{
	{"string_hit", "myKey", "myKey", true},
	{"string_miss", "myKey", "nonsense", false},
	{"simple_struct_hit", simpleStruct{1, "two"}, simpleStruct{1, "two"}, true},
	{"simeple_struct_miss", simpleStruct{1, "two"}, simpleStruct{0, "noway"}, false},
	{"complex_struct_hit", complexStruct{1, simpleStruct{2, "three"}},
		complexStruct{1, simpleStruct{2, "three"}}, true},
}

func TestNewLRUCache(t *testing.T) {
	c := NewLRUCache(0)
	assert.NotNil(t, c)
}

func TestNewLRUCacheWithHook(t *testing.T) {
	c := NewLRUCacheWithHook(0, nil)
	assert.NotNil(t, c)
}

func TestCache_Get(t *testing.T) {
	cache := NewLRUCache(0)
	value := 1234
	for _, tc := range getTests {
		cache.Add(tc.keyToAdd, value)
		v, ok := cache.Get(tc.keyToAdd)

		assert.Equal(t, true, ok)
		assert.Equal(t, value, v)
	}
}

func TestCache_Remove(t *testing.T) {
	cache := NewLRUCache(0)
	cache.Add("key", 1234)
	cache.Remove("key")

	_, ok := cache.Get("key")
	assert.Equal(t, false, ok)
}

func TestEvict(t *testing.T) {
	evictedKeys := make([]Key, 0)
	onEvictedFunc := func(key Key, value interface{}) {
		evictedKeys = append(evictedKeys, key)
	}

	cache := NewLRUCacheWithHook(20, onEvictedFunc)
	for i := 0; i < 22; i++ {
		cache.Add(fmt.Sprintf("key%d", i), 123)
	}

	assert.Equal(t, 20, cache.Len())
	assert.Equal(t, 2, len(evictedKeys))
	assert.Equal(t, "key0", evictedKeys[0])
	assert.Equal(t, "key1", evictedKeys[1])
}

func TestCache_Clear(t *testing.T) {

	evictedKeys := make([]Key, 0)
	onEvictedFunc := func(key Key, value interface{}) {
		evictedKeys = append(evictedKeys, key)
	}

	cache := NewLRUCacheWithHook(20, onEvictedFunc)
	for i := 0; i < 22; i++ {
		cache.Add(fmt.Sprintf("key%d", i), 123)
	}

	cache.Clear()

	assert.Equal(t, 0, cache.Len())
	assert.Equal(t, 22, len(evictedKeys))
	assert.Equal(t, "key0", evictedKeys[0])
	assert.Equal(t, "key1", evictedKeys[1])
}
