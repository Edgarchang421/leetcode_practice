package main

/*
Example 1:

Input
["NumArray", "sumRange", "sumRange", "sumRange"]
[[[-2, 0, 3, -5, 2, -1]], [0, 2], [2, 5], [0, 5]]
Output
[null, 1, -1, -3]

Explanation
NumArray numArray = new NumArray([-2, 0, 3, -5, 2, -1]);
numArray.sumRange(0, 2); // return (-2) + 0 + 3 = 1
numArray.sumRange(2, 5); // return 3 + (-5) + 2 + (-1) = -1
numArray.sumRange(0, 5); // return (-2) + 0 + 3 + (-5) + 2 + (-1) = -3


Constraints:

1 <= nums.length <= 104
-105 <= nums[i] <= 105
0 <= left <= right < nums.length
At most 104 calls will be made to sumRange.

	d := Constructor([]int{-2, 0, 3, -5, 2, -1})
	ans := d.SumRange(0, 5)
	fmt.Printf("ans: %v\n", ans)
*/

// Runtime: 38 ms, faster than 58.22% of Go online submissions for Range Sum Query - Immutable.
// Memory Usage: 8.8 MB, less than 79.81% of Go online submissions for Range Sum Query - Immutable.
type NumArray struct {
	DP   []int
	Nums []int
}

func Constructor(nums []int) NumArray {
	data := NumArray{Nums: nums}
	data.DP = make([]int, len(nums))
	data.DP[0] = nums[0]
	for i := 1; i < len(nums); i++ {
		data.DP[i] = data.DP[i-1] + nums[i]
	}
	return data
}

func (this *NumArray) SumRange(left int, right int) int {
	if left == 0 {
		return this.DP[right]
	}
	return this.DP[right] - this.DP[left-1]
}

// 只存結果
// Runtime: 51 ms, faster than 37.09% of Go online submissions for Range Sum Query - Immutable.
// Memory Usage: 8.4 MB, less than 92.02% of Go online submissions for Range Sum Query - Immutable.
type NumArray2 struct {
	Nums []int
}

func Constructor2(nums []int) NumArray2 {
	var data NumArray2
	data.Nums = make([]int, len(nums))
	data.Nums[0] = nums[0]
	for i := 1; i < len(nums); i++ {
		data.Nums[i] = data.Nums[i-1] + nums[i]
	}
	return data
}

func (this *NumArray2) SumRange(left int, right int) int {
	if left == 0 {
		return this.Nums[right]
	}
	return this.Nums[right] - this.Nums[left-1]
}
