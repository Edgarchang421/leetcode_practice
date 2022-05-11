package main

/*
Example 1:

Input: nums = [0,1,0,3,12]
Output: [1,3,12,0,0]
Example 2:

Input: nums = [0]
Output: [0]

	l := []int{0, 1, 0, 3, 12}
	moveZeroesV2(l)
	fmt.Printf("result: %v\n", l)

	l2 := []int{0, 0}
	moveZeroesV2(l2)
	fmt.Printf("result: %v\n", l2)
*/

// Runtime: 26 ms, faster than 62.78% of Go online submissions for Move Zeroes.
// Memory Usage: 6.6 MB, less than 72.80% of Go online submissions for Move Zeroes.
func moveZeroes(nums []int) {
	for i := len(nums) - 1; i >= 0; {
		if nums[i] == 0 {
			nums = append(nums[:i], nums[i+1:]...)
			nums = append(nums, 0)
			if nums[i] == 0 {
				i--
			}
		} else {
			i--
		}
	}
}

// Runtime: 31 ms, faster than 43.85% of Go online submissions for Move Zeroes.
// Memory Usage: 6.8 MB, less than 68.58% of Go online submissions for Move Zeroes.
func moveZeroesV2(nums []int) {
	count := 0
	for i := 0; i <= len(nums)-1; {
		if nums[i] == 0 {
			nums = append(nums[:i], nums[i+1:]...)
			count++
		} else {
			i++
		}
	}
	for i := 0; i <= count; i++ {
		nums = append(nums, 0)
	}
}
