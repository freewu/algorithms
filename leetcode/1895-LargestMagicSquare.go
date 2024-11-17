package main

// 1895. Largest Magic Square
// A k x k magic square is a k x k grid filled with integers 
// such that every row sum, every column sum, and both diagonal sums are all equal. 
// The integers in the magic square do not have to be distinct. 
// Every 1 x 1 grid is trivially a magic square.

// Given an m x n integer grid, 
// return the size (i.e., the side length k) of the largest magic square that can be found within this grid.

// Example 1:
// <img src="https://assets.leetcode.com/uploads/2021/05/29/magicsquare-grid.jpg" />
// Input: grid = [[7,1,4,5,6],[2,5,1,6,4],[1,5,4,3,2],[1,2,7,3,4]]
// Output: 3
// Explanation: The largest magic square has a size of 3.
// Every row sum, column sum, and diagonal sum of this magic square is equal to 12.
// - Row sums: 5+1+6 = 5+4+3 = 2+7+3 = 12
// - Column sums: 5+5+2 = 1+4+7 = 6+3+3 = 12
// - Diagonal sums: 5+4+3 = 6+4+2 = 12

// Example 2:
// <img src="https://assets.leetcode.com/uploads/2021/05/29/magicsquare2-grid.jpg" />
// Input: grid = [[5,1,3,1],[9,3,3,1],[1,3,3,8]]
// Output: 2

// Constraints:
//     m == grid.length
//     n == grid[i].length
//     1 <= m, n <= 50
//     1 <= grid[i][j] <= 10^6

import "fmt"

func largestMagicSquare(grid [][]int) int {
    m, n := len(grid), len(grid[0])
    rowsum, colsum := make([][]int, m + 1), make([][]int, m + 1)
    for i := 0; i <= m; i++ {
        rowsum[i], colsum[i] = make([]int, n + 1), make([]int, n + 1)
    }
    for i := 1; i < m + 1; i++ {
        for j := 1; j < n + 1; j++ {
            rowsum[i][j], colsum[i][j] = rowsum[i][j-1] + grid[i-1][j-1], colsum[i-1][j] + grid[i-1][j-1]
        }
    }
    check := func(x1, y1, x2, y2 int) bool {
        v := rowsum[x1 + 1][y2 + 1] - rowsum[x1 + 1][y1]
        for i := x1 + 1; i < x2 + 1; i++ {
            if rowsum[i+1][y2+1]- rowsum[i+1][y1] != v { return false }
        }
        for j := y1; j < y2+1; j++ {
            if colsum[x2+1][j+1]-colsum[x1][j+1] != v { return false }
        }
        sum := 0
        for i, j := x1, y1; i <= x2; i, j = i+1, j+1 {
            sum += grid[i][j]
        }
        if sum != v { return false }
        sum = 0
        for i, j := x1, y2; i <= x2; i, j = i+1, j-1 {
            sum += grid[i][j]
        }
        if sum != v { return false }
        return true
    }
    min := func (x, y int) int { if x < y { return x; }; return y; }
    for k := min(m, n); k > 1; k-- {
        for i := 0; i + k - 1 < m; i++ {
            for j := 0; j + k - 1 < n; j++ {
                if check(i, j, i + k - 1, j + k - 1) { return k }
            }
        }
    }
    return 1
}

func largestMagicSquare1(grid [][]int) int {
    m, n := len(grid), len(grid[0])
    min := func (x, y int) int { if x < y { return x; }; return y; }
    for i := min(m, n); i > 1; i-- {
        for j := 0; j + i <= m; j++ {
            for k := 0; k + i <= n; k++ {
                target, p, q := 0, j, k
                for ;p < j + i; p++ {
                    target += grid[p][q]
                    q++
                }
                sum := 0
                for p, q = j, k + i - 1; p < j + i; p++ {
                    sum += grid[p][q]
                    q--
                }
                if sum != target { continue }
                for p = j; p < j + i; p++ {
                    sum = 0
                    for q = k; q < k + i;q++ {
                        sum += grid[p][q]
                    }
                    if sum != target { break }
                }
                if p < j + i { continue }
                for p = k; p < k + i; p++ {
                    sum = 0
                    for q = j; q < j + i; q++ {
                        sum += grid[q][p]
                    }
                    if sum != target { break }
                }
                if p == k + i { return i }
            }
        }
    }
    return 1
}

func main() {
    // Example 1:
    // <img src="https://assets.leetcode.com/uploads/2021/05/29/magicsquare-grid.jpg" />
    // Input: grid = [[7,1,4,5,6],[2,5,1,6,4],[1,5,4,3,2],[1,2,7,3,4]]
    // Output: 3
    // Explanation: The largest magic square has a size of 3.
    // Every row sum, column sum, and diagonal sum of this magic square is equal to 12.
    // - Row sums: 5+1+6 = 5+4+3 = 2+7+3 = 12
    // - Column sums: 5+5+2 = 1+4+7 = 6+3+3 = 12
    // - Diagonal sums: 5+4+3 = 6+4+2 = 12
    fmt.Println(largestMagicSquare([][]int{{7,1,4,5,6},{2,5,1,6,4},{1,5,4,3,2},{1,2,7,3,4}})) // 3
    // Example 2:
    // <img src="https://assets.leetcode.com/uploads/2021/05/29/magicsquare2-grid.jpg" />
    // Input: grid = [[5,1,3,1],[9,3,3,1],[1,3,3,8]]
    // Output: 2
    fmt.Println(largestMagicSquare([][]int{{5,1,3,1},{9,3,3,1},{1,3,3,8}})) // 2

    fmt.Println(largestMagicSquare1([][]int{{7,1,4,5,6},{2,5,1,6,4},{1,5,4,3,2},{1,2,7,3,4}})) // 3
    fmt.Println(largestMagicSquare1([][]int{{5,1,3,1},{9,3,3,1},{1,3,3,8}})) // 2
}