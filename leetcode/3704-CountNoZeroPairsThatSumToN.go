package main

// 3704. Count No-Zero Pairs That Sum to N
// A no-zero integer is a positive integer that does not contain the digit 0 in its decimal representation.

// Given an integer n, count the number of pairs (a, b) where:
//     1. a and b are no-zero integers.
//     2. a + b = n

// Return an integer denoting the number of such pairs.

// Example 1:
// Input: n = 2
// Output: 1
// Explanation:
// The only pair is (1, 1).

// Example 2:
// Input: n = 3
// Output: 2
// Explanation:
// The pairs are (1, 2) and (2, 1).

// Example 3:
// Input: n = 11
// Output: 8
// Explanation:
// The pairs are (2, 9), (3, 8), (4, 7), (5, 6), (6, 5), (7, 4), (8, 3), and (9, 2). 
// Note that (1, 10) and (10, 1) do not satisfy the conditions because 10 contains 0 in its decimal representation.

// Constraints:
//     2 <= n <= 10^15

import "fmt"
import "strconv"

func countNoZeroPairs(n int64) int64 {
    s := strconv.FormatInt(n, 10)
    m := len(s)
    memo := make([][2][2]int, m)
    for i := range memo {
        memo[i] = [2][2]int{{-1, -1}, {-1, -1}} // -1 表示没有计算过
    }
    // 返回两个 1~9 的整数和为 target 的方案数
    twoSumWays := func(target int) int { return max(min(target-1, 19 - target), 0) } // 保证结果非负
    // 返回 d 是否为负数
    isNeg := func(d int) int {
        if d < 0 { return 1 }
        return 0
    }
    // borrow = 1 表示被低位（i+1）借了个 1
    // isNum = 1 表示之前填的数位，两个数都无前导零
    var dfs func(i, borrowed, isNum int) int
    dfs = func(i, borrowed, isNum int) int {
        res := 0
        if i < 0 { return borrowed ^ 1 } // borrowed 必须为 0
        p := &memo[i][borrowed][isNum]
        if *p >= 0 { return *p } // 之前计算过
        defer func() { *p = res }() // 记忆化
        d := int(s[i]-'0') - borrowed
        // 情况一：两个数位都不为 0
        if isNum > 0 {
            res =  dfs(i - 1, 0, 1) * twoSumWays(d)     // 不向高位借 1
            res += dfs(i - 1, 1, 1) * twoSumWays(d+10)  // 向高位借 1
        }
        // 情况二：其中一个数位填前导零
        if i < m - 1 { // 不能是最低位
            if d != 0 {
                // 如果 d < 0，必须向高位借 1
                // 如果 isNum = 1，根据对称性，方案数要乘以 2
                res += dfs(i-1, isNeg(d), 0) * (isNum + 1)
            } else if i == 0 { // 两个数位都填 0，只有当 i = 0 的时候才有解
                res++
            }
        }
        return res
    }
    return int64(dfs(m - 1, 0, 1))
}

func main() {
    // Example 1:
    // Input: n = 2
    // Output: 1
    // Explanation:
    // The only pair is (1, 1).
    fmt.Println(countNoZeroPairs(2)) // 1
    // Example 2:
    // Input: n = 3
    // Output: 2
    // Explanation:
    // The pairs are (1, 2) and (2, 1).
    fmt.Println(countNoZeroPairs(3)) // 2
    // Example 3:
    // Input: n = 11
    // Output: 8
    // Explanation:
    // The pairs are (2, 9), (3, 8), (4, 7), (5, 6), (6, 5), (7, 4), (8, 3), and (9, 2). 
    // Note that (1, 10) and (10, 1) do not satisfy the conditions because 10 contains 0 in its decimal representation.
    fmt.Println(countNoZeroPairs(11)) // 8

    fmt.Println(countNoZeroPairs(64)) // 51
    fmt.Println(countNoZeroPairs(999)) // 656
    fmt.Println(countNoZeroPairs(1024)) // 637
    fmt.Println(countNoZeroPairs(1_000_000_007)) // 188743664
    fmt.Println(countNoZeroPairs(999_999_999_999_999)) // 45237049828496
    fmt.Println(countNoZeroPairs(1_000_000_000_000_000)) // 50891681057058
}