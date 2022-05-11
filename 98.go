package main

// [5,4,6,null,null,3,7]

/*
	node3 := TreeNode{
		Val: 3,
	}
	node2 := TreeNode{
		Val: 1,
	}
	node1 := TreeNode{
		Val:   2,
		Left:  &node2,
		Right: &node3,
	}

	r := isValidBST(&node1)
	fmt.Printf("result: %v\n", r)
*/

func isValidBST(root *TreeNode) bool {
	if root == nil {
		return true
	}
	if root.Left != nil {
		if root.Left.Val >= root.Val {
			return false
		}
	}
	if root.Right != nil {
		if root.Right.Val <= root.Val {
			return false
		}
	}

	return isValidBST(root.Left) && isValidBST(root.Right)
}

// Runtime: 0 ms, faster than 100.00% of Go online submissions for Validate Binary Search Tree.
// Memory Usage: 5.1 MB, less than 94.33% of Go online submissions for Validate Binary Search Tree.
func isValidBSTV2(root *TreeNode) bool {
	if root == nil {
		return true
	}

	return recuIsValidBST(root.Left, &root.Val, nil) && recuIsValidBST(root.Right, nil, &root.Val)
}

func recuIsValidBST(root *TreeNode, max, min *int) bool {
	if root == nil {
		return true
	}
	if max != nil {
		if root.Val >= *max {
			return false
		}
	}
	if min != nil {
		if root.Val <= *min {
			return false
		}
	}

	return recuIsValidBST(root.Left, &root.Val, min) && recuIsValidBST(root.Right, max, &root.Val)
}

// to do stack使用迭代解法
