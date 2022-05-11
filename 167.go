package main

/*
Input: numbers = [2,7,11,15], target = 9
Output: [1,2]

Input: numbers = [2,3,4], target = 6
Output: [1,3]

Input: numbers = [-1,0], target = -1
Output: [1,2]

	r := twoSum([]int{-1, 0}, -1)
	fmt.Printf("result: %v\n", r)
*/

// Runtime: 17 ms, faster than 45.40% of Go online submissions for Two Sum II - Input Array Is Sorted.
// Memory Usage: 5.5 MB, less than 51.06% of Go online submissions for Two Sum II - Input Array Is Sorted.
func twoSum167(numbers []int, target int) []int {
	left := 0
	right := len(numbers) - 1
	for right > left {
		switch {
		case numbers[left]+numbers[right] > target:
			right -= 1
		case numbers[left]+numbers[right] < target:
			left += 1
		case numbers[left]+numbers[right] == target:
			return []int{left + 1, right + 1}
		}
	}
	return nil
}

// Runtime: 14 ms, faster than 62.31% of Go online submissions for Two Sum II - Input Array Is Sorted.
// Memory Usage: 5.4 MB, less than 92.79% of Go online submissions for Two Sum II - Input Array Is Sorted.
func twoSum167Ver2(numbers []int, target int) []int {
	left := 0
	right := len(numbers) - 1
	for right > left {
		ans := numbers[left] + numbers[right]
		switch {
		case ans > target:
			right -= 1
		case ans < target:
			left += 1
		case ans == target:
			return []int{left + 1, right + 1}
		}
	}
	return nil
}

// Runtime: 16 ms, faster than 51.34% of Go online submissions for Two Sum II - Input Array Is Sorted.
// Memory Usage: 5.4 MB, less than 92.79% of Go online submissions for Two Sum II - Input Array Is Sorted.
func twoSum167Ver3(numbers []int, target int) []int {
	left := 0
	right := len(numbers) - 1
	for {
		ans := numbers[left] + numbers[right]
		switch {
		case ans > target:
			right -= 1
		case ans < target:
			left += 1
		case ans == target:
			return []int{left + 1, right + 1}
		}
	}
}

// Runtime: 13 ms, faster than 65.77% of Go online submissions for Two Sum II - Input Array Is Sorted.
// Memory Usage: 5.4 MB, less than 51.06% of Go online submissions for Two Sum II - Input Array Is Sorted.
func twoSum167Ver4(numbers []int, target int) []int {
	left := 0
	right := len(numbers) - 1
	ans := numbers[left] + numbers[right]
	for ans != target {
		if ans > target {
			right -= 1
		} else {
			left += 1
		}
		ans = numbers[left] + numbers[right]
	}
	return []int{left + 1, right + 1}
}

// Runtime: 12 ms, faster than 73.76% of Go online submissions for Two Sum II - Input Array Is Sorted.
// Memory Usage: 5.4 MB, less than 51.06% of Go online submissions for Two Sum II - Input Array Is Sorted.
func twoSumVer5(numbers []int, target int) []int {
	left := 0
	right := len(numbers) - 1

	for ans := numbers[left] + numbers[right]; ans != target; {
		if ans > target {
			right--
		} else {
			left++
		}
		ans = numbers[left] + numbers[right]
	}
	return []int{left + 1, right + 1}
}
