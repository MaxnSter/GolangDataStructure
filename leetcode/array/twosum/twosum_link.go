package twosum

/**
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

// https://leetcode-cn.com/problems/add-two-numbers/description/
// 想象加法运算的过程,就很好解答了
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil && l2 == nil {
		return nil
	}

	root := new(ListNode)
	l := root
	carry := 0
	add1, add2 := 0, 0

	for {

		// 三元表达式什么时候能有???
		if l1 == nil {
			add1 = 0
		} else {
			add1 = l1.Val
		}

		if l2 == nil {
			add2 = 0
		} else {
			add2 = l2.Val
		}

		//加算过程,考虑进位
		sum := add1 + add2 + carry
		carry = sum / 10
		l.Val = sum % 10

		// 三元表达式什么时候能有???,再问一遍
		if l1 != nil {
			l1 = l1.Next
		}

		if l2 != nil {
			l2 = l2.Next
		}

		if l1 == nil && l2 == nil {
			l.Next = nil
			break
		} else {
			l.Next = new(ListNode)
			l = l.Next
		}
	}

	// 相加数遍历完了之后,考虑进位
	if carry != 0 {
		l.Next = &ListNode{
			Val:  carry,
			Next: nil,
		}
	}

	return root
}

func listToNumber(l *ListNode) (number int) {
	digit := 1
	for l != nil {
		number = digit*l.Val + number
		l = l.Next
		digit *= 10
	}

	return
}

func numberToList(number int) *ListNode {
	root := new(ListNode)

	l := root
	for {
		l.Val = number % 10

		number /= 10
		if number > 0 {
			l.Next = new(ListNode)
		} else {
			l.Next = nil
			break
		}
	}

	return root
}
