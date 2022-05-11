package main

/*
	t6 := TreeNode{Val: 1000}
	t5 := TreeNode{Val: 15}
	t4 := TreeNode{Val: 7}
	t3 := TreeNode{Val: 20, Left: &t5, Right: &t4}
	t2 := TreeNode{Val: 9, Left: &t6}
	t1 := TreeNode{Val: -10, Left: &t2, Right: &t3}
	r := isBalanced(&t1)
	fmt.Printf("result: %v\n", r)
*/

// Runtime: 4 ms, faster than 90.26% of Go online submissions for Balanced Binary Tree.
// Memory Usage: 6.2 MB, less than 21.21% of Go online submissions for Balanced Binary Tree.
func isBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}
	leftHeight := bfs(root.Left)
	// fmt.Printf("one side over\n")
	rightHeight := bfs(root.Right)

	result := leftHeight - rightHeight

	if result > 1 {
		return false
	} else if result < -1 {
		return false
	}
	return isBalanced(root.Left) && isBalanced(root.Right)
}

func bfs(n *TreeNode) (h int) {
	if n == nil {
		return
	}
	nodes := []*TreeNode{n}
	for len(nodes) != 0 {
		times := len(nodes)
		for i := 0; i < times; i++ {
			if nodes[i].Left != nil {
				nodes = append(nodes, nodes[i].Left)
			}
			if nodes[i].Right != nil {
				nodes = append(nodes, nodes[i].Right)
			}
			// fmt.Printf("node info: %v\n", nodes[i])
		}
		// fmt.Printf("one layer over\n")
		nodes = nodes[times:]
		h++
	}
	return
}

// Runtime: 10 ms, faster than 41.34% of Go online submissions for Balanced Binary Tree.
// Memory Usage: 5.8 MB, less than 21.21% of Go online submissions for Balanced Binary Tree.
func isBalancedRecuVer(root *TreeNode) bool {
	if root == nil {
		return true
	}

	l := dfsRecu(root.Left)
	r := dfsRecu(root.Right)
	if l-r > 1 {
		return false
	} else if l-r < -1 {
		return false
	}

	return isBalancedRecuVer(root.Left) && isBalancedRecuVer(root.Right)
}

func dfsRecu(root *TreeNode) int {
	if root == nil {
		return 0
	}
	l := dfsRecu(root.Left)
	r := dfsRecu(root.Right)

	if l > r {
		return l + 1
	}
	return r + 1
}

// Runtime: 0 ms, faster than 100.00% of Go online submissions for Balanced Binary Tree.
// Memory Usage: 5.7 MB, less than 70.35% of Go online submissions for Balanced Binary Tree.
func isBalancedRecuVer2(root *TreeNode) bool {
	return dfsRecu2(root) != -1
}

func dfsRecu2(root *TreeNode) int {
	if root == nil {
		return 0
	}
	l := dfsRecu2(root.Left)
	r := dfsRecu2(root.Right)
	if l-r > 1 || l-r < -1 || l == -1 || r == -1 {
		return -1
	}
	if l > r {
		return l + 1
	}
	return r + 1
}
