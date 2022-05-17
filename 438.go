package main

/*
Constraints:
1 <= s.length, p.length <= 3 * 10^4
s and p consist of lowercase English letters.

Example 1:

Input: s = "cbaebabacd", p = "abc"
Output: [0,6]
Explanation:
The substring with start index = 0 is "cba", which is an anagram of "abc".
The substring with start index = 6 is "bac", which is an anagram of "abc".
Example 2:

Input: s = "abab", p = "ab"
Output: [0,1,2]
Explanation:
The substring with start index = 0 is "ab", which is an anagram of "ab".
The substring with start index = 1 is "ba", which is an anagram of "ab".
The substring with start index = 2 is "ab", which is an anagram of "ab".
*/

// Runtime: 476 ms, faster than 8.24% of Go online submissions for Find All Anagrams in a String.
// Memory Usage: 5 MB, less than 82.42% of Go online submissions for Find All Anagrams in a String.
func findAnagrams(s string, p string) []int {
	result := make([]int, 0)
	for i := 0; i <= len(s)-len(p); i++ {
		str := s[i : i+len(p)]
		if isAnagram438(str, p) {
			result = append(result, i)
		}
	}

	return result
}

func isAnagram438(s string, t string) bool {
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

// Runtime: 32 ms, faster than 31.59% of Go online submissions for Find All Anagrams in a String.
// Memory Usage: 5.2 MB, less than 43.13% of Go online submissions for Find All Anagrams in a String.
func findAnagrams2(s string, p string) []int {
	plen := len(p)
	slen := len(s)
	m := make(map[rune]int, 26)
	aNum := rune('a')
	for _, str := range p {
		if _, ok := m[str-aNum]; ok {
			m[str-aNum]++
		} else {
			m[str-aNum] = 1
		}
	}
	// 沒有使用的字
	shouldNotExist := make(map[rune]int)

	result := make([]int, 0)
	var sStrRune rune
	for i := 0; i <= slen-1; i++ {
		// fmt.Printf("str: %v\n", rune(s[i])-aNum)
		sStrRune = rune(s[i]) - aNum
		if i >= plen {
			checkStrRune := rune(s[i-plen]) - aNum
			if _, ok := m[checkStrRune]; ok {
				m[checkStrRune]++
			}
			if _, ok := shouldNotExist[checkStrRune]; ok {
				shouldNotExist[checkStrRune]--
			}
		}
		// fmt.Printf("check i-plen: %v\n", m)
		if _, ok := m[sStrRune]; ok {
			m[sStrRune]--
		} else {
			if _, ok := shouldNotExist[sStrRune]; ok {
				shouldNotExist[sStrRune]++
			} else {
				shouldNotExist[sStrRune] = 1
			}
			continue
		}

		isAnagram := true
		for j := range m {
			if m[j] != 0 {
				isAnagram = false
				break
			}
		}
		for j := range shouldNotExist {
			if shouldNotExist[j] != 0 {
				isAnagram = false
				break
			}
		}
		if isAnagram {
			result = append(result, i-plen+1)
		}
		// fmt.Println(m)
	}

	return result
}

// Runtime: 17 ms, faster than 50.55% of Go online submissions for Find All Anagrams in a String.
// Memory Usage: 5 MB, less than 82.42% of Go online submissions for Find All Anagrams in a String.
func findAnagrams3(s string, p string) []int {
	plen := len(p)
	slen := len(s)
	m := make(map[rune]int, 26)
	aNum := rune('a')
	for _, str := range p {
		strNum := str - aNum
		if _, ok := m[strNum]; ok {
			m[strNum]++
		} else {
			m[strNum] = 1
		}
	}
	result := make([]int, 0)
	var shouldNotExistCounter int
	var sStrRune rune
	for i := 0; i <= slen-1; i++ {
		sStrRune = rune(s[i]) - aNum
		if i >= plen {
			checkStrRune := rune(s[i-plen]) - aNum
			if _, ok := m[checkStrRune]; ok {
				m[checkStrRune]++
			} else {
				shouldNotExistCounter--
			}
		}
		if _, ok := m[sStrRune]; ok {
			m[sStrRune]--
		} else {
			shouldNotExistCounter++
			continue
		}
		if shouldNotExistCounter > 0 {
			continue
		}
		isAnagram := true
		for j := range m {
			if m[j] != 0 {
				isAnagram = false
				break
			}
		}
		if isAnagram {
			result = append(result, i-plen+1)
		}
	}

	return result
}

// Runtime: 3 ms, faster than 96.15% of Go online submissions for Find All Anagrams in a String.
// Memory Usage: 5.1 MB, less than 82.42% of Go online submissions for Find All Anagrams in a String.
func findAnagramsExample1(s string, p string) []int {
	if len(s) < len(p) {
		return nil
	}
	var pCounter, sCounter [26]int
	for i := range p {
		pCounter[p[i]-'a']++
		sCounter[s[i]-'a']++
	}
	var out []int
	for i := 0; i < len(s)-len(p)+1; i++ {
		if pCounter == sCounter {
			out = append(out, i)
		}
		if i+len(p) < len(s) {
			sCounter[s[i]-'a']--
			sCounter[s[i+len(p)]-'a']++
		}
	}
	return out
}
