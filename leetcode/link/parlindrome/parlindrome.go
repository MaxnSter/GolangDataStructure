package parlindrome

/*
234. 回文链表
请判断一个链表是否为回文链表。

示例 1:

输入: 1->2
输出: false
示例 2:

输入: 1->2->2->1
输出: true
进阶：
你能否用 O(n) 时间复杂度和 O(1) 空间复杂度解决此题？
*/
/**
* Definition for singly-linked list.
*
 */
type ListNode struct {
	Val  int
	Next *ListNode
}

func isPalindrome(head *ListNode) bool {
	var output []int
	dumpListTo(head, &output)
	idxB, idxE := 0, len(output)-1

	for idxB < idxE {
		if output[idxB] != output[idxE] {
			return false
		}

		idxB++
		idxE--
	}
	return true
}

func dumpListTo(head *ListNode, output *[]int) {
	for head != nil {
		*output = append(*output, head.Val)
		head = head.Next
	}
}

func isPalindromeEffective(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return true
	}

	slow, fast := head, head.Next
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	//分离开
	reverseHead := slow.Next
	slow.Next = nil

	reverseHead = reverseList(reverseHead)
	l1, l2 := head, reverseHead
	for l1 != nil && l2 != nil {
		if l1.Val != l2.Val {
			//还原
			reverseHead = reverseList(reverseHead)
			slow.Next = reverseHead
			return false
		}
		l1, l2 = l1.Next, l2.Next
	}

	//还原
	reverseHead = reverseList(reverseHead)
	slow.Next = reverseHead
	return true
}

func reverseList(head *ListNode) (reverse *ListNode) {
	for head != nil {
		tmp := head.Next
		head.Next = reverse
		reverse = head
		head = tmp
	}
	return
}
