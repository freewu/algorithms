package main

// 799. Champagne Tower
// We stack glasses in a pyramid, where the first row has 1 glass, the second row has 2 glasses, and so on until the 100th row.  
// Each glass holds one cup of champagne.

// Then, some champagne is poured into the first glass at the top.  
// When the topmost glass is full, any excess liquid poured will fall equally to the glass immediately to the left and right of it.  
// When those glasses become full, any excess champagne will fall equally to the left and right of those glasses, and so on.  (A glass at the bottom row has its excess champagne fall on the floor.)

// For example, after one cup of champagne is poured, the top most glass is full.  
// After two cups of champagne are poured, the two glasses on the second row are half full.  
// After three cups of champagne are poured, those two cups become full - there are 3 full glasses total now.  
// After four cups of champagne are poured, the third row has the middle glass half full, and the two outside glasses are a quarter full, as pictured below.

// <img src="https://s3-lc-upload.s3.amazonaws.com/uploads/2018/03/09/tower.png" />

// Now after pouring some non-negative integer cups of champagne, return how full the jth glass in the ith row is (both i and j are 0-indexed.)

// Example 1:
// Input: poured = 1, query_row = 1, query_glass = 1
// Output: 0.00000
// Explanation: We poured 1 cup of champange to the top glass of the tower (which is indexed as (0, 0)). There will be no excess liquid so all the glasses under the top glass will remain empty.

// Example 2:
// Input: poured = 2, query_row = 1, query_glass = 1
// Output: 0.50000
// Explanation: We poured 2 cups of champange to the top glass of the tower (which is indexed as (0, 0)). There is one cup of excess liquid. The glass indexed as (1, 0) and the glass indexed as (1, 1) will share the excess liquid equally, and each will get half cup of champange.

// Example 3:
// Input: poured = 100000009, query_row = 33, query_glass = 17
// Output: 1.00000
 
// Constraints:
//     0 <= poured <= 10^9
//     0 <= query_glass <= query_row < 100

import "fmt"

func champagneTower(poured int, query_row int, query_glass int) float64 {
    dp := make([][]float64, query_row + 1)
    for i := 0; i < query_row + 1; i++ {
        dp[i] = make([]float64, query_glass + 1)
        for j := 0; j < query_glass + 1; j++ {
            dp[i][j] = -1
        }
    }
    max := func (x, y float64) float64 { if x > y { return x; }; return y; }
    var dfs func(row, col, poured int) float64
    dfs = func(row, col, poured int) float64 {
        if col < 0 || col > row { return 0 }
        if row == 0 && col == 0 { return float64(poured) }
        if dp[row][col] != -1   { return dp[row][col] }
        dp[row][col] = max(dfs(row-1, col-1, poured) - 1, 0) / 2 + max(dfs(row-1, col, poured) - 1, 0) / 2
        return dp[row][col]
    }
    res := dfs(query_row, query_glass, poured)
    if res > 1 { 
        return 1
    }
    return res
}

func champagneTower1(poured int, query_row int, query_glass int) float64 {
    dp := make([][]float64, query_row+1)
    for i := 0; i < query_row + 1; i++ {
        dp[i] = make([]float64, i+1)
    }
    // 初始化
    dp[0][0] = float64(poured)
    for i := 0; i < query_row; i++ {
        for j := 0; j <= i; j++ {
            if dp[i][j] >= 1 {
                // 上一杯平均分到下面两个杯子
                dp[i+1][j] = dp[i+1][j] + (dp[i][j] - 1)/2.0
                dp[i+1][j+1] = dp[i+1][j+1] + (dp[i][j] - 1)/2.0
            }
        }
    }
    min := func (x, y float64) float64 { if x < y { return x; }; return y; }
    return min(1.0, dp[query_row][query_glass])
}

func main() {
    // Example 1:
    // Input: poured = 1, query_row = 1, query_glass = 1
    // Output: 0.00000
    // Explanation: We poured 1 cup of champange to the top glass of the tower (which is indexed as (0, 0)). There will be no excess liquid so all the glasses under the top glass will remain empty.
    fmt.Println(champagneTower(1,1,1)) // 0.00000
    // Example 2:
    // Input: poured = 2, query_row = 1, query_glass = 1
    // Output: 0.50000
    // Explanation: We poured 2 cups of champange to the top glass of the tower (which is indexed as (0, 0)). There is one cup of excess liquid. The glass indexed as (1, 0) and the glass indexed as (1, 1) will share the excess liquid equally, and each will get half cup of champange.
    fmt.Println(champagneTower(2,1,1)) // 0.50000
    // Example 3:
    // Input: poured = 100000009, query_row = 33, query_glass = 17
    // Output: 1.00000
    fmt.Println(champagneTower(100000009,33,17)) // 1.00000

    fmt.Println(champagneTower(0,0,0)) // 1
    fmt.Println(champagneTower(0,99,0)) // 0
    fmt.Println(champagneTower(0,99,99)) // 0
    fmt.Println(champagneTower(1_000_000_000,0,0)) // 1
    fmt.Println(champagneTower(1_000_000_000,99,0)) // 0
    fmt.Println(champagneTower(1_000_000_000,99,99)) // 0

    fmt.Println(champagneTower1(1,1,1)) // 0.00000
    fmt.Println(champagneTower1(2,1,1)) // 0.50000
    fmt.Println(champagneTower1(100000009,33,17)) // 1.00000
    fmt.Println(champagneTower1(0,0,0)) // 1
    fmt.Println(champagneTower1(0,99,0)) // 0
    fmt.Println(champagneTower1(0,99,99)) // 0
    fmt.Println(champagneTower1(1_000_000_000,0,0)) // 1
    fmt.Println(champagneTower1(1_000_000_000,99,0)) // 0
    fmt.Println(champagneTower1(1_000_000_000,99,99)) // 0
}