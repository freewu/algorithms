package main

// 1260. Shift 2D Grid
// Given a 2D grid of size m x n and an integer k. You need to shift the grid k times.

// In one shift operation:
//     Element at grid[i][j] moves to grid[i][j + 1].
//     Element at grid[i][n - 1] moves to grid[i + 1][0].
//     Element at grid[m - 1][n - 1] moves to grid[0][0].

// Return the 2D grid after applying shift operation k times.

// Example 1:
// <img src="https://assets.leetcode.com/uploads/2019/11/05/e1.png" />
// Input: grid = [[1,2,3],[4,5,6],[7,8,9]], k = 1
// Output: [[9,1,2],[3,4,5],[6,7,8]]

// Example 2:
// <img src="https://assets.leetcode.com/uploads/2019/11/05/e2.png" />
// Input: grid = [[3,8,1,9],[19,7,2,5],[4,6,11,10],[12,0,21,13]], k = 4
// Output: [[12,0,21,13],[3,8,1,9],[19,7,2,5],[4,6,11,10]]

// Example 3:
// Input: grid = [[1,2,3],[4,5,6],[7,8,9]], k = 9
// Output: [[1,2,3],[4,5,6],[7,8,9]]

// Constraints:
//     m == grid.length
//     n == grid[i].length
//     1 <= m <= 50
//     1 <= n <= 50
//     -1000 <= grid[i][j] <= 1000
//     0 <= k <= 100

import "fmt"

func shiftGrid(grid [][]int, k int) [][]int {
    m, n, cur := len(grid), len(grid[0]), 0
    arr := make([]int, m * n)
    for i := 0; i < m; i++ {
        for j := 0; j < n; j++ {
            arr[cur] = grid[i][j]
            cur++
        }
    }
    if k > len(arr) { 
        k = k % len(arr) 
    }
    arr = append(arr[len(arr) - k:], arr[:len(arr) - k]...)
    cur = 0
    for i := 0; i < m; i++ {
        for j := 0; j < n; j++ {
            grid[i][j] = arr[cur]
            cur++
        }
    }
    return grid
}

func shiftGrid1(grid [][]int, k int) [][]int {
    m, n := len(grid), len(grid[0])
    product := m * n
    arr := make([]int, product)
    for i := range grid {
        for j := range grid[0] {
            arr[i * n+j] = grid[i][j]
        }
    }
    k = k % product
    arr = append(arr[product - k:], arr[:product - k]...)
    for i := range grid {
        for j := range grid[0] {
            grid[i][j] = arr[i * n + j]
        }
    }
    return grid
}

func shiftGrid2(grid [][]int, k int) [][]int {
    m, n, p := len(grid), len(grid[0]), 0
    product := m * n
    arr := make([]int, product)
    for i := range grid {
        for j := range grid[0] {
            arr[p] = grid[i][j]
            p++
        }
    }
    p = (-(k % product) + product) % product
    for i := range grid {
        for j := range grid[0] {
            grid[i][j] = arr[p]
            p = (p + 1) % product
        }
    }
    return grid
}

func main() {
    // Example 1:
    // <img src="https://assets.leetcode.com/uploads/2019/11/05/e1.png" />
    // Input: grid = [[1,2,3],[4,5,6],[7,8,9]], k = 1
    // Output: [[9,1,2],[3,4,5],[6,7,8]]
    fmt.Println(shiftGrid([][]int{{1,2,3},{4,5,6},{7,8,9}}, 1)) // [[9,1,2],[3,4,5],[6,7,8]]
    // Example 2:
    // <img src="https://assets.leetcode.com/uploads/2019/11/05/e2.png" />
    // Input: grid = [[3,8,1,9],[19,7,2,5],[4,6,11,10],[12,0,21,13]], k = 4
    // Output: [[12,0,21,13],[3,8,1,9],[19,7,2,5],[4,6,11,10]]
    fmt.Println(shiftGrid([][]int{{3,8,1,9},{19,7,2,5},{4,6,11,10},{12,0,21,13}}, 4)) // [[12,0,21,13],[3,8,1,9],[19,7,2,5],[4,6,11,10]]
    // Example 3:
    // Input: grid = [[1,2,3],[4,5,6],[7,8,9]], k = 9
    // Output: [[1,2,3],[4,5,6],[7,8,9]]
    fmt.Println(shiftGrid([][]int{{1,2,3},{4,5,6},{7,8,9}}, 9)) // [[1,2,3],[4,5,6],[7,8,9]]

    fmt.Println(shiftGrid1([][]int{{1,2,3},{4,5,6},{7,8,9}}, 1)) // [[9,1,2],[3,4,5],[6,7,8]]
    fmt.Println(shiftGrid1([][]int{{3,8,1,9},{19,7,2,5},{4,6,11,10},{12,0,21,13}}, 4)) // [[12,0,21,13],[3,8,1,9],[19,7,2,5],[4,6,11,10]]
    fmt.Println(shiftGrid1([][]int{{1,2,3},{4,5,6},{7,8,9}}, 9)) // [[1,2,3],[4,5,6],[7,8,9]]

    fmt.Println(shiftGrid2([][]int{{1,2,3},{4,5,6},{7,8,9}}, 1)) // [[9,1,2],[3,4,5],[6,7,8]]
    fmt.Println(shiftGrid2([][]int{{3,8,1,9},{19,7,2,5},{4,6,11,10},{12,0,21,13}}, 4)) // [[12,0,21,13],[3,8,1,9],[19,7,2,5],[4,6,11,10]]
    fmt.Println(shiftGrid2([][]int{{1,2,3},{4,5,6},{7,8,9}}, 9)) // [[1,2,3],[4,5,6],[7,8,9]]
}