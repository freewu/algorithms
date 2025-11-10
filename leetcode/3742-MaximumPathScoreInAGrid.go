package main

// 3742. Maximum Path Score in a Grid
// You are given an m x n grid where each cell contains one of the values 0, 1, or 2. You are also given an integer k.

// You start from the top-left corner (0, 0) and want to reach the bottom-right corner (m - 1, n - 1) by moving only right or down.

// Each cell contributes a specific score and incurs an associated cost, according to their cell values:
//     0: adds 0 to your score and costs 0.
//     1: adds 1 to your score and costs 1.
//     2: adds 2 to your score and costs 1. ​​​​​​​

// Return the maximum score achievable without exceeding a total cost of k, or -1 if no valid path exists.

// Note: If you reach the last cell but the total cost exceeds k, the path is invalid.

// Example 1:
// Input: grid = [[0, 1],[2, 0]], k = 1
// Output: 2
// Explanation:​​​​​​​
// The optimal path is:
// | Cell    | grid[i][j]	| Score	| Total Score | Cost | Total Cost |
// | (0, 0)  |  0         | 0	    | 0	          | 0	 | 0          |
// | (1, 0)  |  2         | 2	    | 2	          | 1	 | 1          |
// | (1, 1)  |  0         | 0	    | 2	          | 0	 | 1          |
// Thus, the maximum possible score is 2.

// Example 2:
// Input: grid = [[0, 1],[1, 2]], k = 1
// Output: -1
// Explanation:
// There is no path that reaches cell (1, 1)​​​​​​​ without exceeding cost k. Thus, the answer is -1.

// Constraints:
//     1 <= m, n <= 200
//     0 <= k <= 10^3​​​​​​​
//     ​​​​​​​grid[0][0] == 0
//     0 <= grid[i][j] <= 2

import "fmt"
import "slices"

func maxPathScore(grid [][]int, k int) int {
    n, m := len(grid[0]), len(grid)
    f := make([][][]int, m + 1)
    for i := range f {
        f[i] = make([][]int, n + 1)
        for j := range f[i] {
            f[i][j] = make([]int, k + 2)
            for p := range f[i][j] {
                f[i][j][p] = -1 << 31
            }
        }
    }
    for l := 1; l < k + 2; l++ {
        f[0][1][l] = 0 // 保证 f[1][1][] 计算正确
    }
    for i, row := range grid {
        for j, x := range row {
            for l := 0; l < k + 1; l++ {
                nl := l 
                if x > 0 {
                    nl--
                }
                f[i+1][j+1][l+1] = max(f[i][j+1][nl+1], f[i+1][j][nl + 1]) + x
            }
        }
    }
    res := f[m][n][k + 1] 
    if res < 0 { return -1 }
    return res
}

func maxPathScore1(grid [][]int, k int) int {
    m, n := len(grid), len(grid[0])
    arr := make([][][]int, m)
    for i := range m {
        arr[i] = make([][]int, n)
        for j := range n {
            arr[i][j] = make([]int, k+1)
        }
    }
    switch grid[0][0] {
    case 0:
        arr[0][0][0] = 0 + 1
    case 1:
        arr[0][0][1] = 1 + 1
    default:
        arr[0][0][1] = 2 + 1
    }
    for i := range m {
        for j := range n {
            var a, b int
            switch grid[i][j] {
            case 1:
                a, b = 1, 1
            case 2:
                a, b = 2, 1
            }
            if i-1 >= 0 {
                for l := range k + 1 {
                    if arr[i-1][j][l] > 0 && l+b <= k {
                        arr[i][j][l+b] = max(arr[i][j][l+b], arr[i-1][j][l]+a)
                    }
                }
            }
            if j-1 >= 0 {
                for l := range k + 1 {
                    if arr[i][j-1][l] > 0 && l+b <= k {
                        arr[i][j][l+b] = max(arr[i][j][l+b], arr[i][j-1][l]+a)
                    }
                }
            }
        }
    }
    res := slices.Max(arr[m-1][n-1])
    if res > 0 {
        return res - 1
    }
    return -1
}

func main() {
    // Example 1:
    // Input: grid = [[0, 1],[2, 0]], k = 1
    // Output: 2
    // Explanation:​​​​​​​
    // The optimal path is:
    // | Cell    | grid[i][j]	| Score	| Total Score | Cost | Total Cost |
    // | (0, 0)  |  0         | 0	    | 0	          | 0	 | 0          |
    // | (1, 0)  |  2         | 2	    | 2	          | 1	 | 1          |
    // | (1, 1)  |  0         | 0	    | 2	          | 0	 | 1          |
    // Thus, the maximum possible score is 2.
    fmt.Println(maxPathScore([][]int{{0, 1},{2, 0}}, 1)) // 2
    // Example 2:
    // Input: grid = [[0, 1],[1, 2]], k = 1
    // Output: -1
    // Explanation:
    // There is no path that reaches cell (1, 1)​​​​​​​ without exceeding cost k. Thus, the answer is -1.
    fmt.Println(maxPathScore([][]int{{0, 1},{1, 2}}, 1)) // -1

    fmt.Println(maxPathScore1([][]int{{0, 1},{2, 0}}, 1)) // 2
    fmt.Println(maxPathScore1([][]int{{0, 1},{1, 2}}, 1)) // -1
}