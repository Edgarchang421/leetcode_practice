package main

import (
	"bytes"
	"math/big"
	"strings"
)

/*

Example 1:

Input: a = "11", b = "1"
Output: "100"
Example 2:

Input: a = "1010", b = "1011"
Output: "10101"

*/

// Runtime: 0 ms, faster than 100.00% of Go online submissions for Add Binary.
// Memory Usage: 6.8 MB, less than 5.83% of Go online submissions for Add Binary.
func addBinary(a string, b string) string {
	aStrs := strings.Split(a, "")
	bStrs := strings.Split(b, "")

	i := len(aStrs) - 1
	j := len(bStrs) - 1

	// var l int
	// if i > j {
	// 	l = i + 1
	// } else {
	// 	l = j + 1
	// }

	// 110 011 101 進位 留零
	// 000 不進位 留零
	// 111 進位 留一
	// 001 010 100 不進位 留一
	var carry bool
	result := make([]string, 0)
	for i >= 0 && j >= 0 {
		switch {
		case aStrs[i] == "0" && bStrs[j] == "0":
			if carry {
				result = append([]string{"1"}, result...)
			} else {
				result = append([]string{"0"}, result...)
			}
			carry = false
		case aStrs[i] == "1" && bStrs[j] == "1":
			if carry {
				result = append([]string{"1"}, result...)
			} else {
				result = append([]string{"0"}, result...)
			}
			carry = true
		default:
			if carry {
				result = append([]string{"0"}, result...)
				carry = true
			} else {
				result = append([]string{"1"}, result...)
				carry = false
			}
		}

		i--
		j--
	}

	for i >= 0 {
		switch aStrs[i] {
		case "1":
			if carry {
				result = append([]string{"0"}, result...)
				carry = true
			} else {
				result = append([]string{"1"}, result...)
			}
		case "0":
			if carry {
				result = append([]string{"1"}, result...)
				carry = false
			} else {
				result = append([]string{"0"}, result...)
			}
		}
		i--
	}
	for j >= 0 {
		switch bStrs[j] {
		case "1":
			if carry {
				result = append([]string{"0"}, result...)
				carry = true
			} else {
				result = append([]string{"1"}, result...)
			}
		case "0":
			if carry {
				result = append([]string{"1"}, result...)
				carry = false
			} else {
				result = append([]string{"0"}, result...)
			}
		}
		j--
	}
	if carry {
		result = append([]string{"1"}, result...)
	}

	// if i >= 0 {
	// 	result = append(aStrs[:i+1], result...)
	// }
	// if j >= 0 {
	// 	result = append(bStrs[:j+1], result...)
	// }
	// fmt.Printf("str array: %v\n", result)

	var r bytes.Buffer
	for i := 0; i <= len(result)-1; i++ {
		r.Write([]byte(result[i]))
	}

	return r.String()
}

func addBinaryClean(a string, b string) string {
	aStrs := strings.Split(a, "")
	bStrs := strings.Split(b, "")

	i := len(aStrs) - 1
	j := len(bStrs) - 1

	var carry bool
	result := make([]string, 0)
	for i >= 0 && j >= 0 {
		switch {
		case aStrs[i] == "0" && bStrs[j] == "0":
			if carry {
				result = append([]string{"1"}, result...)
			} else {
				result = append([]string{"0"}, result...)
			}
			carry = false
		case aStrs[i] == "1" && bStrs[j] == "1":
			if carry {
				result = append([]string{"1"}, result...)
			} else {
				result = append([]string{"0"}, result...)
			}
			carry = true
		default:
			if carry {
				result = append([]string{"0"}, result...)
				carry = true
			} else {
				result = append([]string{"1"}, result...)
				carry = false
			}
		}

		i--
		j--
	}

	for i >= 0 {
		switch aStrs[i] {
		case "1":
			if carry {
				result = append([]string{"0"}, result...)
				carry = true
			} else {
				result = append([]string{"1"}, result...)
			}
		case "0":
			if carry {
				result = append([]string{"1"}, result...)
				carry = false
			} else {
				result = append([]string{"0"}, result...)
			}
		}
		i--
	}
	for j >= 0 {
		switch bStrs[j] {
		case "1":
			if carry {
				result = append([]string{"0"}, result...)
				carry = true
			} else {
				result = append([]string{"1"}, result...)
			}
		case "0":
			if carry {
				result = append([]string{"1"}, result...)
				carry = false
			} else {
				result = append([]string{"0"}, result...)
			}
		}
		j--
	}
	if carry {
		result = append([]string{"1"}, result...)
	}

	var r bytes.Buffer
	for i := 0; i <= len(result)-1; i++ {
		r.Write([]byte(result[i]))
	}

	return r.String()
}

func addBinaryVer2(a string, b string) string {
	aInt, bInt, sum := new(big.Int), new(big.Int), new(big.Int)
	aInt.SetString(a, 2)
	bInt.SetString(b, 2)
	sum.Add(aInt, bInt)
	return sum.Text(2)
}
