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

Input: nums = [1,1,2]
Output:
[[1,1,2],
 [1,2,1],
 [2,1,1]]

 3 3 0 3
 [[0,3,3,3],[3,0,3,3],[3,3,0,3],[3,3,3,0]]
*/

// 選擇 含有所有element的array
// 限制 排序要不相同，因為會有相同數字，所以同一個位置，已經使用過的數字用set(無序不重複)儲存
// 目標 所有不同排列的array
// Runtime: 7 ms, faster than 49.00% of Go online submissions for Permutations II.
// Memory Usage: 5.5 MB, less than 26.77% of Go online submissions for Permutations II.
func permute47(nums []int) [][]int {
	subArray := make([]int, 0)
	result := make([][]int, 0)
	backtracking47(&result, subArray, nums)
	return result
}
func backtracking47(result *[][]int, subArray []int, nums []int) {
	if len(nums) == 0 {
		*result = append(*result, subArray)
		return
	}
	used := make(map[int]struct{})
	for i, v := range nums {
		if _, ok := used[v]; !ok {
			used[v] = struct{}{}
		} else {
			continue
		}
		newSubarray := copySlice(subArray)
		newSubarray = append(newSubarray, v)
		newNums := copySlice((nums)[:i])
		newNums = append(newNums, (nums)[i+1:]...)
		backtracking47(result, newSubarray, newNums)
	}
}
