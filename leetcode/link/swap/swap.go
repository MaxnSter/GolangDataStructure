package swap

/*
24. 两两交换链表中的节点
给定一个链表，两两交换其中相邻的节点，并返回交换后的链表。

示例:

给定 1->2->3->4, 你应该返回 2->1->4->3.
说明:

你的算法只能使用常数的额外空间。
你不能只是单纯的改变节点内部的值，而是需要实际的进行节点交换。
*/
/**
* Definition for singly-linked list.
* type ListNode struct {
*     Val int
*     Next *ListNode
* }
 */
type ListNode struct {
	Val  int
	Next *ListNode
}

//链表问题,画图就可以了,看下面那个
func swapPairs(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	//先处理头部
	last := head.Next
	last.Next, head.Next = head, last.Next
	head = last

	prev, slow, fast := head.Next, head.Next, head.Next
	if fast.Next == nil || fast.Next.Next == nil {
		//没有多余可以交换的节点了
		return head
	}

	slow, fast = slow.Next, fast.Next.Next
	for {
		//交换
		prev.Next = fast
		slow.Next = fast.Next
		fast.Next = slow

		//交换过来以便前进
		fast, slow = slow, fast
		//跟进
		prev = fast

		if fast.Next == nil || fast.Next.Next == nil {
			break
		}
		fast, slow = fast.Next.Next, slow.Next.Next
	}

	return head
}

//没那么麻烦
func swapPairsClean(head *ListNode) *ListNode {
	dummyHead := &ListNode{Next: head}
	head = dummyHead

	//把head作为prev指针,并且每次遍历head就好了,不需要维护fast和slow
	for head.Next != nil && head.Next.Next != nil {
		fast, slow := head.Next.Next, head.Next

		head.Next = fast
		slow.Next = fast.Next
		fast.Next = slow

		//跟进
		head = slow
	}

	return dummyHead.Next
}
