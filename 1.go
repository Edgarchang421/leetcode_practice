package main

import "fmt"

/*
資料無序，且一定會有解，所以使用map，有重複也可以直接回傳
nums := []int{3, 2, 4}
*/

// Runtime: 63 ms, faster than 8.28% of Go online submissions for Two Sum.
// Memory Usage: 3.5 MB, less than 94.20% of Go online submissions for Two Sum.
func twoSum(nums []int, target int) []int {
	for i, v := range nums {
		// restElements := make([]int, 0, len(nums)-1)
		// restElements = append(restElements, nums[:i]...)
		// restElements = append(restElements, nums[i+1:]...)
		for i2, v2 := range nums {
			if i2 == i {
				continue
			}
			if v+v2 == target {
				return []int{i, i2}
			}
		}
	}
	return nil
}

// 死在重複
func twoSumVer2(nums []int, target int) []int {
	m := make(map[int]int)
	for i, v := range nums {
		m[v] = i
	}

	fmt.Print(m)
	fmt.Print("\n")

	for i, v := range nums {
		ans := target - v
		fmt.Print(ans)
		fmt.Print("\n")

		// 先刪除，檢查完再加回來
		delete(m, v)

		if i2, ok := m[ans]; ok {
			return []int{i, i2}
		}

		m[v] = i
	}

	return nil
}

// Runtime: 11 ms, faster than 53.92% of Go online submissions for Two Sum.
// Memory Usage: 6.3 MB, less than 5.40% of Go online submissions for Two Sum.
func twoSumVer3(nums []int, target int) []int {
	m := make(map[int][]int)
	for i, v := range nums {
		m[v] = append(m[v], i)
	}

	for i, v := range nums {
		ans := target - v

		// 先看有沒有
		if i2, ok := m[ans]; ok {
			// 檢查重複, 且不使用自己
			if len(m[ans]) == 1 && i2[0] != i {
				return []int{i, i2[0]}
				// 有重複的
			} else if len(m[ans]) > 1 {
				for _, index := range m[ans] {
					// 不使用重複的
					if index == i {
						continue
					}
					return []int{i, index}
				}
			}
		}
	}

	return nil
}

// Runtime: 8 ms, faster than 69.83% of Go online submissions for Two Sum.
// Memory Usage: 4.2 MB, less than 53.68% of Go online submissions for Two Sum.
func twoSumVer4(nums []int, target int) []int {
	indexMap := make(map[int]int)
	for index, value := range nums {
		// 讀取過的就存起來
		// 往後的只要往前找是否有此資料即可，使用map無序查詢比較快
		if i, ok := indexMap[target-value]; ok {
			return []int{i, index}
		}
		indexMap[value] = index
	}
	return []int{}
}
