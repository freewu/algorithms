package main

// 3242. Design Neighbor Sum Service
// You are given a n x n 2D array grid containing distinct elements in the range [0, n2 - 1].

// Implement the NeighborSum class:
//     NeighborSum(int [][]grid) 
//         initializes the object.
//     int adjacentSum(int value) 
//         returns the sum of elements which are adjacent neighbors of value, 
//         that is either to the top, left, right, or bottom of value in grid.
//     int diagonalSum(int value) 
//         returns the sum of elements which are diagonal neighbors of value, 
//         that is either to the top-left, top-right, bottom-left, or bottom-right of value in grid.

// <img src="https://assets.leetcode.com/uploads/2024/06/24/design.png" />

// Example 1:
// Input:
// ["NeighborSum", "adjacentSum", "adjacentSum", "diagonalSum", "diagonalSum"]
// [[[[0, 1, 2], [3, 4, 5], [6, 7, 8]]], [1], [4], [4], [8]]
// Output: [null, 6, 16, 16, 4]
// Explanation:
// <img src="https://assets.leetcode.com/uploads/2024/06/24/designexample0.png" />
// The adjacent neighbors of 1 are 0, 2, and 4.
// The adjacent neighbors of 4 are 1, 3, 5, and 7.
// The diagonal neighbors of 4 are 0, 2, 6, and 8.
// The diagonal neighbor of 8 is 4.

// Example 2:
// Input:
// ["NeighborSum", "adjacentSum", "diagonalSum"]
// [[[[1, 2, 0, 3], [4, 7, 15, 6], [8, 9, 10, 11], [12, 13, 14, 5]]], [15], [9]]
// Output: [null, 23, 45]
// Explanation:
// <img src="https://assets.leetcode.com/uploads/2024/06/24/designexample2.png" />
// The adjacent neighbors of 15 are 0, 10, 7, and 6.
// The diagonal neighbors of 9 are 4, 12, 14, and 15.

// Constraints:
//     3 <= n == grid.length == grid[0].length <= 10
//     0 <= grid[i][j] <= n2 - 1
//     All grid[i][j] are distinct.
//     value in adjacentSum and diagonalSum will be in the range [0, n2 - 1].
//     At most 2 * n2 calls will be made to adjacentSum and diagonalSum.

import "fmt"

type NeighborSum  struct {
    grid [][]int
}

func Constructor(grid [][]int) NeighborSum  {
    return NeighborSum { grid }
}

func (this *NeighborSum) AdjacentSum(value int) int {
    r, c := -1, -1
    for i := range this.grid {
        for j := range this.grid[i] {
            if this.grid[i][j] == value { // 找到值所在的位置
                r, c = i, j
            }
        }
    }
    res, n := 0, len(this.grid)
    directions := [][]int{{-1, 0}, {0, 1}, {1, 0}, {0, -1}} // left right up down
    for _, dir := range directions {
        row, col := r + dir[0], c + dir[1]
        if row >= 0 && row < n && col >= 0 && col < n {
            res +=  this.grid[row][col]
        }
    }
    return res
}

func (this *NeighborSum) DiagonalSum(value int) int {
    r, c := -1, -1
    for i := range this.grid {
        for j := range this.grid[i] {
            if this.grid[i][j] == value { // 找到值所在的位置
                r, c = i, j
            }
        }
    }
    res, n := 0, len(this.grid)
    directions := [][]int{{-1, -1}, {-1, 1}, {1, 1}, {1, -1}} // 四个角
    for _, dir := range directions {
        row, col := r + dir[0], c + dir[1]
        if row >= 0 && row < n && col >= 0 && col < n {
            res +=  this.grid[row][col]
        }
    }
    return res
}


type NeighborSum1 struct {
    grid [][]int
    l    int
    r    int
}

func Constructor1(grid [][]int) NeighborSum1 {
    return NeighborSum1{grid, len(grid), len(grid[0])}
}

func (this *NeighborSum1) AdjacentSum(value int) int {
    for i := 0; i < this.l; i++ {
        for j := 0; j < this.r; j++ {
            sum := 0
            if this.grid[i][j] == value {
                if i > 0 { // 左
                    sum += this.grid[i-1][j]
                }
                if i < this.l-1 { // 右
                    sum += this.grid[i+1][j]
                }
                if j > 0 { // 上
                    sum += this.grid[i][j-1]
                }
                if j < this.r-1 { // 下
                    sum += this.grid[i][j+1]
                }
                return sum
            }
        }
    }
    return 0
}

func (this *NeighborSum1) DiagonalSum(value int) int {
    for i := 0; i < this.l; i++ {
        for j := 0; j < this.r; j++ {
            sum := 0
            if this.grid[i][j] == value {
                if i > 0 && j > 0 { // 左上角
                    sum += this.grid[i-1][j-1]
                }
                if i < this.l-1 && j > 0 { // 右上角
                    sum += this.grid[i+1][j-1]
                }
                if i > 0 && j < this.r-1 { // 左下角
                    sum += this.grid[i-1][j+1]
                }
                if i < this.l-1 && j < this.r-1 { // 右下角
                    sum += this.grid[i+1][j+1]
                }
                return sum
            }
        }
    }
    return 0
}

/**
 * Your NeighborSum object will be instantiated and called as such:
 * obj := Constructor(grid);
 * param_1 := obj.AdjacentSum(value);
 * param_2 := obj.DiagonalSum(value);
 */

func main() {
    // Example 1:
    // Input:
    // ["NeighborSum", "adjacentSum", "adjacentSum", "diagonalSum", "diagonalSum"]
    // [[[[0, 1, 2], [3, 4, 5], [6, 7, 8]]], [1], [4], [4], [8]]
    // Output: [null, 6, 16, 16, 4]
    // Explanation:
    // <img src="https://assets.leetcode.com/uploads/2024/06/24/designexample0.png" />
    // The adjacent neighbors of 1 are 0, 2, and 4.
    // The adjacent neighbors of 4 are 1, 3, 5, and 7.
    // The diagonal neighbors of 4 are 0, 2, 6, and 8.
    // The diagonal neighbor of 8 is 4.
    obj1 := Constructor([][]int{{0, 1, 2}, {3, 4, 5}, {6, 7, 8}})
    fmt.Println(obj1)
    fmt.Println(obj1.AdjacentSum(1)) // 6
    fmt.Println(obj1.AdjacentSum(4)) // 16
    fmt.Println(obj1.DiagonalSum(4)) // 16
    fmt.Println(obj1.DiagonalSum(8)) // 4
    // Example 2:
    // Input:
    // ["NeighborSum", "adjacentSum", "diagonalSum"]
    // [[[[1, 2, 0, 3], [4, 7, 15, 6], [8, 9, 10, 11], [12, 13, 14, 5]]], [15], [9]]
    // Output: [null, 23, 45]
    // Explanation:
    // <img src="https://assets.leetcode.com/uploads/2024/06/24/designexample2.png" />
    // The adjacent neighbors of 15 are 0, 10, 7, and 6.
    // The diagonal neighbors of 9 are 4, 12, 14, and 15.
    obj2 := Constructor([][]int{{1, 2, 0, 3}, {4, 7, 15, 6}, {8, 9, 10, 11}, {12, 13, 14, 5}})
    fmt.Println(obj2)
    fmt.Println(obj2.AdjacentSum(15)) // 23
    fmt.Println(obj2.DiagonalSum(9)) // 45

    obj11 := Constructor1([][]int{{0, 1, 2}, {3, 4, 5}, {6, 7, 8}})
    fmt.Println(obj11)
    fmt.Println(obj11.AdjacentSum(1)) // 6
    fmt.Println(obj11.AdjacentSum(4)) // 16
    fmt.Println(obj11.DiagonalSum(4)) // 16
    fmt.Println(obj11.DiagonalSum(8)) // 4
    obj12 := Constructor1([][]int{{1, 2, 0, 3}, {4, 7, 15, 6}, {8, 9, 10, 11}, {12, 13, 14, 5}})
    fmt.Println(obj12)
    fmt.Println(obj12.AdjacentSum(15)) // 23
    fmt.Println(obj12.DiagonalSum(9)) // 45
}