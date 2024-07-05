package main

// LCR 168. 丑数
// 给你一个整数 n ，请你找出并返回第 n 个 丑数 。
// 说明：丑数是只包含质因数 2、3 和/或 5 的正整数；1 是丑数。

// 示例 1：
// 输入: n = 10
// 输出: 12
// 解释: 1, 2, 3, 4, 5, 6, 8, 9, 10, 12 是前 10 个丑数。

// 提示： 
//     1 <= n <= 1690

// 丑数 就是质因子只包含 2、3 和 5 的正整数
// [1, 2, 3, 4, 5, 6, 8, 9, 10, 12] 是由前 10 个丑数组成的序列
// 1 2 3     2*2 5 2*3 2*2*2 3*3 2*5 2*2*3

import "fmt"

func nthUglyNumber(n int) int {
    p2, p3, p5 := 0, 0, 0
    uglyNumbers := []int{1}
    for len(uglyNumbers) < n {
        v2, v3, v5 := uglyNumbers[p2] * 2,uglyNumbers[p3] * 3,uglyNumbers[p5] * 5
        var v int
        if v2 <= v3 && v2 <= v5 {
            p2++
            v = v2
        } else if v3 <= v2 && v3 <= v5 {
            p3++
            v = v3
        } else {
            p5++
            v = v5
        }
        if uglyNumbers[len(uglyNumbers)-1] != v {
            uglyNumbers = append(uglyNumbers, v)
        }
    }
    return uglyNumbers[len(uglyNumbers) - 1]
}

func nthUglyNumber1(n int) int {
    min := func (a, b int) int { if a < b { return a; }; return b; }
    dp, p2, p3, p5 := make([]int, n + 1), 1, 1, 1
    dp[0], dp[1] = 0, 1
    for i := 2; i <= n; i++ {
        x2, x3, x5 := dp[p2]*2, dp[p3]*3, dp[p5]*5
        dp[i] = min(min(x2, x3), x5)
        if dp[i] == x2 { p2++ }
        if dp[i] == x3 { p3++ }
        if dp[i] == x5 { p5++ }
    }
    return dp[n]
}

func nthUglyNumber2(n int) int {
    dp := []int{1}
    min := func (a, b int) int { if a < b { return a; }; return b; }
    for i, j, k := 0, 0, 0; len(dp) < n; {
        t := min(dp[i] * 2, min(dp[j] * 3, dp[k] * 5))
        dp = append(dp, t)
        if dp[i] * 2 == t {
            i ++ 
        }
        if dp[j] * 3 == t {
            j ++ 
        }
        if dp[k] * 5 == t {
            k ++ 
        }
    }
    return dp[n - 1]
}

func main() {
    // Explanation: [1, 2, 3, 4, 5, 6, 8, 9, 10, 12] is the sequence of the first 10 ugly numbers.
    fmt.Println(nthUglyNumber(10)) // 12
    // Explanation: 1 has no prime factors, therefore all of its prime factors are limited to 2, 3, and 5.
    fmt.Println(nthUglyNumber(1)) // 1

    fmt.Println(nthUglyNumber1(10)) // 12
    fmt.Println(nthUglyNumber1(1)) // 1

    fmt.Println(nthUglyNumber2(10)) // 12
    fmt.Println(nthUglyNumber2(1)) // 1
}