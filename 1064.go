package main

// Question:
// Given an array A of distinct integers sorted in ascending order, return the smallest index i // that satisfies A[i] == i.  Return -1 if no such i exists.
// 給一個裡面都包含不同數字的陣列 A，排序已經由小到大排好了。請找出符合 A[i] = i 的值。
// 重複則回傳最小

/* Example:
Input: [-10,-5,0,3,7]
Output: 3
Explanation:
For the given array, A[0] = -10, A[1] = -5, A[2] = 0, A[3] = 3, thus the output is 3.
Example 2:

Input: [0,2,5,8,17]
Output: 0
Explanation:
A[0] = 0, thus the output is 0.
Example 3:

Input: [-10,-5,3,4,7,9]
Output: -1
Explanation:
There is no such i that A[i] = i, thus the output is -1.

	d := []int{-10, -5, 0, 3, 7}
	r := searchIndex(d)
	fmt.Printf("result: %v\n", r)

	d = []int{0, 2, 5, 8, 17}
	r = searchIndex(d)
	fmt.Printf("result: %v\n", r)

	d = []int{-10, -5, 3, 4, 7, 9}
	r = searchIndex(d)
	fmt.Printf("result: %v\n", r)

	d = []int{-1, 0, 2, 3, 8, 17, 20}
	r = searchIndex(d)
	fmt.Printf("result: %v\n", r)
*/

func searchIndex(data []int) int {
	left := 0
	right := len(data) - 1

	smallestIndex := -1

	for right >= left {
		m := (left + right) / 2
		switch {
		case data[m] == m:
			smallestIndex = m
			right = m - 1
		case data[m] > m:
			right = m - 1
		case data[m] < m:
			left = m + 1
		}
	}

	return smallestIndex
}
