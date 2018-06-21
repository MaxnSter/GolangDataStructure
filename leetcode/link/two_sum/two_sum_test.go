package two_sum

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"testing"

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
	wr.WriteString("->nil\n")
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

func Test(t *testing.T) {
	l1 := makeList(7, 2, 4, 3)
	l2 := makeList(5, 6, 4)
	l := addTwoNumbers(l1, l2)

	assert.Equal(t, true, listEqual(l, 7, 8, 0, 7))
}
