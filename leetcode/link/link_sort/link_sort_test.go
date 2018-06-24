package link_sort

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"sort"
	"strconv"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func dumpList(l *ListNode) {
	wr := bufio.NewWriter(os.Stdout)
	for l != nil {
		if l.Next != nil {
			wr.WriteString(fmt.Sprintf("%d->", l.Val))
		} else {
			wr.WriteString(strconv.Itoa(l.Val))
		}
		l = l.Next
	}
	wr.WriteByte('\n')
	wr.Flush()
}

func makeList(vals ...int) *ListNode {
	var head *ListNode

	for i := len(vals) - 1; i >= 0; i-- {
		node := &ListNode{Val: vals[i]}
		node.Next = head
		head = node
	}

	if !listEqual(head, vals...) {
		panic("makeList error")
	}

	return head
}

func listEqual(l *ListNode, vals ...int) bool {
	head := l
	for _, val := range vals {
		if l == nil {
			fmt.Println("list shortted")
			dumpList(head)
			return false
		}

		if l.Val != val {
			fmt.Printf("list:%d, should be %d\n", l.Val, val)
			dumpList(head)
			return false
		}

		l = l.Next
	}
	return true
}

func TestInsertSort(t *testing.T) {
	rand.Seed(time.Now().UnixNano())
	var nums []int
	for i := 0; i < 100000; i++ {
		nums = append(nums, rand.Intn(100))
	}

	l := makeList(nums...)
	start := time.Now()
	ls := insertionSortList(l)
	fmt.Printf("take:%g s\n", time.Now().Sub(start).Seconds())
	sort.Ints(nums)

	assert.Equal(t, true, listEqual(ls, nums...))
}

func TestMergeSort(t *testing.T) {
	rand.Seed(time.Now().UnixNano())
	var nums []int
	for i := 0; i < 100000; i++ {
		nums = append(nums, rand.Intn(100))
	}

	l := makeList(nums...)
	start := time.Now()
	ls := sortList(l)
	fmt.Printf("take:%g s\n", time.Now().Sub(start).Seconds())
	sort.Ints(nums)

	assert.Equal(t, true, listEqual(ls, nums...))
}
