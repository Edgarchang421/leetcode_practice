package main

/*
	n5 := ListNode{Val: 3, Next: nil}
	n4 := ListNode{Val: 3, Next: &n5}
	n3 := ListNode{Val: 2, Next: &n4}
	n2 := ListNode{Val: 1, Next: &n3}
	n1 := ListNode{Val: 1, Next: &n2}
	r := deleteDuplicatesIterateV2(&n1)
	for r != nil {
		fmt.Printf("result: %v\n", r.Val)
		r = r.Next
	}
*/

// Runtime: 6 ms, faster than 41.48% of Go online submissions for Remove Duplicates from Sorted List.
// Memory Usage: 3.6 MB, less than 5.52% of Go online submissions for Remove Duplicates from Sorted List.
func deleteDuplicatesIterateV1(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	numMap := make(map[int]struct{})

	firstHead := head

	lastNode := head
	for head != nil {
		if _, ok := numMap[head.Val]; !ok {
			numMap[head.Val] = struct{}{}
			lastNode = head
		} else {
			lastNode.Next = head.Next
			head = head.Next
		}
	}

	return firstHead
}

//Runtime: 6 ms, faster than 41.48% of Go online submissions for Remove Duplicates from Sorted List.
// Memory Usage: 3 MB, less than 76.18% of Go online submissions for Remove Duplicates from Sorted List.
func deleteDuplicatesIterateV2(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	firstHead := head
	for head != nil && head.Next != nil {
		if head.Val == head.Next.Val {
			head.Next = head.Next.Next
		} else {
			head = head.Next
		}
	}
	return firstHead
}

// Runtime: 3 ms, faster than 82.02% of Go online submissions for Remove Duplicates from Sorted List.
// Memory Usage: 3 MB, less than 76.18% of Go online submissions for Remove Duplicates from Sorted List.

// Runtime: 0 ms, faster than 100.00% of Go online submissions for Remove Duplicates from Sorted List.
// Memory Usage: 3.1 MB, less than 33.28% of Go online submissions for Remove Duplicates from Sorted List.
func deleteDuplicatesIterateV3(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	firstHead := head
	for head.Next != nil {
		if head.Val == head.Next.Val {
			head.Next = head.Next.Next
		} else {
			head = head.Next
		}
	}
	return firstHead
}

//Runtime: 3 ms, faster than 82.02% of Go online submissions for Remove Duplicates from Sorted List.
// Memory Usage: 3.2 MB, less than 33.28% of Go online submissions for Remove Duplicates from Sorted List.
func deleteDuplicatesRecuV2(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	head.Next = deleteDuplicatesRecuV2(head.Next)
	// head.val == head.next.val ? head.next : head;
	if head.Val == head.Next.Val {
		return head.Next
	} else {
		return head
	}
}

// Runtime: 8 ms, faster than 22.71% of Go online submissions for Remove Duplicates from Sorted List.
// Memory Usage: 3.2 MB, less than 33.28% of Go online submissions for Remove Duplicates from Sorted List.
func deleteDuplicatesRecuV3Example(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	} else if head.Val == head.Next.Val {
		head = deleteDuplicatesIterateV1(head.Next)
		return head
	} else {
		head.Next = deleteDuplicatesRecuV3Example(head.Next)
		return head
	}
}
