package main

/*
Example 1:

Input: n = 1
Output: 5
Explanation: The 5 sorted strings that consist of vowels only are ["a","e","i","o","u"].
Example 2:

Input: n = 2
Output: 15
Explanation: The 15 sorted strings that consist of vowels only are
["aa","ae","ai","ao","au","ee","ei","eo","eu","ii","io","iu","oo","ou","uu"].
Note that "ea" is not a valid string since 'e' comes after 'a' in the alphabet.
Example 3:

Input: n = 33
Output: 66045
*/
func countVowelStrings(n int) int {
	dp := make([]int, n+1)
	dp[1] = 5
	a, e, i, o, u := 1, 1, 1, 1, 1
	for j := 2; j <= n; j++ {
		// 結尾是a的可以有五種組合 a e i o u
		// e 4                     e i o u
		// i 3                       i o u
		// o 2                         o u
		// u 1                           u
		// 新的結尾 a*1 e*2 i*3 o*4 u*5
		// dp[i] = dp[i-1]中a結尾的數量*5 + e*4 + i*3 + o*2 + u*1
		dp[j] = a*5 + e*4 + i*3 + o*2 + u*1
		e, i, o, u = a+e, a+e+i, a+e+i+o, a+e+i+o+u
	}

	return dp[n]
}

// Runtime: 1 ms, faster than 71.05% of Go online submissions for Count Sorted Vowel Strings.
// Memory Usage: 1.9 MB, less than 86.84% of Go online submissions for Count Sorted Vowel Strings.
func refactor(n int) int {
	a, e, i, o, u := 1, 1, 1, 1, 1
	for j := 1; j < n; j++ {
		e, i, o, u = a+e, a+e+i, a+e+i+o, a+e+i+o+u
	}
	return a + e + i + o + u
}
