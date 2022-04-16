package main

import "fmt"

/**
122. Best Time to Buy and Sell Stock II
You are given an integer array prices where prices[i] is the price of a given stock on the ith day.
On each day, you may decide to buy and/or sell the stock. You can only hold at most one share of the stock at any time.
However, you can buy it then immediately sell it on the same day.
Find and return the maximum profit you can achieve.

Constraints:

	1 <= prices.length <= 3 * 10^4
	0 <= prices[i] <= 10^4

Example 1:

	Input: prices = [7,1,5,3,6,4]
	Output: 7
	Explanation: Buy on day 2 (price = 1) and sell on day 3 (price = 5), profit = 5-1 = 4.
	Then buy on day 4 (price = 3) and sell on day 5 (price = 6), profit = 6-3 = 3.
	Total profit is 4 + 3 = 7.

Example 2:

	Input: prices = [1,2,3,4,5]
	Output: 4
	Explanation: Buy on day 1 (price = 1) and sell on day 5 (price = 5), profit = 5-1 = 4.
	Total profit is 4.

Example 3:

	Input: prices = [7,6,4,3,1]
	Output: 0
	Explanation: There is no way to make a positive profit, so we never buy the stock to achieve the maximum profit of 0.

# 解题思路
	要求输出最大收益，这一题不止买卖一次，可以买卖多次，买卖不能在同一天内操作
	必然是每次跌了就买入，涨到顶峰的时候就抛出。
	只要有涨峰就开始计算赚的钱，连续涨可以用两两相减累加来计算，两两相减累加，相当于涨到波峰的最大值减去谷底的值。
 */

func maxProfit(prices []int) int {
	profit := 0
	for i := 0; i < len(prices)-1; i++ {
		if prices[i+1] > prices[i] { // 如果第二天涨，今天就买入 计算出最高收益  如果跌(为负) 就不动
			profit += prices[i+1] - prices[i]
		}
	}
	return profit
}

// best solution
func maxProfitBest(prices []int) int {
	if len(prices) < 2 {
		return 0
	}
	profit:=0
	j:=1
	for  i := j; i < len(prices); i++ {
		if prices[i] > prices[i-1] {
			profit += prices[i] - prices[i-1]
			j += 2
		}
	}
	return profit
}

func main() {
	fmt.Printf("maxProfit([]int{7,1,5,3,6,4}) = %v\n",maxProfit([]int{7,1,5,3,6,4})) // 7  (5 - 1) + (6 - 3)
	fmt.Printf("maxProfit([]int{1,2,3,4,5}) = %v\n",maxProfit([]int{1,2,3,4,5})) // 4  (2-1) + (3-2) + (4-3) + (5-4)
	fmt.Printf("maxProfit([]int{7,6,4,3,1}) = %v\n",maxProfit([]int{7,6,4,3,1})) // 0

	fmt.Printf("maxProfitBest([]int{7,1,5,3,6,4}) = %v\n",maxProfitBest([]int{7,1,5,3,6,4})) // 7  (5 - 1) + (6 - 3)
	fmt.Printf("maxProfitBest([]int{1,2,3,4,5}) = %v\n",maxProfitBest([]int{1,2,3,4,5})) // 4  (2-1) + (3-2) + (4-3) + (5-4)
	fmt.Printf("maxProfitBest([]int{7,6,4,3,1}) = %v\n",maxProfitBest([]int{7,6,4,3,1})) // 0
}