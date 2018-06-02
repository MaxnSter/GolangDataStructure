// 寻找动态数据中的中位数
package median

import "container/heap"

// from 九章算法
// https://zhuanlan.zhihu.com/p/29595813
// 中位数定义:
// 中位数是一串数中大小排在中间的数,若共有奇数个数字，中位数只有一个.
// 若共有偶数个数字,中位数定义为大小排在中间的两个数的平均值.例如：
// [2,3,4], 中位数是3
// [2,3], 中位数是(2 + 3) / 2 = 2.5

// 静态数据中的中位数: 排序
// tips: 对于重复某几个操作的问题,正确的算法中每一次操作的复杂度
// 都不会很高,思考的时候最好往O(1)和O(log n)的方向思考

// 分析过程:
// 首先从中位数定义入手,从它的定义可以看出,它均分了大于它的元素和小于
// 它的元素.因此我们可以弄两个集合,一个最大堆,里面所有元素都小于中位数
// 一个最小堆,里面所有元素都大于中位数
// 当前数组个数为偶数时,中位数=(最大堆最大值+最小堆最小值)/2
// 当前数组个数为奇数时,中位数=最大堆最大值
//
// 分析完成,理清涉及到的数据结构的关系:
// 集合A最大堆,集合B为最小堆
// 1.集合A元素都小于集合B中元素
// 2.集合A元素个数和B相等或只多一个
//
// 代码实现上述关系:
// 1.插入数据时,进入A,然后把A最大值加入B
// 2.若B的个数大于A,则把B的最小值加入A

type minHeap []int

func (h minHeap) Len() int            { return len(h) }
func (h minHeap) Less(i, j int) bool  { return h[i] < h[j] }
func (h minHeap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *minHeap) Push(x interface{}) { pushFunc((*[]int)(h), x) }
func (h *minHeap) Pop() interface{}   { return popFunc((*[]int)(h)) }

type maxHeap []int

func (h maxHeap) Len() int            { return len(h) }
func (h maxHeap) Less(i, j int) bool  { return h[i] > h[j] }
func (h maxHeap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *maxHeap) Push(x interface{}) { pushFunc((*[]int)(h), x) }
func (h *maxHeap) Pop() interface{}   { return popFunc((*[]int)(h)) }

func pushFunc(h *[]int, x interface{}) {
	*h = append(*h, x.(int))
}

func popFunc(h *[]int) interface{} {
	n := len(*h)
	x := (*h)[n-1]
	*h = (*h)[:n-1]
	return x
}


/**
 * Your MedianFinder object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddNum(num);
 * param_2 := obj.FindMedian();
 */

// leetCode: https://leetcode-cn.com/problems/find-median-from-data-stream/description/
type MedianFinder struct {
	// minH所有元素均大于maxH中的元素
	minH minHeap
	maxH maxHeap
}

/** initialize your data structure here. */
func Constructor() MedianFinder {
	m := MedianFinder{
		minH: make([]int, 0),
		maxH: make([]int, 0),
	}

	heap.Init(&m.minH)
	heap.Init(&m.maxH)

	return m
}

func (this *MedianFinder) AddNum(num int) {

	// 数据进入最大堆A
	heap.Push(&this.maxH, num)

	// 最大堆A中元素加入最小堆B
	// 保证最小堆B中所有元素均大于最大堆A中的元素
	x := heap.Remove(&this.maxH, 0)
	heap.Push(&this.minH, x)

	// 保证最大堆A的个数始终等于或最多只比B多一个元素
	if this.minH.Len() > this.maxH.Len() {
		x1 := heap.Remove(&this.minH, 0)
		heap.Push(&this.maxH, x1)
	}
}

func (this *MedianFinder) FindMedian() float64 {
	l := this.maxH.Len() + this.minH.Len()

	if l%2 == 0 {
		// 元素个数为偶数,中间二字平均值
		x := float64(this.minH[0])
		y := float64(this.maxH[0])
		return (x + y) / 2
		//return float64((this.minH[0] + this.maxH[0]) / 2)
	} else {
		return float64(this.maxH[0])
	}
}
