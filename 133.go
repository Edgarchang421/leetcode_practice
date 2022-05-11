package main

/*
Input: adjList = [[2,4],[1,3],[2,4],[1,3]]
Output: [[2,4],[1,3],[2,4],[1,3]]
Explanation: There are 4 nodes in the graph.
1st node (val = 1)'s neighbors are 2nd node (val = 2) and 4th node (val = 4).
2nd node (val = 2)'s neighbors are 1st node (val = 1) and 3rd node (val = 3).
3rd node (val = 3)'s neighbors are 2nd node (val = 2) and 4th node (val = 4).
4th node (val = 4)'s neighbors are 1st node (val = 1) and 3rd node (val = 3).

	n2 := Node{Val: 2}
	n4 := Node{Val: 4}
	n3 := Node{Val: 3}
	n1 := Node{Val: 1}
	n1.Neighbors = []*Node{&n2, &n4}
	n2.Neighbors = []*Node{&n1, &n3}
	n3.Neighbors = []*Node{&n2, &n4}
	n4.Neighbors = []*Node{&n1, &n3}

	r := cloneGraph(&n1)
	fmt.Print(r)
*/

// Runtime: 5 ms, faster than 29.20% of Go online submissions for Clone Graph.
// Memory Usage: 5.8 MB, less than 17.24% of Go online submissions for Clone Graph.
func cloneGraph(node *Node) *Node {
	if node == nil {
		return nil
	}
	m := make(map[int][]int)
	queue := []*Node{node}
	for len(queue) != 0 {
		n := queue[0]
		queue = queue[1:]
		if _, ok := m[n.Val]; !ok {
			m[n.Val] = make([]int, 0)
			for i := range n.Neighbors {
				m[n.Val] = append(m[n.Val], n.Neighbors[i].Val)
				queue = append(queue, n.Neighbors...)
			}
		} else {
			continue
		}
	}
	ansNodeMap := make(map[int]*Node)
	for nodeVal := range m {
		ansNodeMap[nodeVal] = &Node{Val: nodeVal}
	}
	for nodeVal, nodeNeighbors := range m {
		for _, neighbor := range nodeNeighbors {
			ansNodeMap[nodeVal].Neighbors = append(ansNodeMap[nodeVal].Neighbors, ansNodeMap[neighbor])
		}
	}
	return ansNodeMap[node.Val]
}

// Runtime: 11 ms, faster than 11.72% of Go online submissions for Clone Graph.
// Memory Usage: 5.7 MB, less than 17.24% of Go online submissions for Clone Graph.
func cloneGraphV2(node *Node) *Node {
	if node == nil {
		return nil
	}
	// 儲存連結
	m := make(map[int][]int)
	// return使用
	ansNodeMap := make(map[int]*Node)
	queue := []*Node{node}
	for len(queue) != 0 {
		n := queue[0]
		queue = queue[1:]
		if _, ok := m[n.Val]; !ok {
			m[n.Val] = make([]int, 0)
			// ansNodeMap[n.Val] = &Node{Val: n.Val}
			if _, ok := ansNodeMap[n.Val]; !ok {
				ansNodeMap[n.Val] = &Node{Val: n.Val}
			}
			for i := range n.Neighbors {
				m[n.Val] = append(m[n.Val], n.Neighbors[i].Val)
				queue = append(queue, n.Neighbors...)
			}
		} else {
			continue
		}
	}
	for nodeVal, nodeNeighbors := range m {
		for _, neighbor := range nodeNeighbors {
			ansNodeMap[nodeVal].Neighbors = append(ansNodeMap[nodeVal].Neighbors, ansNodeMap[neighbor])
		}
	}
	return ansNodeMap[node.Val]
}

// solution example
// Runtime: 0 ms, faster than 100.00% of Go online submissions for Clone Graph.
// Memory Usage: 2.9 MB, less than 17.24% of Go online submissions for Clone Graph.
func dfs(node *Node, visitedNode map[*Node]*Node) *Node {
	if node == nil {
		return nil
	}
	// 當所有的點都儲存過後，就不會再繼續往下遞迴
	if n, ok := visitedNode[node]; ok {
		return n
	}
	// 用舊的node的pointer作為key，儲存新的node pointer作為value
	visitedNode[node] = &Node{
		Val:       node.Val,
		Neighbors: []*Node{},
	}
	for _, neighbor := range node.Neighbors {
		visitedNode[node].Neighbors = append(visitedNode[node].Neighbors, dfs(neighbor, visitedNode))
	}
	// 回傳input node對應的新建立的graph的node
	return visitedNode[node]
}

func cloneGraphExample(node *Node) *Node {
	// map用來記錄input node: new node
	visited := make(map[*Node]*Node)
	return dfs(node, visited)
}
