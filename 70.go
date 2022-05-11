package main

/*
Example 1:

Input: n = 2
Output: 2
Explanation: There are two ways to climb to the top.
1. 1 step + 1 step
2. 2 steps
Example 2:

Input: n = 3
Output: 3
Explanation: There are three ways to climb to the top.
1. 1 step + 1 step + 1 step
2. 1 step + 2 steps
3. 2 steps + 1 step

4
1 1 1 1
2 1 1
1 2 1
---
2 2
1 1 2


5
1 1 1 1 1
2 1 1 1
1 2 1 1
1 1 2 1
2 2 1

1 1 1 2
2 1 2
1 2 2

*/

// dp[n] = dp[n-1] + 可以加二的情況(就是現在階梯層數減二的總數)
// Runtime: 0 ms, faster than 100.00% of Go online submissions for Climbing Stairs.
// Memory Usage: 1.9 MB, less than 89.18% of Go online submissions for Climbing Stairs.
func climbStairs(n int) int {
	dp := make([]int, n+1)

	for i := 0; i <= n; i++ {
		switch i {
		case 0:
			continue
		case 1, 2:
			dp[i] = dp[i-1] + 1
		default:
			dp[i] = dp[i-1] + dp[i-2]
		}
	}

	return dp[n]
}
