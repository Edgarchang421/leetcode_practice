package main

/*
	i := []int{1, 2}
	rotate(i, 3)
	fmt.Printf("result: %v\n", i)
*/

// func rotate2(nums []int, k int) {
// 	for i := 0; i < k; i++ {
// 		nums = append(nums[len(nums)-1:], nums[0:len(nums)-1]...)
// 		fmt.Printf("rotate %v: %v\n", i, nums)
// 	}
// }

// func rotate3(nums []int, k int) {
// 	nums = append(nums[len(nums)-k:], nums[0:len(nums)-k]...)
// 	fmt.Printf("rotate: %v\n", nums)
// }

// Runtime: 31 ms, faster than 79.07% of Go online submissions for Rotate Array.
// Memory Usage: 9.1 MB, less than 17.86% of Go online submissions for Rotate Array.
func rotate(nums []int, k int) {
	if k > len(nums) {
		k = k % len(nums)
	}
	twoPointer(nums, 0, len(nums)-1)
	twoPointer(nums, 0, k-1)
	twoPointer(nums, k, len(nums)-1)
}

func twoPointer(d []int, l, r int) {
	for r >= l {
		d[l], d[r] = d[r], d[l]
		l++
		r--
	}
}

// Runtime: 38 ms, faster than 58.89% of Go online submissions for Rotate Array.
// Memory Usage: 7.9 MB, less than 95.52% of Go online submissions for Rotate Array.
func rotate4(nums []int, k int) {
	k = k % len(nums)
	twoPointer(nums, 0, len(nums)-1)
	twoPointer(nums, 0, k-1)
	twoPointer(nums, k, len(nums)-1)
}

// Runtime: 32 ms, faster than 75.50% of Go online submissions for Rotate Array.
// Memory Usage: 8.2 MB, less than 67.19% of Go online submissions for Rotate Array.
func rotate5(nums []int, k int) {
	n := len(nums)
	k = k % n
	twoPointer(nums, 0, n-1)
	twoPointer(nums, 0, k-1)
	twoPointer(nums, k, n-1)
}

// Runtime: 46 ms, faster than 42.28% of Go online submissions for Rotate Array.
// Memory Usage: 8.4 MB, less than 49.67% of Go online submissions for Rotate Array.
func rotate6(nums []int, k int) {
	if len(nums) != 0 {
		copy(nums, append(nums[len(nums)-k%len(nums):], nums[0:len(nums)-k%len(nums)]...))
	}
}
