package reverse2

/*
92. 反转链表 II
反转从位置 m 到 n 的链表。请使用一趟扫描完成反转。

说明:
1 ≤ m ≤ n ≤ 链表长度。

示例:

输入: 1->2->3->4->5->NULL, m = 2, n = 4
输出: 1->4->3->2->5->NULL
*/
/**
* Definition for singly-linked list.
 */
type ListNode struct {
	Val  int
	Next *ListNode
}

//遍历过程同时反转,关键点是找到反转点的起始位置和结束位置
//画图出来就好
//1->2->3->4->5->NULL
// 1->NULL reverse:4->3->2(遍历过程同时手机) head:5->NULL
// 1->4->3->2->5->NULL
func reverseBetween(head *ListNode, m int, n int) *ListNode {
	if head == nil || head.Next == nil || m == n {
		return head
	}

	dummyHead := &ListNode{Next: head}
	head = dummyHead
	var reverse, reverseTail *ListNode
	var subPrev *ListNode

	for i := 1; i <= n+1; i++ {
		if i == m {
			//记录开始reverse地方
			subPrev = head
			head = head.Next
			subPrev.Next = nil
		} else if i > m {
			tmp := head.Next
			head.Next = reverse
			reverse = head
			head = tmp

			if reverseTail == nil {
				reverseTail = reverse
			}
		} else {
			//记录结束reverse的地方
			head = head.Next
		}
	}

	//合并
	subPrev.Next = reverse
	reverseTail.Next = head
	return dummyHead.Next
}
