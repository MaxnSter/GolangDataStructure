package mergek_sorted

import "github.com/MaxnSter/GolangDataStructure/nmerge"

/*
23. 合并K个排序链表
合并 k 个排序链表，返回合并后的排序链表。请分析和描述算法的复杂度。

示例:

输入:
[
  1->4->5,
  1->3->4,
  2->6
]
输出: 1->1->2->3->4->4->5->6
*/

//典型mergeK问题,不难,只是操作链表有点复杂而已
type ListNode struct {
	Val  int
	Next *ListNode
}

type source struct {
	list *ListNode
}

func (s *source) Less(oth nmerge.MergeSource) bool {
	return s.list.Val < oth.(*source).list.Val
}

func (s *source) Next() bool {
	if s.list.Next == nil {
		return false
	}

	s.list = s.list.Next
	return true
}

func (s *source) OutputTo(to interface{}) {
	l := to.(**ListNode)
	(*l).Next = &ListNode{Val: s.list.Val}
	*l = (*l).Next
}

func mergeKLists(lists []*ListNode) *ListNode {
	head := &ListNode{}
	merge := head

	newList := make([]nmerge.MergeSource, 0)
	for _, l := range lists {
		if l != nil {
			newList = append(newList, &source{list: l})
		}
	}

	nmerge.Merge(newList, &merge)
	return head.Next
}
