package main

import (
	"bytes"
	"strings"
)

/*
	t := []string{"gin", "zen", "gig", "msg"}

	fmt.Print(uniqueMorseRepresentations(t))
*/

// reduce space, 2.2MB
func uniqueMorseRepresentationsVersion2(words []string) int {
	mcslice := [...]string{".-", "-...", "-.-.", "-..", ".", "..-.", "--.", "....", "..", ".---", "-.-", ".-..", "--", "-.", "---", ".--.", "--.-", ".-.", "...", "-", "..-", "...-", ".--", "-..-", "-.--", "--.."}
	rm := make(map[string]struct{})
	var moresCode bytes.Buffer
	var str string
	for _, v := range words {
		for _, w := range v {
			moresCode.WriteString(mcslice[w-'a'])
		}
		str = moresCode.String()
		rm[str] = struct{}{}
		moresCode.Reset()
	}
	return len(rm)
}

// reduce time
func uniqueMorseRepresentations(words []string) int {
	morseCodeAlphabetMap := map[string]string{
		"a": ".-",
		"b": "-...",
		"c": "-.-.",
		"d": "-..",
		"e": ".",
		"f": "..-.",
		"g": "--.",
		"h": "....",
		"i": "..",
		"j": ".---",
		"k": "-.-",
		"l": ".-..",
		"m": "--",
		"n": "-.",
		"o": "---",
		"p": ".--.",
		"q": "--.-",
		"r": ".-.",
		"s": "...",
		"t": "-",
		"u": "..-",
		"v": "...-",
		"w": ".--",
		"x": "-..-",
		"y": "-.--",
		"z": "--..",
	}
	strSet := make(map[string]struct{})
	var splitString []string
	var morseCodeBytes bytes.Buffer
	var morseCodeStr string
	for _, word := range words {
		splitString = strings.Split(word, "")
		for _, letter := range splitString {
			morseCodeBytes.WriteString(morseCodeAlphabetMap[letter])
		}

		morseCodeStr = morseCodeBytes.String()
		strSet[morseCodeStr] = struct{}{}
		morseCodeBytes.Reset()
	}
	return len(strSet)
}

// 二進制 https://leetcode.com/problems/unique-morse-code-words/discuss/550682/The-coolest-solution-using-bit-manipulation-in-Go
