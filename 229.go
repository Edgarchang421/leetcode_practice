package main

/*
	r := majorityElement(nil)
	fmt.Printf("result: %v\n", r)
*/

// Runtime: 13 ms, faster than 55.13% of Go online submissions for Majority Element II.
// Memory Usage: 4.9 MB, less than 100.00% of Go online submissions for Majority Element II.
func majorityElement229(nums []int) []int {
	standard := len(nums) / 3
	n1, n2, c1, c2 := 0, 0, 0, 0
	for i := range nums {
		num := nums[i]
		switch {
		case num == n1:
			c1++
		case num == n2:
			c2++
		case c1 == 0:
			n1 = num
			c1 = 1
		case c2 == 0:
			n2 = num
			c2 = 1
		default:
			c1--
			c2--
		}
	}
	c1, c2 = 0, 0
	for i := range nums {
		num := nums[i]
		if num == n1 {
			c1++
		} else if num == n2 {
			c2++
		}
	}

	ans := make([]int, 0)
	if c1 > standard {
		ans = append(ans, n1)
	}
	if c2 > standard {
		ans = append(ans, n2)
	}

	return ans
}
