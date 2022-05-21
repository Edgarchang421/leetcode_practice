package main

import "fmt"

/*
Example 1:
Input: n = 4, connections = [[0,1],[1,2],[2,0],[1,3]]
Output: [[1,3]]
Explanation: [[3,1]] is also accepted.

Example 2:
Input: n = 2, connections = [[0,1]]
Output: [[0,1]]

func main() {
	// fmt.Printf("ans: %v\n", StronglyConnectedComponents(10))
	fmt.Printf("ans: %v\n", criticalConnections(4, [][]int{{0, 1}, {1, 2}, {2, 0}, {1, 3}}))
}
*/

// Runtime: 616 ms, faster than 40.00% of Go online submissions for Critical Connections in a Network.
// Memory Usage: 33.8 MB, less than 100.00% of Go online submissions for Critical Connections in a Network.
func criticalConnections(n int, connections [][]int) [][]int {
	graph := make([][]int, n)
	for i := range connections {
		graph[connections[i][0]] = append(graph[connections[i][0]], connections[i][1])
		graph[connections[i][1]] = append(graph[connections[i][1]], connections[i][0])
	}
	// for i := range graph {
	// 	fmt.Printf("graph %v: %v\n", i, graph[i])
	// }
	return dfs1192(graph, make([]*int, n), 0, 0, 0)
}

// 深度優先
// Tarjan's algorithm，比較透過dfs搜尋得出的深度，判斷是否存在critical Connection
func dfs1192(graph [][]int, visited []*int, prev int, current int, depth int) (res [][]int) {
	visited[current] = &depth
	fmt.Printf("current node: %v\n", current)
	for i := range visited {
		if visited[i] != nil {
			fmt.Printf("node %v depth: %v ", i, *visited[i])
		} else {
			fmt.Printf("node %v depth: %v ", i, visited[i])
		}
	}
	fmt.Println("")

	for i := range graph[current] {
		neighborNode := graph[current][i]
		// 不會往回走
		if neighborNode == prev {
			continue
		}

		if visited[neighborNode] == nil {
			res = append(res, dfs1192(graph, visited, current, neighborNode, depth+1)...)
		}

		fmt.Printf("current node: %v, compare\n", current)
		// 已經透過遞迴的dfs走過，比較深度
		// 如果next Node深度較小，則更新現在此點的depth
		if *visited[current] > *visited[neighborNode] {
			// 更新depth，降低深度
			visited[current] = visited[neighborNode]
		}

		// 如果next Node深度較大
		// 代表沒有其他點有路徑可以通過，path可以判斷為關鍵路徑
		if *visited[neighborNode] > depth {
			res = append(res, []int{current, neighborNode})
		}
	}

	return res
}

// 使用 Tarjan's algorithm 找到的 strongly connected components，無向圖，任意刪除一條path，都不會造成node被孤立
// 所以要找到 非strongly connected components 的部分
// 定義一個node在深度優先搜索中的次序編號為DFN(u)，定義u或u的往後需要搜尋的子圖可以搜尋到最早的次序編號為Low(u)，
// DFN(u) == Low(u)時，以u為source的深度搜尋可以找到的所有點就會是strongly connected components

// BFS、DFS 走完所有點，如果visited < n，缺少的就是關鍵連接，該如何選擇起始點? 可以選擇刪除某條連結後，可以連接最多的node。
// 只有一個連接的可以直接判斷為關鍵連接

// 兩點相連，則為strongly connected
// 一個圖之中的子圖，如果所有點都是strongly connected，則稱為strongly connected components
// strongly connected components 任取任意數量的node都會形成cycle、strongly connected components
// 有向圖之中，兩個點至少要有一條路徑，三個點至少要有三條，四個點至少要有六條，五個點至少要有十條，n個點至少要有dp[n-1]+n-1
func StronglyConnectedComponents(n int) int {
	if n < 2 {
		return 0
	}
	last := 1
	if n == 2 {
		return last
	}
	var result int
	for i := 3; i <= n; i++ {
		result = last + i - 1
		last = result
	}
	return result
}
