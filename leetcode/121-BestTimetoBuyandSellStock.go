package main

import "fmt"

/**
121. Best Time to Buy and Sell Stock
You are given an array prices where prices[i] is the price of a given stock on the ith day.
You want to maximize your profit by choosing a single day to buy one stock and choosing a different day in the future to sell that stock.
Return the maximum profit you can achieve from this transaction. If you cannot achieve any profit, return 0.

Constraints:

	1 <= prices.length <= 10^5
	0 <= prices[i] <= 10^4

Example 1:

	Input: prices = [7,1,5,3,6,4]
	Output: 5
	Explanation: Buy on day 2 (price = 1) and sell on day 5 (price = 6), profit = 6-1 = 5.
	Note that buying on day 2 and selling on day 1 is not allowed because you must buy before you sell.

Example 2:

	Input: prices = [7,6,4,3,1]
	Output: 0
	Explanation: In this case, no transactions are done and the max profit = 0.

# 解题思路
	找出股票中能赚的钱最多的差价
	DP

 */

// 解法一 模拟 DP
func maxProfit(prices []int) int {
	if len(prices) < 1 {
		return 0
	}
	min, maxProfit := prices[0], 0 // 先把第一天设置为买入价格
	for i := 1; i < len(prices); i++ {
		if prices[i] - min > maxProfit { // 如果当天的价格 - 买入价格 大于 最大利润   则设置新的利润差价
			maxProfit = prices[i] - min
		}
		if prices[i] < min { // 如果当天价格 小于 买入价格
			min = prices[i] // 设置为使用 prices[i] 买入
		}
	}
	return maxProfit
}

// 解法二 单调栈
func maxProfit1(prices []int) int {
	if len(prices) == 0 {
		return 0
	}
	stack, res := []int{ prices[0] }, 0
	for i := 1; i < len(prices); i++ {
		if prices[i] > stack[len(stack)-1] {
			stack = append(stack, prices[i])
		} else {
			index := len(stack) - 1
			for ; index >= 0; index-- {
				if stack[index] < prices[i] {
					break
				}
			}
			stack = stack[:index+1]
			stack = append(stack, prices[i])
		}
		res = max(res, stack[len(stack)-1]-stack[0])
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// best solution
func maxProfitBest(prices []int) int {
	if len(prices) == 1 {
		return 0
	}
	max := -1
	start := 0
	for i := 1; i < len(prices); i++ {
		if prices[start] > prices[i] {
			start = i
		}
		delta := prices[i] - prices[start]
		if delta > max {
			max = delta
		}
	}
	return max
}

func main() {
	fmt.Printf("maxProfit([]int{7,1,5,3,6,4}) = %v\n",maxProfit([]int{7,1,5,3,6,4})) // 5    (6 - 1)
	fmt.Printf("maxProfit([]int{7,6,4,3,1}) = %v\n",maxProfit([]int{7,6,4,3,1})) // 0   当天买入 当天卖出

	fmt.Printf("maxProfit1([]int{7,1,5,3,6,4}) = %v\n",maxProfit1([]int{7,1,5,3,6,4})) // 5    (6 - 1)
	fmt.Printf("maxProfit1([]int{7,6,4,3,1}) = %v\n",maxProfit1([]int{7,6,4,3,1})) // 0   当天买入 当天卖出

	fmt.Printf("maxProfitBest([]int{7,1,5,3,6,4}) = %v\n",maxProfitBest([]int{7,1,5,3,6,4})) // 5    (6 - 1)
	fmt.Printf("maxProfitBest([]int{7,6,4,3,1}) = %v\n",maxProfitBest([]int{7,6,4,3,1})) // 0   当天买入 当天卖出
}
