package main

// 1411. Number of Ways to Paint N × 3 Grid
// You have a grid of size n x 3 and you want to paint each cell of the grid 
// with exactly one of the three colors: Red, Yellow, or Green while making sure that no two adjacent cells have the same color (i.e., no two cells that share vertical or horizontal sides have the same color).

// Given n the number of rows of the grid, return the number of ways you can paint this grid. 
// As the answer may grow large, the answer must be computed modulo 10^9 + 7.

// Example 1:
// <img src="https://assets.leetcode.com/uploads/2020/03/26/e1.png" />
// Input: n = 1
// Output: 12
// Explanation: There are 12 possible way to paint the grid as shown.

// Example 2:
// Input: n = 5000
// Output: 30228214

// Constraints:
//     n == grid.length
//     1 <= n <= 5000

import "fmt"

func numOfWays(n int) int {
    a, b, mod := 6, 6, 1_000_000_007
    for i := 1; i < n; i++ {
        a, b = (3 * a + 2 * b) % mod, (2 * a + 2 * b) % mod
    }
    return (a + b) % mod
}

func numOfWays1(n int) int {
    // 状压dp + 整行考虑(而不是考虑每个格子)  + 三进制 + 预处理
    // 如果使用普通的dfs, 需要5个参数 x,y表示坐标, pre:上一行的涂色, left:左侧的涂色, line:整行已经处理过的部分的涂色(方便给下一行提供涂色)
    // 使用三进制(大小3) 而不是 3个bit位来确定颜色(大小8). 因为一个位置只能图一个颜色,无需考虑 集合
    // 预处理一: 生成所有合理的 一行涂色 (所有方案只有3*3,合理的<27)
    // 预处理二: 因为只有相邻依赖,枚举当前行的涂色,寻找到上一行的可行的涂色 ( <27*27≈1000), 这样 O(10^3 *5000)< 10^7是可以过的
    types := []int{}
    for i := 0; i < 3; i++ {
        for j := 0; j < 3; j++ {
            for k := 0; k < 3; k++ {
                if i != j && j != k {
                    types = append(types, 9 * i + 3 * j + k)
                }
            }
        }
    }
    // 预处理,计算当前颜色为t1时,它相邻的合法的颜色有那些t2
    m, mod := len(types), 1_000_000_007
    valid := make([][]int, m)
    for i, t1 := range types {
        // trick!! 三进制如何取得各个位置的值
        i1, j1, k1 := t1 / 9, (t1 % 9)/3, t1%3 // 取第二位的值 先消除掉前一位应该用 t%9而不是 t/9
        for j, t2 := range types {
            i2, j2, k2 := t2/9, (t2%9)/3, t2%3
            if i1 != i2 && j1 != j2 && k1 != k2 {
                valid[i] = append(valid[i], j)
            }
        }
    }
    dp := make([][]int, n) // dp[i][type]在 第i排的颜色状态为type时(使用的是type的索引),可行的方案数
    for i := range dp {
        dp[i] = make([]int, m)
    }
    for j := 0; j < m; j++ { // base case 第一排可以任意选颜色
        dp[0][j] = 1
    }
    for i := 1; i < n; i++ {
        for j, _ := range types { 
            for _, pre := range valid[j] {
                dp[i][j] = (dp[i][j] + dp[i-1][pre]) % mod
            }
        }
    }
    res := 0
    for _, v := range dp[n-1] {
        res = (res + v) % mod
    }
    return res
}

func main() {
    // Example 1:
    // <img src="https://assets.leetcode.com/uploads/2020/03/26/e1.png" />
    // Input: n = 1
    // Output: 12
    // Explanation: There are 12 possible way to paint the grid as shown.
    fmt.Println(numOfWays(1)) // 12
    // Example 2:
    // Input: n = 5000
    // Output: 30228214
    fmt.Println(numOfWays(5000)) // 30228214

    fmt.Println(numOfWays(2)) // 54
    fmt.Println(numOfWays(100)) // 905790447
    fmt.Println(numOfWays(999)) // 672393158
    fmt.Println(numOfWays(1000)) // 650420578
    fmt.Println(numOfWays(1024)) // 317135560
    fmt.Println(numOfWays(4999)) // 134620719

    fmt.Println(numOfWays1(1)) // 12
    fmt.Println(numOfWays1(5000)) // 30228214
    fmt.Println(numOfWays1(2)) // 54
    fmt.Println(numOfWays1(100)) // 905790447
    fmt.Println(numOfWays1(999)) // 672393158
    fmt.Println(numOfWays1(1000)) // 650420578
    fmt.Println(numOfWays1(1024)) // 317135560
    fmt.Println(numOfWays1(4999)) // 134620719
}