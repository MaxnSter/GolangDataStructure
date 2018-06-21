package merge

/*
 * https://leetcode-cn.com/problems/merge-two-sorted-lists/description/
 * 21. 合并两个有序链表
 * 将两个有序链表合并为一个新的有序链表并返回。新链表是通过拼接给定的两个链表的所有节点组成的。
 * 示例：
 * 输入：1->2->4, 1->3->4
 * 输出：1->1->2->3->4->4
 *
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

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil && l2 == nil {
		return nil
	}

	if l1 == nil && l2 != nil {
		return l2
	}

	if l1 != nil && l2 == nil {
		return l1
	}

	root := new(ListNode)
	l := root

	for l1 != nil && l2 != nil {
		if l1.Val <= l2.Val {
			// 直接利用现有节点,无需再new
			l.Next = l1
			l1 = l1.Next
		} else {
			l.Next = l2
			l2 = l2.Next
		}

		l = l.Next
	}

	// 不需要for
	if l1 != nil {
		l.Next = l1
	}

	if l2 != nil {
		l.Next = l2
	}

	return root.Next
}
