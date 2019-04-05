package addTwoNums

import (
	"strconv"
)

// ListNode describes a node in a linked list
type ListNode struct {
	Val  int
	Next *ListNode
}

// AddTwoLists adds two linked lists of numbers
func AddTwoLists(l1, l2 *ListNode) *ListNode {
	var sum, carry int
	curr := &ListNode{}
	head := curr

	for carry != 0 || l1 != nil || l2 != nil {
		curr.Next = &ListNode{}
		curr = curr.Next
		sum = 0

		if l1 != nil {
			sum += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			sum += l2.Val
			l2 = l2.Next
		}
		if carry != 0 {
			sum += carry
		}

		carry = sum / 10
		curr.Val = sum % 10
	}

	return head.Next
}

// NewList builds a singly linked list from slice of ints
func NewList(nums []int) *ListNode {
	l := ListNode{}
	l.new(nums)
	return &l
}

func (l *ListNode) new(nums []int) {
	for _, v := range nums {
		l.append(v)
	}
}

func (l *ListNode) append(n int) {
	for {
		if l.Next == nil {
			l.Val = n
			l.Next = &ListNode{}
			break
		}

		l = l.Next
	}
}

func (l *ListNode) debug() string {
	var r string
	for {
		if l.Next == nil {
			break
		}
		r += strconv.Itoa(l.Val)
		l = l.Next
	}
	return r
}
