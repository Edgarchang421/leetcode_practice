package main

import "fmt"

/*
Input: nums = [1,1,2]
Output: 2, nums = [1,2,_]

Input: nums = [0,0,1,1,1,2,2,3,3,4]
Output: 5, nums = [0,1,2,3,4,_,_,_,_,_]

[]int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
*/

// Runtime: 70 ms, faster than 7.05% of Go online submissions for Remove Duplicates from Sorted Array.
// Memory Usage: 4.7 MB, less than 7.30% of Go online submissions for Remove Duplicates from Sorted Array.
func removeDuplicates(nums []int) int {
	m := make(map[int]struct{})
	// r := make([]int, 0, len(nums))
	for i := 0; i <= len(nums)-1; i++ {
		fmt.Printf("i: %v\n", i)
		if _, ok := m[nums[i]]; !ok {
			m[nums[i]] = struct{}{}
			// r = append(r, nums[i])
		} else {
			if i < len(nums)-1 {
				fmt.Printf("nums: %v\n", nums)
				nums = append(nums[:i], nums[i+1:]...)
			} else {
				nums = nums[:i]
			}
			i -= 1
		}
	}

	// nums = r
	fmt.Printf("result nums: %v\n", nums)
	return len(m)
}

// Runtime: 20 ms, faster than 16.94% of Go online submissions for Remove Duplicates from Sorted Array.
// Memory Usage: 4.4 MB, less than 24.24% of Go online submissions for Remove Duplicates from Sorted Array.
func removeDuplicatesVer2(nums []int) int {
	ln := len(nums)
	if ln <= 1 {
		return ln
	}

	j := 0
	for i := 1; i <= ln-1; i++ {
		if nums[i] != nums[j] {
			j++
			nums[j] = nums[i]
		}
	}
	fmt.Printf("result nums: %v\n", nums)
	return j + 1
}
