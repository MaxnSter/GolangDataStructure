// Package consistenthash provides an implementation of a ring hash.
package consistenthash

import (
	"hash/crc32"
	"sort"
	"strconv"
)

// Hash是用于计算hash值的函数
type Hash func(data []byte) uint32

// 一致性哈希表
type Map struct {
	hash     Hash
	replicas int
	keys     []int // sorted
	hashMap  map[int]string
}

// New创建并且返回一致性哈希表
// replicas指定每个节点的复用个数(为了分布更均匀?)
// fn是用于计算hash值的方法,若传nil,则默认为crc32.ChecksumIEEE
func New(replicas int, fn Hash) *Map {
	m := &Map{
		replicas: replicas,
		hash:     fn,
		hashMap:  make(map[int]string),
	}

	if m.hash == nil {
		m.hash = crc32.ChecksumIEEE
	}
	return m
}

// IsEmpty检查节点个数是否为0
func (m *Map) IsEmpty() bool {
	return len(m.keys) == 0
}

// Add增加节点至表中
func (m *Map) Add(keys ...string) {
	for _, key := range keys {
		for i := 0; i < m.replicas; i++ {
			hash := int(m.hash([]byte(strconv.Itoa(i) + key)))
			m.keys = append(m.keys, hash)
			m.hashMap[hash] = key
		}
	}

	sort.Ints(m.keys)
}

// 根据key的hash值,顺时针旋转寻找与其最近的节点
func (m *Map) Get(key string) string {
	if m.IsEmpty() {
		return ""
	}

	hash := int(m.hash([]byte(key)))
	idx := sort.Search(len(m.keys), func(i int) bool {
		// 二分查找
		// if !f(middle); low = middle + 1
		return hash <= m.keys[i]
	})

	// 回到第一个节点,环形的体现
	if idx == len(m.keys) {
		idx = 0
	}

	return m.hashMap[m.keys[idx]]
}
