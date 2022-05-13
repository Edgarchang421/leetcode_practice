package main

import "fmt"

/*
Example 1:
Input: root = [1,2,3,null,5]
Output: ["1->2->5","1->3"]

Example 2:
Input: root = [1]
Output: ["1"]

	t5 := TreeNode{Val: 5}
	t2 := TreeNode{Val: 2}
	t3 := TreeNode{Val: 3}
	t1 := TreeNode{Val: 1}
	t1.Left = &t2
	t1.Right = &t3
	t2.Right = &t5

	r := binaryTreePaths(&t1)
	for i := range r {
		fmt.Printf("possible path: %v\n", r[i])
	}
*/

// Runtime: 3 ms, faster than 49.31% of Go online submissions for Binary Tree Paths.
// Memory Usage: 2.4 MB, less than 70.51% of Go online submissions for Binary Tree Paths.
func binaryTreePaths(root *TreeNode) []string {
	if root == nil {
		return nil
	}
	possiblePath := fmt.Sprintf("%v", root.Val)
	result := make([]string, 0)
	if root.Left == nil && root.Right == nil {
		result = append(result, possiblePath)
		return result
	}

	dfs257(&result, possiblePath, root.Right)
	dfs257(&result, possiblePath, root.Left)
	return result
}

// 選擇: 所有可能路徑
// 限制: 節點不為nil
// 目標: left == right == null
func dfs257(result *[]string, possiblePath string, root *TreeNode) {
	if root == nil {
		return
	}
	possiblePath = fmt.Sprintf("%v->%v", possiblePath, root.Val)
	if root.Left == nil && root.Right == nil {
		*result = append(*result, possiblePath)
		return
	}
	dfs257(result, possiblePath, root.Left)
	dfs257(result, possiblePath, root.Right)
}
