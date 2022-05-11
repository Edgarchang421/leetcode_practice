package main

/*
Example 1:

Input: nums = [2,2,1]
Output: 1
Example 2:

Input: nums = [4,1,2,1,2]
Output: 4
Example 3:

Input: nums = [1]
Output: 1

	r := singleNumberV2([]int{4, 1, 2, 1, 2})
	fmt.Printf("result: %v\n", r)

	r = singleNumberV2([]int{1})
	fmt.Printf("result: %v\n", r)

	r = singleNumberV2([]int{2, 2, 1})
	fmt.Printf("result: %v\n", r)
*/

// Runtime: 32 ms, faster than 21.14% of Go online submissions for Single Number.
// Memory Usage: 6.7 MB, less than 32.14% of Go online submissions for Single Number.
func singleNumber(nums []int) int {
	// value: index
	m := map[int]struct{}{}
	for _, v := range nums {
		if _, ok := m[v]; ok {
			delete(m, v)
		} else {
			m[v] = struct{}{}
		}
	}
	for value := range m {
		return value
	}
	return -1
}

// Runtime: 16 ms, faster than 84.06% of Go online submissions for Single Number.
// Memory Usage: 6.6 MB, less than 41.26% of Go online submissions for Single Number.
func singleNumberV2(nums []int) int {
	ans := nums[0]
	for i := 1; i <= len(nums)-1; i++ {
		ans = ans ^ nums[i]
	}
	return ans
}
