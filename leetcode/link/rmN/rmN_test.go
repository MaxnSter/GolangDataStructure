package rmN

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func makeList(vals ...int) *ListNode {
	var head *ListNode

	for i := len(vals) - 1; i >= 0; i-- {
		node := &ListNode{Val: vals[i]}
		if head == nil {
			head = node
		} else {
			node.Next = head
			head = node
		}
	}

	if !ListEqual(head, vals...) {
		panic("makeList error")
	}

	return head
}

func ListEqual(l *ListNode, vals ...int) bool {
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
	l := makeList(1, 2, 3, 4, 5)
	assert.Equal(t, true, ListEqual(removeNthFromEndEffective(l, 1), 1, 2, 3, 4))
}
