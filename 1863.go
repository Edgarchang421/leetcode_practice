package main

import "fmt"

/*
Example 1:

Input: nums = [1,3]
Output: 6
Explanation: The 4 subsets of [1,3] are:
- The empty subset has an XOR total of 0.
- [1] has an XOR total of 1.
- [3] has an XOR total of 3.
- [1,3] has an XOR total of 1 XOR 3 = 2.
0 + 1 + 3 + 2 = 6


Example 2:

Input: nums = [5,1,6]
Output: 28
Explanation: The 8 subsets of [5,1,6] are:
- The empty subset has an XOR total of 0.
- [5] has an XOR total of 5.
- [1] has an XOR total of 1.
- [6] has an XOR total of 6.
- [5,1] has an XOR total of 5 XOR 1 = 4.
- [5,6] has an XOR total of 5 XOR 6 = 3.
- [1,6] has an XOR total of 1 XOR 6 = 7.
- [5,1,6] has an XOR total of 5 XOR 1 XOR 6 = 2.
0 + 5 + 1 + 6 + 4 + 3 + 7 + 2 = 28


Example 3:

Input: nums = [3,4,5,6,7,8]
Output: 480
Explanation: The sum of all XOR totals for every subset is 480.

	// fmt.Printf("result: %v\n", subsetXORSum([]int{3, 4, 5, 6, 7, 8}))
	// fmt.Printf("result: %v\n", subsetXORSum([]int{1, 3}))
	fmt.Printf("result: %v\n", subsetXORSum2([]int{5, 1, 6}))
	// fmt.Printf("result: %v\n", subsetXORSum2([]int{5, 5}))

	// a := []int{1, 3}
	// b := append(a[:1], a[2:]...)
	// fmt.Printf("b: %v\n", b)
*/

// Runtime: 18 ms, faster than 6.67% of Go online submissions for Sum of All Subset XOR Totals.
// Memory Usage: 7.7 MB, less than 6.67% of Go online submissions for Sum of All Subset XOR Totals.
func subsetXORSum(nums []int) int {
	result := 0
	for i := 0; i <= len(nums); i++ {
		subset := []int{}
		backTracking(&result, subset, i, nums, 0)
	}
	return result
}

// 選擇 sub set
// 限制	set無序，所以 [1,3] == [3,1]
// 目標 所有 sub set
func backTracking(result *int, subSet []int, elementNeeded int, nums []int, position int) {
	if elementNeeded == 0 {
		fmt.Printf("subSet: %v\n", subSet)
		subArrayXOR := 0
		for i := range subSet {
			subArrayXOR ^= subSet[i]
		}
		*result += subArrayXOR
		return
	}

	for i := position; i <= len(nums)-1; i++ {
		newSubSet := append(subSet, nums[i])
		backTracking(result, newSubSet, elementNeeded-1, nums, i+1)
	}
}

// Runtime: 6 ms, faster than 30.00% of Go online submissions for Sum of All Subset XOR Totals.
// Memory Usage: 1.9 MB, less than 50.00% of Go online submissions for Sum of All Subset XOR Totals.
func subsetXORSum2(nums []int) int {
	result := 0
	// i為subset所需elements數量
	for i := 0; i <= len(nums); i++ {
		backTracking2(&result, nums, i, 0)
	}
	return result
}

func backTracking2(result *int, nums []int, elementNeeded int, current int) {
	if elementNeeded == 0 {
		*result += current
		return
	}
	for i := range nums {
		// backtracking 限制: set是無序的，故 [5,1,6] == [1,6,5], nums[i+1:]使得選用過的不會再選取
		// elementNeeded = elementNeeded -1
		backTracking2(result, nums[i+1:], elementNeeded-1, current^nums[i])
	}
}

// Runtime: 2 ms, faster than 60.00% of Go online submissions for Sum of All Subset XOR Totals.
// Memory Usage: 1.8 MB, less than 86.67% of Go online submissions for Sum of All Subset XOR Totals.
func subsetXORSumExample(nums []int) int {
	res := 0
	dfs1863(nums, 0, &res)
	return res
}

func dfs1863(nums []int, cur int, res *int) {
	if len(nums) == 0 {
		*res += cur
		return
	}

	dfs1863(nums[1:], cur^nums[0], res)
	dfs1863(nums[1:], cur, res)
}
