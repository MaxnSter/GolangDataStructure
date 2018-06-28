package dup2

/*
82. 删除排序链表中的重复元素 II
给定一个排序链表，删除所有含有重复数字的节点，只保留原始链表中 没有重复出现 的数字。

示例 1:

输入: 1->2->3->3->4->4->5
输出: 1->2->5
示例 2:

输入: 1->1->1->2->3
输出: 2->3
*/
/**
* Definition for singly-linked list.
*
 */
type ListNode struct {
	Val  int
	Next *ListNode
}

//也是一直移动,找出那一重复的块,整个next
// 1->2->3->3->4->4->5
//    |  |  |
//    V  V  V
//  head slow fast
//
// head.Next = fast.Next
//
// 1->2->4->4->5
//    |
//    V
//   head
//
//这个过程
func deleteDuplicates(head *ListNode) *ListNode {
	dummyHead := &ListNode{Next: head}
	head = dummyHead

	for head.Next != nil && head.Next.Next != nil {
		slow, fast := head.Next, head.Next.Next
		if slow.Val != fast.Val {
			head = head.Next
			continue
		}

		//一直往前走,直到下一个结点的Val与自己不相等或到达尾结点
		for fast.Next != nil && fast.Next.Val == fast.Val {
			fast = fast.Next
		}

		//移除相等的区域
		head.Next = fast.Next
	}

	return dummyHead.Next
}
