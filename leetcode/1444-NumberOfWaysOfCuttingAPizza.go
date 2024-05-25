package main

// 1444. Number of Ways of Cutting a Pizza
// Given a rectangular pizza represented as a rows x cols matrix containing the following characters: 
//     'A' (an apple) and '.' (empty cell) and given the integer k. 
// You have to cut the pizza into k pieces using k-1 cuts. 

// For each cut you choose the direction: vertical or horizontal, 
// then you choose a cut position at the cell boundary and cut the pizza into two pieces. 
// If you cut the pizza vertically, give the left part of the pizza to a person. 
// If you cut the pizza horizontally, give the upper part of the pizza to a person. 
// Give the last piece of pizza to the last person.

// Return the number of ways of cutting the pizza such that each piece contains at least one apple. 
// Since the answer can be a huge number, return this modulo 10^9 + 7.

// Example 1:
// <img src="https://assets.leetcode.com/uploads/2020/04/23/ways_to_cut_apple_1.png" />
// Input: pizza = ["A..","AAA","..."], k = 3
// Output: 3 
// Explanation: The figure above shows the three ways to cut the pizza. Note that pieces must contain at least one apple.

// Example 2:
// Input: pizza = ["A..","AA.","..."], k = 3
// Output: 1

// Example 3:
// Input: pizza = ["A..","A..","..."], k = 1
// Output: 1
 
// Constraints:
//     1 <= rows, cols <= 50
//     rows == pizza.length
//     cols == pizza[i].length
//     1 <= k <= 10
//     pizza consists of characters 'A' and '.' only.

import "fmt"

// dp
func ways(pizza []string, k int) int {
    rows, cols, mod := len(pizza), len(pizza[0]), 1_000_000_007
    apples, dp:= make([][]int, rows + 1), make([][][]int, k)
    for row := 0; row < rows+1; row++ { // init apples
        apples[row] = make([]int, cols+1)
    }
    for remain := 0; remain < k; remain++ { // init dp
        dp[remain] = make([][]int, rows)
        for i := 0; i < rows; i++ {
            dp[remain][i] = make([]int, cols)
        }
    }
    for row := rows - 1; row >= 0; row-- {
        for col := cols - 1; col >= 0; col-- {
            if apples[row] == nil {
                apples[row] = make([]int, cols+1)
            }
            if pizza[row][col] == 'A' {
                apples[row][col]++
            }
            apples[row][col] = apples[row][col] + apples[row][col+1] + apples[row+1][col] - apples[row+1][col+1]
            if apples[row][col] > 0 {
                dp[0][row][col] = 1
            }
        }
    }
    for remain := 1; remain < k; remain++ {
        for row := 0; row < rows; row++ {
            for col := 0; col < cols; col++ {
                for i := row + 1; i < rows; i++ {
                    if apples[row][col] > apples[i][col] {
                        dp[remain][row][col] = (dp[remain][row][col] + dp[remain-1][i][col]) % mod
                    }
                }
                for j := col + 1; j < cols; j++ {
                    if apples[row][col] > apples[row][j] {
                        dp[remain][row][col] = (dp[remain][row][col] + dp[remain-1][row][j]) % mod
                    }
                }
            }
        }
    }
    return dp[k-1][0][0]
}

func ways1(pizza []string, k int) int {
    m, n, mod := len(pizza), len(pizza[0]),1_000_000_007
    // sum[i][j] 为从左上角坐标为 (i，j) 的矩形中苹果的个数，这样
    //  1 可以很快的计算出切掉的矩形中是否含有苹果
    //  2 可以得知剩下的矩形中是否含有苹果
    sum := make([][]int, m+1)
    for i := 0; i < m+1; i++ {
        sum[i] = make([]int, n+1)
    }
    // dp[k][i][j] 意为 剩余 k 刀，从左上角坐标为 (i，j) 的矩阵中，能够切出的方案数
    dp := make([][][]int, k)
    for i := 0; i < k; i++ {
        dp[i] = make([][]int, m+1)
        for j := 0; j <= m; j++ {
            dp[i][j] = make([]int, n+1)
        }
    }
    for i := m - 1; i >= 0; i-- { // 预处理，得到 sum 矩阵
        for j := n - 1; j >= 0; j-- {
            sum[i][j] = sum[i+1][j] + sum[i][j+1] - sum[i+1][j+1]
            if pizza[i][j] == 'A' { sum[i][j]++; }
            // 如果此矩阵有苹果，则切完最后一刀剩下此矩阵时，内含苹果，方案数+1
            // 此矩阵指 （i，j）为左上角，原始矩阵右下角为右下角的矩阵
            if sum[i][j] > 0 { dp[0][i][j] = 1; }
        }
    }
    for k2 := 1; k2 < k; k2++ { // 枚举切割刀数
        for i := 0; i < m; i++ { // 枚举左上角坐标
            for j := 0; j < n; j++ {
                for i2 := i + 1; i2 < m; i2++ { // 枚举横着切
                    if sum[i][j]-sum[i2][j] > 0 { // 判断上一排是否有苹果
                        dp[k2][i][j] = (dp[k2][i][j] + dp[k2-1][i2][j]) % mod
                    }
                }
                for j2 := j + 1; j2 < n; j2++ { // 枚举竖着切
                    if sum[i][j]-sum[i][j2] > 0 { // 判断左一列是否有苹果
                        dp[k2][i][j] = (dp[k2][i][j] + dp[k2-1][i][j2]) % mod
                    }
                }
            }
        }
    }
    return dp[k-1][0][0]
}

func main() {
    // Example 1:
    // <img src="https://assets.leetcode.com/uploads/2020/04/23/ways_to_cut_apple_1.png" />
    // Input: pizza = ["A..","AAA","..."], k = 3
    // Output: 3 
    // Explanation: The figure above shows the three ways to cut the pizza. Note that pieces must contain at least one apple.
    fmt.Println(ways([]string{"A..","AAA","..."}, 3)) // 3
    // Example 2:
    // Input: pizza = ["A..","AA.","..."], k = 3
    // Output: 1
    fmt.Println(ways([]string{"A..","AA.","..."}, 3)) // 1
    // Example 3:
    // Input: pizza = ["A..","A..","..."], k = 1
    // Output: 1
    fmt.Println(ways([]string{"A..","A..","..."}, 1)) // 1

    fmt.Println(ways1([]string{"A..","AAA","..."}, 3)) // 3
    fmt.Println(ways1([]string{"A..","AA.","..."}, 3)) // 1
    fmt.Println(ways1([]string{"A..","A..","..."}, 1)) // 1
}