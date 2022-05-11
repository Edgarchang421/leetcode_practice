package main

import (
	"fmt"
	"sort"
)

/*
Example 1:
Input: nums = [-1,0,1,2,-1,-4]
Output: [[-1,-1,2],[-1,0,1]]

Example 2:
Input: nums = []
Output: []

Example 3:
Input: nums = [0]
Output: []

[0,0,0,0]
[0,0,0]

回傳三個數字的index都不相等

答案不會重複

[]int{0, 0, 0, 0}
[]int{-1, 0, 1, 2, -1, -4}
*/
func threeSum(nums []int) [][]int {
	l := len(nums)
	if l < 3 {
		return nil
	}

	// 排列過，才能利用two pointer減少時間複雜度
	sort.Slice(nums, func(i, j int) bool { return nums[i] < nums[j] })

	ans := [][]int{}

	// 每一個數字都要搭配two pointer，len-2是因為right point永遠都會使用到最後一個num
	for i := 0; i <= l-2; i++ {
		fmt.Printf("\ni: %v\n", i)
		// 從i=1開始，判斷nums[i]是不是和nums[i-1]相同，如果有相同要直接跳過，答案中不能有重複
		if i > 0 && nums[i] == nums[i-1] {
			fmt.Printf("nums[%v] %v == nums[%v-1] %v , continue\n", i, nums[i], i, nums[i-1])
			continue
		}

		// left point 從 i 的下一個開始
		leftIndex := i + 1
		// right point　從最後一個num開始
		rightIndex := l - 1
		fmt.Printf("i init: left right: %v %v %v\n", i, leftIndex, rightIndex)

		// two pointer尋找符合 nums[i] + nums[leftIndex] + nums[rightIndex] == 0 的目標index
		// right point
		for rightIndex > leftIndex {
			fmt.Printf("i left right: %v %v %v\n", i, leftIndex, rightIndex)
			sum := nums[i] + nums[leftIndex] + nums[rightIndex]
			if sum == 0 {
				ans = append(ans, []int{nums[i], nums[leftIndex], nums[rightIndex]})
				fmt.Printf("i: %v, left: %v, right: %v\n", i, leftIndex, rightIndex)
				fmt.Printf("ans: %v\n", ans)

				rightIndex--
				// 判斷 right point -- 之後，是否和上一個num相同，如果相同就繼續往下 -- ，因為答案不能重複
				for rightIndex > leftIndex && nums[rightIndex] == nums[rightIndex+1] {
					fmt.Printf("nums[rightIndex: %v] %v == nums[rightIndex: %v+1] %v , continue\n", rightIndex, nums[rightIndex], rightIndex, nums[rightIndex+1])
					rightIndex--
				}
			} else if sum > 0 {
				fmt.Printf("sum > 0: %v, continue\n", sum)
				rightIndex--
			} else if sum < 0 {
				fmt.Printf("sum < 0: %v, continue\n", sum)
				leftIndex++
			}
		}
	}

	return ans
}

// Runtime: 49 ms, faster than 48.95% of Go online submissions for 3Sum.
// Memory Usage: 7.4 MB, less than 71.53% of Go online submissions for 3Sum.
func threeSumClean(nums []int) [][]int {
	l := len(nums)
	if l < 3 {
		return nil
	}

	sort.Slice(nums, func(i, j int) bool { return nums[i] < nums[j] })

	ans := [][]int{}

	for i := 0; i <= l-2; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		leftIndex := i + 1
		rightIndex := l - 1
		for rightIndex > leftIndex {
			sum := nums[i] + nums[leftIndex] + nums[rightIndex]
			if sum == 0 {
				ans = append(ans, []int{nums[i], nums[leftIndex], nums[rightIndex]})

				rightIndex--
				for rightIndex > leftIndex && nums[rightIndex] == nums[rightIndex+1] {
					rightIndex--
				}
			} else if sum > 0 {
				rightIndex--
			} else if sum < 0 {
				leftIndex++
			}
		}
	}

	return ans
}
