package main

import (
	"fmt"
	"sort"
)

/*
Example 1:

Input: coins = [1,2,5], amount = 11
Output: 3
Explanation: 11 = 5 + 5 + 1
Example 2:

Input: coins = [2], amount = 3
Output: -1
Example 3:

Input: coins = [1], amount = 0
Output: 0

[186,419,83,408]
6249

Constraints:

1 <= coins.length <= 12
1 <= coins[i] <= 231 - 1
0 <= amount <= 104

	// r := coinChangeMyVer([]int{1, 2, 5}, 11)
	// fmt.Printf("result: %v\n", r)

	// r := coinChangeMyVer([]int{2}, 3)
	// fmt.Printf("result: %v\n", r)

	// r := coinChangeMyVer([]int{1}, 0)
	// fmt.Printf("result: %v\n", r)

	// r := coinChangeMyVer([]int{1, 2, 5, 10}, 27)
	// fmt.Printf("result: %v\n", r)

	// r = coinChangeMyVer([]int{2, 5, 10, 1}, 27)
	// fmt.Printf("result: %v\n", r)

	r := coinChangeMyVer([]int{186, 419, 83, 408}, 6249)
	fmt.Printf("result: %v\n", r)

	// r := Fibonacci(20)
	// fmt.Printf("result: %v\n", r)
*/

func coinChange(coins []int, amount int) int {
	// 初始化一個每一個amount所需的硬幣數
	INF := 65535

	// 用來記錄每一個amount的slice的長度，加上0
	size := amount + 1

	// slice紀錄每一種amount所需要的零錢數
	change := make([]int, size)

	// 初始化所有amount都需要一個前面設定的數字的硬幣數
	for i := range change {
		change[i] = INF
	}

	// Initialization for $0
	// 零元需要零個
	change[0] = 0

	// 從最小的amount開始，計算每個amount需要多少硬幣數
	for value := 1; value <= amount; value += 1 {
		// 可以使用的面額
		fmt.Printf("\n")
		fmt.Printf("\n")
		fmt.Printf("-----*****-----\n")
		fmt.Printf("amount: %v\n", value)
		fmt.Printf("-----*****-----\n")

		for _, coin := range coins {
			// 面額大於amount則不使用
			if coin > value {
				// coin value is to big, can not make change with current coin
				continue
			}

			// 最初始的硬幣數為前面設定很大的數字65535
			// 用原本的數字65535，和前面已經算出來過可以使用的(前面小的先算出來)比較
			// +1為加上此時的硬幣面額
			// update dp_table, try to make change with coin
			fmt.Printf("value-coin: %v\n", value-coin)
			fmt.Printf("value: %v, coin: %v, change[value-coin]+1: %v, change[value]: %v\n", value, coin, change[value-coin]+1, change[value])
			change[value] = min(change[value-coin]+1, change[value])
		}
	}

	if change[amount] == INF {
		// Reject, no solution
		return -1

	} else {

		// Accept, return total count of coin change
		return change[amount]
	}
}

// Runtime: 6 ms, faster than 94.34% of Go online submissions for Coin Change.
// Memory Usage: 6.4 MB, less than 67.68% of Go online submissions for Coin Change.
func coinChangeExample(coins []int, amount int) int {
	dp := make([]int, amount+1)

	for i := 1; i < amount+1; i++ {
		dp[i] = -1
	}

	for _, c := range coins {
		fmt.Printf("coin: %v\n", c)
		for i := 1; i < amount+1; i++ {
			fmt.Printf("amount: %v\n", i)
			// 硬幣面額大於目標值，或者此amount在前面的計算中沒有答案
			if i-c < 0 || dp[i-c] < 0 {
				continue
			}
			// 前面已經算出過解答
			if dp[i] > -1 {
				// 比較哪個使用的硬幣比較少
				dp[i] = min(dp[i], dp[i-c]+1)
				// fmt.Printf("min(dp[i], dp[i-c]+1)\n")
				// fmt.Printf("%v %v\n\n", dp[i], dp[i-c]+1)
			} else { // 前面沒有解答
				dp[i] = dp[i-c] + 1
				fmt.Printf("dp[i] = dp[i-c] + 1\n")
				fmt.Printf("%v = %v + 1\n\n", dp[i], dp[i-c])
			}
		}
	}
	return dp[amount]
}

func min322(x, y int) int {
	if x < y {
		return x
	}
	return y
}

// error
// 最大的可以使用，但是用小的才能除盡
func coinChangeVer0(coins []int, amount int) int {
	if amount == 0 {
		return 0
	}

	sort.Slice(coins, func(i, j int) bool { return coins[i] < coins[j] })

	result := 0
	for i := len(coins) - 1; i >= 0; i-- {
		if amount/coins[i] > 0 {
			result += amount / coins[i]
			amount -= coins[i] * (amount / coins[i])
			fmt.Printf("coin %v use %v, amount: %v\n", coins[i], result, amount)
		}
	}

	if amount != 0 {
		return -1
	}
	return result
}

func Fibonacci(index int) []int {
	// 0 1 1 2 3 5 8 13 21 34 55 89 144...
	result := make([]int, 0, index)
	for i := 0; i <= index; i++ {
		switch {
		case i == 0:
			result = append(result, 0)
		case i == 1:
			result = append(result, 1)
		default:
			result = append(result, result[i-1]+result[i-2])
		}
	}
	return result
}

// Runtime: 12 ms, faster than 77.72% of Go online submissions for Coin Change.
// Memory Usage: 6.3 MB, less than 92.80% of Go online submissions for Coin Change.
func coinChangeMyVer(coins []int, amount int) int {
	dp := make([]int, amount+1)
	for i := range dp {
		dp[i] = -1
	}

	dp[0] = 0

	for amount := range dp {
		for _, coin := range coins {
			// 面額超過目標價 or 此目標價減掉這個面額是沒有解的 就跳過
			if amount-coin < 0 || dp[amount-coin] == -1 {
				// fmt.Printf("amount: %v, coin: %v, continue\n", amount, coin)
				continue
			}

			// 如果已經有解了，就比較使用此面額的硬幣數(減掉使用的面額，加上這個面額的一個硬幣)，以及原本的硬幣數，取比較小的
			if dp[amount] > -1 {
				dp[amount] = min(dp[amount-coin]+1, dp[amount])
			} else {
				// 此目標價沒有解，但是此目標價減掉這個硬幣面額是有解的，就使用那個解加上這個硬幣面額*1
				dp[amount] = dp[amount-coin] + 1
			}
			// fmt.Printf("dp[amount]: %v, amount: %v, coin: %v\n", dp[amount], amount, coin)
		}
	}

	return dp[amount]
}

func coinChangeMyVerClean(coins []int, amount int) int {
	dp := make([]int, amount+1)
	for i := range dp {
		dp[i] = -1
	}
	dp[0] = 0

	for amount := range dp {
		for _, coin := range coins {
			if amount-coin < 0 || dp[amount-coin] == -1 {
				continue
			}

			if dp[amount] > -1 {
				dp[amount] = min(dp[amount-coin]+1, dp[amount])
			} else {
				dp[amount] = dp[amount-coin] + 1
			}
		}
	}

	return dp[amount]
}
