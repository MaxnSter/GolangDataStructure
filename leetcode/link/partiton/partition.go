package partiton

/*
86. 分隔链表
给定一个链表和一个特定值 x，对链表进行分隔，使得所有小于 x 的节点都在大于或等于 x 的节点之前。

你应当保留两个分区中每个节点的初始相对位置。

示例:

输入: head = 1->4->3->2->5->2, x = 3
输出: 1->2->2->4->3->5
*/
/**
* Definition for singly-linked list.
 */
type ListNode struct {
	Val  int
	Next *ListNode
}

//遍历一遍,把小于那个数的节点收集起来,再合并
func partition(head *ListNode, x int) *ListNode {
	dummyHead, dummySmallerHead := &ListNode{Next: head}, &ListNode{}
	head, smallerHead := dummyHead, dummySmallerHead

	for head.Next != nil {
		if head.Next.Val < x {
			smallerHead.Next = head.Next
			head.Next = smallerHead.Next.Next
			smallerHead.Next.Next = nil
			smallerHead = smallerHead.Next
		} else {
			head = head.Next
		}
	}

	//组合
	smallerHead.Next = dummyHead.Next
	return dummySmallerHead.Next
}
