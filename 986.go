package main

/*
Example 1:
Input: firstList = [[0,2],[5,10],[13,23],[24,25]], secondList = [[1,5],[8,12],[15,24],[25,26]]
Output: [[1,2],[5,5],[8,10],[15,23],[24,24],[25,25]]

Example 2:
Input: firstList = [[1,3],[5,9]], secondList = []
Output: []

firstList = [[3,5],[9,20]]
secondList = [[4,5],[7,10],[11,12],[14,15],[16,20]]
Expected nas: [[4,5],[9,10],[11,12],[14,15],[16,20]]

Constraints:

0 <= firstList.length, secondList.length <= 1000
firstList.length + secondList.length >= 1
0 <= starti < endi <= 109
endi < starti+1
0 <= startj < endj <= 109
endj < startj+1

	fmt.Printf("ans [[1,2],[5,5],[8,10],[15,23],[24,24],[25,25]]:\n    %v\n",
		intervalIntersection(
			[][]int{{0, 2}, {5, 10}, {13, 23}, {24, 25}},
			[][]int{{1, 5}, {8, 12}, {15, 24}, {25, 26}},
		))
	// fmt.Printf("ans [[4,5],[9,10],[11,12],[14,15],[16,20]]:\n    %v\n",
	// 	intervalIntersection(
	// 		[][]int{{3, 5}, {9, 20}},
	// 		[][]int{{4, 5}, {7, 10}, {11, 12}, {14, 15}, {16, 20}},
	// 	))
*/

// Runtime: 16 ms, faster than 92.06% of Go online submissions for Interval List Intersections.
// Memory Usage: 6.8 MB, less than 92.52% of Go online submissions for Interval List Intersections.
func intervalIntersection(firstList [][]int, secondList [][]int) [][]int {
	result := make([][]int, 0)
	var i, j int
	for i < len(firstList) && j < len(secondList) {
		interval := []int{
			max(secondList[j][0], firstList[i][0]),
			min(secondList[j][1], firstList[i][1]),
		}
		if interval[0] <= interval[1] {
			result = append(result, interval)
		}
		if firstList[i][1] < secondList[j][1] {
			i++
		} else {
			j++
		}
	}
	return result
}
