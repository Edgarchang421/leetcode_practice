package main

import (
	"fmt"
)

/*
	n3 := TreeNode{Val: -3}
	// n2 := TreeNode{Val: 2}
	// n1 := TreeNode{Val: 1, Left: &n2, Right: &n3}
	r := maxPathSum(&n3)
	fmt.Printf("result: %v\n", r)

	// t6 := TreeNode{Val: 1000}
	// t5 := TreeNode{Val: 15}
	// t4 := TreeNode{Val: 7}
	// t3 := TreeNode{Val: 20, Left: &t5, Right: &t4}
	// t2 := TreeNode{Val: 9, Left: &t6}
	// t1 := TreeNode{Val: -10, Left: &t2, Right: &t3}
	// r := maxPathSum(&t1)
	// fmt.Printf("result: %v\n", r)
*/

var ans124 int

// Runtime: 19 ms, faster than 70.52% of Go online submissions for Binary Tree Maximum Path Sum.
// Memory Usage: 8.1 MB, less than 17.69% of Go online submissions for Binary Tree Maximum Path Sum.
func maxPathSum(root *TreeNode) int {
	ans124 = root.Val
	getSubTreeMaxSum(root)
	return ans124
}
func getSubTreeMaxSum(node *TreeNode) int {
	var leftSubTreeMaxSum int
	var rightSubTreeMaxSum int
	if node.Left != nil {
		leftSubTreeMaxSum = getSubTreeMaxSum(node.Left)
	} else {
		leftSubTreeMaxSum = 0
	}
	if node.Right != nil {
		rightSubTreeMaxSum = getSubTreeMaxSum(node.Right)
	} else {
		rightSubTreeMaxSum = 0
	}
	nodeVal := node.Val
	nodeMaxPathSum := maxNum(nodeVal, leftSubTreeMaxSum+nodeVal, rightSubTreeMaxSum+nodeVal)
	ans124 = maxNum(ans124, nodeMaxPathSum, leftSubTreeMaxSum+rightSubTreeMaxSum+nodeVal)
	return nodeMaxPathSum
}
func maxNum(num ...int) int {
	max := num[0]
	for i := range num {
		if num[i] > max {
			max = num[i]
		}
	}
	return max
}

// Runtime: 37 ms, faster than 13.38% of Go online submissions for Binary Tree Maximum Path Sum.
// Memory Usage: 8.2 MB, less than 12.70% of Go online submissions for Binary Tree Maximum Path Sum.
func gg(n *TreeNode) int {
	if n == nil {
		return 0
	}
	l := gg(n.Left)
	r := gg(n.Right)
	v := n.Val

	m := big124(big124(l, r)+v, v)
	ans124 = big124(big124(ans124, m), l+r+v)
	return m
}

func big124(x, y int) int {
	if x > y {
		return x
	}
	return y
}

// 路徑必包含兩點個點，值為兩點相加
// return不能為0，非空路徑

// 錯誤邏輯，無法從最小問題開始解
// 一層一層看，第一層和第二層 dp[2] = max( root+left , root+right , root+left+right )

// 一個點要找到一個路徑距離只有一個node的，就看附近三個
// 不能走回頭路，最長路徑 = depth * 2 - 1
// 一個點一個點看
// dp[node] = max ( dp[root]+node.Val , dp[left]+node.Val , dp[right]+node.Val )
// dp[root] = max( -10+20 , -10+9 ) = 10
// dp[left] = max( 35 )
// dp[right] = max ( 27 )

// dp[node:t5] = max ( t5.val + dp[t3] ) = 15 + 20
// dp[t3] = max ( dp[root]+node t3.Val , dp[right]+node t3.Val ,(left node重複不使用) ) = max ( 20 , -10 ) = 20

// 每一個節點出發的最大值
// dp[root] = maxNum(dp[root.Left]+root.Val, dp[root.Right]+root.Val, root.Val)

// 紀錄每個點出發的最大路經
var dp124 map[*TreeNode]int

// 無法解決root不往下走，以及無法跨左右
func maxPathSumVerFail(root *TreeNode) int {
	if root == nil {
		return 0
	}
	fmt.Printf("root info: %v\n", root)
	dp124 = make(map[*TreeNode]int)
	switch {
	case root.Left != nil && root.Right != nil:
		dp124[root] = maxNum(
			root.Val+root.Left.Val,
			root.Val+root.Right.Val,
			root.Val+root.Left.Val+root.Right.Val,
		)
	case root.Left != nil && root.Right == nil:
		dp124[root] = root.Left.Val + root.Val
	case root.Left == nil && root.Right != nil:
		dp124[root] = root.Right.Val + root.Val
	default:
		return root.Val
	}
	nodeMaxPathIncludRoot(root.Left, root)
	nodeMaxPathIncludRoot(root.Right, root)
	var max int
	for _, value := range dp124 {
		if value > max {
			max = value
		}
	}
	return max
}

func nodeMaxPathIncludRoot(node *TreeNode, root *TreeNode) {
	if node == nil {
		return
	}
	fmt.Printf("node info: %v\n", node)
	switch {
	case node.Left != nil && node.Right != nil:
		dp124[node] = maxNum(
			dp124[root],
			dp124[root]+node.Left.Val,
			dp124[root]+node.Right.Val,
			node.Val+node.Right.Val+node.Left.Val,
		)
	case node.Left != nil && node.Right == nil:
		dp124[node] = maxNum(
			dp124[root],
			dp124[root]+node.Left.Val,
			node.Val+node.Left.Val,
		)
	case node.Left == nil && node.Right != nil:
		dp124[node] = maxNum(
			dp124[root],
			dp124[root]+node.Right.Val,
			node.Val+node.Right.Val,
		)
	default:
		dp124[node] = dp124[root]
	}
	nodeMaxPathIncludRoot(node.Left, node)
	nodeMaxPathIncludRoot(node.Right, node)
}
