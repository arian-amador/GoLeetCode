package mergeTwoSortedLists

import (
	"strconv"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(l1, l2 *ListNode) *ListNode {
	if l1 == nil && l2 == nil {
		return nil
	}

	sorted := &ListNode{}
	result := sorted

	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			sorted.Next = l1
			l1 = l1.Next
			sorted = sorted.Next
		} else {
			sorted.Next = l2
			l2 = l2.Next
			sorted = sorted.Next
		}
	}

	if l1 == nil {
		sorted.Next = l2
	}
	if l2 == nil {
		sorted.Next = l1
	}

	return result.Next
}

func newList(s []int) *ListNode {
	r := &ListNode{}
	l := r

	for i, v := range s {
		l.Val = v
		if i != len(s)-1 {
			l.Next = &ListNode{}
			l = l.Next
		}
	}

	return r
}

func (l *ListNode) output() []string {
	s := []string{}

	for l != nil {
		s = append(s, strconv.Itoa(l.Val))
		l = l.Next
	}

	return s
}
