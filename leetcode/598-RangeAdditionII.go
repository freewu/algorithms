package main

// 598. Range Addition II
// You are given an m x n matrix M initialized with all 0's and an array of operations ops, 
// where ops[i] = [ai, bi] means M[x][y] should be incremented by one for all 0 <= x < ai and 0 <= y < bi.

// Count and return the number of maximum integers in the matrix after performing all the operations.

// Example 1:
// <img src="https://assets.leetcode.com/uploads/2020/10/02/ex1.jpg" />
// Input: m = 3, n = 3, ops = [[2,2],[3,3]]
// Output: 4
// Explanation: The maximum integer in M is 2, and there are four of it in M. So return 4.

// Example 2:
// Input: m = 3, n = 3, ops = [[2,2],[3,3],[3,3],[3,3],[2,2],[3,3],[3,3],[3,3],[2,2],[3,3],[3,3],[3,3]]
// Output: 4

// Example 3:
// Input: m = 3, n = 3, ops = []
// Output: 9

// Constraints:
//     1 <= m, n <= 4 * 10^4
//     0 <= ops.length <= 10^4
//     ops[i].length == 2
//     1 <= ai <= m
//     1 <= bi <= n

import "fmt"

func maxCount(m int, n int, ops [][]int) int {
    if len(ops) == 0 {
        return m * n
    }
    min := func (x, y int) int { if x < y { return x; }; return y; }
    row, col := 1 << 32 - 1, 1 << 32 - 1
    for _, v := range ops {
        row = min(row,v[0])
        col = min(col,v[1])
    }
    return row * col
}

func maxCount1(m int, n int, ops [][]int) int {
    for _, v := range ops {
        if m > v[0] {
            m = v[0]
        }
        if n > v[1] {
            n = v[1]
        }
    }
    return m * n
}

func main() {
    // Example 1:
    // <img src="https://assets.leetcode.com/uploads/2020/10/02/ex1.jpg" />
    // Input: m = 3, n = 3, ops = [[2,2],[3,3]]
    // Output: 4
    // Explanation: The maximum integer in M is 2, and there are four of it in M. So return 4.
    fmt.Println(maxCount(3,3,[][]int{{2,2},{3,3}})) // 4
    // Example 2:
    // Input: m = 3, n = 3, ops = [[2,2],[3,3],[3,3],[3,3],[2,2],[3,3],[3,3],[3,3],[2,2],[3,3],[3,3],[3,3]]
    // Output: 4
    fmt.Println(maxCount(3,3,[][]int{{2,2},{3,3},{3,3},{3,3},{2,2},{3,3},{3,3},{3,3},{2,2},{3,3},{3,3},{3,3}})) // 4
    // Example 3:
    // Input: m = 3, n = 3, ops = []
    // Output: 9
    fmt.Println(maxCount(3,3,[][]int{})) // 9

    fmt.Println(maxCount1(3,3,[][]int{{2,2},{3,3}})) // 4
    fmt.Println(maxCount1(3,3,[][]int{{2,2},{3,3},{3,3},{3,3},{2,2},{3,3},{3,3},{3,3},{2,2},{3,3},{3,3},{3,3}})) // 4
    fmt.Println(maxCount1(3,3,[][]int{})) // 9
}