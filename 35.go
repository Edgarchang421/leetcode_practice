package main

/*
Example 1:

Input: nums = [1,3,5,6], target = 5
Output: 2
Example 2:

Input: nums = [1,3,5,6], target = 2
Output: 1
Example 3:

Input: nums = [1,3,5,6], target = 7
Output: 4

	r := searchInsert([]int{1, 3, 5, 6}, 5)

	r = searchInsert([]int{1, 3, 5, 6}, 2)

	r = searchInsert([]int{1, 3, 5, 6}, 7)
*/

// Runtime: 2 ms, faster than 73.99% of Go online submissions for Search Insert Position.
// Memory Usage: 2.9 MB, less than 100.00% of Go online submissions for Search Insert Position.
func searchInsert(nums []int, target int) int {
	l, r := 0, len(nums)-1

	for r >= l {
		m := (l + r) / 2
		switch {
		case nums[m] == target:
			return m
		case nums[m] < target:
			l = m + 1
		case nums[m] > target:
			r = m - 1
		}
	}

	return l
}
