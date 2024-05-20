package main

// 2352. Equal Row and Column Pairs
// Given a 0-indexed n x n integer matrix grid, 
// return the number of pairs (ri, cj) such that row ri and column cj are equal.
// A row and column pair is considered equal if they contain the same elements in the same order (i.e., an equal array).

// Example 1:
// <img src="https://assets.leetcode.com/uploads/2022/06/01/ex1.jpg" />
// Input: grid = [[3,2,1],[1,7,6],[2,7,7]]
// Output: 1
// Explanation: There is 1 equal row and column pair:
// - (Row 2, Column 1): [2,7,7]

// Example 2:
// <img src="https://assets.leetcode.com/uploads/2022/06/01/ex2.jpg" />
// Input: grid = [[3,1,2,2],[1,4,4,5],[2,4,2,2],[2,4,2,2]]
// Output: 3
// Explanation: There are 3 equal row and column pairs:
// - (Row 0, Column 0): [3,1,2,2]
// - (Row 2, Column 2): [2,4,2,2]
// - (Row 3, Column 2): [2,4,2,2]
 
// Constraints:
//     n == grid.length == grid[i].length
//     1 <= n <= 200
//     1 <= grid[i][j] <= 10^5

import "fmt"
import "strconv"

func equalPairs(grid [][]int) int {
    count,rowMap := 0, make(map[string]int)
    for i := 0; i < len(grid); i++ {
        rowStr := ""
        for j := 0; j < len(grid); j++ {
            rowStr = rowStr + strconv.Itoa(grid[i][j]) + ","
        }
        rowMap[rowStr]++
    }
    for i := 0; i < len(grid); i++ {
        colStr := ""
        for j := 0; j < len(grid); j++ {
            colStr = colStr + strconv.Itoa(grid[j][i]) + ","
        }
        count = count + rowMap[colStr]
    }
    return count
}

func equalPairs1(grid [][]int) int {
    res, mapRow, row, col := 0,map[[200]int]int{}, len(grid), len(grid[0])
    for i := 0; i < row; i++ {
        key := [200]int{}
        for j :=0; j < col; j++ {
            key[j] = grid[i][j]
        }
        mapRow[key]++
    }
    for i := 0; i < col; i++ {
        key := [200]int{}
        for j := 0; j < row; j++ {
            key[j] = grid[j][i]
        }
        res += mapRow[key]
    }
    return res
}

func main() {
    // Example 1:
    // <img src="https://assets.leetcode.com/uploads/2022/06/01/ex1.jpg" />
    // Input: grid = [[3,2,1],[1,7,6],[2,7,7]]
    // Output: 1
    // Explanation: There is 1 equal row and column pair:
    // - (Row 2, Column 1): [2,7,7]
    fmt.Println(equalPairs([][]int{{3,2,1},{1,7,6},{2,7,7}})) // 1
    // Example 2:
    // <img src="https://assets.leetcode.com/uploads/2022/06/01/ex2.jpg" />
    // Input: grid = [[3,1,2,2],[1,4,4,5],[2,4,2,2],[2,4,2,2]]
    // Output: 3
    // Explanation: There are 3 equal row and column pairs:
    // - (Row 0, Column 0): [3,1,2,2]
    // - (Row 2, Column 2): [2,4,2,2]
    // - (Row 3, Column 2): [2,4,2,2]
    fmt.Println(equalPairs([][]int{{3,1,2,2},{1,4,4,5},{2,4,2,2},{2,4,2,2}})) // 3

    fmt.Println(equalPairs1([][]int{{3,2,1},{1,7,6},{2,7,7}})) // 1
    fmt.Println(equalPairs1([][]int{{3,1,2,2},{1,4,4,5},{2,4,2,2},{2,4,2,2}})) // 3
}