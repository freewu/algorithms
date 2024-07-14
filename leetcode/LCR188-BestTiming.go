package main

// LCR 188. 买卖芯片的最佳时机
// 数组 prices 记录了某芯片近期的交易价格，其中 prices[i] 表示的 i 天该芯片的价格。
// 你只能选择 某一天 买入芯片，并选择在 未来的某一个不同的日子 卖出该芯片。
// 请设计一个算法计算并返回你从这笔交易中能获取的最大利润。

// 如果你不能获取任何利润，返回 0。

// 示例 1：
// 输入：prices = [3, 6, 2, 9, 8, 5]
// 输出：7
// 解释：在第 3 天（芯片价格 = 2）买入，在第 4 天（芯片价格 = 9）卖出，最大利润 = 9 - 2 = 7。

// 示例 2：
// 输入：prices = [8, 12, 15, 7, 3, 10]
// 输出：7
// 解释：在第 5 天（芯片价格 = 3）买入，在第 6 天（芯片价格 = 10）卖出，最大利润 = 10 - 3 = 7。

// 提示：
//     0 <= prices.length <= 10^5
//     0 <= prices[i] <= 10^4

import "fmt"

// 模拟 DP
func bestTiming(prices []int) int {
    if len(prices) < 1 {
        return 0
    }
    mn, res := prices[0], 0 // 先把第一天设置为买入价格
    for i := 1; i < len(prices); i++ {
        if prices[i] - mn > res { // 如果当天的价格 - 买入价格 大于 最大利润   则设置新的利润差价
            res = prices[i] - mn
        }
        if prices[i] < mn { // 如果当天价格 小于 买入价格
            mn = prices[i] // 设置为使用 prices[i] 买入
        }
    }
    return res
}

// 单调栈
func bestTiming1(prices []int) int {
    if len(prices) == 0 {
        return 0
    }
    stack, res := []int{ prices[0] }, 0
    max := func (x, y int) int { if x > y { return x; }; return y; }
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
        res = max(res, stack[len(stack)-1] - stack[0])
    }
    return res
}

func bestTiming2(prices []int) int {
    if len(prices) == 1 {
        return 0
    }
    res, start := -1, 0
    for i := 1; i < len(prices); i++ {
        if prices[start] > prices[i] {
            start = i
        }
        delta := prices[i] - prices[start]
        if delta > res {
            res = delta
        }
    }
    return res
}

func main() {
    // Explanation: Buy on day 2 (price = 1) and sell on day 5 (price = 6), profit = 6-1 = 5.
    // Note that buying on day 2 and selling on day 1 is not allowed because you must buy before you sell.
    fmt.Printf("bestTiming([]int{7,1,5,3,6,4}) = %v\n",bestTiming([]int{7,1,5,3,6,4})) // 5    (6 - 1)
    // Explanation: In this case, no transactions are done and the max profit = 0.
    fmt.Printf("bestTiming([]int{7,6,4,3,1}) = %v\n",bestTiming([]int{7,6,4,3,1})) // 0   当天买入 当天卖出

    fmt.Printf("bestTiming1([]int{7,1,5,3,6,4}) = %v\n",bestTiming1([]int{7,1,5,3,6,4})) // 5    (6 - 1)
    fmt.Printf("bestTiming1([]int{7,6,4,3,1}) = %v\n",bestTiming1([]int{7,6,4,3,1})) // 0   当天买入 当天卖出

    fmt.Printf("bestTiming2([]int{7,1,5,3,6,4}) = %v\n",bestTiming2([]int{7,1,5,3,6,4})) // 5    (6 - 1)
    fmt.Printf("bestTiming2([]int{7,6,4,3,1}) = %v\n",bestTiming2([]int{7,6,4,3,1})) // 0   当天买入 当天卖出
}
