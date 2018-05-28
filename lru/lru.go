// from groupCache
package lru

import "container/list"

// Cache是LRU cache,非线程安全
type Cache struct {
	// MaxEntries 表示最大的元素个数
	// 0 表示个数无限制
	MaxEntries int

	// OnEvicted 提供元素被清除时的hook
	OnEvicted func(key Key, value interface{})

	ll    *list.List
	cache map[interface{}]*list.Element
}

// Key是任何可以比较的类型
type Key interface{}

type entry struct {
	key   Key
	value interface{}
}

// NewLRUCache创建并返回一个Cache
// maxEntry指定Cache中最大元素个数
// 如果maxEntry为0,则元素个数不受限制
func NewLRUCache(maxEntry int) *Cache {
	return &Cache{
		MaxEntries: maxEntry,
		ll:         list.New(),
		cache:      make(map[interface{}]*list.Element),
	}
}

// NewLRUCacheWithHook与NewLRUCache相同,但调用者还可以指定元素被清除时的回调
func NewLRUCacheWithHook(maxEntry int, onEvicted func(key Key, value interface{})) *Cache {
	c := NewLRUCache(maxEntry)
	c.OnEvicted = onEvicted
	return c
}

// Add添加一个新元素至cache中
// 如果key已经存在,直接替换
func (c *Cache) Add(key Key, value interface{}) {
	if c.cache == nil {
		c.cache = make(map[interface{}]*list.Element)
		c.ll = list.New()
	}

	if ee, ok := c.cache[key]; ok {
		c.ll.MoveToFront(ee)
		ee.Value.(*entry).value = value
		return
	}

	ele := c.ll.PushFront(&entry{key: key, value: value})
	c.cache[key] = ele

	// 插入新元素,检查是否需要删除过时元素
	if c.MaxEntries != 0 && c.ll.Len() > c.MaxEntries {
		c.RemoveOldest()
	}
}

// Get检查并返回key对应的value
func (c *Cache) Get(key Key) (value interface{}, ok bool) {
	if c.cache == nil {
		return
	}

	if ele, hit := c.cache[key]; hit {
		c.ll.MoveToFront(ele)
		return ele.Value.(*entry).value, true
	}

	return
}

// Remove将指定key从cache中删除
func (c *Cache) Remove(key Key) {
	if c.cache == nil {
		return
	}

	if ele, hit := c.cache[key]; hit {
		c.removeElement(ele)
	}
}

// RemoveOldest删除最"过时"的元素
func (c *Cache) RemoveOldest() {
	if c.cache == nil {
		return
	}

	ele := c.ll.Back()
	if ele != nil {
		c.removeElement(ele)
	}
}

func (c *Cache) removeElement(e *list.Element) {
	c.ll.Remove(e)
	kv := e.Value.(*entry)
	delete(c.cache, kv.key)

	if c.OnEvicted != nil {
		c.OnEvicted(kv.key, kv.value)
	}
}

// Len返回cache中元素的个数
func (c *Cache) Len() int {
	if c.cache == nil {
		return 0
	}

	return c.ll.Len()
}

// Clear清除cache中所有元素
func (c *Cache) Clear() {
	if c.OnEvicted != nil {
		for _, e := range c.cache {
			kv := e.Value.(*entry)
			c.OnEvicted(kv.key, kv.value)
		}
	}

	c.ll = nil
	c.cache = nil
}
