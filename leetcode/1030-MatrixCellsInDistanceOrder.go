package main

// 1030. Matrix Cells in Distance Order
// You are given four integers row, cols, rCenter, and cCenter. T
// here is a rows x cols matrix and you are on the cell with the coordinates (rCenter, cCenter).

// Return the coordinates of all cells in the matrix, 
// sorted by their distance from (rCenter, cCenter) from the smallest distance to the largest distance. 
// You may return the answer in any order that satisfies this condition.

// The distance between two cells (r1, c1) and (r2, c2) is |r1 - r2| + |c1 - c2|.

// Example 1:
// Input: rows = 1, cols = 2, rCenter = 0, cCenter = 0
// Output: [[0,0],[0,1]]
// Explanation: The distances from (0, 0) to other cells are: [0,1]

// Example 2:
// Input: rows = 2, cols = 2, rCenter = 0, cCenter = 1
// Output: [[0,1],[0,0],[1,1],[1,0]]
// Explanation: The distances from (0, 1) to other cells are: [0,1,1,2]
// The answer [[0,1],[1,1],[0,0],[1,0]] would also be accepted as correct.

// Example 3:
// Input: rows = 2, cols = 3, rCenter = 1, cCenter = 2
// Output: [[1,2],[0,2],[1,1],[0,1],[1,0],[0,0]]
// Explanation: The distances from (1, 2) to other cells are: [0,1,1,2,2,3]
// There are other answers that would also be accepted as correct, such as [[1,2],[1,1],[0,2],[1,0],[0,1],[0,0]].

// Constraints:
//     1 <= rows, cols <= 100
//     0 <= rCenter < rows
//     0 <= cCenter < cols

import "fmt"

func allCellsDistOrder(rows int, cols int, rCenter int, cCenter int) [][]int {
    buckets, res := [][][]int{}, [][]int{}
    for i := 0; i < (rows + cols - 1); i++ {
        buckets = append(buckets, [][]int{})
    }
    abs := func(x int) int { if x < 0 { return -x; }; return x; }
    for r := 0; r < rows; r++ {
        for c := 0; c < cols; c++ {
            distance := abs(r - rCenter) + abs(c - cCenter)
            buckets[distance] = append(buckets[distance], []int{r, c}) // 单元格(r1, c1) 和 (r2, c2) 之间的距离为|r1 - r2| + |c1 - c2|
        }
    }
    for _, bucket := range buckets {
        res = append(res, bucket...)
    }
    return res
}


func allCellsDistOrder1(rows, cols, rCenter, cCenter int) [][]int {
    directions := [][2]int{{1, 1}, {1, -1}, {-1, -1}, {-1, 1}}
    max := func (x, y int) int { if x > y { return x; }; return y; }
    res, row, col := make([][]int, 1, rows * cols) , rCenter, cCenter
    res[0] = []int{ rCenter, cCenter }
    maxDist := max(rCenter, rows - 1 - rCenter) + max(cCenter, cols - 1 - cCenter)
    for dist := 1; dist <= maxDist; dist++ {
        row--
        for i, dir := range directions {
            for i % 2 == 0 && row != rCenter || i % 2 == 1 && col != cCenter {
                if 0 <= row && row < rows && 0 <= col && col < cols { // 边界检测
                    res = append(res, []int{row, col})
                }
                row += dir[0]
                col += dir[1]
            }
        }
    }
    return res
}

func main() {
    // Example 1:
    // Input: rows = 1, cols = 2, rCenter = 0, cCenter = 0
    // Output: [[0,0],[0,1]]
    // Explanation: The distances from (0, 0) to other cells are: [0,1]
    fmt.Println(allCellsDistOrder(1,2,0,0)) // [[0,0],[0,1]]
    // Example 2:
    // Input: rows = 2, cols = 2, rCenter = 0, cCenter = 1
    // Output: [[0,1],[0,0],[1,1],[1,0]]
    // Explanation: The distances from (0, 1) to other cells are: [0,1,1,2]
    // The answer [[0,1],[1,1],[0,0],[1,0]] would also be accepted as correct.
    fmt.Println(allCellsDistOrder(2,2,0,1)) // [[0,1],[0,0],[1,1],[1,0]]
    // Example 3:
    // Input: rows = 2, cols = 3, rCenter = 1, cCenter = 2
    // Output: [[1,2],[0,2],[1,1],[0,1],[1,0],[0,0]]
    // Explanation: The distances from (1, 2) to other cells are: [0,1,1,2,2,3]
    // There are other answers that would also be accepted as correct, such as [[1,2],[1,1],[0,2],[1,0],[0,1],[0,0]].
    fmt.Println(allCellsDistOrder(2,3,1,2)) // [[1,2],[0,2],[1,1],[0,1],[1,0],[0,0]]

    //fmt.Println(allCellsDistOrder(100,100,0,0)) // [[5 6] [3 4]]

    fmt.Println(allCellsDistOrder1(1,2,0,0)) // [[0,0],[0,1]]
    fmt.Println(allCellsDistOrder1(2,2,0,1)) // [[0,1],[0,0],[1,1],[1,0]]
    fmt.Println(allCellsDistOrder1(2,3,1,2)) // [[1,2],[0,2],[1,1],[0,1],[1,0],[0,0]]
}