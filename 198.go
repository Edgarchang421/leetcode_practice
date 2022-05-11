package main

/*
Example 1:

Input: nums = [1,2,3,1]
Output: 4
Explanation: Rob house 1 (money = 1) and then rob house 3 (money = 3).
Total amount you can rob = 1 + 3 = 4.
Example 2:

Input: nums = [2,7,9,3,1]
Output: 12
Explanation: Rob house 1 (money = 2), rob house 3 (money = 9) and rob house 5 (money = 1).
Total amount you can rob = 2 + 9 + 1 = 12.

Constraints:
1 <= nums.length <= 100
0 <= nums[i] <= 400

	r := rob2([]int{2, 7, 9, 3, 1})
*/

// dp[i] = dp[i-1] + 如果有搶最後一間，那最新的就不能加入，沒有則否
// dp[i] = 一種情況是最新增加屋子的有搶，那就不能搶倒數第二間，就要使用dp[j-2] + nums[i] 和 dp[j-1] 比大小
// Runtime: 0 ms, faster than 100.00% of Go online submissions for House Robber.
// Memory Usage: 1.9 MB, less than 78.52% of Go online submissions for House Robber.
func rob(nums []int) int {
	l := len(nums)
	if l == 0 {
		return 0
	} else if l == 1 {
		return nums[0]
	}
	dp := make([]int, l)
	dp[0] = nums[0]
	last := true
	if nums[1] > nums[0] {
		dp[1] = nums[1]
	} else {
		dp[1] = nums[0]
		last = false
	}
	for i := 2; i <= l-1; i++ {
		if last {
			if dp[i-2]+nums[i] > dp[i-1] {
				dp[i] = dp[i-2] + nums[i]
				last = true
			} else {
				dp[i] = dp[i-1]
				last = false
			}
		} else {
			dp[i] = dp[i-1] + nums[i]
			last = true
		}
	}
	return dp[l-1]
}

// Runtime: 1 ms, faster than 45.52% of Go online submissions for House Robber.
// Memory Usage: 2 MB, less than 32.78% of Go online submissions for House Robber.
func rob2(nums []int) int {
	l := len(nums)
	if l == 0 {
		return 0
	} else if l == 1 {
		return nums[0]
	}
	dp := make([]int, l)
	dp[0] = nums[0]
	if nums[1] > nums[0] {
		dp[1] = nums[1]
	} else {
		dp[1] = nums[0]
	}
	for i := 2; i <= l-1; i++ {
		dp[i] = big198(dp[i-2]+nums[i], dp[i-1])
	}
	return dp[l-1]
}

func big198(x, y int) int {
	if x > y {
		return x
	}
	return y
}
