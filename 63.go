package main

/*
	dp := make([][]int, 3)
	for i := range dp {
		dp[i] = make([]int, 3)
	}
	dp[1][1] = 1
	r := uniquePathsWithObstacles(dp)
*/

// Runtime: 3 ms, faster than 56.22% of Go online submissions for Unique Paths II.
// Memory Usage: 2.4 MB, less than 67.74% of Go online submissions for Unique Paths II.
func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	if obstacleGrid[0][0] == 1 {
		return 0
	}
	m := len(obstacleGrid)
	n := len(obstacleGrid[0])

	res := make([][]int, m)
	for i := range res {
		res[i] = make([]int, n)
	}
	res[0][0] = 1
	for i := 0; i <= m-1; i++ {
		for j := 0; j <= n-1; j++ {
			if obstacleGrid[i][j] != 1 {
				if i-1 >= 0 {
					res[i][j] += res[i-1][j]
				}
				if j-1 >= 0 {
					res[i][j] += res[i][j-1]
				}
			}
		}
	}
	return res[m-1][n-1]
}
