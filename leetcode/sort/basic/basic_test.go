package basic

import (
	"fmt"
	"math/rand"
	"sort"
	"testing"
	"time"
)

var (
	N     = 1000000
	NHard = 100000

	//保证多组测试数据一致
	testData = makeArray(N)
)

type sortFunc func([]int)

func makeArray(N int) (a []int) {
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < N; i++ {
		a = append(a, rand.Intn(N))
	}
	return a
}

func SortTest(t *testing.T, N int, fnName string, fn sortFunc) {
	data := make([]int, N)
	copy(data, testData)

	start := time.Now()
	fn(data)
	take := time.Now().Sub(start).Seconds()

	if sort.IsSorted(sort.IntSlice(data)) {
		fmt.Printf("sortFunc:%s, sort:%d take:%f sec\n", fnName, N, take)
	} else {
		t.Fatalf("not sorted")
	}
}

func TestBubbleSort(t *testing.T) {
	if N >= NHard {
		//臣妾做不到啊
		return
	}
	SortTest(t, N, "bubble", bubbleSort)
}

func TestInsertSort(t *testing.T) {
	if N >= NHard {
		return
	}
	SortTest(t, N, "insert", insertSort)
}

func TestSelectSort(t *testing.T) {
	if N >= NHard {
		return
	}
	SortTest(t, N, "select", selectSort)
}

func TestShellSort(t *testing.T) {
	SortTest(t, N, "shell", shellSort)
}

func TestHeapSort(t *testing.T) {
	SortTest(t, N, "heap", heapSort)
}

func TestMergeSort(t *testing.T) {
	SortTest(t, N, "merge", mergeSort)
}

func TestQuickSort(t *testing.T) {
	SortTest(t, N, "quick", quickSort)
}

func TestGo(t *testing.T) {
	SortTest(t, N, "golang", func(ints []int) {
		sort.Ints(ints)
	})
}

func log(testName string) func() {
	fmt.Println(testName + "start")
	return func() {
		fmt.Println(testName + "end")
	}
}

func BenchmarkBubbleSort(b *testing.B) {
	if N >= NHard/10 {
		return
	}

	defer log("bubbleSort")()
	for i := 0; i < b.N; i++ {
		data := make([]int, N)
		copy(data, testData)
		bubbleSort(data)
	}
}

func BenchmarkInsertSort(b *testing.B) {
	if N >= NHard/10 {
		return
	}

	defer log("insertSort")()
	for i := 0; i < b.N; i++ {
		data := make([]int, N)
		copy(data, testData)
		insertSort(data)
	}
}
func BenchmarkSelectSort(b *testing.B) {
	if N >= NHard/10 {
		return
	}

	for i := 0; i < b.N; i++ {
		data := make([]int, N)
		copy(data, testData)
		selectSort(data)
	}
}
func BenchmarkShellSort(b *testing.B) {
	for i := 0; i < b.N; i++ {
		data := make([]int, N)
		copy(data, testData)
		shellSort(data)
	}
}
func BenchmarkHeapSort(b *testing.B) {
	for i := 0; i < b.N; i++ {
		data := make([]int, N)
		copy(data, testData)
		heapSort(data)
	}
}
func BenchmarkMergeSort(b *testing.B) {
	for i := 0; i < b.N; i++ {
		data := make([]int, N)
		copy(data, testData)
		mergeSort(data)
	}
}
func BenchmarkQuickSort(b *testing.B) {
	for i := 0; i < b.N; i++ {
		data := make([]int, N)
		copy(data, testData)
		quickSort(data)
	}
}
func BenchmarkGoSort(b *testing.B) {
	for i := 0; i < b.N; i++ {
		data := make([]int, N)
		copy(data, testData)
		sort.Ints(data)
	}
}
