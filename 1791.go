package main

/*
Input: edges = [[1,2],[2,3],[4,2]]
Output: 2
Explanation: As shown in the figure above, node 2 is connected to every other node, so 2 is the center.
Example 2:

Input: edges = [[1,2],[5,1],[1,3],[1,4]]
Output: 1

	r := findCenter([][]int{{1, 2}, {5, 1}, {1, 3}, {1, 4}})
	fmt.Printf("result: %v\n", r)
*/

// Runtime: 131 ms, faster than 78.90% of Go online submissions for Find Center of Star Graph.
// Memory Usage: 15.2 MB, less than 94.50% of Go online submissions for Find Center of Star Graph.
func findCenter(edges [][]int) int {
	l := len(edges)
	graph := make([]int, l+2)
	for i := range edges {
		graph[edges[i][0]]++
		graph[edges[i][1]]++
	}
	for i := 1; i <= l+1; i++ {
		if graph[i] == l {
			return i
		}
	}
	return 1
}

// Runtime: 139 ms, faster than 75.23% of Go online submissions for Find Center of Star Graph.
// Memory Usage: 15.5 MB, less than 86.24% of Go online submissions for Find Center of Star Graph.
func findCenterV2(edges [][]int) int {
	if edges[0][0] == edges[1][0] || edges[0][0] == edges[1][1] {
		return edges[0][0]
	}
	return edges[0][1]
}
