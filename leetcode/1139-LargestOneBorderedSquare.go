package main

// 1139. Largest 1-Bordered Square
// Given a 2D grid of 0s and 1s, 
// return the number of elements in the largest square subgrid that has all 1s on its border, 
// or 0 if such a subgrid doesn't exist in the grid.

// Example 1:
// Input: grid = [[1,1,1],[1,0,1],[1,1,1]]
// Output: 9

// Example 2:
// Input: grid = [[1,1,0,0]]
// Output: 1

// Constraints:
//     1 <= grid.length <= 100
//     1 <= grid[0].length <= 100
//     grid[i][j] is 0 or 1

import "fmt"

func largest1BorderedSquare(grid [][]int) int {
    res, m, n := 0, len(grid), len(grid[0])
    horizontal, vertical := make([][]int, m), make([][]int, m)
    for i := range horizontal {
        horizontal[i] = make([]int, n)
    }
    for i := range vertical {
        vertical[i] = make([]int, n)
    }
    for i := 0; i < m; i++ {
        for j := 0; j < n; j++ {
            if grid[i][j] == 1 {
                if i == 0 {
                    vertical[i][j] = 1
                } else {
                    vertical[i][j] = vertical[i-1][j] + 1
                }
                if j == 0 {
                    horizontal[i][j] = 1
                } else {
                    horizontal[i][j] = horizontal[i][j-1] + 1
                }
            }
        }
    }
    min := func (x, y int) int { if x < y { return x; }; return y; }
    for i := m-1; i >= 0; i-- {
        for j := n-1; j >= 0; j-- {
            mn := min(horizontal[i][j], vertical[i][j])
            for mn > res {
                if horizontal[i - mn + 1][j] >= mn && vertical[i][j - mn + 1] >= mn {
                    res = mn
                }
                mn--
            }
        }
    }
    return res * res
}

func largest1BorderedSquare1(grid [][]int) int {
    if len(grid) == 0 || len(grid[0]) == 0 {
        return 0
    }
    res, m, n := 0, len(grid), len(grid[0])
    left, up := make([][]int, m), make([][]int, m)
    min := func (x, y int) int { if x < y { return x; }; return y; }
    for i := 0; i < m; i++ {
        left[i], up[i] = make([]int, n), make([]int, n)
        for j := 0; j < n; j++ {
            if grid[i][j] == 0 { continue }
            if j == 0 {
                left[i][j] = 1
            } else {
                left[i][j] = left[i][j-1] + 1
            }
            if i == 0 {
                up[i][j] = 1
            } else {
                up[i][j] = up[i-1][j] + 1
            }
            border := min(left[i][j], up[i][j])
            for border > res {
                if min(left[i-border+1][j], up[i][j-border+1]) >= border {
                    res = border
                    break
                }
                border--
            }
        }
    }
    return res * res
}


func main() {
    // Example 1:
    // Input: grid = [[1,1,1],[1,0,1],[1,1,1]]
    // Output: 9
    fmt.Println(largest1BorderedSquare([][]int{{1,1,1},{1,0,1},{1,1,1}})) // 9
    // Example 2:
    // Input: grid = [[1,1,0,0]]
    // Output: 1
    fmt.Println(largest1BorderedSquare([][]int{{1,1,0,0}})) // 1

    fmt.Println(largest1BorderedSquare1([][]int{{1,1,1},{1,0,1},{1,1,1}})) // 9
    fmt.Println(largest1BorderedSquare1([][]int{{1,1,0,0}})) // 1
}