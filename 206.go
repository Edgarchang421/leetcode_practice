package main

import "fmt"

/*
	a3 := &ListNode{3, nil}
	a2 := &ListNode{4, a3}
	a1 := &ListNode{2, a2}

	r := reverseListVer2(a1)

	fmt.Printf("result: %v %v %v\n", r, r.Next, r.Next.Next)
	// fmt.Printf("result: %v\n", r)
*/

func reverseList(head *ListNode) *ListNode {
	nextListNode := &ListNode{
		Val: head.Val,
	}
	head = head.Next
	if head == nil {
		return nextListNode
	}

	for {
		fmt.Printf("before for loop process next node: %v\n", nextListNode)
		currentListNode := ListNode{
			Val:  head.Val,
			Next: nextListNode,
		}
		nextListNode = &currentListNode
		fmt.Printf("new next node: %v\n", nextListNode)
		fmt.Printf("current node: %v\n", currentListNode)

		head = head.Next
		fmt.Printf("head next node: %v\n", head)
		if head == nil {
			return &currentListNode
		}
	}
}

// Runtime: 3 ms, faster than 51.88% of Go online submissions for Reverse Linked List.
// Memory Usage: 2.6 MB, less than 79.19% of Go online submissions for Reverse Linked List.
func reverseListClean(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	nextListNode := &ListNode{
		Val: head.Val,
	}
	head = head.Next
	if head == nil {
		return nextListNode
	}

	for {
		currentListNode := &ListNode{
			Val:  head.Val,
			Next: nextListNode,
		}
		nextListNode = currentListNode

		head = head.Next
		if head == nil {
			return currentListNode
		}
	}
}

// Runtime: 0 ms, faster than 100.00% of Go online submissions for Reverse Linked List.
// Memory Usage: 2.6 MB, less than 79.19% of Go online submissions for Reverse Linked List.
func reverseListVer2(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	var nextListNode *ListNode

	for {
		currentListNode := &ListNode{
			Val:  head.Val,
			Next: nextListNode,
		}
		nextListNode = currentListNode

		head = head.Next
		if head == nil {
			return currentListNode
		}
	}
}

// Runtime: 3 ms, faster than 51.88% of Go online submissions for Reverse Linked List.
// Memory Usage: 2.5 MB, less than 100.00% of Go online submissions for Reverse Linked List.
func reverseListVer3(head *ListNode) (prev *ListNode) {
	for head != nil {
		head.Next, prev, head = prev, head, head.Next
	}
	return
}
