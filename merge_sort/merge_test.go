package merge_sort

import (
	"github.com/stretchr/testify/assert"
	"math/rand"
	"sort"
	"testing"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

const (
	n = 4000000
)

type ss []int

func (s ss) Len() int {
	return len(s)
}

func (s ss) Less(i, j int) bool {
	return s[i] < s[j]
}

func (s ss) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func genArray(n int) ss {
	a := make([]int, n)
	for i := 0; i < n; i++ {
		a[i] = rand.Intn(n)
	}
	return a
}

func TestMergeSort(t *testing.T) {
	a := genArray(n)
	assert.Equal(t, false, sort.IsSorted(a), "array should not be sorted")
	MergeSort(a)
	assert.Equal(t, true, sort.IsSorted(a), "array should be sorted")
}

func TestMergeSortParallel(t *testing.T) {
	a := genArray(n)
	assert.Equal(t, false, sort.IsSorted(a), "array should not be sorted")
	MergeSortParallel(a)
	assert.Equal(t, true, sort.IsSorted(a), "array should be sorted")
}

func BenchmarkMergeSort(b *testing.B) {
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		a := genArray(n)
		b.StartTimer()
		MergeSort(a)
	}
}

func BenchmarkMergeSortParallel(b *testing.B) {
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		a := genArray(n)
		b.StartTimer()
		MergeSortParallel(a)
	}
}

func BenchmarkStdSort(b *testing.B) {
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		a := genArray(n)
		b.StartTimer()
		sort.Sort(a)
	}
}
