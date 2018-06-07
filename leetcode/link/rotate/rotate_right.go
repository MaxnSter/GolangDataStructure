package rotate

/**
给定一个链表，旋转链表，将链表每个节点向右移动 k 个位置，其中 k 是非负数。

示例 1:

输入: 1->2->3->4->5->NULL, k = 2
输出: 4->5->1->2->3->NULL
解释:
向右旋转 1 步: 5->1->2->3->4->NULL
向右旋转 2 步: 4->5->1->2->3->NULL
示例 2:

输入: 0->1->2->NULL, k = 4
输出: 2->0->1->NULL
解释:
向右旋转 1 步: 2->0->1->NULL
向右旋转 2 步: 1->2->0->NULL
向右旋转 3 步: 0->1->2->NULL
向右旋转 4 步: 2->0->1->NULL

https://leetcode-cn.com/problems/rotate-list/description/
*/
type ListNode struct {
	Val  int
	Next *ListNode
}

// 我把链表的优势丢了, 反思.
func rotateRight1(head *ListNode, k int) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	listLen := getListLen(head)
	k = k % listLen
	if k == 0 {
		return head
	}

	lHead, lTail := head, head

	for i := 0; i < listLen-k-1; i++ {
		lHead = lHead.Next
	}

	lTail = lHead.Next
	lTailTmp := lTail
	lHead.Next = nil

	for lTail.Next != nil {
		lTail = lTail.Next
	}
	lTail.Next = head
	head = lTailTmp

	return head

}

func rotateRight(head *ListNode, k int) *ListNode {

	if head == nil || head.Next == nil {
		return head
	}

	var listLen int
	l := head

	for l != nil {
		listLen++
		l = l.Next
	}

	k = k % listLen
	if k == 0 {
		return head
	}

	tmp := make([]int, listLen)
	l = head

	// 数组存放转移后的元素
	for i := 0; i < listLen; i++ {
		tmp[(i+k)%listLen] = l.Val
		l = l.Next
	}

	// 链表重新赋值
	l = head
	for i := 0; i < listLen; i++ {
		l.Val = tmp[i]
		l = l.Next
	}

	return head
}

// 解法2: 链表长度为n,向右移动k次 == 向左移动n-k次
func rotateRight2(head *ListNode, k int) *ListNode {

	if head == nil || head.Next == nil {
		return head
	}

	var listLen int
	l := head

	for l != nil {
		listLen++
		l = l.Next
	}

	k = k % listLen
	if k == 0 {
		return head
	}

	for i := 0; i < listLen-k; i++ {
		head = moveListToLeft(head)
	}

	return head
}

func moveListToLeft(head *ListNode) *ListNode {
	tmp := head.Val
	l := head

	for l != nil && l.Next != nil {
		l.Val = l.Next.Val
		l = l.Next
	}

	l.Val = tmp
	return head
}


func getListLen(head *ListNode) (listLen int) {
	for head != nil {
		listLen++
		head = head.Next
	}
	return
}
