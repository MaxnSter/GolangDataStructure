package kreverse

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

func Test(t *testing.T) {
	l := makeList(1,2,3,4,5)
	lr := reverseKGroup(l,5)

	assert.Equal(t, true, listEqual(lr,5,4,3,2,1))
}