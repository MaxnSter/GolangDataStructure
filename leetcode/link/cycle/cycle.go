package cycle

/*
141. 环形链表
给定一个链表，判断链表中是否有环。

进阶：
你能否不使用额外空间解决此题？
*/
type ListNode struct {
	Val  int
	Next *ListNode
}

func hashCycle(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return false
	}

	fast, slow := head, head
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next

		if fast == slow {
			return true
		}
	}

	return false
}

/*
给定一个链表，返回链表开始入环的第一个节点。 如果链表无环，则返回 null。

说明：不允许修改给定的链表。

进阶：
你是否可以不用额外空间解决此题？
*/
func detectCycle(head *ListNode) *ListNode {
	//若有环,相遇后快指针从头开始,步数改为1,慢指针接着往下走,下一次相遇点就是入环点
	if head == nil || head.Next == nil {
		return nil
	}

	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next

		//相遇,肯定有解
		if fast == slow {
			fast = head
			//从头开始,步数改为1
			for fast != slow {
				fast = fast.Next
				slow = slow.Next
			}
			return fast
		}
	}
	return nil
}

/*
编写一个程序，找到两个单链表相交的起始节点。



例如，下面的两个链表：

A:          a1 → a2
                   ↘
                     c1 → c2 → c3
                   ↗
B:     b1 → b2 → b3
在节点 c1 开始相交。



注意：

如果两个链表没有交点，返回 null.
在返回结果后，两个链表仍须保持原有的结构。
可假定整个链表结构中没有循环。
程序尽量满足 O(n) 时间复杂度，且仅用 O(1) 内存。
*/
func getIntersectionNode(headA *ListNode, headB *ListNode, endA *ListNode, endB *ListNode) *ListNode {
	//len(listA) > len(listB)
	//listA先走lenMax-lenMin步,然后同时一起走
	//如果相交,一起走的过程一定会到达某同一个结点
	if headA == endA || headB == endB {
		return nil
	}

	lenA, lenB := listLen(headA, endA), listLen(headB, endB)
	if lenA > lenB {
		for i := 0; i < lenA-lenB; i++ {
			headA = headA.Next
		}
	} else {
		for i := 0; i < lenB-lenA; i++ {
			headB = headB.Next
		}
	}

	for headA != endA && headB != endB {
		if headA == headB {
			return headA
		}
		headA, headB = headA.Next, headB.Next
	}
	return nil
}

func listLen(head *ListNode, end *ListNode) (count int) {
	for head != end {
		count++
		head = head.Next
	}
	return
}

//判断两个环节点是否相交
func genInterSectionNodeInCycleList(headA *ListNode, headB *ListNode) *ListNode {
	//先找出各自的入环节点
	inLoopA, inLoopB := detectCycle(headA), detectCycle(headB)

	//如果入环点相同,说明二者肯定相交,并且可能在入环之前已经相交,可以转换为求两个无环链表的节点
	//形式为终止节点为入环点
	if inLoopA == inLoopB {
		node := getIntersectionNode(headA, headB, inLoopA, inLoopB)

		if node == nil {
			return inLoopA
		} else {
			return node
		}
	} else {
		//如果入环点不相同
		//可能为以下形式
		//      list1               listB
		//       |                  |
		//       V:node1            V:node2
		//      ----------------------------
		//      |                          |
		//      |                          |
		//      ----------------------------
		//
		// 或者
		//
		//    list1                list2
		//     |                     |
		//     V                     V
		//   -----                 -----
		//   |   |                 |   |
		//   -----                 -----
		//
		// 此时只要list1一直走,如果在回到头节点前没有碰到node2,说明是第二种情况,此时不相交
		// 否则,二者相交,返回node1或者node2都可以
		tmpA := headA.Next
		if tmpA != headA {
			if tmpA == inLoopB {
				//相交
				return inLoopB
			}
			tmpA = tmpA.Next
		}
		//没有碰到
		return nil
	}
}

//终极问题
//判断两个链表是否相交?
//1.判断两个链表是否有环
//2.如果一个有环,一个没有环,不可能相交
//3.如果两个都没有环,转换为无环节点相交的问题
//4.如果两个都有环,转换未有环节点相交的问题
