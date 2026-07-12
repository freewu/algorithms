package main

// 3989. Maximum Consistent Columns in a Grid
// You are given a 2D integer array grid of size m x n, and an integer limit.

// You may remove zero or more columns from the grid, but at least one column must remain. 
// The relative order of the remaining columns must be preserved.

// A grid is called consistent if for every row i, and for every pair of adjacent remaining columns a and b with a < b, the following holds: |grid[i][b] - grid[i][a]| <= limit.

// Return the maximum number of columns that can remain such that the resulting grid is consistent.

// Example 1:
// Input: grid = [[-2,0,3]], limit = 2
// Output: 2
// Explanation:
// Remove column 2 and keep columns 0 and 1, which gives |grid[0][1] − grid[0][0]| = |0 − (−2)| = 2 <= limit.
// Thus, the maximum number of columns that can remain is 2.

// Example 2:
// Input: grid = [[1,-1,1],[2,2,2]], limit = 1
// Output: 2
// Explanation:
// Remove column 1 and keep columns 0 and 2, which gives
// |grid[0][2] − grid[0][0]| = |1 − 1| = 0 <= limit and
// |grid[1][2] − grid[1][0]| = |2 − 2| = 0 <= limit.
// Thus, the maximum number of columns that can remain is 2.

// Example 3:
// Input: grid = [[-5,5]], limit = 9
// Output: 1
// Explanation:
// Remove either column 0 or column 1, since |grid[0][1] − grid[0][0]| = |5 − (−5)| = 10 > limit.
// Thus, the maximum number of columns that can remain is 1.

// Constraints:
//     1 <= m == grid.length <= 250
//     1 <= n == grid[i].length <= 250
//     -10^5 <= grid[i][j] <= 10^5
//     0 <= limit <= 10^5

import "fmt"

func maxConsistentColumns(grid [][]int, limit int) int {
    abs := func(x int) int { if x < 0 { return -x; }; return x; }
    res, n := 0, len(grid[0])
    f := make([]int, n)
    for i := range n {
    next:
        for j := i - 1; j >= 0; j-- { // 枚举上一个保留的列
            if f[j] <= f[i] {
                continue
            }
            for _, row := range grid {
                if abs(row[i]-row[j]) > limit {
                    continue next // 列 i 和列 j 不是一致的
                }
            }
            f[i] = f[j]
        }
        f[i]++
        res = max(res, f[i])
    }
    return res
}

func maxConsistentColumns1(grid [][]int, limit int) int {
    res, n := 1, len(grid[0])
    memo := make([]int, n)
    var dfs func(int) int
    dfs = func(i int) int {
        p := &memo[i]
        if *p > 0 {
            return *p
        }
        val := 1
        for j := i + 1; j < n; j++ {
            ok := true
            for _, row := range grid {
                d := row[j] - row[i]
                if d < -limit || d > limit {
                    ok = false
                    break
                }
            }
            if ok {
                val = max(val, dfs(j)+1)
            }
        }
        *p = val
        return val
    }
    for i := range n {
        res = max(res, dfs(i))
    }
    return res
}

func main() {
    // Example 1:
    // Input: grid = [[-2,0,3]], limit = 2
    // Output: 2
    // Explanation:
    // Remove column 2 and keep columns 0 and 1, which gives |grid[0][1] − grid[0][0]| = |0 − (−2)| = 2 <= limit.
    // Thus, the maximum number of columns that can remain is 2.
    fmt.Println(maxConsistentColumns([][]int{{-2,0,3}}, 2)) // 2
    // Example 2:
    // Input: grid = [[1,-1,1],[2,2,2]], limit = 1
    // Output: 2
    // Explanation:
    // Remove column 1 and keep columns 0 and 2, which gives
    // |grid[0][2] − grid[0][0]| = |1 − 1| = 0 <= limit and
    // |grid[1][2] − grid[1][0]| = |2 − 2| = 0 <= limit.
    // Thus, the maximum number of columns that can remain is 2.
    fmt.Println(maxConsistentColumns([][]int{{1,-1,1},{2,2,2}}, 1)) // 2
    // Example 3:
    // Input: grid = [[-5,5]], limit = 9
    // Output: 1
    // Explanation:
    // Remove either column 0 or column 1, since |grid[0][1] − grid[0][0]| = |5 − (−5)| = 10 > limit.
    // Thus, the maximum number of columns that can remain is 1.  
    fmt.Println(maxConsistentColumns([][]int{{-5,5}}, 9)) // 1

    fmt.Println(maxConsistentColumns1([][]int{{-2,0,3}}, 2)) // 2
    fmt.Println(maxConsistentColumns1([][]int{{1,-1,1},{2,2,2}}, 1)) // 2
    fmt.Println(maxConsistentColumns1([][]int{{-5,5}}, 9)) // 1
}