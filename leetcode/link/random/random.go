package random

/*
138. 复制带随机指针的链表
给定一个链表，每个节点包含一个额外增加的随机指针，该指针可以指向链表中的任何节点或空节点。

要求返回这个链表的深度拷贝。
*/

type ListNode struct {
	Val    int
	Next   *ListNode
	Random *ListNode
}

// 很有意思的一道题
//
//      |----|
//      |    V
// A -> B -> C -> nil     ->   A -> A1-> B -> B1 -> C -> C1 -> nil
// |    ^    |     ^
// |    |    |     |
// |----     |-----|
//
// 然后就可以有 A1->random = A->random->next,当然考虑random指向nil的情况
//最后再把A1,B1,C1分离出来
func copyRandomList(head *ListNode) *ListNode {
	copyList, root := &ListNode{}, head

	//复制各个节点
	for head != nil {
		tmp := head.Next
		dup := &ListNode{Val: head.Val}
		dup.Next = head.Next
		head.Next = dup

		head = tmp
	}

	//复制random
	for head = root; head != nil; {
		if head.Random == nil {
			//这里注意
			head.Next.Random = nil
		} else {
			head.Next.Random = head.Random.Next
		}
	}

	//分离
	head, copyHead := root, copyList
	for head != nil {
		dup := head.Next
		head.Next = head.Next.Next //list数量必定为偶数,所以这里也肯定不会panic
		head = head.Next

		copyHead.Next = dup
		copyHead = copyHead.Next
	}

	return copyList.Next
}
