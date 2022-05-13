package main

import "fmt"

/*
Example 1:

Input: s = "aab"
Output: [["a","a","b"],["aa","b"]]
Example 2:

Input: s = "a"
Output: [["a"]]

abba e cddc
abba、a bb a: e c d d c、e cddc、e c dd c

*/

// Runtime: 280 ms, faster than 84.09% of Go online submissions for Palindrome Partitioning.
// Memory Usage: 29 MB, less than 13.26% of Go online submissions for Palindrome Partitioning.
func partition(s string) [][]string {
	result := [][]string{}
	recPart(s, []string{}, &result)
	return result
}

func recPart(inputString string, possiblePartition []string, result *[][]string) {
	// 遞迴中止，已經跑完所有string中的字
	if len(inputString) == 0 {
		*result = append(*result, possiblePartition)
		fmt.Printf("len(inputString) == 0, result: %v\n", *result)
		return
	}
	for i := 1; i <= len(inputString); i++ {
		subStr := inputString[:i]
		if isPal(subStr) {
			newPossiblePartition := make([]string, len(possiblePartition)+1)
			copy(newPossiblePartition, possiblePartition)
			// 是回文，新增到possible Partition。
			newPossiblePartition[len(newPossiblePartition)-1] = subStr
			fmt.Printf("full str: %v, pal sub str: %v, repart: %v, new Possible Partition: %v\n", inputString, subStr, inputString[i:], newPossiblePartition)
			// inputString[:i]是回文，再把剩下的字串遞迴，繼續找回文以及生成所有可能字串
			recPart(inputString[i:], newPossiblePartition, result)
		} else {
			fmt.Printf("full str: %v, unpal sub str: %v\n", inputString, subStr)
		}
	}
}

func isPal(s string) bool {
	for i := 0; i < len(s)/2; i++ {
		if s[i] != s[len(s)-1-i] {
			return false
		}
	}
	return true
}

// wrong answer
func partitionerror(s string) [][]string {
	dp := make([]string, len(s))
	result := make([][][]string, len(s))
	for i := 0; i < len(s); i++ {
		// 要判斷前一個是否有回文，i要大於1
		if i-1 >= 0 {
			// 檢查是否為兩個字的回文
			if s[i] == s[i-1] {
				dp[i] = s[i-1 : i+1]
			} else if i-2 >= 0 {
				if s[i] == s[i-2] {
					dp[i] = s[i-2 : i+1]
				}
			} else if len(dp[i-1]) > 0 && i-len(dp[i-1])-1 >= 0 {
				// 檢查前面是否有回文，有的話就找到應該檢查的index，檢查是否形成回文
				index := i - len(dp[i-1]) - 1
				if s[i] == s[index] {
					dp[i] = s[index : i+1]
				}
			}
			// if Palindrome(s[i-1 : i+1]) {
			// 	dp[i] = s[i-1 : i+1]
			// } else if Palindrome(s[i-2 : i+1]) {
			// 	dp[i] = s[i-2 : i+1]
			// }

			// index := i - len(dp[i-1]) - 1
			// if len(dp[i-1]) > 0 && index >= 0 {
			// 	if Palindrome(s[index : i+1]) {
			// 		dp[i] = s[index : i+1]
			// 	}
			// }
		}
		fmt.Printf("dp[%v]: %v\n", i, dp[i])

		singleStr := string(s[i])
		last := make([][]string, 0)
		if i > 0 {
			last = result[i-1]
			fmt.Printf("last result[%v-1]: %v\n", i, last)
		}
		PalindromeStr := dp[i]
		if len(PalindromeStr) > 0 {
			// 有回文
			index := i - len(PalindromeStr)
			if index < 0 {
				// 從第i個字到第0個字都是回文，所以直接使用回文作為[]string{}
				possiblePartition := []string{PalindromeStr}
				result[i] = append(result[i], possiblePartition)
			} else {
				resultReduceIndexSlice := result[index]
				for j := 0; j < len(resultReduceIndexSlice); j++ {
					// 把現在的字串減掉回文的長度，全部加上PalindromeStr
					possiblePartition := append(resultReduceIndexSlice[j], PalindromeStr)
					result[i] = append(result[i], possiblePartition)
					fmt.Printf("last result[%v-1]: %v\n", i, last)
				}
			}

			for j := 0; j < len(last); j++ {
				fmt.Printf("i: %v, now: %v, last: %v, str: %v\n", i, result[i], last, string(s[i]))
				// 最新的一個字做為單一一個字的回文，
				possiblePartition := append(last[j], singleStr)
				result[i] = append(result[i], possiblePartition)
			}
		} else if len(PalindromeStr) == 0 {
			if i-1 >= 0 {
				for j := 0; j < len(last); j++ {
					possiblePartition := append(last[j], singleStr)
					result[i] = append(result[i], possiblePartition)
				}
			} else {
				possiblePartition := []string{singleStr}
				result[i] = append(result[i], possiblePartition)
			}
		}

		fmt.Printf("result[%v]: %v\n", i, result[i])
	}
	return result[len(s)-1]
}

// wrong answer
func partitionerror2(s string) [][]string {
	dp := make([]string, len(s))
	for i := 0; i <= len(s)-1; i++ {
		// 要判斷前一個是否有回文，i要大於1
		if i-1 >= 0 {
			// 檢查是否為兩個字的回文
			if s[i] == s[i-1] {
				dp[i] = s[i-1 : i+1]
				continue
			}
			// 檢查前面是否有回文，有的話就找到應該檢查的index，檢查是否形成回文
			if len(dp[i-1]) > 0 && i-len(dp[i-1])-1 >= 0 {
				index := i - len(dp[i-1]) - 1
				if s[i] == s[index] {
					dp[i] = s[index : i+1]
					// dp[i-1] = ""
				}
			}
		}
	}
	fmt.Printf("dp: %v\n", dp)

	return nil
}

func isPalindrome(s string) bool {
	left := 0
	right := len(s) - 1
	for left < right {
		if s[left] != s[right] {
			return false
		}
		left++
		right--
	}
	return true
}

// wrong answer
func partitionerror3(s string) [][]string {
	dp := make([]string, len(s))
	result := make([][][]string, len(s))
	for i := 0; i < len(s); i++ {
		if i-1 >= 0 {
			if s[i] == s[i-1] {
				dp[i] = s[i-1 : i+1]
			} else if len(dp[i-1]) > 0 && i-len(dp[i-1])-1 >= 0 {
				index := i - len(dp[i-1]) - 1
				if s[i] == s[index] {
					dp[i] = s[index : i+1]
				}
			} else if i-2 >= 0 {
				if s[i] == s[i-2] {
					dp[i] = s[i-2 : i+1]
				}
			}
		}
		singleStr := string(s[i])
		last := make([][]string, 0)
		if i > 0 {
			last = result[i-1]
		}
		PalindromeStr := dp[i]
		if len(PalindromeStr) > 0 {
			index := i - len(PalindromeStr)
			if index < 0 {
				possiblePartition := []string{PalindromeStr}
				result[i] = append(result[i], possiblePartition)
			} else {
				resultReduceIndexSlice := result[index]
				for j := 0; j < len(resultReduceIndexSlice); j++ {
					possiblePartition := append(resultReduceIndexSlice[j], PalindromeStr)
					result[i] = append(result[i], possiblePartition)
				}
			}
			for j := 0; j < len(last); j++ {
				possiblePartition := append(last[j], singleStr)
				result[i] = append(result[i], possiblePartition)
			}
		} else if len(PalindromeStr) == 0 {
			if i-1 >= 0 {
				for j := 0; j < len(last); j++ {
					possiblePartition := append(last[j], singleStr)
					result[i] = append(result[i], possiblePartition)
				}
			} else {
				possiblePartition := []string{singleStr}
				result[i] = append(result[i], possiblePartition)
			}
		}
	}
	return result[len(s)-1]
}
