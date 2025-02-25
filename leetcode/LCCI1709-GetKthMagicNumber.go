package main

// 面试题 17.09. Get Kth Magic Number LCCI
// Design an algorithm to find the kth number such that the only prime factors are 3, 5, and 7. 
// Note that 3, 5, and 7 do not have to be factors, but it should not have any other prime factors. 
// For example, the first several multiples would be (in order) 1, 3, 5, 7, 9, 15, 21.

// Example 1:
// Input: k = 5
// Output: 9

import "fmt"

func getKthMagicNumber(k int) int {
    dp := make([]int, k + 1)
    dp[0] = 1 // 初始默认值为 1
    p3, p5, p7 := 0, 0, 0 // 设置因子的基础下标(3,5,7)
    min := func (x, y int) int { if x < y { return x; }; return y; }
    for i := 1; i < k; i++ { // 开始遍历并计算从下标1到下标 k - 1 的值
        dp[i] = min(dp[p3] * 3, min(dp[p5] * 5, dp[p7] * 7)) // 计算三乘以三个因子之后谁最小，并将其赋值给dp[i]
        // 维护各因子的下标值
        if dp[i] % 3 == 0 { p3++ }
        if dp[i] % 5 == 0 { p5++ }
        if dp[i] % 7 == 0 { p7++ }
    }
    return dp[k - 1]
}

func main() {
    fmt.Println(getKthMagicNumber(1)) // 1 
    fmt.Println(getKthMagicNumber(2)) // 3
    fmt.Println(getKthMagicNumber(3)) // 5
    fmt.Println(getKthMagicNumber(4)) // 7 
    fmt.Println(getKthMagicNumber(5)) // 9
    fmt.Println(getKthMagicNumber(6)) // 21
    fmt.Println(getKthMagicNumber(99)) // 32805
    fmt.Println(getKthMagicNumber(100)) // 33075
    fmt.Println(getKthMagicNumber(1024)) // 98876953125
}