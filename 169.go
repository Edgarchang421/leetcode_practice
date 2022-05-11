package main

/*
	r := majorityElement([]int{2, 2, 1, 1, 1, 2, 2})
	fmt.Printf("result: %v\n", r)
*/

// Runtime: 35 ms, faster than 17.00% of Go online submissions for Majority Element.
// Memory Usage: 6.3 MB, less than 15.15% of Go online submissions for Majority Element.
func majorityElement(nums []int) int {
	// value times
	m := make(map[int]int)
	l := len(nums)
	for i := range nums {
		if n, ok := m[nums[i]]; ok {
			if n+1 > l {
				return nums[i]
			}
			m[nums[i]] += 1
		} else {
			m[nums[i]] = 1
		}
	}
	max := m[nums[0]]
	key := nums[0]
	for k, v := range m {
		if v > max {
			max = v
			key = k
		}
	}
	return key
}

// Runtime: 8 ms, faster than 99.85% of Go online submissions for Majority Element.
// Memory Usage: 7.2 MB, less than 5.10% of Go online submissions for Majority Element.
func majorityElementV2(nums []int) int {
	res := nums[0]
	count := 0
	for i := range nums {
		if nums[i] == res {
			count++
		} else if count > 0 {
			count--
		} else {
			res = nums[i]
			count++
		}
	}

	return res
}
