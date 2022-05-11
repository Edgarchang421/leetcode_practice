package main

// Runtime: 2 ms, faster than 45.74% of Go online submissions for Unique Binary Search Trees.
// Memory Usage: 2 MB, less than 30.49% of Go online submissions for Unique Binary Search Trees.
func numTrees(n int) int {
	dp := make([]int, n+1)
	dp[0] = 1
	for i := 1; i <= n; i++ {
		for j := 0; j < i; j++ {
			dp[i] += dp[j] * dp[i-j-1]
		}
	}
	return dp[n]
}
