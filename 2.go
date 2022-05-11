package main

import "fmt"

/*
// a3 := ListNode{3, nil}
	// a2 := ListNode{4, &a3}
	// a1 := ListNode{2, &a2}

	// a1 := ListNode{0, nil}

	a2 := ListNode{6, nil}
	a1 := ListNode{5, &a2}

	// b3 := ListNode{4, nil}
	// b2 := ListNode{6, &b3}
	// b1 := ListNode{5, &b2}

	// b1 := ListNode{0, nil}

	b3 := ListNode{9, nil}
	b2 := ListNode{4, &b3}
	b1 := ListNode{5, &b2}

	c := addTwoReverseLinkedList(&a1, &b1)
	fmt.Printf("first node: %v\n", c)
	// fmt.Printf("n1: %v, n2: %v, n3: %v\n", c, c.Next, c.Next.Next)
	fmt.Printf("\nans: %v\n", getNum(c))
*/
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	a := getNum(l1)
	b := getNum(l2)

	ans := a + b

	return getListNode(ans)
}

//  溢位無法處理
func getNum(node *ListNode) int {
	var ans int

	var multiple int = 1

	for {
		ans += node.Val * multiple
		if node.Next == nil {
			break
		}
		multiple *= 10
		node = node.Next
		// fmt.Printf("ans: %v\n", ans)
	}

	return ans
}

func getListNode(num int) *ListNode {
	var firstNode *ListNode
	var oldNode *ListNode
	var newNode *ListNode
	for {
		// fmt.Printf("num: %v\n", num)
		// 新節點
		newNode = &ListNode{num % 10, nil}
		// fmt.Printf("newNode: %v\n", newNode)
		// 下一次會被處理的數字等於自己除以10
		num /= 10
		// fmt.Printf("new num: %v\n", num)

		// 沒有舊節點，就會是第一個節點，直接設定即可
		if oldNode == nil {
			// fmt.Printf("oldNode nil process\n")
			firstNode = newNode
			oldNode = newNode
		} else { // 有舊節點，舊節點的Next要設定
			// fmt.Printf("oldNode not nil process\n")
			oldNode.Next = newNode
			oldNode = newNode
		}
		// fmt.Printf("after list node process newNode: %v\n", newNode)
		// 沒有餘數就break迴圈
		if num == 0 {
			break
		}
	}
	return firstNode
}

// Runtime: 39 ms, faster than 5.76% of Go online submissions for Add Two Numbers.
// Memory Usage: 5.4 MB, less than 6.00% of Go online submissions for Add Two Numbers.
func addTwoReverseLinkedList(l1 *ListNode, l2 *ListNode) *ListNode {
	var firstNode *ListNode
	var oldNode *ListNode
	var newNode *ListNode

	var carry bool
	var r int

	var v1 int
	var v2 int

	for {
		// new list init
		if oldNode == nil {
			// 初始node沒有進位問題
			r = l1.Val + l2.Val

			// 相加結果大於10要進位
			if r >= 10 {
				carry = true
				r = r % 10
			} else {
				carry = false
			}
			newNode = &ListNode{
				Val:  r,
				Next: nil,
			}
			firstNode = newNode
			oldNode = newNode
			l1 = l1.Next
			l2 = l2.Next
		} else if oldNode != nil {
			if l1 == nil && l2 == nil {
				break
			}
			fmt.Printf("old node not nil process\n")

			// 判斷是否兩個list都有新node要相加
			if l1 != nil {
				v1 = l1.Val
			} else {
				v1 = 0
			}
			if l2 != nil {
				v2 = l2.Val
			} else {
				v2 = 0
			}

			// 判斷進位
			if carry {
				r = v1 + v2 + 1
			} else {
				r = v1 + v2
			}
			// 相加結果大於10要進位
			if r >= 10 {
				carry = true
				r = r % 10
			} else {
				carry = false
			}
			newNode = &ListNode{
				Val:  r,
				Next: nil,
			}
			oldNode.Next = newNode
			fmt.Printf("old node: %v\n", oldNode)
			fmt.Printf("new node: %v\n", newNode)
			oldNode = newNode
			// 還有一個list有element就要繼續相加
			if l1 != nil {
				l1 = l1.Next
			}
			if l2 != nil {
				l2 = l2.Next
			}

		}
	}

	if carry {
		newNode = &ListNode{
			Val:  1,
			Next: nil,
		}
		oldNode.Next = newNode
	}

	return firstNode
}

// 有fmt print效能差很多
// Runtime: 28 ms, faster than 5.76% of Go online submissions for Add Two Numbers.
// Memory Usage: 5.4 MB, less than 6.00% of Go online submissions for Add Two Numbers.

// 沒有fmt
// Runtime: 15 ms, faster than 44.92% of Go online submissions for Add Two Numbers.
// Memory Usage: 4.5 MB, less than 78.08% of Go online submissions for Add Two Numbers.
func addTwoReverseLinkedListVer2(l1 *ListNode, l2 *ListNode) *ListNode {
	// 進位
	var carry int

	// init list process
	// 初始node沒有進位問題
	r := l1.Val + l2.Val
	// 相加結果大於10要進位
	if r >= 10 {
		carry = 1
		r = r % 10
	} else {
		carry = 0
	}
	// first node init
	newNode := &ListNode{
		Val:  r,
		Next: nil,
	}
	firstNode := newNode
	oldNode := newNode
	l1 = l1.Next
	l2 = l2.Next

	// 都nil了就break
	for !(l1 == nil && l2 == nil) {
		var v1 int
		var v2 int

		// 判斷是否兩個list都有新node要相加
		if l1 != nil {
			v1 = l1.Val
		} else {
			v1 = 0
		}
		if l2 != nil {
			v2 = l2.Val
		} else {
			v2 = 0
		}

		// 判斷進位
		r = v1 + v2 + carry
		// 相加結果大於10要進位
		if r >= 10 {
			carry = 1
			r = r % 10
		} else {
			carry = 0
		}
		newNode = &ListNode{
			Val:  r,
			Next: nil,
		}
		oldNode.Next = newNode
		oldNode = newNode
		// 如果node是nil時，讀取會記憶體錯誤
		if l1 != nil {
			l1 = l1.Next
		}
		if l2 != nil {
			l2 = l2.Next
		}
		// fmt.Printf("old node: %v\n", oldNode)
		// fmt.Printf("new node: %v\n", newNode)
	}

	// 最後一個如果還有進位
	if carry == 1 {
		newNode = &ListNode{
			Val:  1,
			Next: nil,
		}
		oldNode.Next = newNode
	}

	return firstNode
}

// 刪掉註解效能也差很多
// Runtime: 7 ms, faster than 90.48% of Go online submissions for Add Two Numbers.
// Memory Usage: 4.5 MB, less than 78.08% of Go online submissions for Add Two Numbers.
func addTwoReverseLinkedListClear(l1 *ListNode, l2 *ListNode) *ListNode {
	var carry int

	r := l1.Val + l2.Val
	if r >= 10 {
		carry = 1
		r = r % 10
	} else {
		carry = 0
	}
	newNode := &ListNode{
		Val:  r,
		Next: nil,
	}
	firstNode := newNode
	oldNode := newNode
	l1 = l1.Next
	l2 = l2.Next

	for !(l1 == nil && l2 == nil) {
		var v1 int
		var v2 int

		if l1 != nil {
			v1 = l1.Val
			l1 = l1.Next
		} else {
			v1 = 0
		}
		if l2 != nil {
			v2 = l2.Val
			l2 = l2.Next
		} else {
			v2 = 0
		}

		r = v1 + v2 + carry
		if r >= 10 {
			carry = 1
			r = r % 10
		} else {
			carry = 0
		}
		newNode = &ListNode{
			Val:  r,
			Next: nil,
		}
		oldNode.Next = newNode
		oldNode = newNode
	}

	if carry == 1 {
		newNode = &ListNode{
			Val:  1,
			Next: nil,
		}
		oldNode.Next = newNode
	}

	return firstNode
}

// 最簡潔的example，但leetcode上測試的結果比 addTwoReverseLinkedListClear 差
// Runtime: 14 ms, faster than 48.47% of Go online submissions for Add Two Numbers.
// Memory Usage: 4.5 MB, less than 78.08% of Go online submissions for Add Two Numbers.
func final(l1 *ListNode, l2 *ListNode) *ListNode {
	var carry int

	// 這個first node不會有val，只會儲存真正第一個node的pointer
	firstNode := &ListNode{}
	node := firstNode

	for l1 != nil || l2 != nil {
		if l1 != nil {
			carry += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			carry += l2.Val
			l2 = l2.Next
		}

		node.Next = &ListNode{Val: carry % 10}
		node = node.Next

		carry /= 10
	}

	if carry == 1 {
		node.Next = &ListNode{Val: 1}
	}

	return firstNode.Next
}
