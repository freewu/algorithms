package main

// 1536. Minimum Swaps to Arrange a Binary Grid
// Given an n x n binary grid, in one step you can choose two adjacent rows of the grid and swap them.

// A grid is said to be valid if all the cells above the main diagonal are zeros.

// Return the minimum number of steps needed to make the grid valid, or -1 if the grid cannot be valid.

// The main diagonal of a grid is the diagonal that starts at cell (1, 1) and ends at cell (n, n).

// Example 1:
// <img src="https://assets.leetcode.com/uploads/2020/07/28/fw.jpg" />
// Input: grid = [[0,0,1],[1,1,0],[1,0,0]]
// Output: 3

// Example 2:
// <img src="https://assets.leetcode.com/uploads/2020/07/16/e2.jpg" />
// Input: grid = [[0,1,1,0],[0,1,1,0],[0,1,1,0],[0,1,1,0]]
// Output: -1
// Explanation: All rows are similar, swaps have no effect on the grid.

// Example 3:
// <img src="https://assets.leetcode.com/uploads/2020/07/16/e3.jpg" />
// Input: grid = [[1,0,0],[1,1,0],[1,1,1]]
// Output: 0

// Constraints:
//     n == grid.length == grid[i].length
//     1 <= n <= 200
//     grid[i][j] is either 0 or 1

import "fmt"

func minSwaps(grid [][]int) int {
    n := len(grid)
    trailing_zeros := make([]int, n)
    for row, vect := range grid {
        count := 0
        for col := n-1; col>=0 && vect[col]==0; col-- {
            count++
        }
        trailing_zeros[row] = count
    }
    // bookkeeping for the rows:
    //  = 0 - stays at the current position
    //  = m - was pushed m positions down
    //  =-1 - is already used in a parent subproblem
    moved := make([]int,n)
    var helper func(trailing_zeros []int, moved []int, objective int) int
    helper = func(trailing_zeros []int, moved []int, objective int) int {
        if objective == 0 { return 0 } // end of recursion
        swaps := -1 // will stay negative if no good row is found
        for row, v := range trailing_zeros {
            if moved[row] < 0 { continue } // fast forward if current row was already placed to its final destination sometime before
            if v < objective { // push down the row if it doesn't satisfy its objective,
                moved[row]++ // a more lucky row from below would have to be dragged over this one
            } else { // otherwise count swaps and move it out of the way
                swaps = row + moved[row] - (len(moved)-1 - objective) // len(moved) == N
                moved[row] = -1 // used, not needed anymore
                break
            }
        }
        if swaps < 0 {  // no matching row was found?
            return -1
        }
        one_less := helper(trailing_zeros, moved, objective - 1)
        if one_less < 0 {
            return -1
        }
        return one_less + swaps
    }
    return helper(trailing_zeros, moved, n-1)
}

func minSwaps1(grid [][]int) int {
    res, n := 0, len(grid)
    pos := make([]int, n)
    for i := 0; i < n; i++ {
        for j := 0; j < n; j++ {
            if grid[i][j] == 1 {
                pos[i] = j
            }
        }
    }
    for r := 0; r < n; r++ {
        find := -1
        for i := r; i < n; i++ {
            if pos[i] <= r {
                find = i
                break
            }
        }
        if find == -1 {
            return -1
        }
        res += find - r
        for i := find; i > r; i-- {
            pos[i] = pos[i-1]
        }
    }
    return res
}

func main() {
    // Example 1:
    // <img src="https://assets.leetcode.com/uploads/2020/07/28/fw.jpg" />
    // Input: grid = [[0,0,1],[1,1,0],[1,0,0]]
    // Output: 3
    fmt.Println(minSwaps([][]int{{0,0,1},{1,1,0},{1,0,0}})) // 3
    // Example 2:
    // <img src="https://assets.leetcode.com/uploads/2020/07/16/e2.jpg" />
    // Input: grid = [[0,1,1,0],[0,1,1,0],[0,1,1,0],[0,1,1,0]]
    // Output: -1
    // Explanation: All rows are similar, swaps have no effect on the grid.
    fmt.Println(minSwaps([][]int{{0,1,1,0},{0,1,1,0},{0,1,1,0},{0,1,1,0}})) // -1
    // Example 3:
    // <img src="https://assets.leetcode.com/uploads/2020/07/16/e3.jpg" />
    // Input: grid = [[1,0,0],[1,1,0],[1,1,1]]
    // Output: 0
    fmt.Println(minSwaps([][]int{{1,0,0},{1,1,0},{1,1,1}})) // 0

    fmt.Println(minSwaps1([][]int{{0,0,1},{1,1,0},{1,0,0}})) // 3
    fmt.Println(minSwaps1([][]int{{0,1,1,0},{0,1,1,0},{0,1,1,0},{0,1,1,0}})) // -1
    fmt.Println(minSwaps1([][]int{{1,0,0},{1,1,0},{1,1,1}})) // 0
}