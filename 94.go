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
		Val: 2,
	}
	node22 := TreeNode{
		Val: 1,
	}
	node12 := TreeNode{
		Val:   3,
		Left:  &node22,
		Right: &node32,
	}

	r := recuTreeInorderVer3(&node1)
	fmt.Printf("result: %v\n", r)

	r = recuTreeInorderVer3(&node12)
	fmt.Printf("result: %v\n", r)
*/

// Runtime: 1 ms, faster than 41.03% of Go online submissions for Binary Tree Inorder Traversal.
// Memory Usage: 1.9 MB, less than 80.06% of Go online submissions for Binary Tree Inorder Traversal.
func inorderTraversal(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	result := make([]int, 0)
	result = recuTreeInorder(root, result)

	return result
}

func recuTreeInorder(node *TreeNode, nums []int) []int {
	if node.Left != nil {
		nums = recuTreeInorder(node.Left, nums)
	}
	nums = append(nums, node.Val)
	if node.Right != nil {
		nums = recuTreeInorder(node.Right, nums)
	}
	return nums
}

// Runtime: 0 ms, faster than 100.00% of Go online submissions for Binary Tree Inorder Traversal.
// Memory Usage: 2 MB, less than 26.91% of Go online submissions for Binary Tree Inorder Traversal.
func inorderTraversal2(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	result := make([]int, 0)
	result = recuTreeInorder2(root, &result)

	return result
}

func recuTreeInorder2(node *TreeNode, nums *[]int) []int {
	if node.Left != nil {
		*nums = recuTreeInorder2(node.Left, nums)
	}
	*nums = append(*nums, node.Val)
	if node.Right != nil {
		*nums = recuTreeInorder2(node.Right, nums)
	}
	return *nums
}

// Runtime: 2 ms, faster than 35.21% of Go online submissions for Binary Tree Inorder Traversal.
// Memory Usage: 2 MB, less than 80.06% of Go online submissions for Binary Tree Inorder Traversal.
func recuTreeInorderVer3(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	nums := make([]int, 0)
	if root.Left != nil {
		nums = append(nums, recuTreeInorderVer3(root.Left)...)
	}
	nums = append(nums, root.Val)
	if root.Right != nil {
		nums = append(nums, recuTreeInorderVer3(root.Right)...)
	}

	return nums
}

// Runtime: 2 ms, faster than 35.21% of Go online submissions for Binary Tree Inorder Traversal.
// Memory Usage: 1.9 MB, less than 80.06% of Go online submissions for Binary Tree Inorder Traversal.
func inorderTraversalIterate(root *TreeNode) []int {
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

// ????????????????????????root?????????left
func inorderTraversalIterateVer2(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	result := make([]int, 0)
	rootStack := make([]*TreeNode, 0)

	for {
		// ????????????
		if root.Left != nil {
			rootStack = append(rootStack, root)
			root = root.Left
			continue
		}

		// ???????????????????????????????????????result
		result = append(result, root.Val)

		// ????????????????????????????????????
		if root.Right != nil {
			root = root.Right
			continue
		}

		// ???????????????????????????????????????????????????result
		// ???????????????????????????
		if len(rootStack) == 0 {
			break
		}
		// ????????????
		root = rootStack[len(rootStack)-1]
		// ??????????????????
		rootStack = rootStack[:len(rootStack)-1]
	}

	return result
}
