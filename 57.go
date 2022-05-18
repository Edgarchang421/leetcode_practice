package main

/*
Example 1:

Input: intervals = [[1,3],[6,9]], newInterval = [2,5]
Output: [[1,5],[6,9]]

{1 2 3} {6 7 8 9}是互相連接，3456之間沒有連接
新的interval [2,5] 2345之間可以連結，所以 1 原本可以連結到 2 ，2 又可以連結到 5
2在1 3之間， 3在2 5之間，所以可以跳躍間隔
1. 先找到沒有被 newInterval 覆蓋的區間
2. 在被 newInterval 覆蓋的區間，找到最小和最大的start、end
3. 加入剩餘的區間

Example 2:

Input: intervals = [[1,2],[3,5],[6,7],[8,10],[12,16]], newInterval = [4,8]
Output: [[1,2],[3,10],[12,16]]
Explanation: Because the new interval [4,8] overlaps with [3,5],[6,7],[8,10].


Constraints:

0 <= intervals.length <= 104
intervals[i].length == 2
0 <= starti <= endi <= 105
intervals is sorted by starti in ascending order.
newInterval.length == 2
0 <= start <= end <= 105

	fmt.Printf("ans: %v\n", insert([][]int{{1, 3}, {6, 9}}, []int{2, 5}))
	fmt.Printf("ans: %v\n", insert([][]int{{1, 2}, {3, 5}, {6, 7}, {8, 10}, {12, 16}}, []int{4, 8}))
*/

// 1. 先找到沒有被 newInterval 覆蓋的區間
// 2. 在原有 intervals 被 newInterval 覆蓋的區間，找到最小和最大的start、end
// 3. 加入剩餘的區間
// Runtime: 9 ms, faster than 56.00% of Go online submissions for Insert Interval.
// Memory Usage: 4.7 MB, less than 38.33% of Go online submissions for Insert Interval.
func insert(intervals [][]int, newInterval []int) [][]int {
	result := make([][]int, 0)

	// i 只宣告一次
	var i int
	// 先找到連接到 新間隔 之前的位置
	for ; i < len(intervals) && intervals[i][1] < newInterval[0]; i++ {
		result = append(result, intervals[i])
	}

	// 最後 新間隔 會結束於 input newInterval 的 end 已經小於 原有interval的start
	// 因為上面迴圈的 intervals[i][1] < newInterval[0] 判斷式，所以現在 intervals[i][1] > newInterval[0]，因此要比較 newInterval[0] 和 intervals[i][0]找到新間隔的start
	for ; i < len(intervals) && intervals[i][0] <= newInterval[1]; i++ {
		newInterval[0] = min(intervals[i][0], newInterval[0])
		newInterval[1] = max(intervals[i][1], newInterval[1])
	}

	result = append(result, newInterval)

	for ; i < len(intervals); i++ {
		result = append(result, intervals[i])
	}

	return result
}
