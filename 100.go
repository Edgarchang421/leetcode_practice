package main

/*
	node3 := TreeNode{
		Val: 3,
	}
	node2 := TreeNode{
		Val:  2,
		Left: &node3,
	}
	node1 := TreeNode{
		Val:   1,
		Right: &node2,
	}

	node32 := TreeNode{
		Val: 3,
	}
	node22 := TreeNode{
		Val:  2,
		Left: &node32,
	}
	node12 := TreeNode{
		Val:   1,
		Right: &node22,
	}

	r := isSameTree(&node1, &node12)
	fmt.Printf("result: %v\n", r)
*/

// Runtime: 3 ms, faster than 31.26% of Go online submissions for Same Tree.
// Memory Usage: 2.1 MB, less than 69.28% of Go online submissions for Same Tree.
func isSameTree(p *TreeNode, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	} else if (p != nil && q == nil) || (p == nil && q != nil) {
		return false
	}
	if p.Val != q.Val {
		return false
	}

	return isSameTree(p.Right, q.Right) && isSameTree(p.Left, q.Left)
}

// Runtime: 3 ms, faster than 31.26% of Go online submissions for Same Tree.
// Memory Usage: 2 MB, less than 69.28% of Go online submissions for Same Tree.
func isSameTreeV2(p *TreeNode, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	} else if p == nil || q == nil {
		return false
	}
	if p.Val != q.Val {
		return false
	}

	return isSameTreeV2(p.Right, q.Right) && isSameTreeV2(p.Left, q.Left)
}

func isSameTreeV3(p *TreeNode, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	} else if p == nil || q == nil {
		return false
	}
	if p.Val != q.Val {
		return false
	}

	return isSameTree(p.Right, q.Right) && isSameTree(p.Left, q.Left)
}

// Runtime: 0 ms, faster than 100.00% of Go online submissions for Same Tree.
// Memory Usage: 2.1 MB, less than 11.10% of Go online submissions for Same Tree.
func isSameTreeV4(p *TreeNode, q *TreeNode) bool {
	// 先進先出，取list[0]
	qP := []*TreeNode{p}
	qQ := []*TreeNode{q}

	for len(qP) != 0 && len(qQ) != 0 {
		pNode := qP[0]
		qP = qP[1:]

		qNode := qQ[0]
		qQ = qQ[1:]

		if pNode == nil && qNode == nil {
			continue
		}
		if pNode == nil && qNode != nil || pNode != nil && qNode == nil {
			return false
		}
		if pNode.Val != qNode.Val {
			return false
		}

		qP = append(qP, pNode.Left, pNode.Right)
		qQ = append(qQ, qNode.Left, qNode.Right)
	}

	if len(qP) == 0 && len(qQ) == 0 {
		return true
	}
	return false
}
