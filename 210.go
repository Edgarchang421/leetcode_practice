package main

/*
Example 1:

Input: numCourses = 2, prerequisites = [[1,0]]
Output: [0,1]
Explanation: There are a total of 2 courses to take. To take course 1 you should have finished course 0. So the correct course order is [0,1].
Example 2:

Input: numCourses = 4, prerequisites = [[1,0],[2,0],[3,1],[3,2]]
Output: [0,2,1,3]
Explanation: There are a total of 4 courses to take. To take course 3 you should have finished both courses 1 and 2. Both courses 1 and 2 should be taken after you finished course 0.
So one correct course order is [0,1,2,3]. Another correct ordering is [0,2,1,3].
Example 3:

Input: numCourses = 1, prerequisites = []
Output: [0]

fmt.Printf("ans: %v\n", findOrder(4, [][]int{{1, 0}, {2, 0}, {3, 1}, {3, 2}}))
*/
func findOrder(numCourses int, prerequisites [][]int) []int {
	// index代表課程，value紀錄所有需要先修此課程的課
	graph := make([][]int, numCourses)
	// index代表課程，value紀錄有多少課需要先修
	courseNeedPreRequisiteNums := make([]int, numCourses)
	// prerequisites = [[1,0]] 0為1的先修課程
	for i := range prerequisites {
		// b為a的先修，畫作地圖: b-->a
		a := prerequisites[i][0]
		b := prerequisites[i][1]
		// graph記錄此課程為那些課程的先修課，b為a先修課，所以新增a到b
		graph[b] = append(graph[b], a)
		// a需要先修課程數量加一
		courseNeedPreRequisiteNums[a]++
	}
	// fmt.Printf("graph: %v, courseNeedPreRequisiteNums: %v\n", graph, courseNeedPreRequisiteNums)

	// 都沒有先修課程為source出發點
	coursesWithoutPreRequisite := make([]int, 0)
	for i := range courseNeedPreRequisiteNums {
		if courseNeedPreRequisiteNums[i] == 0 {
			coursesWithoutPreRequisite = append(coursesWithoutPreRequisite, i)
		}
	}

	result := make([]int, 0)
	// 沒有先修課程的優先放入到result
	for i := 0; i < len(coursesWithoutPreRequisite); i++ {
		// fmt.Printf("sources: %v\n", sources)
		courseWithoutPreRequisite := coursesWithoutPreRequisite[i]
		// graph紀錄此課程為哪些課的先修課
		// 所以needPreRequisiteCourses會是需要此堂先修課的課程們
		needPreRequisiteCourses := graph[courseWithoutPreRequisite]
		// fmt.Printf("needPreRequisiteCourses: %v\n", needPreRequisiteCourses)
		for j := range needPreRequisiteCourses {
			needPreRequisiteCourse := needPreRequisiteCourses[j]
			// needPreRequisiteCourse所需的先修課程數減一
			courseNeedPreRequisiteNums[needPreRequisiteCourse]--
			// 如果為零，也就是沒有先修課程了，就放入sources
			if courseNeedPreRequisiteNums[needPreRequisiteCourse] == 0 {
				coursesWithoutPreRequisite = append(coursesWithoutPreRequisite, needPreRequisiteCourse)
			}
		}
		// 新增到result
		result = append(result, courseWithoutPreRequisite)
	}

	// 如果有迴圈，就會無法加入到source，就無法新增到result，判斷result長度是否等於全課程數量
	if len(result) != numCourses {
		return nil
	}

	return result
}

// Runtime: 18 ms, faster than 47.34% of Go online submissions for Course Schedule II.
// Memory Usage: 6.2 MB, less than 87.68% of Go online submissions for Course Schedule II.
func findOrderClean(numCourses int, prerequisites [][]int) []int {
	graph := make([][]int, numCourses)
	courseNeedPreRequisiteNums := make([]int, numCourses)
	for i := range prerequisites {
		a := prerequisites[i][0]
		b := prerequisites[i][1]
		graph[b] = append(graph[b], a)
		courseNeedPreRequisiteNums[a]++
	}
	sources := make([]int, 0)
	for i := range courseNeedPreRequisiteNums {
		if courseNeedPreRequisiteNums[i] == 0 {
			sources = append(sources, i)
		}
	}
	result := make([]int, 0)
	for i := 0; i < len(sources); i++ {
		courseWithoutPreRequisite := sources[i]
		needPreRequisiteCourses := graph[courseWithoutPreRequisite]
		for j := range needPreRequisiteCourses {
			needPreRequisiteCourse := needPreRequisiteCourses[j]
			courseNeedPreRequisiteNums[needPreRequisiteCourse]--
			if courseNeedPreRequisiteNums[needPreRequisiteCourse] == 0 {
				sources = append(sources, needPreRequisiteCourse)
			}
		}
		result = append(result, courseWithoutPreRequisite)
	}
	if len(result) != numCourses {
		return nil
	}
	return result
}

// BFS、DFS找到已走過的node就是有cycle，回傳nil
// 檢查所有沒有先修課程的課
