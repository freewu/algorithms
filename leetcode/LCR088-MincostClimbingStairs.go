package main

// LCR 088. 使用最小花费爬楼梯
// 数组的每个下标作为一个阶梯，第 i 个阶梯对应着一个非负数的体力花费值 cost[i]（下标从 0 开始）。
// 每当爬上一个阶梯都要花费对应的体力值，一旦支付了相应的体力值，就可以选择向上爬一个阶梯或者爬两个阶梯。
// 请找出达到楼层顶部的最低花费。在开始时，你可以选择从下标为 0 或 1 的元素作为初始阶梯。

// 示例 1：
// 输入：cost = [10, 15, 20]
// 输出：15
// 解释：最低花费是从 cost[1] 开始，然后走两步即可到阶梯顶，一共花费 15 。

// 示例 2：
// 输入：cost = [1, 100, 1, 1, 1, 100, 1, 1, 100, 1]
// 输出：6
// 解释：最低花费方式是从 cost[0] 开始，逐个经过那些 1 ，跳过 cost[3] ，一共花费 6 。

// 提示：
//     2 <= cost.length <= 1000
//     0 <= cost[i] <= 999

import "fmt"

func minCostClimbingStairs(cost []int) int {
    pay := make([]int, len(cost))
    pay[0], pay[1] = cost[0], cost[1]
    min := func (x, y int) int { if x < y { return x; }; return y; }
    // 计算每一步需要的累计费用
    for i := 2 ; i < len(pay); i++ {
        // 走两步 pay[i-2] 走一步 pay[i-1]
        pay[i] = min(pay[i-2], pay[i-1]) + cost[i]
    }
    return min(pay[len(pay)-2], pay[len(pay)-1])
}

// dfs
func minCostClimbingStairs1(cost []int) int {
    n := len(cost)
    dp := make([]int, n + 1)
    for i := 0; i < n+1; i += 1 {
        dp[i] = -1
    }
    min := func (x, y int) int { if x < y { return x; }; return y; }
    var dfs func(int) int
    dfs = func (i int) int {
        if i < 0 {
            return 0
        } else if i == n {
            return min(dfs(i-1), dfs(i-2))
        }
        res := &dp[i]
        if *res != -1 {
            return dp[i]
        }
        *res = min(dfs(i-1), dfs(i-2)) + cost[i]
        return *res
    }
    return dfs(n)
}

func minCostClimbingStairs2(cost []int) int {
    // 只用三位来保存数据
    dp := make([]int,3)
    min := func (x, y int) int { if x < y { return x; }; return y; }
    for i := 2; i <= len(cost); i++ {
        // 走一步的费用: dp[1] + cost[i-1] ,  走两步的费用:  dp[2]+ cost[i-2]
        dp[0] = min(dp[1] + cost[i-1], dp[2]+ cost[i-2])
        dp[2] = dp[1]
        dp[1] = dp[0]
    }
    return dp[1]
}

func minCostClimbingStairs3(cost []int) int {
    n := len(cost)
    dp := make([]int,  n + 1)
    dp[0], dp[1] = 0, 0
    min := func (x, y int) int { if x < y { return x; }; return y; }
    for i := 2; i <= n; i++ {
        dp[i] = min(dp[i-1] + cost[i-1], dp[i-2] + cost[i-2])
    }
    return dp[n]
}

func main() {
    // Explanation: You will start at index 1.
    // - Pay 15 and climb two steps to reach the top.
    // The total cost is 15.
    fmt.Println(minCostClimbingStairs([]int{10,15,20})) // 15
    // Explanation: You will start at index 0.
    // - Pay 1 and climb two steps to reach index 2.
    // - Pay 1 and climb two steps to reach index 4.
    // - Pay 1 and climb two steps to reach index 6.
    // - Pay 1 and climb one step to reach index 7.
    // - Pay 1 and climb two steps to reach index 9.
    // - Pay 1 and climb one step to reach the top.
    // The total cost is 6.
    fmt.Println(minCostClimbingStairs([]int{1,100,1,1,1,100,1,1,100,1})) // 6

    fmt.Println(minCostClimbingStairs1([]int{10,15,20})) // 15
    fmt.Println(minCostClimbingStairs1([]int{1,100,1,1,1,100,1,1,100,1})) // 6

    fmt.Println(minCostClimbingStairs2([]int{10,15,20})) // 15
    fmt.Println(minCostClimbingStairs2([]int{1,100,1,1,1,100,1,1,100,1})) // 6

    fmt.Println(minCostClimbingStairs3([]int{10,15,20})) // 15
    fmt.Println(minCostClimbingStairs3([]int{1,100,1,1,1,100,1,1,100,1})) // 6
}