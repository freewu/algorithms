package main

// 885. Spiral Matrix III
// You start at the cell (rStart, cStart) of an rows x cols grid facing east. 
// The northwest corner is at the first row and column in the grid, and the southeast corner is at the last row and column.

// You will walk in a clockwise spiral shape to visit every position in this grid. 
// Whenever you move outside the grid's boundary, we continue our walk outside the grid (but may return to the grid boundary later.). 
// Eventually, we reach all rows * cols spaces of the grid.

// Return an array of coordinates representing the positions of the grid in the order you visited them.

// Example 1:
// <img src="https://s3-lc-upload.s3.amazonaws.com/uploads/2018/08/24/example_1.png" />
// Input: rows = 1, cols = 4, rStart = 0, cStart = 0
// Output: [[0,0],[0,1],[0,2],[0,3]]

// Example 2:
// <img src="https://s3-lc-upload.s3.amazonaws.com/uploads/2018/08/24/example_2.png" />
// Input: rows = 5, cols = 6, rStart = 1, cStart = 4
// Output: [[1,4],[1,5],[2,5],[2,4],[2,3],[1,3],[0,3],[0,4],[0,5],[3,5],[3,4],[3,3],[3,2],[2,2],[1,2],[0,2],[4,5],[4,4],[4,3],[4,2],[4,1],[3,1],[2,1],[1,1],[0,1],[4,0],[3,0],[2,0],[1,0],[0,0]]

// Constraints:
//     1 <= rows, cols <= 100
//     0 <= rStart < rows
//     0 <= cStart < cols

import "fmt"

func spiralMatrixIII(rows int, cols int, rStart int, cStart int) [][]int {
    valid := func(rows int, cols int, r0 int, c0 int) bool { // 边界检测
        if r0 < rows && c0 < cols && r0 >= 0 && c0 >= 0 { return true }
        return false
    }
    currRow, currCol, loop := rStart, cStart, 1
    res := [][]int{{ currRow, currCol }}
    for len(res) < rows * cols {
        for i := 0; i < loop; i++ {
            currCol++
            if valid(rows, cols, currRow, currCol) {
                res = append(res, []int{currRow, currCol})
            }
        }
        for i := 0; i < loop; i++ {
            currRow++
            if valid(rows, cols, currRow,currCol) {
                res = append(res, []int{currRow, currCol})
            }
        }
        loop++
        for i := 0; i < loop; i++ {
            currCol--
            if valid(rows, cols, currRow,currCol) {
                res = append(res, []int{currRow, currCol})
            }
        }
        for i := 0; i < loop; i++ {
            currRow--
            if valid(rows, cols, currRow,currCol) { 
                res = append(res, []int{currRow, currCol})
            }
        }   
        loop++
    }
    return res
}

func spiralMatrixIII1(rows int, cols int, rStart int, cStart int) [][]int {
    total := rows * cols
    order := [][]int{}
    index, lineCount, row, col := 0, 1, rStart, cStart
    dirIndex := 0
    dirs := [4][2]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}
    for index < total {
        steps := (lineCount + 1) / 2
        for i := 0; i < steps; i++ {
            if row >= 0 && row < rows && col >= 0 && col < cols {
                order = append(order, []int{row, col})
                index++
            }
            row += dirs[dirIndex][0]
            col += dirs[dirIndex][1]
        }
        lineCount++
        dirIndex = (dirIndex + 1) % 4
    }
    return order
}

func main() {
    // Example 1:
    // <img src="https://s3-lc-upload.s3.amazonaws.com/uploads/2018/08/24/example_1.png" />
    // Input: rows = 1, cols = 4, rStart = 0, cStart = 0
    // Output: [[0,0],[0,1],[0,2],[0,3]]
    fmt.Println(spiralMatrixIII(1,4,0,0)) // [[0,0],[0,1],[0,2],[0,3]]
    // Example 2:
    // <img src="https://s3-lc-upload.s3.amazonaws.com/uploads/2018/08/24/example_2.png" />
    // Input: rows = 5, cols = 6, rStart = 1, cStart = 4
    // Output: [[1,4],[1,5],[2,5],[2,4],[2,3],[1,3],[0,3],[0,4],[0,5],[3,5],[3,4],[3,3],[3,2],[2,2],[1,2],[0,2],[4,5],[4,4],[4,3],[4,2],[4,1],[3,1],[2,1],[1,1],[0,1],[4,0],[3,0],[2,0],[1,0],[0,0]]
    fmt.Println(spiralMatrixIII(5,6,1,4)) //  [[1,4],[1,5],[2,5],[2,4],[2,3],[1,3],[0,3],[0,4],[0,5],[3,5],[3,4],[3,3],[3,2],[2,2],[1,2],[0,2],[4,5],[4,4],[4,3],[4,2],[4,1],[3,1],[2,1],[1,1],[0,1],[4,0],[3,0],[2,0],[1,0],[0,0]]

    fmt.Println(spiralMatrixIII1(1,4,0,0)) // [[0,0],[0,1],[0,2],[0,3]]
    fmt.Println(spiralMatrixIII1(5,6,1,4)) //  [[1,4],[1,5],[2,5],[2,4],[2,3],[1,3],[0,3],[0,4],[0,5],[3,5],[3,4],[3,3],[3,2],[2,2],[1,2],[0,2],[4,5],[4,4],[4,3],[4,2],[4,1],[3,1],[2,1],[1,1],[0,1],[4,0],[3,0],[2,0],[1,0],[0,0]]
}