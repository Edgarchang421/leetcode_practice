package main

/*
	node7 := TreeNode{
		Val: 4,
	}
	node6 := TreeNode{
		Val: 3,
	}
	node5 := TreeNode{
		Val: 4,
	}
	node4 := TreeNode{
		Val: 3,
	}
	node3 := TreeNode{
		Val:   2,
		Left:  &node7,
		Right: &node6,
	}
	node2 := TreeNode{
		Val:   2,
		Left:  &node4,
		Right: &node5,
	}
	node1 := TreeNode{
		Val:   1,
		Left:  &node3,
		Right: &node2,
	}

	// node32 := TreeNode{
	// 	Val: 2,
	// }
	// node22 := TreeNode{
	// 	Val: 1,
	// }
	// node12 := TreeNode{
	// 	Val:   3,
	// 	Left:  &node22,
	// 	Right: &node32,
	// }
	r := isSymmetric(&node1)
	fmt.Printf("result: %v\n", r)
*/

// Runtime: 0 ms, faster than 100.00% of Go online submissions for Symmetric Tree.
// Memory Usage: 2.8 MB, less than 100.00% of Go online submissions for Symmetric Tree.
func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}

	return isSy(root.Left, root.Right)
}

func isSy(l *TreeNode, r *TreeNode) bool {
	if l == nil && r == nil {
		return true
	}
	if l == nil || r == nil {
		return false
	}
	if l.Val != r.Val {
		return false
	}

	return isSy(l.Left, r.Right) && isSy(l.Right, r.Left)
}

// 遇到null會失敗
func isSymmetricVerFail(root *TreeNode) bool {
	if root.Left == nil && root.Right == nil {
		return true
	}
	if root.Left == nil || root.Right == nil {
		return true
	}
	left := inOrderTraversalIterate(root.Left)
	right := inOrderTraversalIterateVer2(root.Right)
	for i := range left {
		if left[i] != right[i] {
			return false
		}
	}
	// 一邊用前序 一邊用後序
	// 存到list中比較
	return true
}

// 中序，左中右
func inOrderTraversalIterate(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	result := make([]int, 0)
	rootStack := make([]*TreeNode, 0)

	for root != nil || len(rootStack) != 0 {
		if root != nil {
			rootStack = append(rootStack, root)
			root = root.Left
		} else {
			root = rootStack[len(rootStack)-1]
			result = append(result, root.Val)
			rootStack = rootStack[:len(rootStack)-1]
			root = root.Right
		}
	}
	return result
}

// 中序，改為右中左
func inOrderTraversalIterateVer2(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	result := make([]int, 0)
	rootStack := make([]*TreeNode, 0)

	for root != nil || len(rootStack) != 0 {
		if root != nil {
			rootStack = append(rootStack, root)
			root = root.Right
		} else {
			root = rootStack[len(rootStack)-1]
			result = append(result, root.Val)
			rootStack = rootStack[:len(rootStack)-1]
			root = root.Left
		}
	}
	return result
}
