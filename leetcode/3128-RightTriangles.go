package main

// 3128. Right Triangles
// You are given a 2D boolean matrix grid.
// Return an integer that is the number of right triangles that can be made with the 3 elements of grid such that all of them have a value of 1.

// Note:
//     A collection of 3 elements of grid is a right triangle if one of its elements is in the same row with another element and in the same column with the third element. 
//     The 3 elements do not have to be next to each other.

// Example 1:
// 0	1	0
// 0	1	1
// 0	1	0
// 0	1	0
// 0	1	1
// 0	1	0
// Input: grid = [[0,1,0],[0,1,1],[0,1,0]]
// Output: 2
// Explanation:
// There are two right triangles.

// Example 2:
// 1	0	0	0
// 0	1	0	1
// 1	0	0	0
// Input: grid = [[1,0,0,0],[0,1,0,1],[1,0,0,0]]
// Output: 0
// Explanation:
// There are no right triangles.

// Example 3:
// 1	0	1
// 1	0	0
// 1	0	0
// 1	0	1
// 1	0	0
// 1	0	0
// Input: grid = [[1,0,1],[1,0,0],[1,0,0]]
// Output: 2
// Explanation:
// There are two right triangles.

// Constraints:
//     1 <= grid.length <= 1000
//     1 <= grid[i].length <= 1000
//     0 <= grid[i][j] <= 1

import "fmt"

func numberOfRightTriangles(grid [][]int) int64 {
    rows, columns := len(grid), len(grid[0])
    horizontalFrequency, verticalFrequency := make([]int, columns), make([]int, rows)
    recordFrequencyOfOnes := func(horizontalFrequency []int, verticalFrequency []int, grid [][]int) {
        for r := 0; r < rows; r++ {
            for c := 0; c < columns; c++ {
                if grid[r][c] == 1 {
                    horizontalFrequency[c]++
                    verticalFrequency[r]++
                }
            }
        }
    }
    countNumberOfRightTriangles := func(horizontalFrequency []int, verticalFrequency []int, grid [][]int) int64 {
        var numberOfRightTriangles int64 = 0
        for r := 0; r < rows; r++ {
            for c := 0; c < columns; c++ {
                if grid[r][c] == 1 {
                    numberOfRightTriangles += (int64)((horizontalFrequency[c] - 1) * (verticalFrequency[r] - 1))
                }
            }
        }
        return numberOfRightTriangles
    }
    recordFrequencyOfOnes(horizontalFrequency, verticalFrequency, grid)
    return countNumberOfRightTriangles(horizontalFrequency, verticalFrequency, grid)
}

func numberOfRightTriangles1(grid [][]int) int64 {
    res, m, n := 0, len(grid), len(grid[0])
    rows, cols := make([]int, m), make([]int, n)
    for i := range rows {
        for j := range cols {
            rows[i] += grid[i][j]
            cols[j] += grid[i][j]
        }
    }
    for i := range rows {
        for j := range cols {
            if grid[i][j] == 1 {
                res += (rows[i] - 1) * (cols[j] - 1)
            }
        }
    }
    return int64(res)
}

func numberOfRightTriangles2(grid [][]int) int64 {
    res, m, n := 0, len(grid), len(grid[0])
    rowSum, colSum := make([]int, m), make([]int, n)
    for i, row := range grid {
        for j, x := range row {
            if x == 0 { continue }
            rowSum[i]++
            colSum[j]++
        }
    }
    for i, row := range grid {
        for j, x := range row {
            if x == 0 { continue }
            res += (rowSum[i] - 1) * (colSum[j] - 1)
        }
    }
    return int64(res)
}

func main() {
    // Example 1:
    // 0	1	0
    // 0	1	1
    // 0	1	0
    // 0	1	0
    // 0	1	1
    // 0	1	0
    // Input: grid = [[0,1,0],[0,1,1],[0,1,0]]
    // Output: 2
    // Explanation:
    // There are two right triangles.
    fmt.Println(numberOfRightTriangles([][]int{{0,1,0},{0,1,1},{0,1,0}})) // 2
    // Example 2:
    // 1	0	0	0
    // 0	1	0	1
    // 1	0	0	0
    // Input: grid = [[1,0,0,0],[0,1,0,1],[1,0,0,0]]
    // Output: 0
    // Explanation:
    // There are no right triangles.
    fmt.Println(numberOfRightTriangles([][]int{{1,0,0,0},{0,1,0,1},{1,0,0,0}})) // 0
    // Example 3:
    // 1	0	1
    // 1	0	0
    // 1	0	0
    // 1	0	1
    // 1	0	0
    // 1	0	0
    // Input: grid = [[1,0,1],[1,0,0],[1,0,0]]
    // Output: 2
    // Explanation:
    // There are two right triangles.
    fmt.Println(numberOfRightTriangles([][]int{{1,0,1},{1,0,0},{1,0,0}})) // 2

    fmt.Println(numberOfRightTriangles1([][]int{{0,1,0},{0,1,1},{0,1,0}})) // 2
    fmt.Println(numberOfRightTriangles1([][]int{{1,0,0,0},{0,1,0,1},{1,0,0,0}})) // 0
    fmt.Println(numberOfRightTriangles1([][]int{{1,0,1},{1,0,0},{1,0,0}})) // 2

    fmt.Println(numberOfRightTriangles2([][]int{{0,1,0},{0,1,1},{0,1,0}})) // 2
    fmt.Println(numberOfRightTriangles2([][]int{{1,0,0,0},{0,1,0,1},{1,0,0,0}})) // 0
    fmt.Println(numberOfRightTriangles2([][]int{{1,0,1},{1,0,0},{1,0,0}})) // 2
}