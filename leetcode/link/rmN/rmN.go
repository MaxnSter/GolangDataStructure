package rmN

/*
给定一个链表，删除链表的倒数第 n 个节点，并且返回链表的头结点。

示例：

给定一个链表: 1->2->3->4->5, 和 n = 2.

当删除了倒数第二个节点后，链表变为 1->2->3->5.
说明：

给定的 n 保证是有效的。

进阶：

你能尝试使用一趟扫描实现吗？
*/
type ListNode struct {
	Val  int
	Next *ListNode
}

//这个才是真正的一趟遍历,下面那个可以扔掉了
func removeNthFromEndEffective(head *ListNode, n int) *ListNode {
	dummyHead := &ListNode{Next: head} //for n == len(head)
	fast, slow := dummyHead, dummyHead

	//巧用快慢指针,J贼得很
	for fast != nil {
		if n <= 0 {
			slow = slow.Next
		}
		fast = fast.Next
		n--
	}

	if slow.Next != nil {
		slow.Next = slow.Next.Next
	}
	return dummyHead.Next
}

func removeNthFromEndN(head *ListNode, n int) *ListNode {
	if head == nil {
		return nil
	}

	m, tmp, count := map[int]*ListNode{}, head, 0
	for tmp != nil {
		count++
		m[count] = tmp

		tmp = tmp.Next
	}

	if n == count {
		return head.Next
	}

	//保证N有效,所以就不做判断了
	m[count-n].Next = m[count-n].Next.Next
	return head
}
