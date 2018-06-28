package two_sum

/*
445. 两数相加 II
给定两个非空链表来代表两个非负整数。数字最高位位于链表开始位置。它们的每个节点只存储单个数字。将这两数相加会返回一个新的链表。



你可以假设除了数字 0 之外，这两个数字都不会以零开头。

进阶:

如果输入链表不能修改该如何处理？换句话说，你不能对列表中的节点进行翻转。

示例:

输入: (7 -> 2 -> 4 -> 3) + (5 -> 6 -> 4)
输出: 7 -> 8 -> 0 -> 7
*/
/**
* Definition for singly-linked list.
 */
type ListNode struct {
	Val  int
	Next *ListNode
}

//绝对不可以使用list-to-int的方法,会溢出!!!!!
//注意进位
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	l1 = reverseList(l1)
	l2 = reverseList(l2)
	var carry, curVal int

	dummyHead := &ListNode{}
	head := dummyHead

	for l1 != nil && l2 != nil {
		sum := l1.Val + l2.Val + carry
		curVal, carry = sum%10, sum/10

		head.Next = &ListNode{Val: curVal}
		head, l1, l2 = head.Next, l1.Next, l2.Next
	}

	for l1 != nil {
		sum := l1.Val + carry
		curVal, carry = sum%10, sum/10

		head.Next = &ListNode{Val: curVal}
		head, l1 = head.Next, l1.Next
	}

	for l2 != nil {
		sum := l2.Val + carry
		curVal, carry = sum%10, sum/10

		head.Next = &ListNode{Val: curVal}
		head, l2 = head.Next, l2.Next

	}

	if carry != 0 {
		head.Next = &ListNode{Val: carry}
	}

	return reverseList(dummyHead.Next)
}

func reverseList(l *ListNode) (reverse *ListNode) {
	for l != nil {
		tmp := l.Next
		l.Next = reverse
		reverse = l
		l = tmp
	}
	return
}
