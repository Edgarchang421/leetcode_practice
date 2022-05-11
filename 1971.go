package main

/*
Input: n = 3, edges = [[0,1],[1,2],[2,0]], source = 0, destination = 2
Output: true
Explanation: There are two paths from vertex 0 to vertex 2:
- 0 → 1 → 2
- 0 → 2

Input: n = 6, edges = [[0,1],[0,2],[3,5],[5,4],[4,3]], source = 0, destination = 5
Output: false
Explanation: There is no path from vertex 0 to vertex 5.

	r := validPath(6, [][]int{{0, 1}, {0, 2}, {3, 5}, {5, 4}, {4, 3}}, 0, 5)
	fmt.Printf("result: %v\n", r)

	r = validPath(3, [][]int{{0, 1}, {1, 2}, {2, 0}}, 0, 2)
	fmt.Printf("result: %v\n", r)
*/

// Runtime: 587 ms, faster than 23.38% of Go online submissions for Find if Path Exists in Graph.
// Memory Usage: 44.4 MB, less than 45.77% of Go online submissions for Find if Path Exists in Graph.
// BFS，DFS則是將stack的pop該為先進先出，優先處理每一個路徑走到終點，stack處理後進先出，廣度優先
func validPath(n int, edges [][]int, source int, destination int) bool {
	// 建立Adjacency Matirx
	// 儲存每一個node有連接的node，使用slice
	ans := make(map[int][]int)
	for i := 1; i <= n; i++ {
		ans[i] = make([]int, 0)
	}
	for i := range edges {
		a := edges[i][0]
		b := edges[i][1]
		// 此題為無向圖，所以路徑連接為雙向
		ans[a] = append(ans[a], b)
		ans[b] = append(ans[b], a)
	}

	// stack儲存接下來要走的點
	stack := []int{source}
	// done標記此點是否走過
	done := map[int]bool{}
	// stack中沒有路可以走就結束迴圈
	for len(stack) != 0 {
		// 取出現在的node，update stack slice
		curr := stack[len(stack)-1]
		stack = stack[0 : len(stack)-1]

		// 如果是終點，就return
		if curr == destination {
			return true
		}

		// 這個node沒有走過
		if !done[curr] {
			// update為true
			done[curr] = true
			// 加入Adjacency Matirx中儲存這個node有連接的node
			stack = append(stack, ans[curr]...)
		}
	}

	return false
}
