package main

import "fmt"

/*
Input: n = 4, connections = [[0,1],[0,2],[1,2]]
Output: 1
Explanation: Remove cable between computer 1 and 2 and place between computers 1 and 3.

Input: n = 6, connections = [[0,1],[0,2],[0,3],[1,2],[1,3]]
Output: 2

Input: n = 6, connections = [[0,1],[0,2],[0,3],[1,2]]
Output: -1
Explanation: There are not enough cables.

100
[[17,51],[33,83],[53,62],[25,34],[35,90],[29,41],[14,53],[40,84],[41,64],[13,68],[44,85],[57,58],[50,74],[20,69],[15,62],[25,88],[4,56],[37,39],
[30,62],[69,79],[33,85],[24,83],[35,77],[2,73],[6,28],[46,98],[11,82],[29,72],[67,71],[12,49],[42,56],[56,65],[40,70],[24,64],[29,51],[20,27],
[45,88],[58,92],[60,99],[33,46],[19,69],[33,89],[54,82],[16,50],[35,73],[19,45],[19,72],[1,79],[27,80],[22,41],[52,61],[50,85],[27,45],[4,84],
[11,96],[0,99],[29,94],[9,19],[66,99],[20,39],[16,85],[12,27],[16,67],[61,80],[67,83],[16,17],[24,27],[16,25],[41,79],[51,95],[46,47],[27,51],
[31,44],[0,69],[61,63],[33,95],[17,88],[70,87],[40,42],[21,42],[67,77],[33,65],[3,25],[39,83],[34,40],[15,79],[30,90],[58,95],[45,56],[37,48],
[24,91],[31,93],[83,90],[17,86],[61,65],[15,48],[34,56],[12,26],[39,98],[1,48],[21,76],[72,96],[30,69],[46,80],[6,29],[29,81],[22,77],[85,90],
[79,83],[6,26],[33,57],[3,65],[63,84],[77,94],[26,90],[64,77],[0,3],[27,97],[66,89],[18,77],[27,43]]
13

	// r := makeConnected(4, [][]int{{0, 1}, {0, 2}, {1, 2}})
	// fmt.Printf("ans should be 1: %v\n", r)

	r := makeConnected(6, [][]int{{0, 1}, {0, 2}, {0, 3}, {1, 2}, {1, 3}})
	fmt.Printf("ans should be 2: %v\n", r)

	// r = makeConnected(6, [][]int{{0, 1}, {0, 2}, {0, 3}, {1, 2}})
	// fmt.Printf("ans should be -1: %v\n", r)

*/

// Runtime: 134 ms, faster than 16.67% of Go online submissions for Number of Operations to Make Network Connected.
// Memory Usage: 15.7 MB, less than 23.81% of Go online submissions for Number of Operations to Make Network Connected.
func makeConnected(n int, connections [][]int) int {
	if len(connections) < n-1 {
		return -1
	}
	if n == 1 {
		return 0
	}
	graph := make([][]int, n)
	for i := range connections {
		graph[connections[i][0]] = append(graph[connections[i][0]], connections[i][1])
		graph[connections[i][1]] = append(graph[connections[i][1]], connections[i][0])
	}
	fmt.Printf("graph: %v\n", graph)

	bfs := make([]int, 0)
	// 分為兩大群就找不到
	for i := range graph {
		if len(graph[i]) == 0 {
			bfs = append(bfs, i)
		}
	}
	queue := []int{connections[0][0]}
	undone := map[int]struct{}{}
	counter := 0
	for i := 0; i < n; i++ {
		undone[i] = struct{}{}
	}
	for len(undone) != 0 {
		if len(queue) == 0 {
			counter++
			for i := range undone {
				queue = append(queue, i)
				break
			}
		}
		pc := queue[0]
		queue = queue[1:]

		if _, ok := undone[pc]; ok {
			queue = append(queue, graph[pc]...)
			delete(undone, pc)
		}
	}
	fmt.Printf("bfs: %v\n", bfs)
	return counter
}

/*
def makeConnected(self, n, connections):
        if len(connections) < n - 1: return -1
        G = [set() for i in xrange(n)]
        for i, j in connections:
            G[i].add(j)
            G[j].add(i)
        seen = [0] * n

        def dfs(i):
            if seen[i]: return 0
            seen[i] = 1
            for j in G[i]: dfs(j)
            return 1

        return sum(dfs(i) for i in xrange(n)) - 1
*/

// Runtime: 84 ms, faster than 72.62% of Go online submissions for Number of Operations to Make Network Connected.
// Memory Usage: 13.7 MB, less than 36.90% of Go online submissions for Number of Operations to Make Network Connected.
func makeConnectedV2(n int, connections [][]int) int {
	if n-1 > len(connections) {
		return -1
	}

	visited := make(map[int]bool)
	graph := make([][]int, n)
	for i := range connections {
		graph[connections[i][0]] = append(graph[connections[i][0]], connections[i][1])
		graph[connections[i][1]] = append(graph[connections[i][1]], connections[i][0])
	}

	var connectedNode int
	connectedNode = 1
	for k := range graph {
		if _, ok := visited[k]; !ok {
			dfs1319(k, graph, visited)
		} else {
			connectedNode++
		}
	}

	return n - connectedNode
}
func dfs1319(node int, graph [][]int, visited map[int]bool) {
	if _, ok := visited[node]; !ok {

		visited[node] = true
		for _, n := range graph[node] {
			dfs1319(n, graph, visited)
		}
	}
	return
}
