package main

import "fmt"

/*
	// a5 := ListNode{5, nil}
	// a4 := ListNode{4, &a5}
	// a3 := ListNode{3, &a4}
	// a2 := ListNode{2, &a3}
	// a1 := ListNode{1, &a2}

	// r := reverseBetween(&a1, 2, 4)
	// fmt.Printf("result: %v %v %v %v %v", r, r.Next, r.Next.Next, r.Next.Next.Next, r.Next.Next.Next.Next)

	// a1 := ListNode{1, nil}
	// r := reverseBetween(&a1, 1, 1)
	// fmt.Printf("result: %v\n", r)

	a2 := ListNode{5, nil}
	a1 := ListNode{3, &a2}
	// r := reverseBetweenVer2(&a1, 1, 2)

	// for r != nil {
	// 	fmt.Printf("node: %v ", r)
	// 	r = r.Next
	// }

	// a := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	// fmt.Print(arrayReverse(a, 0, 9))

	// a5 := ListNode{5, nil}
	// a4 := ListNode{4, &a5}
	// a3 := ListNode{3, &a4}
	// a2 := ListNode{2, &a3}
	// a1 := ListNode{1, &a2}

	r := LinkedListReverseVer2Refactor(&a1, 1, 2)
	for r != nil {
		fmt.Printf("node: %v ", r)
		r = r.Next
	}
*/

// Runtime: 3 ms, faster than 31.67% of Go online submissions for Reverse Linked List II.
// Memory Usage: 2.1 MB, less than 13.57% of Go online submissions for Reverse Linked List II.
func reverseBetween(head *ListNode, left int, right int) *ListNode {
	ans, _, _ := LinkedListReverse(head, left, right)
	return ans
}

// Recursion func ver1
// Runtime: 3 ms, faster than 31.67% of Go online submissions for Reverse Linked List II.
// Memory Usage: 2.1 MB, less than 13.57% of Go online submissions for Reverse Linked List II.
func LinkedListReverse(node *ListNode, left, right int) (*ListNode, int, int) {
	if left == right || left > right {
		return node, left, right
	}
	head := node
	var leftNode *ListNode

	count := 1
	for node != nil {
		if count == left {
			leftNode = node
		}
		if count == right {
			node.Val, leftNode.Val = leftNode.Val, node.Val
			break
		}
		node = node.Next
		count++
	}
	left += 1
	right -= 1

	return LinkedListReverse(head, left, right)
}

// Runtime: 4 ms, faster than 13.12% of Go online submissions for Reverse Linked List II.
// Memory Usage: 2.1 MB, less than 13.57% of Go online submissions for Reverse Linked List II.

// Runtime: 1 ms, faster than 44.34% of Go online submissions for Reverse Linked List II.
// Memory Usage: 2 MB, less than 76.02% of Go online submissions for Reverse Linked List II.
func LinkedListReverseVer2(head *ListNode, left, right int) *ListNode {
	if left == right || left > right {
		return head
	}
	first := head
	var leftNode *ListNode

	count := 1
	for head != nil {
		if count == left {
			leftNode = head
		}
		if count == right {
			head.Val, leftNode.Val = leftNode.Val, head.Val
			break
		}
		head = head.Next
		count++
	}
	left += 1
	right -= 1

	LinkedListReverse(first, left, right)
	return first
}

// final answer?????????fmt??????????????????0ms
// Runtime: 0 ms, faster than 100.00% of Go online submissions for Reverse Linked List II.
// Memory Usage: 2 MB, less than 76.02% of Go online submissions for Reverse Linked List II.
func LinkedListReverseVer2Refactor(head *ListNode, left, right int) *ListNode {
	// ?????????????????????
	if left == right || left > right {
		return head
	}
	// ?????????head?????????????????????????????????linkedlist
	originalHeadNode := head

	// ???????????????????????????left???????????????node
	var leftNode *ListNode

	for count := 1; count <= right; count++ {
		// ?????? ???left??? node
		if count == left {
			leftNode = head
		}
		// ????????? ???right??? node
		// ??????val??????
		if count == right {
			head.Val, leftNode.Val = leftNode.Val, head.Val
			break
		}
		// ??????linkedlist
		head = head.Next
	}

	// ?????????????????????node???????????????
	left += 1
	right -= 1

	// ????????????
	LinkedListReverse(originalHeadNode, left, right)
	return originalHeadNode
}

// error answer

// left right???index???????????????left????????????node???right???????????????node??????????????????node
func reverseBetweenVer2(head *ListNode, left int, right int) *ListNode {
	if head == nil || left > right {
		return nil
	}
	if head.Next == nil || left == right {
		return head
	}
	firstNode := head
	count := 1
	var leftNode *ListNode
	var rightNode *ListNode
	var lastNode *ListNode
	for head != nil {
		if count == left {
			leftNode = head
		}
		if count == right {
			rightNode = head
		}

		if head.Next == nil {
			lastNode = head
		}
		head = head.Next
		count++
	}
	rightNode.Next = nil

	var reverseListHead *ListNode
	if right != count {
		reverseListHead = newReverseList(leftNode, lastNode)
		fmt.Printf("reverse %v\n", reverseListHead)
	} else {
		reverseListHead = newReverseList(leftNode, nil)
		fmt.Printf("reverse %v", reverseListHead)
	}

	if left != 1 {
		firstNode.Next = reverseListHead
	}

	return firstNode
}

func newReverseList(head *ListNode, lastNode *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	for {
		currentListNode := &ListNode{
			Val:  head.Val,
			Next: lastNode,
		}
		lastNode = currentListNode

		head = head.Next
		if head == nil {
			return currentListNode
		}
	}
}

func arrayReverse(data []int, left, right int) ([]int, int, int) {
	if left == right || left > right {
		return data, left, right
	}
	data[left], data[right] = data[right], data[left]
	left += 1
	right -= 1
	return arrayReverse(data, left, right)
}

// ???stack?????????????????????????????????????????????????????????????????????????????????????????????????????????
