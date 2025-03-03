package main

// 3459. Length of Longest V-Shaped Diagonal Segment
// You are given a 2D integer matrix grid of size n x m, where each element is either 0, 1, or 2.

// A V-shaped diagonal segment is defined as:
//     1. The segment starts with 1.
//     2. The subsequent elements follow this infinite sequence: 2, 0, 2, 0, ....
//     3. The segment:
//         3.1 Starts along a diagonal direction (top-left to bottom-right, bottom-right to top-left, top-right to bottom-left, or bottom-left to top-right).
//         3.2 Continues the sequence in the same diagonal direction.
//         3.3 Makes at most one clockwise 90-degree turn to another diagonal direction while maintaining the sequence.

// <img src="https://assets.leetcode.com/uploads/2025/01/11/length_of_longest3.jpg" />

// Return the length of the longest V-shaped diagonal segment. 
// If no valid segment exists, return 0.

// Example 1:
// Input: grid = [[2,2,1,2,2],[2,0,2,2,0],[2,0,1,1,0],[1,0,2,2,2],[2,0,0,2,2]]
// Output: 5
// Explanation:
// <img src="https://assets.leetcode.com/uploads/2024/12/09/matrix_1-2.jpg" />
// The longest V-shaped diagonal segment has a length of 5 and follows these coordinates: (0,2) → (1,3) → (2,4), takes a 90-degree clockwise turn at (2,4), and continues as (3,3) → (4,2).

// Example 2:
// Input: grid = [[2,2,2,2,2],[2,0,2,2,0],[2,0,1,1,0],[1,0,2,2,2],[2,0,0,2,2]]
// Output: 4
// Explanation:
// <img src="https://assets.leetcode.com/uploads/2024/12/09/matrix_2.jpg" />
// The longest V-shaped diagonal segment has a length of 4 and follows these coordinates: (2,3) → (3,2), takes a 90-degree clockwise turn at (3,2), and continues as (2,1) → (1,0).

// Example 3:
// Input: grid = [[1,2,2,2,2],[2,2,2,2,0],[2,0,0,0,0],[0,0,2,2,2],[2,0,0,2,0]]
// Output: 5
// Explanation:
// <img src="https://assets.leetcode.com/uploads/2024/12/09/matrix_3.jpg" />
// The longest V-shaped diagonal segment has a length of 5 and follows these coordinates: (0,0) → (1,1) → (2,2) → (3,3) → (4,4).

// Example 4:
// Input: grid = [[1]]
// Output: 1
// Explanation:
// The longest V-shaped diagonal segment has a length of 1 and follows these coordinates: (0,0).

// Constraints:
//     n == grid.length
//     m == grid[i].length
//     1 <= n, m <= 500
//     grid[i][j] is either 0, 1 or 2.

import "fmt"

func lenOfVDiagonal(grid [][]int) int {
    res, m, n := 0, len(grid), len(grid[0])
    memo := make([][][4][2]int, m)
    for i := range memo {
        memo[i] = make([][4][2]int, n)
    }
    max := func (x, y int) int { if x > y { return x; }; return y; }
    directions := [4][2]int{{1, 1}, {1, -1}, {-1, -1}, {-1, 1}}
    var dfs func(int, int, int, int, int) int
    dfs = func(i, j, k, canTurn, target int) (res int) {
        i += directions[k][0]
        j += directions[k][1]
        if i < 0 || i >= m || j < 0 || j >= n || grid[i][j] != target { return }
        p := &memo[i][j][k][canTurn]
        if *p > 0 {
            return *p
        }
        defer func() { *p = res }()
        res = dfs(i, j, k, canTurn, 2-target)
        if canTurn == 1 {
            maxs := [4]int{ m - i - 1, j, i, n - j - 1} // 理论最大值（走到底）
            k = (k + 1) % 4
            // 优化二：如果理论最大值没有超过 res，那么不递归
            if maxs[k] > res {
                res = max(res, dfs(i, j, k, 0, 2-target))
            }
        }
        return res + 1
    }
    for i, row := range grid {
        for j, x := range row {
            if x != 1 { continue }
            maxs := [4]int{ m - i, j + 1, i + 1, n - j } // 理论最大值（走到底）
            for k, mx := range maxs { // 枚举起始方向
                if mx > res { // 优化一：如果理论最大值没有超过 ans，那么不递归
                    res = max(res, dfs(i, j, k, 1, 2) + 1)
                }
            }
        }
    }
    return res
}

func main() {
    // Example 1:
    // Input: grid = [[2,2,1,2,2],[2,0,2,2,0],[2,0,1,1,0],[1,0,2,2,2],[2,0,0,2,2]]
    // Output: 5
    // Explanation:
    // <img src="https://assets.leetcode.com/uploads/2024/12/09/matrix_1-2.jpg" />
    // The longest V-shaped diagonal segment has a length of 5 and follows these coordinates: (0,2) → (1,3) → (2,4), takes a 90-degree clockwise turn at (2,4), and continues as (3,3) → (4,2).
    fmt.Println(lenOfVDiagonal([][]int{{2,2,1,2,2},{2,0,2,2,0},{2,0,1,1,0},{1,0,2,2,2},{2,0,0,2,2}})) // 5
    // Example 2:
    // Input: grid = [[2,2,2,2,2],[2,0,2,2,0],[2,0,1,1,0],[1,0,2,2,2],[2,0,0,2,2]]
    // Output: 4
    // Explanation:
    // <img src="https://assets.leetcode.com/uploads/2024/12/09/matrix_2.jpg" />
    // The longest V-shaped diagonal segment has a length of 4 and follows these coordinates: (2,3) → (3,2), takes a 90-degree clockwise turn at (3,2), and continues as (2,1) → (1,0).
    fmt.Println(lenOfVDiagonal([][]int{{2,2,2,2,2},{2,0,2,2,0},{2,0,1,1,0},{1,0,2,2,2},{2,0,0,2,2}})) // 4
    // Example 3:
    // Input: grid = [[1,2,2,2,2],[2,2,2,2,0],[2,0,0,0,0],[0,0,2,2,2],[2,0,0,2,0]]
    // Output: 5
    // Explanation:
    // <img src="https://assets.leetcode.com/uploads/2024/12/09/matrix_3.jpg" />
    // The longest V-shaped diagonal segment has a length of 5 and follows these coordinates: (0,0) → (1,1) → (2,2) → (3,3) → (4,4).
    fmt.Println(lenOfVDiagonal([][]int{{1,2,2,2,2},{2,2,2,2,0},{2,0,0,0,0},{0,0,2,2,2},{2,0,0,2,0}})) // 5
    // Example 4:
    // Input: grid = [[1]]
    // Output: 1
    // Explanation:
    // The longest V-shaped diagonal segment has a length of 1 and follows these coordinates: (0,0).
    fmt.Println(lenOfVDiagonal([][]int{{1}})) // 1
}