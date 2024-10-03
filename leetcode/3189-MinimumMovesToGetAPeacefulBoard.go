package main

// 3189. Minimum Moves to Get a Peaceful Board
// Given a 2D array rooks of length n, where rooks[i] = [xi, yi] indicates the position of a rook on an n x n chess board. 
// Your task is to move the rooks 1 cell at a time vertically or horizontally (to an adjacent cell) such that the board becomes peaceful.

// A board is peaceful if there is exactly one rook in each row and each column.

// Return the minimum number of moves required to get a peaceful board.

// Note that at no point can there be two rooks in the same cell.

// Example 1:
// Input: rooks = [[0,0],[1,0],[1,1]]
// Output: 3
// Explanation:
// <img src="https://assets.leetcode.com/uploads/2024/06/17/ex1-edited.gif" />

// Example 2:
// Input: rooks = [[0,0],[0,1],[0,2],[0,3]]
// Output: 6
// Explanation:
// <img src="https://assets.leetcode.com/uploads/2024/06/17/ex2-edited.gif" />

// Constraints:
//     1 <= n == rooks.length <= 500
//     0 <= xi, yi <= n - 1
//     The input is generated such that there are no 2 rooks in the same cell.

import "fmt"
import "slices"

func minMoves(rooks [][]int) int {
    res, n := 0, len(rooks)
    slices.SortFunc(rooks, func(a, b []int) int { 
        return a[0] - b[0]
    })
    abs := func(x int) int { if x < 0 { return -x; }; return x; }
    for i := 0; i < n; i++ {
        res += abs(i - rooks[i][0])
    }
    slices.SortFunc(rooks, func(a, b []int) int { 
        return a[1] - b[1]
    })
    for i := 0; i < n; i++ {
        res += abs(i - rooks[i][1])
    }
    return res
}

func main() {
    // Example 1:
    // Input: rooks = [[0,0],[1,0],[1,1]]
    // Output: 3
    // Explanation:
    // <img src="https://assets.leetcode.com/uploads/2024/06/17/ex1-edited.gif" />
    fmt.Println(minMoves([][]int{{0,0},{1,0},{1,1}})) // 3
    // Example 2:
    // Input: rooks = [[0,0],[0,1],[0,2],[0,3]]
    // Output: 6
    // Explanation:
    // <img src="https://assets.leetcode.com/uploads/2024/06/17/ex2-edited.gif" />
    fmt.Println(minMoves([][]int{{0,0},{0,1},{0,2},{0,3}})) // 6
}