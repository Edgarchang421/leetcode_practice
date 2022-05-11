package main

/*
	r := uniquePaths(3, 7)
*/

// 橫向 i 格 縱向 j 格，假設橫向 i 是固定的，另外縱向 j 格每多一格，就會多 i -1 個解
// 2 * 1 有一個解，2*2 有兩個解，2*3 有三個解
// 3*1
// dp[i][j] = dp[i][j-1] + i + 1
// Runtime: 0 ms, faster than 100.00% of Go online submissions for Unique Paths.
// Memory Usage: 1.9 MB, less than 72.33% of Go online submissions for Unique Paths.
func uniquePaths(m int, n int) int {
	if m == 1 || n == 1 {
		return 1
	}
	dp := make([][]int, m)
	for i := range dp {
		dp[i] = make([]int, n)
		dp[i][0] = 1
	}
	for i := range dp[0] {
		dp[0][i] = 1
	}
	for i := 1; i <= m-1; i++ {
		for j := 1; j <= n-1; j++ {
			dp[i][j] = dp[i][j-1] + dp[i-1][j]
		}
	}
	return dp[m-1][n-1]
}
