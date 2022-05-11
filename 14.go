package main

import (
	"bytes"
	"strings"
)

/*
Input: strs = ["flower","flow","flight"]
Output: "fl"

Input: strs = ["dog","racecar","car"]
Output: ""

["ab", "a"]
a

["aaa","aa","aaa"]
aa

[]string{"flower", "flow", "flight"}

[]string{"dog", "racecar", "car"}

[]string{"aaa", "aa", "aaa"}
*/
// Runtime: 5 ms, faster than 21.72% of Go online submissions for Longest Common Prefix.
// Memory Usage: 3.7 MB, less than 8.88% of Go online submissions for Longest Common Prefix.
func longestCommonPrefix(strs []string) string {
	ans := []string{}
	for i := range strs {
		strArray := strings.Split(strs[i], "")
		if len(strArray) < len(ans) {
			ans = ans[:len(strArray)]
		}
		for j := range strArray {
			// 第一個字全部加進去
			if i == 0 {
				ans = append(ans, strArray[j])
				continue
			}

			// 已經沒有相同prefix
			if len(ans) == 0 || j > len(ans)-1 {
				break
			}

			// 第二個字開始的處理
			if strArray[j] == ans[j] {
				continue
			} else {
				ans = ans[:j]
				break
			}
		}
	}

	var finalPrefix bytes.Buffer
	for i := range ans {
		finalPrefix.WriteString(ans[i])
	}

	return finalPrefix.String()
}

func longestCommonPrefixVer2(strs []string) string {
	var res string

	// 先取第一個字的字母，後面的字都比較過，就判斷是否為字首

	// 只跑第一個字的長度
	for i := 0; i < len(strs[0]); i++ {
		for j := 0; j < len(strs); j++ {
			// 其他的字
			// 當這個字的長度，比現在第一個字for loop的i，第i個英文字母，比較小時就回傳現在的prefix
			// 第j個字的第i的字母不等於第0個字的第i個字母，回傳現在的prefix
			if i >= len(strs[j]) || strs[j][i] != strs[0][i] {
				return res
			}
		}

		// 其他所有字都有這個字母
		res += string(strs[0][i])
	}

	return res
}
