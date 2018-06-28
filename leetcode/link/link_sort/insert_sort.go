package link_sort

import "math"

/*
147. 对链表进行插入排序
插入排序算法：

插入排序是迭代的，每次只移动一个元素，直到所有元素可以形成一个有序的输出列表。
每次迭代中，插入排序只从输入数据中移除一个待排序的元素，找到它在序列中适当的位置，并将其插入。
重复直到所有输入数据插入完为止。


示例 1：

输入: 4->2->1->3
输出: 1->2->3->4
示例 2：

输入: -1->5->3->4->0
输出: -1->0->3->4->5
*/
/**
* Definition for singly-linked list.
 */
type ListNode struct {
	Val  int
	Next *ListNode
}

//折腾了好久,烦烦烦
func insertionSortList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	dummyHead := &ListNode{Next: head, Val: math.MinInt64}
	for endNode := head; endNode.Next != nil; {
		insertNode, insertNodePrev := endNode.Next, endNode

		//只能向遍历,然后再调整节点,还有更好的办法吗?
		hHead := dummyHead
		for ; hHead != insertNodePrev && hHead.Next.Val < insertNode.Val; hHead = hHead.Next {
		}
		//找到第一个可插入点
		if hHead != insertNodePrev {
			insertNodePrev.Next = insertNode.Next
			insertNode.Next = hHead.Next
			hHead.Next = insertNode
		} else {
			endNode = endNode.Next
		}
	}

	return dummyHead.Next
}
