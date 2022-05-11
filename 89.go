package main

import "fmt"

/*
Example 1:

Input: n = 2
Output: [0,1,3,2]
Explanation:
The binary representation of [0,1,3,2] is [00,01,11,10].
- 00 and 01 differ by one bit
- 01 and 11 differ by one bit
- 11 and 10 differ by one bit
- 10 and 00 differ by one bit
[0,2,3,1] is also a valid gray code sequence, whose binary representation is [00,10,11,01].
- 00 and 10 differ by one bit
- 10 and 11 differ by one bit
- 11 and 01 differ by one bit
- 01 and 00 differ by one bit
Example 2:

Input: n = 1
Output: [0,1]

3
[0,1,3,2,6,7,5,4]
*/

// Runtime: 7 ms, faster than 89.66% of Go online submissions for Gray Code.
// Memory Usage: 6.9 MB, less than 55.17% of Go online submissions for Gray Code.
func grayCode(n int) []int {
	result := make([]int, 1)
	// for (int i = 1; i <= n; i++) {
	// 	int len = res.size();
	// 	for (int j = len - 1; j >= 0; j--) {
	// 		res.add(res.get(j) + adder);
	// 	}
	// 	adder *= 2;
	// }

	addr := 1
	for i := 1; i <= n; i++ {
		l := len(result)
		for j := l - 1; j >= 0; j-- {
			result = append(result, result[j]+addr)
		}
		addr = addr << 1
	}

	return result
}

// Runtime: 16 ms, faster than 43.10% of Go online submissions for Gray Code.
// Memory Usage: 7.7 MB, less than 29.31% of Go online submissions for Gray Code.
func grayCodeExample(n int) []int {

	// total 2^n codes for bit length n
	codeCount := 1 << n

	// slice to store gray codes
	grayCode := make([]int, codeCount)

	for i := 0; i < codeCount; i += 1 {

		toggleMask := (i >> 1)
		// ^為XOR運算
		grayCode[i] = i ^ toggleMask
		fmt.Printf("i: %v, toggle_mask: %v, i ^ toggleMask: %v\n", i, toggleMask, grayCode[i])
	}

	return grayCode
}
