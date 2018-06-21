package reverse

/*
206. 反转链表
反转一个单链表。

示例:

输入: 1->2->3->4->5->NULL
输出: 5->4->3->2->1->NULL
进阶:
你可以迭代或递归地反转链表。你能否用两种方法解决这道题？
*/
/**
* Definition for singly-linked list.
 */
type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	var reverse *ListNode
	for head != nil {
		tmp := head.Next
		head.Next = reverse
		reverse= head
		head = tmp
	}
	return reverse
}
