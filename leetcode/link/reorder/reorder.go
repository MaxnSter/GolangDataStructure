package reorder

/*
143. 重排链表

给定一个单链表 L：L0→L1→…→Ln-1→Ln ，
将其重新排列后变为： L0→Ln→L1→Ln-1→L2→Ln-2→…

你不能只是单纯的改变节点内部的值，而是需要实际的进行节点交换。

示例 1:

给定链表 1->2->3->4, 重新排列为 1->4->2->3.
示例 2:

给定链表 1->2->3->4->5, 重新排列为 1->5->2->4->3.
*/
/**
* Definition for singly-linked list.
*
 */
type ListNode struct {
	Val  int
	Next *ListNode
}

func reorderList(head *ListNode) {
	if head == nil || head.Next == nil {
		return
	}

	count, subNode := listLen(head), head
	for i := 0; i < (count-1)/2; i++ {
		subNode = subNode.Next
	}

	tailHead := subNode.Next
	subNode.Next = nil
	tailHead = reverseList(tailHead)

	//merge
	for head != nil && tailHead != nil {
		l1 := head
		l2 := tailHead

		head = head.Next
		tailHead = tailHead.Next

		l2.Next = l1.Next
		l1.Next = l2
	}
}

func listLen(list *ListNode) (count int) {
	for list != nil {
		count++
		list = list.Next
	}
	return count
}

func reverseList(list *ListNode) (reverse *ListNode) {
	for list != nil {
		tmp := list.Next
		list.Next = reverse
		reverse = list
		list = tmp
	}
	return
}
