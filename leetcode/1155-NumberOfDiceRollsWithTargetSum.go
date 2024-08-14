package main

// 1155. Number of Dice Rolls With Target Sum
// You have n dice, and each dice has k faces numbered from 1 to k.

// Given three integers n, k, and target, 
// return the number of possible ways (out of the kn total ways) to roll the dice, 
// so the sum of the face-up numbers equals target. 
// Since the answer may be too large, return it modulo 10^9 + 7.

// Example 1:
// Input: n = 1, k = 6, target = 3
// Output: 1
// Explanation: You throw one die with 6 faces.
// There is only one way to get a sum of 3.

// Example 2:
// Input: n = 2, k = 6, target = 7
// Output: 6
// Explanation: You throw two dice, each with 6 faces.
// There are 6 ways to get a sum of 7: 1+6, 2+5, 3+4, 4+3, 5+2, 6+1.

// Example 3:
// Input: n = 30, k = 30, target = 500
// Output: 222616187
// Explanation: The answer must be returned modulo 10^9 + 7.

// Constraints:
//     1 <= n, k <= 30
//     1 <= target <= 1000

import "fmt"

func numRollsToTarget(n int, k int, target int) int {
    mod := 1_000_000_007
    cache := make(map[[2]int]int)
    var dfs func(n, target int) int
    dfs = func(n, target int) int {
        if n == 0 {
            if target == 0 {
                return 1 
            } else {
                return 0
            }
        }
        if v, ok := cache[[2]int{ n, target }]; ok {
            return v
        }
        res := 0
        for v := 1; v <= k; v++ {
            if target >= v {
                res = (res + dfs(n - 1, target - v)) % mod
            }
        }
        cache[[2]int{ n, target }] = res
        return res
    }
    return dfs(n, target)
}

func numRollsToTarget1(n int, k int, target int) int {
    // 多重背包的视角
    // 有n个物品, 每个物品体积都是1, 但有1->k个,至少选一个,至多选k个, 求体积和为target的方法数
    // trick!! 至少选一个不方便转移, n个骰子至少掷出1, 可以变为掷出0->k-1, 这样 掷出的目标和就是 target-n了,更符合多重背包的定义
    mod := 1_000_000_007
    if target < n {
        return 0
    }
    target -= n //  trick!! 目标是target-n,这样每个物品可以选或者不选,,更符合多重背包定义
    sum, dp := 0, make([]int, target + 1)
    dp[0] = 1
    min := func (x, y int) int { if x < y { return x; }; return y; }
    for i := 1; i <= n; i++ { // 枚举n个物品
        w, cnt := 1, k-1
        sum += w * cnt // trick!! 因为每个物品都一个类型, 可以直接计算出sum= i * cnt
        limit := min(sum, target)
        for j := w; j <= limit; j++ { // 利用同余前缀和对多维背包优化, 不同的是至少选一个
            dp[j] = (dp[j] + dp[j-w]) % mod
        }
        for j := limit; j >= (cnt+1) * w; j-- {
            dp[j] = (dp[j] - dp[j-(cnt+1) * w]) % mod // mod 后相减,可能为负数
        }
    }
    return (dp[target] % mod + mod) % mod
}

func main() {
    // Example 1:
    // Input: n = 1, k = 6, target = 3
    // Output: 1
    // Explanation: You throw one die with 6 faces.
    // There is only one way to get a sum of 3.
    fmt.Println(numRollsToTarget(1,6,3)) // 1
    // Example 2:
    // Input: n = 2, k = 6, target = 7
    // Output: 6
    // Explanation: You throw two dice, each with 6 faces.
    // There are 6 ways to get a sum of 7: 1+6, 2+5, 3+4, 4+3, 5+2, 6+1.
    fmt.Println(numRollsToTarget(2,6,7)) // 6
    // Example 3:
    // Input: n = 30, k = 30, target = 500
    // Output: 222616187
    // Explanation: The answer must be returned modulo 10^9 + 7.
    fmt.Println(numRollsToTarget(30,30,500)) // 222616187

    fmt.Println(numRollsToTarget1(1,6,3)) // 1
    fmt.Println(numRollsToTarget1(2,6,7)) // 6
    fmt.Println(numRollsToTarget1(30,30,500)) // 222616187
}