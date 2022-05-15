package main

/*
Example 1:

Input: nums = [1,2,3]
Output: [[1,2,3],[1,3,2],[2,1,3],[2,3,1],[3,1,2],[3,2,1]]
Example 2:

Input: nums = [0,1]
Output: [[0,1],[1,0]]
Example 3:

Input: nums = [1]
Output: [[1]]

	fmt.Printf("ans: %v\n", permute([]int{1, 2, 3}))
*/

// 選擇 含有所有element的array
// 限制 排序要不相同
// 目標 所有不同排列的array
// Runtime: 3 ms, faster than 45.79% of Go online submissions for Permutations.
// Memory Usage: 3 MB, less than 20.15% of Go online submissions for Permutations.
func permute(nums []int) [][]int {
	subArray := make([]int, 0)
	result := make([][]int, 0)
	backtracking(&result, subArray, nums)
	return result
}
func backtracking(result *[][]int, subArray []int, nums []int) {
	if len(nums) == 0 {
		*result = append(*result, subArray)
		return
	}
	for i, v := range nums {
		newSubarray := copySlice(subArray)
		newSubarray = append(newSubarray, v)
		newNums := copySlice((nums)[:i])
		newNums = append(newNums, (nums)[i+1:]...)
		backtracking(result, newSubarray, newNums)
	}
}
func copySlice(a []int) []int {
	result := make([]int, len(a))
	for i := range a {
		result[i] = a[i]
	}
	return result
}
