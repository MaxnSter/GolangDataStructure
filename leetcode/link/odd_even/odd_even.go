package odd_even

/*
328. 奇偶链表
给定一个单链表，把所有的奇数节点和偶数节点分别排在一起。请注意，这里的奇数节点和偶数节点指的是节点编号的奇偶性，而不是节点的值的奇偶性。

请尝试使用原地算法完成。你的算法的空间复杂度应为 O(1)，时间复杂度应为 O(nodes)，nodes 为节点总数。

示例 1:

输入: 1->2->3->4->5->NULL
输出: 1->3->5->2->4->NULL
示例 2:

输入: 2->1->3->5->6->4->7->NULL
输出: 2->3->6->7->1->5->4->NULL
说明:

应当保持奇数节点和偶数节点的相对顺序。
链表的第一个节点视为奇数节点，第二个节点视为偶数节点，以此类推。
*/
type ListNode struct {
	Val  int
	Next *ListNode
}

//奇偶结点重新收集,然后再合并,直接操作指针就好了,不需要额外空间
func oddEvenList(head *ListNode) *ListNode {
	dummyOddHead, dummyEvenHead := &ListNode{}, &ListNode{}
	dummyHead := &ListNode{Next: head}

	oddHead, evenHead, head := dummyOddHead, dummyEvenHead, dummyHead
	count := 1
	for head.Next != nil {
		if count%2 == 1 {
			oddHead.Next = head.Next
			head.Next = oddHead.Next.Next
			oddHead.Next.Next = nil
			oddHead = oddHead.Next
		} else {
			evenHead.Next = head.Next
			head.Next = evenHead.Next.Next
			evenHead.Next.Next = nil
			evenHead = evenHead.Next
		}
		count++
	}

	//最后把奇偶两个链表结合起来,尾头合并,返回头部
	oddHead.Next = dummyEvenHead.Next
	return dummyOddHead.Next
}
