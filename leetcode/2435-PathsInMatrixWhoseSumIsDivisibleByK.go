package main

// 2435. Paths in Matrix Whose Sum Is Divisible by K
// You are given a 0-indexed m x n integer matrix grid and an integer k. 
// You are currently at position (0, 0) and you want to reach position (m - 1, n - 1) moving only down or right.

// Return the number of paths where the sum of the elements on the path is divisible by k. 
// Since the answer may be very large, return it modulo 10^9 + 7.

// Example 1:
// <img src="https://assets.leetcode.com/uploads/2022/08/13/image-20220813183124-1.png" />
// Input: grid = [[5,2,4],[3,0,5],[0,7,2]], k = 3
// Output: 2
// Explanation: There are two paths where the sum of the elements on the path is divisible by k.
// The first path highlighted in red has a sum of 5 + 2 + 4 + 5 + 2 = 18 which is divisible by 3.
// The second path highlighted in blue has a sum of 5 + 3 + 0 + 5 + 2 = 15 which is divisible by 3.

// Example 2:
// <img src="https://assets.leetcode.com/uploads/2022/08/17/image-20220817112930-3.png" />
// Input: grid = [[0,0]], k = 5
// Output: 1
// Explanation: The path highlighted in red has a sum of 0 + 0 = 0 which is divisible by 5.

// Example 3:
// <img src="https://assets.leetcode.com/uploads/2022/08/12/image-20220812224605-3.png" />
// Input: grid = [[7,3,4,9],[2,3,6,2],[2,3,7,0]], k = 1
// Output: 10
// Explanation: Every integer is divisible by 1 so the sum of the elements on every possible path is divisible by k.

// Constraints:
//     m == grid.length
//     n == grid[i].length
//     1 <= m, n <= 5 * 10^4
//     1 <= m * n <= 5 * 10^4
//     0 <= grid[i][j] <= 100
//     1 <= k <= 50

import "fmt"

// Recursion 超出时间限制 73 / 88 
func numberOfPaths(grid [][]int, k int) int {
    n, m, mod := len(grid), len(grid[0]), 1_000_000_007
    var dfs func(r, c, sum int) int 
    dfs = func(r, c, sum int) int {
        if r == n || c == m { return 0 } // If we are out of the grid just return {no possible paths}
        sum += grid[r][c]
        rem := sum % k
        if r == n - 1 && c == m - 1 {
            if rem == 0 { return 1 } // If we have reached the end of grid check if the sum is divisible by k
            return 0
        }
        return (dfs(r + 1, c, sum) + dfs(r, c + 1, sum)) % mod // Go down and right
    }
    return dfs(0, 0, 0) % mod
}

// dp + dfs
func numberOfPaths1(grid [][]int, k int) int {
    n, m, mod := len(grid), len(grid[0]), 1_000_000_007
    dp := make([][][]int, n)
    for i := range dp {
        dp[i] = make([][]int, m)
        for j := range dp[i] {
            dp[i][j] = make([]int, k)
            for g := range dp[i][j] {
                dp[i][j][g] = -1
            }
        }
    }
    var dfs func(r, c, sum int) int 
    dfs = func(r, c, sum int) int {
        if r == n || c == m { return 0 } // If we are out of the grid just return {no possible paths}
        sum += grid[r][c]
        rem := sum % k
        if dp[r][c][rem] != -1 { return dp[r][c][rem] }
        if r == n - 1 && c == m - 1 {
            dp[r][c][rem] = 0
            if rem == 0 { dp[r][c][rem] = 1 } // If we have reached the end of grid check if the sum is divisible by k
            return dp[r][c][rem]
        }
        dp[r][c][rem] = (dfs(r + 1, c, sum) + dfs(r, c + 1, sum)) % mod // Go down and right
        return dp[r][c][rem]
    }
    return dfs(0, 0, 0) % mod
}

func numberOfPaths2(grid [][]int, k int) int {
    n, mod := len(grid[0]), 1_000_000_007
    dp := make([][]int, n + 1)
    for i := range dp {
        dp[i] = make([]int, k)
    }
    dp[1][0] = 1
    for i := range grid {
        for j, v := range grid[i] {
            next := make([]int, k)
            for d := 0; d < k; d++ {
                next[(d + v) % k] = (dp[j][d] + dp[j + 1][d]) % mod
            }
            dp[j + 1] = next
        }
    }
    return dp[n][0]
}

func numberOfPaths3(grid [][]int, k int) int {
    n, mod := len(grid[0]), 1_000_000_007
    dp := make([][]int, n + 1)
    for i := range dp {
        dp[i] = make([]int, k)
    }
    dp[1][0] = 1
    nf := make([]int, k) // 避免在循环内反复创建 []int
    for i := range grid {
        for j, x := range grid[i] {
            for s := 0; s < k; s++ {
                k1 := (s + x) % k
                nf[s] = (dp[j+1][k1] + dp[j][k1]) % mod
            }
            copy(dp[j+1], nf) // 复制到 dp[j+1] 中
        }
    }
    return dp[n][0] 
}

func main() {
    // Example 1:
    // <img src="https://assets.leetcode.com/uploads/2022/08/13/image-20220813183124-1.png" />
    // Input: grid = [[5,2,4],[3,0,5],[0,7,2]], k = 3
    // Output: 2
    // Explanation: There are two paths where the sum of the elements on the path is divisible by k.
    // The first path highlighted in red has a sum of 5 + 2 + 4 + 5 + 2 = 18 which is divisible by 3.
    // The second path highlighted in blue has a sum of 5 + 3 + 0 + 5 + 2 = 15 which is divisible by 3.
    fmt.Println(numberOfPaths([][]int{{5,2,4},{3,0,5},{0,7,2}}, 3)) // 2
    // Example 2:
    // <img src="https://assets.leetcode.com/uploads/2022/08/17/image-20220817112930-3.png" />
    // Input: grid = [[0,0]], k = 5
    // Output: 1
    // Explanation: The path highlighted in red has a sum of 0 + 0 = 0 which is divisible by 5.
    fmt.Println(numberOfPaths([][]int{{0,0}}, 5)) // 1
    // Example 3:
    // <img src="https://assets.leetcode.com/uploads/2022/08/12/image-20220812224605-3.png" />
    // Input: grid = [[7,3,4,9],[2,3,6,2],[2,3,7,0]], k = 1
    // Output: 10
    // Explanation: Every integer is divisible by 1 so the sum of the elements on every possible path is divisible by k.
    fmt.Println(numberOfPaths([][]int{{7,3,4,9},{2,3,6,2},{2,3,7,0}}, 1)) // 10

    fmt.Println(numberOfPaths1([][]int{{5,2,4},{3,0,5},{0,7,2}}, 3)) // 2
    fmt.Println(numberOfPaths1([][]int{{0,0}}, 5)) // 1
    fmt.Println(numberOfPaths1([][]int{{7,3,4,9},{2,3,6,2},{2,3,7,0}}, 1)) // 10

    fmt.Println(numberOfPaths2([][]int{{5,2,4},{3,0,5},{0,7,2}}, 3)) // 2
    fmt.Println(numberOfPaths2([][]int{{0,0}}, 5)) // 1
    fmt.Println(numberOfPaths2([][]int{{7,3,4,9},{2,3,6,2},{2,3,7,0}}, 1)) // 10

    fmt.Println(numberOfPaths3([][]int{{5,2,4},{3,0,5},{0,7,2}}, 3)) // 2
    fmt.Println(numberOfPaths3([][]int{{0,0}}, 5)) // 1
    fmt.Println(numberOfPaths3([][]int{{7,3,4,9},{2,3,6,2},{2,3,7,0}}, 1)) // 10
}