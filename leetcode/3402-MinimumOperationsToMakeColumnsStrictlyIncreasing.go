package main

// 3402. Minimum Operations to Make Columns Strictly Increasing
// You are given a m x n matrix grid consisting of non-negative integers.

// In one operation, you can increment the value of any grid[i][j] by 1.

// Return the minimum number of operations needed to make all columns of grid strictly increasing.

// Example 1:
// Input: grid = [[3,2],[1,3],[3,4],[0,1]]
// Output: 15
// Explanation:
// To make the 0th column strictly increasing, we can apply 3 operations on grid[1][0], 2 operations on grid[2][0], and 6 operations on grid[3][0].
// To make the 1st column strictly increasing, we can apply 4 operations on grid[3][1].
// <img src="https://assets.leetcode.com/uploads/2024/11/10/firstexample.png" />

// Example 2:
// Input: grid = [[3,2,1],[2,1,0],[1,2,3]]
// Output: 12
// Explanation:
// To make the 0th column strictly increasing, we can apply 2 operations on grid[1][0], and 4 operations on grid[2][0].
// To make the 1st column strictly increasing, we can apply 2 operations on grid[1][1], and 2 operations on grid[2][1].
// To make the 2nd column strictly increasing, we can apply 2 operations on grid[1][2].
// <img src="https://assets.leetcode.com/uploads/2024/11/10/secondexample.png" />

// Constraints:
//     m == grid.length
//     n == grid[i].length
//     1 <= m, n <= 50
//     0 <= grid[i][j] < 2500
 
import "fmt"

func minimumOperations(grid [][]int) int {
    res := 0
    for i := 1; i < len(grid); i++ {
        for j := 0; j <len(grid[i]); j++ {
            if grid[i-1][j] >= grid[i][j] {
                op := grid[i-1][j] - grid[i][j] + 1
                res += op
                grid[i][j] += op
            }
        } 
    }  
    return res
}

func minimumOperations1(grid [][]int) int {
    res, m, n := 0, len(grid), len(grid[0])
    if m == 0 { return 0 }
    abs := func(x int) int { if x < 0 { return -x; }; return x; }
    for i := 1; i < m; i++ {
        for j := 0; j < n; j++ {
            if grid[i][j] - grid[i-1][j] <= 0 {
                res += abs(grid[i][j] - grid[i-1][j]) + 1
                grid[i][j] = grid[i-1][j] + 1
            }
        }
    }
    return res
}

func main() {
    // Example 1:
    // Input: grid = [[3,2],[1,3],[3,4],[0,1]]
    // Output: 15
    // Explanation:
    // To make the 0th column strictly increasing, we can apply 3 operations on grid[1][0], 2 operations on grid[2][0], and 6 operations on grid[3][0].
    // To make the 1st column strictly increasing, we can apply 4 operations on grid[3][1].
    // <img src="https://assets.leetcode.com/uploads/2024/11/10/firstexample.png" />
    fmt.Println(minimumOperations([][]int{{3,2},{1,3},{3,4},{0,1}})) // 15
    // Example 2:
    // Input: grid = [[3,2,1],[2,1,0],[1,2,3]]
    // Output: 12
    // Explanation:
    // To make the 0th column strictly increasing, we can apply 2 operations on grid[1][0], and 4 operations on grid[2][0].
    // To make the 1st column strictly increasing, we can apply 2 operations on grid[1][1], and 2 operations on grid[2][1].
    // To make the 2nd column strictly increasing, we can apply 2 operations on grid[1][2].
    // <img src="https://assets.leetcode.com/uploads/2024/11/10/secondexample.png" />2
    fmt.Println(minimumOperations([][]int{{3,2,1},{2,1,0},{1,2,3}})) // 12

    fmt.Println(minimumOperations1([][]int{{3,2},{1,3},{3,4},{0,1}})) // 15
    fmt.Println(minimumOperations1([][]int{{3,2,1},{2,1,0},{1,2,3}})) // 12
}