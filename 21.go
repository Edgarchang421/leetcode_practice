package main

import "fmt"

/*
Input: list1 = [1,2,4], list2 = [1,3,4]
Output: [1,1,2,3,4,4]

Input: list1 = [], list2 = []
Output: []

Input: list1 = [], list2 = [0]
Output: [0]

	a3 := &ListNode{Val: 4, Next: nil}
	a2 := &ListNode{Val: 2, Next: a3}
	a1 := &ListNode{Val: 1, Next: a2}

	b3 := &ListNode{Val: 4, Next: nil}
	b2 := &ListNode{Val: 3, Next: b3}
	b1 := &ListNode{Val: 1, Next: b2}
*/

// Runtime: 0 ms, faster than 100.00% of Go online submissions for Merge Two Sorted Lists.
// Memory Usage: 2.6 MB, less than 46.24% of Go online submissions for Merge Two Sorted Lists.
func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	if list1 == nil && list2 == nil {
		return nil
	}
	if list1 == nil {
		return list2
	}
	if list2 == nil {
		return list1
	}

	var head *ListNode
	var node *ListNode

	for list1 != nil && list2 != nil {
		fmt.Printf("l1: %v, l2: %v\n", list1.Val, list2.Val)
		if list1.Val >= list2.Val {
			newNode := &ListNode{Val: list2.Val}
			if node == nil {
				head = newNode
				node = newNode
			} else {
				node.Next = newNode
				node = newNode
			}
			list2 = list2.Next
		} else {
			newNode := &ListNode{Val: list1.Val}
			if node == nil {
				head = newNode
				node = newNode
			} else {
				node.Next = newNode
				node = newNode
			}
			list1 = list1.Next
		}
	}

	if list1 != nil {
		node.Next = list1
	}
	if list2 != nil {
		node.Next = list2
	}

	return head
}

func mergeTwoListsVer2(list1 *ListNode, list2 *ListNode) *ListNode {
	if list1 == nil {
		return list2
	}
	if list2 == nil {
		return list1
	}

	if list1.Val <= list2.Val {
		list1.Next = mergeTwoListsVer2(list1.Next, list2)
		return list1
	} else {
		list2.Next = mergeTwoListsVer2(list1, list2.Next)
		return list2
	}
}
