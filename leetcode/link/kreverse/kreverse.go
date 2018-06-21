package kreverse

/*
给出一个链表，每 k 个节点一组进行翻转，并返回翻转后的链表。

k 是一个正整数，它的值小于或等于链表的长度。如果节点总数不是 k 的整数倍，那么将最后剩余节点保持原有顺序。

示例 :

给定这个链表：1->2->3->4->5

当 k = 2 时，应当返回: 2->1->4->3->5

当 k = 3 时，应当返回: 3->2->1->4->5

说明 :

你的算法只能使用常数的额外空间。
你不能只是单纯的改变节点内部的值，而是需要实际的进行节点交换。
*/
/**
* Definition for singly-linked list.
 */
type ListNode struct {
	Val  int
	Next *ListNode
}

//画图画图就会了
func reverseKGroup(head *ListNode, k int) *ListNode {
	if head == nil || head.Next == nil || k == 1 {
		return head
	}

	dummyHead := &ListNode{Next: head}
	prev, nextPrev := dummyHead, dummyHead
	for prev.Next != nil {
		//寻找下一间隔
		for i := 0; nextPrev != nil && i < k; i++ {
			nextPrev = nextPrev.Next
		}
		if nextPrev == nil {
			//走到了尽头
			break
		}

		//找到间隔,以及对应的边界
		nextBegin := nextPrev.Next
		nextPrev.Next = nil
		reverseBegin := prev.Next

		//反转,注意反转后,nextPrev已经变了
		reverseBegin, reverseEnd := reverseList(reverseBegin)

		//再连接起来
		prev.Next, reverseEnd.Next = reverseBegin, nextBegin

		//移动到下一个间隔,对应dummyHead,我们总是在下一个区间之前
		//注意这里nextPrev也要重新赋值,因为原来的值经过reverse之后就不对了
		prev, nextPrev = reverseEnd, reverseEnd
	}

	return dummyHead.Next
}

func reverseList(l *ListNode) (reveres *ListNode, reverseEnd *ListNode) {
	for l != nil {
		tmp := l.Next
		l.Next = reveres
		reveres = l
		l = tmp

		if reverseEnd == nil {
			reverseEnd = reveres
		}
	}
	return
}
