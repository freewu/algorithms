package main

// 1937. Maximum Number of Points with Cost
// You are given an m x n integer matrix points (0-indexed). 
// Starting with 0 points, you want to maximize the number of points you can get from the matrix.

// To gain points, you must pick one cell in each row. 
// Picking the cell at coordinates (r, c) will add points[r][c] to your score.

// However, you will lose points if you pick a cell too far from the cell that you picked in the previous row. 
// For every two adjacent rows r and r + 1 (where 0 <= r < m - 1), picking cells at coordinates (r, c1) and (r + 1, c2) will subtract abs(c1 - c2) from your score.

// Return the maximum number of points you can achieve.

// abs(x) is defined as:
//     x for x >= 0.
//     -x for x < 0.
 
// Example 1:
// <img src="https://assets.leetcode.com/uploads/2021/07/12/screenshot-2021-07-12-at-13-40-26-diagram-drawio-diagrams-net.png" />
// Input: points = [[1,2,3],[1,5,1],[3,1,1]]
// Output: 9
// Explanation:
// The blue cells denote the optimal cells to pick, which have coordinates (0, 2), (1, 1), and (2, 0).
// You add 3 + 5 + 3 = 11 to your score.
// However, you must subtract abs(2 - 1) + abs(1 - 0) = 2 from your score.
// Your final score is 11 - 2 = 9.

// Example 2:
// <img src="https://assets.leetcode.com/uploads/2021/07/12/screenshot-2021-07-12-at-13-42-14-diagram-drawio-diagrams-net.png" />
// Input: points = [[1,5],[2,3],[4,2]]
// Output: 11
// Explanation:
// The blue cells denote the optimal cells to pick, which have coordinates (0, 1), (1, 1), and (2, 0).
// You add 5 + 3 + 4 = 12 to your score.
// However, you must subtract abs(1 - 1) + abs(1 - 0) = 1 from your score.
// Your final score is 12 - 1 = 11.

// Constraints:
//     m == points.length
//     n == points[r].length
//     1 <= m, n <= 10^5
//     1 <= m * n <= 10^5
//     0 <= points[r][c] <= 10^5

import "fmt"

func maxPoints(points [][]int) int64 {
    res, m, n := 0, len(points), len(points[0])
    dp, g := make([]int, n), make([]int, n) //上一个的, 新生成的
    maxLeft, maxRight := make([]int, n), make([]int, n)
    for i := 0; i < n; i++ {
        dp[i] = points[0][i]
    }
    max := func (x, y int) int { if x > y { return x; }; return y; }
    for i := 1; i < m; i++ {
        maxLeft[0], maxRight[n-1] = dp[0], dp[n-1] - (n-1)
        for j := 1; j < n; j++ {
            maxLeft[j] = max(maxLeft[j-1], dp[j] + j)
            maxRight[(n-1)-j] = max(maxRight[n-j], dp[(n-1)-j] - (n-1) + j)
        }
        // 更新dp的值.
        for j := 0; j < n; j++ {
            g[j] = max(points[i][j] - j + maxLeft[j], points[i][j] + j + maxRight[j])
        }
        dp = g
    }
    for i := 0; i < n; i++ {
        res = max(res, dp[i]) //找最大值.
    }
    return int64(res)
}

func maxPoints1(points [][]int) int64 {
    res, n, inf := 0, len(points[0]), -1 << 31
    dp, sufMax := make([][2]int, n), make([]int, n) // 后缀最大值
    for i, row := range points {
        if i == 0 {
            for j, v := range row {
                res = max(res, v)
                dp[j][0] = v + j
                dp[j][1] = v - j
            }
        } else {
            preMax := inf
            for j, v := range row {
                preMax = max(preMax, dp[j][0])
                t := max(v-j+preMax, v+j+sufMax[j]) // 左侧和右侧的最大值即为选择 points[i][j] 时的计算结果
                res = max(res, t) // 直接更新答案，这样下面就不直接存储 t 了，改为存储 t + j 和 t - j
                dp[j][0] = t + j
                dp[j][1] = t - j
            }
        }
        // 计算完一整行 f 后，对于每个位置 j，计算其右侧的所有 f[k] - k 的最大值
        // 这可以通过倒着遍历 f 求出
        sufMax[n-1] = dp[n-1][1]
        for j := n - 2; j >= 0; j-- {
            sufMax[j] = max(sufMax[j+1], dp[j][1])
        }
    }
    return int64(res)
}

func main() {
    // Example 1:
    // <img src="https://assets.leetcode.com/uploads/2021/07/12/screenshot-2021-07-12-at-13-40-26-diagram-drawio-diagrams-net.png" />
    // Input: points = [[1,2,3],[1,5,1],[3,1,1]]
    // Output: 9
    // Explanation:
    // The blue cells denote the optimal cells to pick, which have coordinates (0, 2), (1, 1), and (2, 0).
    // You add 3 + 5 + 3 = 11 to your score.
    // However, you must subtract abs(2 - 1) + abs(1 - 0) = 2 from your score.
    // Your final score is 11 - 2 = 9.
    fmt.Println(maxPoints([][]int{{1,2,3},{1,5,1},{3,1,1}})) // 9
    // Example 2:
    // <img src="https://assets.leetcode.com/uploads/2021/07/12/screenshot-2021-07-12-at-13-42-14-diagram-drawio-diagrams-net.png" />
    // Input: points = [[1,5],[2,3],[4,2]]
    // Output: 11
    // Explanation:
    // The blue cells denote the optimal cells to pick, which have coordinates (0, 1), (1, 1), and (2, 0).
    // You add 5 + 3 + 4 = 12 to your score.
    // However, you must subtract abs(1 - 1) + abs(1 - 0) = 1 from your score.
    // Your final score is 12 - 1 = 11.
    fmt.Println(maxPoints([][]int{{1,5},{2,3},{4,2}})) // 11

    fmt.Println(maxPoints1([][]int{{1,2,3},{1,5,1},{3,1,1}})) // 9
    fmt.Println(maxPoints1([][]int{{1,5},{2,3},{4,2}})) // 11
}