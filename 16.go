package main

import "sort"

/*
Example 1:

Input: nums = [-1,2,1,-4], target = 1
Output: 2
Explanation: The sum that is closest to the target is 2. (-1 + 2 + 1 = 2).
Example 2:

Input: nums = [0,0,0], target = 1
Output: 0

0 1 2, 0
3

0 1 1 1, 100
3

[1,1,-1,-1,3]
-1

1 2 5 10 11 , 12

[1,2,4,8,16,32,64,128]
82

[]int{1, 2, 4, 8, 16, 32, 64, 128}, 82
*/
// Runtime: 10 ms, faster than 49.33% of Go online submissions for 3Sum Closest.
// Memory Usage: 2.8 MB, less than 72.87% of Go online submissions for 3Sum Closest.
func threeSumClosest(nums []int, target int) int {
	sort.Slice(nums, func(i, j int) bool { return nums[i] < nums[j] })
	l := len(nums)
	left := 1
	right := l - 1
	ans := nums[0] + nums[left] + nums[right]
	for i := 0; i <= l-2; i++ {
		for right > left {
			sum := nums[i] + nums[left] + nums[right]
			ans = closest(target, sum, ans)
			if sum == target {
				return ans
			} else if sum > target {
				right--
			} else if sum < target {
				left++
			}
		}
		left = i + 2
		right = l - 1
	}
	return ans
}

func myABS(n int) int {
	if n > 0 {
		return n
	}
	return 0 - n
}

func closest(target, x, y int) int {
	if myABS(target-x) < myABS(target-y) {
		return x
	}
	return y
}
