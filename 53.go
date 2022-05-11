package main

import "fmt"

/*
Example 1:
Input: nums = [-2,1,-3,4,-1,2,1,-5,4]
Output: 6
Explanation: [4,-1,2,1] has the largest sum = 6.

Example 2:
Input: nums = [1]
Output: 1

Example 3:
Input: nums = [5,4,-1,7,8]
Output: 23

[]int{-2, 1, -3, 4, -1, 2, 1, -5, 4}
*/

// 從最小的陣列(只有一個)開始算，每次多加一個元素再看最大子陣列的植有沒有變
func maxSubArray(nums []int) int {
	dp := make([]int, len(nums))
	dp[0] = nums[0]
	for i := 1; i <= len(nums)-1; i++ {
		// list := nums[:i]
		// fmt.Printf("list: %v\n", list)
		// 多加了一個元素後，總共i個元素，會多出i個子陣列，只包含自己，往前加一加二...，所以會包含最新元素的子陣列一定會包含i和i-1
		// i個之中最大的再和前面的所有i-1個元素中最大的比較
		if dp[i-1] < 0 && nums[i] > dp[i-1] {
			dp[i] = nums[i]
			continue
		}

		dp[i] = big2(dp[i-1]+nums[i], dp[i-1])
		fmt.Printf("dp: %v\n", dp)
	}

	return dp[len(nums)-1]
}

func big2(x, y int) int {
	if x < y {
		return y
	}
	return x
}

// Input: nums = [-2,1,-3,4,-1,2,1,-5,4]
// Output: 6
// Runtime: 100 ms, faster than 77.58% of Go online submissions for Maximum Subarray.
// Memory Usage: 8.4 MB, less than 99.31% of Go online submissions for Maximum Subarray.
func maxSubArrayVer2(nums []int) int {
	all := nums[0]    // 陣列中最大的子陣列總和
	newest := nums[0] // 現在陣列中最大的子陣列總和，要判斷是否要加上最新element，或者原本的最大子陣列總和是否小於零
	for i := 1; i <= len(nums)-1; i++ {
		fmt.Printf("b res: %v, now: %v, nums[%v]: %v, nums: %v\n", all, newest, i, nums[i], nums[:i+1])
		// 如果now(dp[i-1])已經小於0，就直接用最新的element作為now
		if newest < 0 {
			newest = nums[i]
		} else { // now大於0，先相加，等等再比較
			// 這邊會暫存i-1，直到某次減到變成負的，下個i再變正的就直接開始新的subarray
			newest += nums[i]
		}

		// 如果有了新元素的最大子陣列總和，大於原本沒有最新元素的最大子陣列總和
		if newest > all {
			all = newest
		}
		fmt.Printf("a res: %v, now: %v, nums[%v]: %v, nums: %v\n\n", all, newest, i, nums[i], nums[:i+1])
	}

	return all
}

func maxSubArrayTest(nums []int) int {
	dp := make([]int, len(nums))
	dp[0] = nums[0]
	includNewestElement := make([]int, len(nums))
	includNewestElement[0] = nums[0]

	for i := 1; i <= len(nums)-1; i++ {
		if includNewestElement[i-1] < 0 {
			includNewestElement[i] = nums[i]
		} else {
			includNewestElement[i] = includNewestElement[i-1] + nums[i]
		}

		dp[i] = big2(dp[i-1], includNewestElement[i])
	}

	return dp[len(nums)-1]
}
