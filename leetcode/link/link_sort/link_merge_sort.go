package link_sort

/*
在 O(n log n) 时间复杂度和常数级空间复杂度下，对链表进行排序。

示例 1:

输入: 4->2->1->3
输出: 1->2->3->4
示例 2:

输入: -1->5->3->4->0
输出: -1->0->3->4->5
*/
func sortList(head *ListNode) *ListNode {
	//基准情况,不用递归就能求解
	if head == nil || head.Next == nil {
		return head
	}

	//不断推进
	mid := midden(head)
	left, right := head, mid.Next
	mid.Next = nil

	//每一次递归调用的情况都必须要使求解状况朝接近基准情况的方向推进
	//所以要不断更新left,right的值
	left = sortList(left)
	right = sortList(right)

	//设计法则,假设所有递归调用都能运行,到这里我们已经拿到了想要的结果
	//不要去纠结过程
	return merge(left, right)
}

func midden(head *ListNode) *ListNode {
	slow, fast := head, head.Next
	for fast != nil && fast.Next != nil {
		slow, fast = slow.Next, fast.Next.Next
	}
	return slow
}

func merge(l1 *ListNode, l2 *ListNode) *ListNode {
	dummyHead := &ListNode{}
	head := dummyHead

	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			head.Next = l1
			l1, head = l1.Next, head.Next
		} else {
			head.Next = l2
			l2, head = l2.Next, head.Next
		}
	}

	if l1 != nil {
		head.Next = l1
	}

	if l2 != nil {
		head.Next = l2
	}

	return dummyHead.Next
}
