package main

// LCR 131. 砍竹子 I
// 现需要将一根长为正整数 bamboo_len 的竹子砍为若干段，每段长度均为正整数。请返回每段竹子长度的最大乘积是多少。

// 示例 1：
// 输入: bamboo_len = 12
// 输出: 81

// 提示：
//     2 <= bamboo_len <= 58

import "fmt"

// dp
func cuttingBamboo(bamboo_len int) int {
    dp := make([]int, bamboo_len + 1)
    dp[0], dp[1] = 1, 1
    max := func (x, y int) int { if x > y { return x; }; return y; }
    for i := 1; i <= bamboo_len; i++ {
        for j := 1; j < i; j++ {
            dp[i] = max(dp[i], j * max(dp[i-j], i-j))
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
            res *= min(bamboo_len, 3)
            bamboo_len -= 3
        } else {
            res *= bamboo_len
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


    fmt.Println(cuttingBamboo1(2)) // 1
    fmt.Println(cuttingBamboo1(10)) // 36
    fmt.Println(cuttingBamboo1(12)) // 81
}