package main

// 3446. Sort Matrix by Diagonals
// You are given an n x n square matrix of integers grid. Return the matrix such that:
//     1. The diagonals in the bottom-left triangle (including the middle diagonal) are sorted in non-increasing order.
//     2. The diagonals in the top-right triangle are sorted in non-decreasing order.

// Example 1:
// Input: grid = [[1,7,3],[9,8,2],[4,5,6]]
// Output: [[8,2,3],[9,6,7],[4,5,1]]
// Explanation:
// <img src="https://assets.leetcode.com/uploads/2024/12/29/4052example1drawio.png" />
// The diagonals with a black arrow (bottom-left triangle) should be sorted in non-increasing order:
// [1, 8, 6] becomes [8, 6, 1].
// [9, 5] and [4] remain unchanged.
// The diagonals with a blue arrow (top-right triangle) should be sorted in non-decreasing order:
// [7, 2] becomes [2, 7].
// [3] remains unchanged.

// Example 2:
// Input: grid = [[0,1],[1,2]]
// Output: [[2,1],[1,0]]
// Explanation:
// <img src="https://assets.leetcode.com/uploads/2024/12/29/4052example2adrawio.png" />
// The diagonals with a black arrow must be non-increasing, so [0, 2] is changed to [2, 0]. 
// The other diagonals are already in the correct order.

// Example 3:
// Input: grid = [[1]]
// Output: [[1]]
// Explanation:
// Diagonals with exactly one element are already in order, so no changes are needed.

// Constraints:
//     grid.length == grid[i].length == n
//     1 <= n <= 10
//     -10^5 <= grid[i][j] <= 10^5

import "fmt"
import "sort"
import "slices"

func sortMatrix(grid [][]int) [][]int {
    n := len(grid)
    sortList := make([]int, 0, n)
    for r := 0; r < n; r++ {
        sortList = sortList[0:0]
        for i, j := r, 0; i < n && j < n; i, j = i+1, j+1 {
            sortList = append(sortList, grid[i][j])
        }
        sort.Slice(sortList, func(i, j int)bool{
            return sortList[i] > sortList[j]
        })
        k := 0
        for i, j := r, 0; i < n && j < n; i, j = i+1, j+1 {
            grid[i][j] = sortList[k]
            k++
        }
    }
    for c := 1; c < n; c++ {
        sortList = sortList[0:0]
        for i, j := 0, c; i < n && j < n; i, j = i+1, j+1 {
            sortList = append(sortList, grid[i][j])
        }
        sort.Ints(sortList)
        k := 0
        for i, j := 0, c; i < n && j < n; i, j = i+1, j+1 {
            grid[i][j] = sortList[k]
            k++
        }
    }
    return grid
}

func sortMatrix1(grid [][]int) [][]int {
    n := len(grid)
    min := func (x, y int) int { if x < y { return x; }; return y; }
    max := func (x, y int) int { if x > y { return x; }; return y; }
    for k := 2; k < n * 2 - 1; k++ {
        mn, mx := max(n - k, 0), min(n * 2 - 1 - k, n - 1)
        arr := []int{}
        for j := mn; j <= mx; j++ {
            arr = append(arr, grid[k + j - n][j])
        }
        if mn > 0 {
            slices.Sort(arr)
        } else {
            slices.SortFunc(arr, func(a, b int) int {
                return b - a
            })
        }
        for j := mn; j <= mx; j++ {
            grid[k + j - n][j] = arr[j - mn]
        }
    }
    return grid
}

func main() {
    // Example 1:
    // Input: grid = [[1,7,3],[9,8,2],[4,5,6]]
    // Output: [[8,2,3],[9,6,7],[4,5,1]]
    // Explanation:
    // <img src="https://assets.leetcode.com/uploads/2024/12/29/4052example1drawio.png" />
    // The diagonals with a black arrow (bottom-left triangle) should be sorted in non-increasing order:
    // [1, 8, 6] becomes [8, 6, 1].
    // [9, 5] and [4] remain unchanged.
    // The diagonals with a blue arrow (top-right triangle) should be sorted in non-decreasing order:
    // [7, 2] becomes [2, 7].
    // [3] remains unchanged.
    fmt.Println(sortMatrix([][]int{{0,1},{1,2}})) // [[8,2,3],[9,6,7],[4,5,1]]
    // Example 2:
    // Input: grid = [[0,1],[1,2]]
    // Output: [[2,1],[1,0]]
    // Explanation:
    // <img src="https://assets.leetcode.com/uploads/2024/12/29/4052example2adrawio.png" />
    // The diagonals with a black arrow must be non-increasing, so [0, 2] is changed to [2, 0]. 
    // The other diagonals are already in the correct order.
    fmt.Println(sortMatrix([][]int{{1,7,3},{9,8,2},{4,5,6}})) // [[2,1],[1,0]]
    // Example 3:
    // Input: grid = [[1]]
    // Output: [[1]]
    // Explanation:
    // Diagonals with exactly one element are already in order, so no changes are needed.
    fmt.Println(sortMatrix([][]int{{1}})) // [[1]]

    fmt.Println(sortMatrix1([][]int{{0,1},{1,2}})) // [[8,2,3],[9,6,7],[4,5,1]]
    fmt.Println(sortMatrix1([][]int{{1,7,3},{9,8,2},{4,5,6}})) // [[2,1],[1,0]]
    fmt.Println(sortMatrix1([][]int{{1}})) // [[1]]
}