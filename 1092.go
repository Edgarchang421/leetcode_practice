package main

/*
Example 1:

Input: str1 = "abac", str2 = "cab"
Output: "cabac"
Explanation:
str1 = "abac" is a subsequence of "cabac" because we can delete the first "c".
str2 = "cab" is a subsequence of "cabac" because we can delete the last "ac".
The answer provided is the shortest such string that satisfies these properties.


Example 2:

Input: str1 = "aaaaaaaa", str2 = "aaaaaaaa"
Output: "aaaaaaaa"


Constraints:

1 <= str1.length, str2.length <= 1000
str1 and str2 consist of lowercase English letters.

	fmt.Printf("ans cabac, my ans: %v\n", shortestCommonSupersequence("abac", "cab"))
*/

// Runtime: 4 ms, faster than 92.86% of Go online submissions for Shortest Common Supersequence .
// Memory Usage: 12.1 MB, less than 71.43% of Go online submissions for Shortest Common Supersequence .
func shortestCommonSupersequence(str1 string, str2 string) string {
	// 先找到兩個字串的最長共同子字串
	// 二維陣列的DP，因為會減一求解，所以dp[i][0]、dp[0][j]都是0
	// 從最小的字串，逐步推算
	LCSDP := make([][]int, len(str1)+1)
	for i := range LCSDP {
		LCSDP[i] = make([]int, len(str2)+1)
	}
	for i := 1; i <= len(str1); i++ {
		for j := 1; j <= len(str2); j++ {
			// 比對for loop到的字是否相同
			if str1[i-1] == str2[j-1] {
				// 如果相同，就是兩個字串各自減少一個字的LCS長度加一
				LCSDP[i][j] = LCSDP[i-1][j-1] + 1
			} else {
				// 不同，找到目前為止LCS的長度
				LCSDP[i][j] = max(LCSDP[i-1][j], LCSDP[i][j-1])
			}
		}
	}

	// make SCS
	// SCS可以視為LCS加上兩組字串各自剩餘的字
	// 所以使用LCS的DP來推算，找到LCS以及未使用的字
	var SCS string
	i := len(str1)
	j := len(str2)
	for i > 0 && j > 0 {
		// 會從DP[i][j]來開始推算
		// 從這邊開始，根據DP儲存的資料，判斷兩個字串目前最後一個字是否相同
		if str1[i-1] == str2[j-1] {
			// 相同，則取這個字加入SCS
			SCS += string(str1[i-1])
			// 再繼續往前搜尋，因為相同所以兩個字串都需要往前
			i--
			j--

			// 在兩個字串最後的字兩兩不相同的情況下
			// 在前面LCS會取max(LCSDP[i-1][j], LCSDP[i][j-1])
			// 因為前面定義SCS可以視為LCS加上兩組字串各自剩餘的字，所以要找DP儲存較大數字(代表LCS較長)的那一格資料
			// 例 str1: abcde str2: ace
			/*
				dp:
					a  c  e
				a  1  1  1
				b  1  1  1
				c  1  2  2
				d  1  2  2
				e  1  2  3
				取完e之後，移動至(abcd,ac)，需判斷取出d or c之後，LCS的長度會較長，取出d之後的(abc,ac)的LCS較長，故取出d放到SCS中
			*/
		} else if LCSDP[i-1][j] > LCSDP[i][j-1] {
			SCS += string(str1[i-1])
			i--
		} else {
			SCS += string(str2[j-1])
			j--
		}
	}

	// 其中一個字串已經取出完畢，取出另外一個字串剩餘的字
	for i > 0 {
		SCS += string(str1[i-1])
		i--
	}
	for j > 0 {
		SCS += string(str2[j-1])
		j--
	}

	// 因為是從判斷最後一個字開始，所以要反轉
	return reverseString(SCS)
}
