package mergek_sorted

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func makeList(vals ...int) *ListNode {
	head := &ListNode{}
	tmp := head

	for _, val := range vals {
		tmp.Next = &ListNode{Val:val}
		tmp = tmp.Next
	}

	return head.Next
}

func listEqual(l *ListNode, vals ...int) bool {
	for _, val := range vals {
		if l == nil {
			fmt.Println("list shortted")
			return false
		}

		if l.Val != val {
			fmt.Printf("list:%d, should be %d\n", l.Val, val)
			return false
		}

		l = l.Next
	}
	return true
}

func Test(t *testing.T) {
	l1 := makeList(1,4,5)
	l2 := makeList(1,3,4)
	l3 := makeList(2,6)

	merge := mergeKLists([]*ListNode{l1,l2,l3})
	assert.Equal(t, true, listEqual(merge, 1,1,2,3,4,4,5,6))
}
