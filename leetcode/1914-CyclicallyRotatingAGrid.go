package main

// 1914. Cyclically Rotating a Grid
// You are given an m x n integer matrix grid​​​, where m and n are both even integers, and an integer k.
// The matrix is composed of several layers, which is shown in the below image, where each color is its own layer:
//     <img src="https://assets.leetcode.com/uploads/2021/06/10/ringofgrid.png" />

// A cyclic rotation of the matrix is done by cyclically rotating each layer in the matrix. 
// To cyclically rotate a layer once, each element in the layer will take the place of the adjacent element in the counter-clockwise direction. 
// An example rotation is shown below:
//     <img src="https://assets.leetcode.com/uploads/2021/06/22/explanation_grid.jpg" />

// Return the matrix after applying k cyclic rotations to it.

// Example 1:
// <img src="https://assets.leetcode.com/uploads/2021/06/19/rod2.png" />
// Input: grid = [[40,10],[30,20]], k = 1
// Output: [[10,20],[40,30]]
// Explanation: The figures above represent the grid at every state.

// Example 2:
// <img src="https://assets.leetcode.com/uploads/2021/06/10/ringofgrid5.png" />
// <img src="https://assets.leetcode.com/uploads/2021/06/10/ringofgrid6.png" />
// <img src="https://assets.leetcode.com/uploads/2021/06/10/ringofgrid7.png" />
// Input: grid = [[1,2,3,4],[5,6,7,8],[9,10,11,12],[13,14,15,16]], k = 2
// Output: [[3,4,8,12],[2,11,10,16],[1,7,6,15],[5,9,13,14]]
// Explanation: The figures above represent the grid at every state.

// Constraints:
//     m == grid.length
//     n == grid[i].length
//     2 <= m, n <= 50
//     Both m and n are even integers.
//     1 <= grid[i][j] <= 5000
//     1 <= k <= 10^9

import "fmt"

func rotateGrid(grid [][]int, k int) [][]int {
    min := func (x, y int) int { if x < y { return x; }; return y; }
    m, n := len(grid), len(grid[0])
    layers := min(m, n) / 2
    elements := m * 2 + n * 2 - 4
    buff := make([]int, elements * 2)

    getSetLayer := func (b []int, grid [][]int, s, m, n, size, k int, get bool) {
        for p := 0; p < size; p++ {
            if p < m { // left and going down, increase i
                if get {
                    b[p] = grid[s+p][s]
                } else {
                    grid[s+p][s] = b[p+size-k]
                }
            } else if p < m-1+n { // Down and going right, increase j
                if get {
                    b[p] = grid[s+m-1][s+p-m+1]
                } else {
                    grid[s+m-1][s+p-m+1] = b[p+size-k]
                }
            } else if p < 2*m+n-2 { // right and going up, decrease i
                if get {
                    b[p] = grid[s+2*m+n-3-p][s+n-1]
                } else {
                    grid[s+2*m+n-3-p][s+n-1] = b[p+size-k]
                }
            } else { // up and going left, descrease j
                if get {
                    b[p] = grid[s][s+2*m+2*n-4-p]
                } else {
                    grid[s][s+2*m+2*n-4-p] = b[p+size-k]
                }
            }
        }
        if get {
            //copy the end of array to make is circular
            for p := 0; p < size; p++ {
                b[p+size] = b[p]
            }
        }
    }

    for i := 0; i < layers; i++ {
        // Rotate each layer
        rot := k % elements
        getSetLayer(buff, grid, i, m, n, elements, rot, true)
        getSetLayer(buff, grid, i, m, n, elements, rot, false)
        m -= 2
        n -= 2
        elements = m * 2 + n * 2 - 4
    }
    return grid
}

func main() {
    // Example 1:
    // <img src="https://assets.leetcode.com/uploads/2021/06/19/rod2.png" />
    // Input: grid = [[40,10],[30,20]], k = 1
    // Output: [[10,20],[40,30]]
    // Explanation: The figures above represent the grid at every state.
    grid1 := [][]int{
        {40,10},
        {30,20},
    }
    fmt.Println(rotateGrid(grid1,1)) // [[10,20],[40,30]]
    // Example 2:
    // <img src="https://assets.leetcode.com/uploads/2021/06/10/ringofgrid5.png" />
    // <img src="https://assets.leetcode.com/uploads/2021/06/10/ringofgrid6.png" />
    // <img src="https://assets.leetcode.com/uploads/2021/06/10/ringofgrid7.png" />
    // Input: grid = [[1,2,3,4],[5,6,7,8],[9,10,11,12],[13,14,15,16]], k = 2
    // Output: [[3,4,8,12],[2,11,10,16],[1,7,6,15],[5,9,13,14]]
    // Explanation: The figures above represent the grid at every state.
    grid2 := [][]int{
        {1, 2,  3,  4},
        {5, 6,  7,  8},
        {9, 10, 11, 12},
        {13,14, 15, 16},
    }
    fmt.Println(rotateGrid(grid2, 2)) // [[3,4,8,12],[2,11,10,16],[1,7,6,15],[5,9,13,14]]
}