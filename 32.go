package main

import (
	"fmt"
	"strings"
)

/*
(: 40
): 41
Example 1:

Input: s = "(()"
Output: 2
Explanation: The longest valid parentheses substring is "()".
Example 2:

Input: s = ")()())"
Output: 4
Explanation: The longest valid parentheses substring is "()()".
Example 3:

Input: s = ""
Output: 0

"()(()"
2

	// 被右邊打斷
	// r := longestValidParentheses("())())(())")
	// fmt.Printf("result 4: %v\n", r)

	// r = longestValidParentheses("())")
	// fmt.Printf("result 2: %v\n", r)

	// 被左邊打斷
	// 右括弧不夠的情況下，會被多的左括弧打斷有效括弧

		1 0
		1 1
		2 1
		3 1
		3 2
		()(())(()

	r := ans("(())(((()")
	fmt.Printf("result: %v\n", r)

*/

// Runtime: 0 ms, faster than 100.00% of Go online submissions for Longest Valid Parentheses.
// Memory Usage: 2.8 MB, less than 64.80% of Go online submissions for Longest Valid Parentheses.
func ans(s string) int {
	// strList := strings.Split(s, "")
	dp := make([]int, len(s))
	res := 0
	for i := 0; i <= len(s)-1; i++ {
		// 一定要是 ) 才會形成正確有效括號
		if i > 0 && s[i] == 41 {
			// 前一個剛好是 ( ，直接取dp[i-2]，這邊的最大值
			if s[i-1] == 40 {
				var val int
				if i-2 >= 0 {
					val = dp[i-2]
				}
				dp[i] = val + 2
			} else {
				index := i - dp[i-1] - 1
				if index >= 0 && s[index] == 40 {
					var val int
					if index > 0 {
						val = dp[index-1]
					}
					dp[i] = 2 + dp[i-1] + val
				}
			}
		}

		if dp[i] > res {
			res = dp[i]
		}

		fmt.Printf("dp: %v\n", dp)
	}
	return res
}

// error
// dp[i] = max(dp[0], dp[1] ,... ,dp[i-1])
// 加上以下情況
// 如果是左括弧 則直接忽略
// 如果是右括弧 檢查前面字串 是否有左括弧可以搭配 有的話 最大有效括弧子字串是否有連接 left = right +1
func dp(s string) int {
	strList := strings.Split(s, "")
	dp := make([]int, len(strList))
	leftCounter := 0
	rightCounter := 0
	for i := 1; i <= len(strList)-1; i++ {
		switch strList[i] {
		case "(":
			leftCounter++
			dp[i] = dp[i-1]
			continue
		case ")":
			// rightCounter++

			if leftCounter > rightCounter {
				rightCounter++
				// 檢查是否新的有效括弧，有被包含在最長子字串括弧中，有的話就直接+1，沒有的話比較最長子字串和最後面的括弧子字串長度
				// if leftCounter == rightCounter {
				// 	dp[i]=rightCounter*2
				// } else {

				// }
				// dp[i] = bigger(dp[i-1], dp[i-1]+1)
				// if strList[i-1] == "(" {
				// 	dp[i] = dp[i-1] + 1
				// }
				if leftCounter > rightCounter {

				}

			} else if leftCounter < rightCounter {
				leftCounter = 0
				rightCounter = 0
				dp[i] = dp[i-1]
			}
		}
	}
	return dp[len(s)-1]
}
func bigger(x, y int) int {
	if x > y {
		return x
	}
	return y
}
func longestValidParentheses(s string) int {
	leftCounter := 0
	rightCounter := 0
	sub := make([]int, 0)
	for _, str := range s {
		switch str {
		case 40:
			leftCounter++
		case 41:
			rightCounter++
		}
		if rightCounter > leftCounter {
			sub = append(sub, (rightCounter-1)*2)
			leftCounter, rightCounter = 0, 0
		}
	}
	if leftCounter == rightCounter {
		sub = append(sub, rightCounter*2)
	}
	biggestNum := sub[0]
	for i := range sub {
		if sub[i] > biggestNum {
			biggestNum = sub[i]
		}
	}
	return biggestNum
}

// func stack(s string) int {
// 	stack := make([]string, len(s))
// 	for _,str:=range s{

// 	}
// }
