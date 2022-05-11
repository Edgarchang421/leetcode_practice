package main

import (
	"fmt"
	"strings"
)

/*
    // s := "()[]{}"
	// fmt.Printf("input: %v, result: %v\n", s, isValid(s))

	// s = "{]"
	// fmt.Printf("input: %v, result: %v\n", s, isValid(s))

	// s = "{"
	// fmt.Printf("input: %v, result: %v\n", s, isValid(s))

	// s = "]"
	// fmt.Printf("input: %v, result: %v\n", s, isValid(s))

	// s = "(){[]}"
	// fmt.Printf("input: %v, result: %v\n", s, isValid(s))

	s := "a b [d] { (g) (h) i [-[ {;} ]-] }"
	fmt.Printf("input: %v, result: %v\n", s, isValidVer2(s))
*/
func isValid(s string) bool {
	left := []string{}

	strSplitArray := strings.Split(s, "")
	fmt.Printf("strSplitArray: %v\n", strSplitArray)

	for _, v := range strSplitArray {
		fmt.Printf("process string: %s\n", v)
		switch v {
		case "{", "[", "(":
			left = append(left, v)
		case ")":
			if len(left) == 0 {
				return false
			}
			if left[len(left)-1] != "(" {
				return false
			}
			left = left[:len(left)-1]
		case "]":
			if len(left) == 0 {
				return false
			}
			if left[len(left)-1] != "[" {
				return false
			}
			left = left[:len(left)-1]
		case "}":
			if len(left) == 0 {
				return false
			}
			if left[len(left)-1] != "{" {
				return false
			}
			left = left[:len(left)-1]
		}
	}

	if len(left) != 0 {
		return false
	}

	return true
}

// Runtime: 2 ms, faster than 41.36% of Go online submissions for Valid Parentheses.
// Memory Usage: 2.1 MB, less than 25.91% of Go online submissions for Valid Parentheses.
func isValidClean(s string) bool {
	strSplitArray := strings.Split(s, "")

	if len(strSplitArray)%2 != 0 {
		return false
	}
	tryMap := map[string]string{
		")": "(",
		"]": "[",
		"}": "{",
	}
	left := make([]string, 0, len(strSplitArray))
	l := len(left)

	for _, v := range strSplitArray {
		switch v {
		case "{", "[", "(":
			left = append(left, v)
			l = len(left)

		case ")", "]", "}":
			if l == 0 {
				return false
			}
			if left[l-1] != tryMap[v] {
				return false
			}
			left = left[:l-1]

			l = len(left)
		}
	}

	if len(left) != 0 {
		return false
	}

	return true
}

// Runtime: 0 ms, faster than 100.00% of Go online submissions for Valid Parentheses.
// Memory Usage: 2.1 MB, less than 25.91% of Go online submissions for Valid Parentheses.
func isValidVer2(s string) bool {
	strSplitArray := strings.Split(s, "")
	// if len(strSplitArray)%2 != 0 || len(strSplitArray) == 0 {
	// 	return false
	// }

	tryMap := map[string]string{
		")": "(",
		"]": "[",
		"}": "{",
	}
	left := make([]string, 0, len(strSplitArray))
	l := len(left)

	for _, v := range strSplitArray {
		switch v {
		case "{", "[", "(":
			left = append(left, v)
			l = len(left)
			fmt.Printf("left: %v\n", left)
		case ")", "]", "}":
			if l == 0 || left[l-1] != tryMap[v] {
				return false
			}
			left = left[:l-1]
			l = len(left)
		}
	}

	return len(left) == 0
}

// Runtime: 1 ms, faster than 42.63% of Go online submissions for Valid Parentheses.
// Memory Usage: 1.9 MB, less than 99.36% of Go online submissions for Valid Parentheses.
func isValidVer3(s string) bool {
	if len(s)%2 != 0 || len(s) == 0 {
		return false
	}

	tryMap := map[rune]rune{
		')': '(',
		']': '[',
		'}': '{',
	}
	left := make([]rune, 0, len(s))
	l := len(left)

	for _, v := range s {
		switch v {
		case '{', '[', '(':
			left = append(left, v)
			l = len(left)
		case ')', ']', '}':
			if l == 0 || left[l-1] != tryMap[v] {
				return false
			}
			left = left[:l-1]
			l = len(left)
		}
	}

	return len(left) == 0
}
