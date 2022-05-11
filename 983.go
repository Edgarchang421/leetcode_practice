package main

import (
	"fmt"
)

/*
costs中為1天、7天、30天的票價
[1,4,6,7,8,20]
[7,2,15]
6

[4,5,9,11,14,16,17,19,21,22,24]
[1,4,18]
10

Input: days = [1,4,6,7,8,20], costs = [2,7,15]
Output: 11

Input: days = [1,2,3,4,5,6,7,8,9,10,30,31], costs = [2,7,15]
Output: 17

Constraints:

1 <= days.length <= 365
1 <= days[i] <= 365
days is in strictly increasing order.
costs.length == 3
1 <= costs[i] <= 1000

	// r := mincostTickets([]int{1, 4, 6, 7, 8, 20}, []int{2, 7, 15})
	// fmt.Printf("result: %v\n", r)

	// r = mincostTickets([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 30, 31}, []int{2, 7, 15})
	// fmt.Printf("result: %v\n", r)

	// r := mincostTickets([]int{1, 4, 6, 7, 8, 20}, []int{7, 2, 15})
	// fmt.Printf("result: %v\n", r)

	r := mincostTicketsVer2([]int{4, 5, 9, 11, 14, 16, 17, 19, 21, 22, 24}, []int{1, 4, 18})
	fmt.Printf("result: %v\n", r)
*/

// Runtime: 0 ms, faster than 100.00% of Go online submissions for Minimum Cost For Tickets.
// Memory Usage: 2.3 MB, less than 77.97% of Go online submissions for Minimum Cost For Tickets.
func mincostTicketsVer2(days []int, costs []int) int {
	l := len(days)
	dp := make([]int, l+1)

	for i := 1; i <= l; i++ {
		dp[i] = dp[i-1] + costs[0]
		weekPassIndex := indexCheck(days, days[i-1]-7)
		dp[i] = min(dp[weekPassIndex+1]+costs[1], dp[i])
		monthPassIndex := indexCheck(days, days[i-1]-30)
		dp[i] = min(dp[monthPassIndex+1]+costs[2], dp[i])
	}

	return dp[l]
}

func mincostTickets(days []int, costs []int) int {
	costMap := map[int]int{}
	costMap[1] = costs[0]
	costMap[7] = costs[1]
	costMap[30] = costs[2]

	l := len(days)

	// 記錄所有day的最低花費
	dp := make([]int, len(days)+1)
	// for i := range dp {
	// 	dp[i] = 65535
	// }

	// dp[0] = 0

	// 跑完每一個有計畫旅遊的day
	// dp[0] 是沒有出遊日的花費，dp[i-1]對應到day[i]的累積最低花費
	// days[i] to days[i-1], dp[i] to dp[i-1]
	for i := 1; i <= l; i++ {
		fmt.Printf("\n\nday: %v\n\n", days[i-1])
		for duration, cost := range costMap {
			fmt.Printf("\nduration %v, cost %v\n", duration, cost)
			fmt.Printf("dp: %v\n", dp)
			// 單日票，直接加總，不用比較
			if duration == 1 {
				dp[i] = dp[i-1] + cost
				// fmt.Printf("duration: %v, cost: %v, day: %v, dp[%v]: %v\n", duration, cost, days[i], i, dp[i])
				fmt.Printf("duration: %v, cost: %v, dp[%v]: %v\n", duration, cost, i, dp[i])
				continue
			}

			// 7天或30天週期，開始要比較
			// dp(20) = min ( dp(19)+costs[0] , dp(13)+costs[1] , dp(-10)+cost[2] )
			// 要找出正確的index值，不能使用沒有旅遊的day
			index := indexCheck(days, days[i-1]-duration)
			// if index == -1 {
			// 	fmt.Printf("index check: %v can't found\n", days[i-1]-duration)
			// 	continue
			// }
			// index原本是days[i]的，回到這邊要+1，才會對應到dp中的資料
			dp[i] = min(dp[index+1]+cost, dp[i])
			fmt.Printf("day: %v, dp[%v]+cost %v: %v, dp[%v]: %v\n", days[i-1], index+1, cost, dp[index+1]+cost, i, dp[i])
		}
	}

	return dp[l]
}

func min983(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func indexCheck(days []int, targetDay int) (index int) {
	// 1 4 6 7 8 20
	// 20-7 = 13
	// 應該回傳8
	// 看8是在days中第幾個index
	// 回傳
	for i := len(days) - 1; i >= 0; i-- {
		if days[i] <= targetDay {
			fmt.Printf("target day: %v, final index: %v\n", targetDay, i)
			return i
		}
	}
	return -1
}

// example
func getIndex(days []int, index, num int) int {
	// 這只是先傳一個day index，不用從最大的開始比較
	i := index - 1
	for ; i >= 0; i-- {
		if days[i] <= num {
			// fmt.Printf("i %v, days[%v] %v, num %v\n", i, i, days[i], num)
			return i
		}
	}
	return -1
}

func mincostTicketsExample(days []int, costs []int) int {
	n := len(days)
	dp := make([]int, n+1)
	for i := 1; i <= n; i++ {
		dp[i] = costs[0] + dp[i-1]
		// i 是day和dp的index，因為有dp[0]，所以從1開始，所以放到slice的index要減一
		weekIndex := getIndex(days, i-1, days[i-1]-7) + 1
		monthIndex := getIndex(days, i-1, days[i-1]-30) + 1
		// fmt.Printf("week %v month %v\n", weekIndex, monthIndex)
		// fmt.Printf("dp[i] %v, costs[1]+dp[weekIndex] %v\n", dp[i], costs[1]+dp[weekIndex])
		dp[i] = min(dp[i], costs[1]+dp[weekIndex])
		dp[i] = min(dp[i], costs[2]+dp[monthIndex])
		fmt.Printf("dp %v\n", dp)
	}
	return dp[n]
}
