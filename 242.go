package main

/*
Example 1:

Input: s = "anagram", t = "nagaram"
Output: true
Example 2:

Input: s = "rat", t = "car"
Output: false

fmt.Printf("ans: %v\n", isAnagram("anagram", "nagaram"))
fmt.Printf("ans: %v\n", isAnagram("rat", "car"))
*/

// Runtime: 6 ms, faster than 68.94% of Go online submissions for Valid Anagram.
// Memory Usage: 2.8 MB, less than 76.19% of Go online submissions for Valid Anagram.
func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	m := make(map[rune]int)
	for _, str := range s {
		if _, ok := m[str]; ok {
			m[str]++
		} else {
			m[str] = 1
		}
	}
	for _, str := range t {
		if _, ok := m[str]; ok {
			m[str]--
		} else {
			return false
		}
	}
	for _, v := range m {
		if v != 0 {
			return false
		}
	}
	return true
}

// Runtime: 0 ms, faster than 100.00% of Go online submissions for Valid Anagram.
// Memory Usage: 2.8 MB, less than 76.19% of Go online submissions for Valid Anagram.
func isAnagram2(s string, t string) bool {
	m := make([]int, 26)
	aNum := rune('a')
	for _, str := range s {
		m[str-aNum]++
	}
	for _, str := range t {
		m[str-aNum]--
	}
	for _, val := range m {
		if val != 0 {
			return false
		}
	}
	return true
}
