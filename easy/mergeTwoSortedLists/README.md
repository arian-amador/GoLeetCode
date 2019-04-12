# EASY - Merge Two Sorted Lists [![Build Status](https://api.travis-ci.org/arian-amador/GoLeetCode.svg)](https://travis-ci.org/arian-amador/GoLeetCode)

### Problem

Merge two sorted linked lists and return it as a new list. The new list should be made by splicing together the nodes of the first two lists.

##### Example

```Go
Input: 1->2->4, 1->3->4

Output: 1->1->2->3->4->4
```

### Solutions

To merge two sorted linked lists we start with a blank link list node. Looping while both of the current list nodes exist. Each loop iteration we check which list node has a smaller value. We set the smaller node to the next element in its list and add the element to the sorted list.

```Go
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
```

After we iterate over both lists or find a nil list node. The loop ends but we still have to check if either of the list nodes still has an node with a value. This represents the largest value from both lists. We check both list nodes and set the last sorted linked list to that value.

```Go
	if l1 == nil {
		sorted.Next = l2
	}
	if l2 == nil {
		sorted.Next = l1
	}
```

---

#### Hope you find this useful!
