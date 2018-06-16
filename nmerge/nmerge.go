//n路归并
package nmerge

import "container/heap"

type MergeSource interface {
	Less(oth MergeSource) bool

	//如果还有下个元素,则返回true,并更新source状态,将下个元素更新为当前元素
	Next() bool

	//将当前元素记录至任意可以记录的地方
	OutputTo(to interface{})
}

type SourceHeap []MergeSource

func (sh SourceHeap) Len() int           { return len(sh) }
func (sh SourceHeap) Less(i, j int) bool { return sh[i].Less(sh[j]) }
func (sh SourceHeap) Swap(i, j int)      { sh[i], sh[j] = sh[j], sh[i] }

func (sh *SourceHeap) Push(x interface{}) {
	*sh = append(*sh, x.(MergeSource))
}

func (sh *SourceHeap) Pop() interface{} {
	l := len(*sh)
	x := (*sh)[l-1]
	*sh = (*sh)[:l-1]
	return x
}

//根据len(srcs)进行N路归并
//例如len(srcs)==8,则进行8路归并
func Merge(srcs []MergeSource, output interface{}) {
	if len(srcs) < 1 {
		return
	}

	srcHeap := SourceHeap(srcs)
	heap.Init(&srcHeap)
	for len(srcHeap) > 0 {
		maxSource := heap.Pop(&srcHeap).(MergeSource)
		maxSource.OutputTo(output)

		if maxSource.Next() {
			heap.Push(&srcHeap, maxSource)
		}
	}
}
