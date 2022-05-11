package main

/*
	// t6 := TreeNode{Val: 1000}
	t5 := TreeNode{Val: 15}
	t4 := TreeNode{Val: 7}
	t3 := TreeNode{Val: 20, Left: &t5, Right: &t4}
	// t2 := TreeNode{Val: 9, Left: &t6}
	t2 := TreeNode{Val: 9}
	t1 := TreeNode{Val: -10, Left: &t2, Right: &t3}

	r := minDepth(&t1)
	fmt.Printf("result: %v\n", r)
*/

// 兩邊分開走，再比較大小
// Runtime: 205 ms, faster than 75.94% of Go online submissions for Minimum Depth of Binary Tree.
// Memory Usage: 20.3 MB, less than 68.16% of Go online submissions for Minimum Depth of Binary Tree.
func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	l := minDepth(root.Left)
	r := minDepth(root.Right)

	if l == 0 || r == 0 {
		return l + r + 1
	}

	return min(l, r) + 1
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

// 兩邊分開走，for迴圈先走到底部時就直接回傳，使用list儲存要繼續往下走的點
// Runtime: 279 ms, faster than 41.51% of Go online submissions for Minimum Depth of Binary Tree.
// Memory Usage: 22.5 MB, less than 39.62% of Go online submissions for Minimum Depth of Binary Tree.
func minDepthV2(root *TreeNode) int {
	if root == nil {
		return 0
	}

	nodes := []*TreeNode{root}
	depth := 1
	for len(nodes) != 0 {
		nodeNeedMove := len(nodes)
		for i := range nodes {
			if nodes[i].Left == nil && nodes[i].Right == nil {
				return depth
			}
			if nodes[i].Left != nil {
				nodes = append(nodes, nodes[i].Left)
			}
			if nodes[i].Right != nil {
				nodes = append(nodes, nodes[i].Right)
			}
		}
		nodes = nodes[nodeNeedMove:]
		depth++
	}

	return depth
}
