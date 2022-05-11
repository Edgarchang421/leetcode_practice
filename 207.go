package main

import "fmt"

/*
Example 1:

Input: numCourses = 2, prerequisites = [[1,0]]
Output: true
Explanation: There are a total of 2 courses to take.
To take course 1 you should have finished course 0. So it is possible.
Example 2:

Input: numCourses = 2, prerequisites = [[1,0],[0,1]]
Output: false
Explanation: There are a total of 2 courses to take.
To take course 1 you should have finished course 0, and to take course 0 you should also have finished course 1. So it is impossible.

[[0,10],[3,18],[5,5],[6,11],[11,14],[13,1],[15,1],[17,4]] 20
false

5
[[1,4],[2,4],[3,1],[3,2]]
true

8
[[1,0],[2,6],[1,7],[6,4],[7,0],[0,5]]
true

	// r := canFinishV4(2, [][]int{{1, 0}})
	// fmt.Printf("result should be true : %v\n", r)

	// r = canFinishV4(2, [][]int{{1, 0}, {0, 1}})
	// fmt.Printf("result should be false: %v\n", r)

	// r = canFinishV4(20, [][]int{{0, 10}, {3, 18}, {5, 5}, {6, 11}, {11, 14}, {13, 1}, {15, 1}, {17, 4}})
	// fmt.Printf("result should be false: %v\n", r)

	// r := canFinishV4(5, [][]int{{1, 4}, {2, 4}, {3, 1}, {3, 2}})
	// fmt.Printf("result should be true : %v\n", r)

	r := canFinishAns(8, [][]int{{1, 0}, {2, 6}, {1, 7}, {6, 4}, {7, 0}, {0, 5}})
	fmt.Printf("result should be true : %v\n", r)

	r = canFinishAns(3, [][]int{{1, 0}, {1, 2}, {0, 1}})
	fmt.Printf("result should be false: %v\n", r)
*/

// 主要紀錄每一個node往前一步node需求，確認所有node最終都可以從一個沒有前一步node的地方開始
func canFinishAns(numCourses int, prerequisites [][]int) bool {
	// graph: 紀錄此課程為哪些課程的先修課，index紀錄先修課程，value紀錄有那些課程需要先修此門課程
	graph := make([][]int, numCourses)
	// degree: 紀錄每門課程的先修課需求，index紀錄課程num，value紀錄每一門課程有多少先修課程需求
	degree := make([]int, numCourses)

	for _, prerequisite := range prerequisites {
		graph[prerequisite[1]] = append(graph[prerequisite[1]], prerequisite[0])
		// 第prerequisite[0]個課程的先修課數目要 ++
		degree[prerequisite[0]] += 1
	}

	// 紀律沒有先修課程的課
	bfs := make([]int, 0)
	for course, d := range degree {
		// 沒有先修課需求的課程，新增到bfs
		if d == 0 {
			bfs = append(bfs, course)
		}
	}
	fmt.Printf("graph: %v ,degree: %v, bfs: %v\n", graph, degree, bfs)
	// for loop沒有先修的課程的次數
	for i := 0; i < len(bfs); i++ {
		// 沒有先修課程的value
		course := bfs[i]
		fmt.Printf("bfs[%v] course: %v\n", i, course)
		// for loop所有先修課程
		for _, j := range graph[course] {
			fmt.Printf("graph[%v]:%v ,j: %v\n", course, graph[course], j)
			// 此course沒有先修課程了，所以可以直接將需要先修此課程的課程的先修課數目減一
			degree[j] -= 1
			// 為0的時候，代表此課程沒有先修課了，所以加入至bfs，最外層for loop長度更改。
			if degree[j] == 0 {
				bfs = append(bfs, j)
			}
		}
	}

	if len(bfs) == numCourses {
		return true
	}

	return false
}

func canFinish(numCourses int, prerequisites [][]int) bool {
	if len(prerequisites) == 0 {
		return true
	}
	// 記錄每一門課程所需先修
	m := make([][]int, numCourses)
	for i := range m {
		m[i] = make([]int, 0)
	}
	for i := range prerequisites {
		a := prerequisites[i][0]
		b := prerequisites[i][1]
		m[a] = append(m[a], b)
	}
	// 後進先出，後來被找到的先修的課要先被解決
	// done := make(map[int]bool)
	stack := []int{prerequisites[0][0]}
	// for len(stack) != 0 {
	for {
		lesson := stack[len(stack)-1]
		// if lesson == stack[0] {
		// 	return false
		// }
		// stack = stack[:len(stack)-1]

		// 把先修課加入stack
		if len(m[lesson]) != 0 {
			for i := range m[lesson] {
				if m[lesson][i] == stack[0] {
					return false
				}
			}
			stack = append(stack, m[lesson]...)
		} else {
			break
		}
	}
	return true
}

func canFinishV2(numCourses int, prerequisites [][]int) bool {
	if len(prerequisites) == 0 {
		return true
	}
	m := make([][]int, numCourses)
	for i := range m {
		m[i] = make([]int, 0)
	}
	for i := range prerequisites {
		a := prerequisites[i][0]
		b := prerequisites[i][1]
		m[a] = append(m[a], b)
	}
	stack := []int{prerequisites[0][0]}
	visited := make(map[int]struct{})
	for {
		lesson := stack[0]
		visited[lesson] = struct{}{}
		stack = stack[1:]
		if len(m[lesson]) != 0 {
			for i := range m[lesson] {
				if _, ok := visited[m[lesson][i]]; ok {
					return false
				}
			}
			stack = append(stack, m[lesson]...)
		} else {
			break
		}
	}
	return true
}

func canFinishV3(numCourses int, prerequisites [][]int) bool {
	if len(prerequisites) == 0 {
		return true
	}
	m := make([][]int, numCourses)
	for i := range m {
		m[i] = make([]int, 0)
	}
	stack := make([]int, len(prerequisites))
	for i := range prerequisites {
		a := prerequisites[i][0]
		b := prerequisites[i][1]
		m[a] = append(m[a], b)
		stack[i] = prerequisites[i][0]
	}

	visited := make(map[int]struct{})
	for {
		lesson := stack[0]
		visited[lesson] = struct{}{}
		stack = stack[1:]
		if len(m[lesson]) != 0 {
			for i := range m[lesson] {
				if _, ok := visited[m[lesson][i]]; ok {
					return false
				}
			}
			stack = append(stack, m[lesson]...)
		} else {
			break
		}
	}
	return true
}

func canFinishV4(numCourses int, prerequisites [][]int) bool {
	if len(prerequisites) == 0 {
		return true
	}
	graph := make([][]int, numCourses)
	for i := range graph {
		graph[i] = make([]int, 0)
	}
	for i := range prerequisites {
		a := prerequisites[i][0]
		b := prerequisites[i][1]
		graph[a] = append(graph[a], b)
	}
	for i := range prerequisites {
		visited := make(map[int]struct{})
		mustBeDone := []int{prerequisites[i][0]}
		for len(mustBeDone) != 0 {
			fmt.Printf("i: %v, visited: %v, mustBeDone: %v\n", i, visited, mustBeDone)
			course := mustBeDone[0]
			mustBeDone = mustBeDone[1:]
			if len(graph[course]) == 0 {
				continue
			}

			if _, ok := visited[course]; ok {
				return false
			} else {
				visited[course] = struct{}{}
			}
			mustBeDone = append(mustBeDone, graph[course]...)
		}
	}
	return true
}
