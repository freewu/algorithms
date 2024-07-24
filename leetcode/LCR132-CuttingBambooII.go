package main

// LCR 132. 砍竹子 II
// 现需要将一根长为正整数 bamboo_len 的竹子砍为若干段，每段长度均为 正整数。请返回每段竹子长度的 最大乘积 是多少。
// 答案需要取模 1e9+7（1000000007），如计算初始结果为：1000000008，请返回 1。

// 示例 1：
// 输入：bamboo_len = 12
// 输出：81

// 提示：
//     2 <= bamboo_len <= 1000

import "fmt"

// dp
func cuttingBamboo(bamboo_len int) int {
    dp := make([]int, bamboo_len + 1)
    dp[0], dp[1] = 1, 1
    max := func (x, y int) int { if x > y { return x; }; return y; }
    for i := 1; i <= bamboo_len; i++ {
        for j := 1; j < i; j++ {
            dp[i] = max(dp[i], (j * max(dp[i-j], i-j)) % 1_000_000_007 ) 
        }
    }
    return dp[bamboo_len]
}

func cuttingBamboo1(bamboo_len int) int {
    if bamboo_len <= 3 {
        return bamboo_len - 1
    }
    min := func (x, y int) int { if x < y { return x; }; return y; }
    res := 1
    for bamboo_len > 0 {
        if bamboo_len > 4 {
            res = (res * min(bamboo_len, 3)) % 1_000_000_007
            bamboo_len -= 3
        } else {
            res = (res * bamboo_len) % 1_000_000_007
            bamboo_len = 0
        }
    }
    return res
}

func main() {
    // Example 1:
    // Input: n = 2
    // Output: 1
    // Explanation: 2 = 1 + 1, 1 × 1 = 1.
    fmt.Println(cuttingBamboo(2)) // 1
    // Example 2:
    // Input: n = 10
    // Output: 36
    // Explanation: 10 = 3 + 3 + 4, 3 × 3 × 4 = 36.
    fmt.Println(cuttingBamboo(10)) // 36
    fmt.Println(cuttingBamboo(12)) // 81
    fmt.Println(cuttingBamboo(120)) // 953271190

    fmt.Println(cuttingBamboo1(2)) // 1
    fmt.Println(cuttingBamboo1(10)) // 36
    fmt.Println(cuttingBamboo1(12)) // 81
    fmt.Println(cuttingBamboo1(120)) // 953271190
}