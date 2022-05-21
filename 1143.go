package main

/*
Example 1:
Input: text1 = "abcde", text2 = "ace"
Output: 3
Explanation: The longest common subsequence is "ace" and its length is 3.

Example 2:
Input: text1 = "abc", text2 = "abc"
Output: 3
Explanation: The longest common subsequence is "abc" and its length is 3.

Example 3:
Input: text1 = "abc", text2 = "def"
Output: 0
Explanation: There is no such common subsequence, so the result is 0.

"psnw"
"vozsh"
1

"ezupkr"
"ubmrapg"
2

"oxcpqrsvwf"
"shmtulqrypy"
2

"bl"
"yby"
1

	fmt.Printf("ans: %v\n", longestCommonSubsequence("abcde", "ace"))
	// fmt.Printf("ans: %v\n", longestCommonSubsequence("psnw", "vozsh"))
	// fmt.Printf("ans: %v\n", longestCommonSubsequence("ezupkr", "ubmrapg"))
	// fmt.Printf("ans: %v\n", longestCommonSubsequence("oxcpqrsvwf", "shmtulqrypy"))
	// fmt.Printf("ans: %v\n", longestCommonSubsequence("bl", "yby"))
*/

// dp[i][j] = max(dp[i][j-1] (+ 1 or not) , dp[i-1][j] (+ 1 or not) )

// if text1[i] == text2[j]
// DP[i][j] = DP[i - 1][j - 1] + 1

// DP[i][j] = max(DP[i - 1][j], DP[i][j - 1]) , otherwise
// Runtime: 12 ms, faster than 43.13% of Go online submissions for Longest Common Subsequence.
// Memory Usage: 11 MB, less than 67.86% of Go online submissions for Longest Common Subsequence.
func longestCommonSubsequence(text1 string, text2 string) int {
	dp := make([][]int, len(text1)+1)
	for i := range dp {
		dp[i] = make([]int, len(text2)+1)
	}
	for i := 1; i <= len(text1); i++ {
		for j := 1; j <= len(text2); j++ {
			if text1[i-1] == text2[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
			} else {
				dp[i][j] = max(dp[i-1][j], dp[i][j-1])
			}
		}
	}
	// fmt.Printf("dp done\n")
	// for i := range dp {
	// 	fmt.Printf("dp: %v\n", dp[i])
	// }
	return dp[len(text1)][len(text2)]
}
