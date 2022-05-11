package main

import (
	"fmt"
)

/*
Example 1:

Input: n = 2, trust = [[1,2]]
Output: 2
Example 2:

Input: n = 3, trust = [[1,3],[2,3]]
Output: 3
Example 3:

Input: n = 3, trust = [[1,3],[2,3],[3,1]]
Output: -1

4 [[1,3],[1,4],[2,3],[2,4],[4,3]]
3

	r := findJudge(3, [][]int{{1, 3}, {2, 3}, {3, 1}})
	fmt.Printf("result: %v\n", r)

	r = findJudge(3, [][]int{{1, 3}, {2, 3}})
	fmt.Printf("result: %v\n", r)
*/

// Runtime: 168 ms, faster than 19.44% of Go online submissions for Find the Town Judge.
// Memory Usage: 8.1 MB, less than 24.44% of Go online submissions for Find the Town Judge.
func findJudge(n int, trust [][]int) int {
	// people : trust[]
	people := make(map[int][]int)
	for i := 1; i <= n; i++ {
		people[i] = make([]int, 0)
	}
	for i := range trust {
		a := trust[i][0]
		b := trust[i][1]
		// 此人信任的人的list
		people[a] = append(people[a], b)
	}
	fmt.Printf("people: %v\n", people)
	var judge int
	for k, v := range people {
		// 法官不信任任何人
		if len(v) == 0 {
			judge = k
			break
		}
	}
	// 每個人都有信任人，所以沒有法官
	if judge == 0 {
		return -1
	}

	// 所有人都要信任法官
	for i := 1; i <= n; i++ {
		// 不用檢查法官的信任人
		if i == judge {
			continue
		}
		// 如果村民的信任人有judge，為true，否則false回傳-1
		tag := false
		for _, person := range people[i] {
			if person == judge {
				tag = true
				break
			}
		}
		if tag != true {
			return -1
		}
	}
	return judge
}

// 信任人數為多個時會錯誤
func findJudgeVF1(n int, trust [][]int) int {
	m := make(map[int]int)
	for i := 1; i <= n; i++ {
		m[i] = 0
	}
	for i := range trust {
		a := trust[i][0]
		b := trust[i][1]

		m[a] = b
	}
	// 檢查信任人數
	var j int
	for k, v := range m {
		if v == 0 {
			j = k
		}
	}
	if j == 0 {
		return -1
	}
	// 檢查是否所有人都信任
	for i := range m {
		if i == j {
			continue
		}

		if m[i] != j {
			return -1
		}
	}
	return j
}

// Runtime: 195 ms, faster than 9.44% of Go online submissions for Find the Town Judge.
// Memory Usage: 8.2 MB, less than 17.78% of Go online submissions for Find the Town Judge.
func findJudgeV2(n int, trust [][]int) int {
	// 取所有trust[i]的第二項信任人
	// judgePossbile: count trust person
	judgePossbile := map[int]int{}
	for i := range trust {
		judgeNum := trust[i][1]
		if _, ok := judgePossbile[judgeNum]; !ok {
			judgePossbile[judgeNum] = 1
		} else {
			judgePossbile[judgeNum]++
		}
		// 相信某個人的人數到達標準，所有人減掉法官一人
		if judgePossbile[judgeNum] == n-1 {
			// 檢查法官有沒有信任別人
			for j := range trust {
				if trust[j][0] == judgeNum {
					return -1
				}
			}
			return judgeNum
		}
	}
	if len(trust) == 0 && n == 1 {
		return 1
	}
	return -1
}

// Runtime: 129 ms, faster than 46.67% of Go online submissions for Find the Town Judge.
// Memory Usage: 8 MB, less than 39.44% of Go online submissions for Find the Town Judge.
// 把每一個pair視為graph的一條路徑，法官最終會等於 n-1 ，不會有進來
func findJudgeV3(n int, trust [][]int) int {
	judgePossbile := make([]int, n+1)
	for i := range trust {
		judgePossbile[trust[i][0]]--
		judgePossbile[trust[i][1]]++
	}
	for i := 1; i <= n; i++ {
		if judgePossbile[i] == n-1 {
			return i
		}
	}
	return -1
}
