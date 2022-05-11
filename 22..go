package main

import "fmt"

/*
["(())(())","()()(())",]
4

(((()))) ((()))() ()(())() ((())()) ((()())) (()())() (()()()) ()((())) (()(())) ()()(()) (())()() ()(()()) ()()()()

(((()))) ((()))() ()(())() ((())()) ((()())) (()())() (()()()) ()((())) (()(())) ()()(()) (())()() ()(()()) ()()()()

Input: n = 3
Output: ["((()))","(()())","(())()","()(())","()()()"]

三個分開 兩個分開 全部括弧

()() 遇到一個左括弧加一個 遇到一個又括弧加一個 用counter
(()) ()()

Example 2:

Input: n = 1
Output: ["()"]
*/
// 錯誤
func generateParenthesisFail(n int) []string {
	dp := make([][]string, n+1)
	dp[1] = []string{"()"}
	for i := 2; i <= n; i++ {
		r := make([]string, 0)
		for j := 0; j < len(dp[i-1]); j++ {
			str := dp[i-1][j]
			a := fmt.Sprintf("(%v)", str)
			b := fmt.Sprintf("()%v", str)
			c := fmt.Sprintf("%v()", str)
			r = append(r, []string{a, b, c}...)
		}
		dp[i] = r
		fmt.Printf("dp: %v\n", dp)
	}
	return dp[n]
}

// 錯誤
func generateParenthesisFailV2(n int) []string {
	dp := make(map[int]map[string]struct{})
	num1 := make(map[string]struct{})
	num1["()"] = struct{}{}
	dp[1] = num1

	for i := 2; i <= n; i++ {
		result := make(map[string]struct{})
		for str := range dp[i-1] {
			a := fmt.Sprintf("(%v)", str)
			b := fmt.Sprintf("()%v", str)
			c := fmt.Sprintf("%v()", str)
			result[a] = struct{}{}
			result[b] = struct{}{}
			result[c] = struct{}{}
		}
		dp[i] = result
	}

	ans := make([]string, 0)
	for str := range dp[n] {
		ans = append(ans, str)
	}
	return ans
}

// Runtime: 3 ms, faster than 51.75% of Go online submissions for Generate Parentheses.
// Memory Usage: 2.7 MB, less than 92.59% of Go online submissions for Generate Parentheses.
func generateParenthesis(n int) []string {
	//output at each stage
	op := ""

	//final output
	var s []string

	//passing by reference
	solve(n, n, &s, op)
	return s
}

//taking pointer to *[]string because it needs update
func solve(left, right int, ans *[]string, str string) {
	fmt.Printf("o: %v, c: %v, s: %v, op: %v ,recu\n", left, right, *ans, str)
	// base case
	// 已經加入正確數量的左右括弧了，將str加入倒ans中
	if left == 0 && right == 0 {
		//update the pointer for base case, just add whatever is the output

		*ans = append(*ans, str)
		fmt.Printf("s: %v\n", *ans)
		fmt.Printf("o: %v, c: %v, s: %v, op: %v ,left==right==0 return\n", left, right, *ans, str)
		return
	}

	// ( case, when we have opening backet we add '(' to output
	// 一定要先加入左括弧，才能和右括弧形成正確的排序
	if left != 0 {
		op1 := str + "("
		fmt.Printf("o: %v, op: %v\n", left, op1)
		solve(left-1, right, ans, op1)
	}

	// ) case , when we have opening backet we add ')' to output only if closing count is greater
	// 左括弧的數量已經使用完了，如果右括弧的數量還多於左括弧，就需要加進去
	if right > left {
		op2 := str + ")"
		fmt.Printf("c: %v, op: %v\n", right, op2)
		solve(left, right-1, ans, op2)
	}

	fmt.Printf("o: %v, c: %v, s: %v, op: %v return\n", left, right, *ans, str)
	return
}
