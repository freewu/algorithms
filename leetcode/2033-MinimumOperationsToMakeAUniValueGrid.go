package main

// 2033. Minimum Operations to Make a Uni-Value Grid
// You are given a 2D integer grid of size m x n and an integer x. 
// In one operation, you can add x to or subtract x from any element in the grid.

// A uni-value grid is a grid where all the elements of it are equal.

// Return the minimum number of operations to make the grid uni-value. 
// If it is not possible, return -1.

// Example 1:
// <img src="https://assets.leetcode.com/uploads/2021/09/21/gridtxt.png" />
// Input: grid = [[2,4],[6,8]], x = 2
// Output: 4
// Explanation: We can make every element equal to 4 by doing the following: 
// - Add x to 2 once.
// - Subtract x from 6 once.
// - Subtract x from 8 twice.
// A total of 4 operations were used.

// Example 2:
// <img src="https://assets.leetcode.com/uploads/2021/09/21/gridtxt-1.png" />
// Input: grid = [[1,5],[2,3]], x = 1
// Output: 5
// Explanation: We can make every element equal to 3.

// Example 3:
// <img src="https://assets.leetcode.com/uploads/2021/09/21/gridtxt-2.png" />
// Input: grid = [[1,2],[3,4]], x = 2
// Output: -1
// Explanation: It is impossible to make every element equal.

// Constraints:
//     m == grid.length
//     n == grid[i].length
//     1 <= m, n <= 10^5
//     1 <= m * n <= 10^5
//     1 <= x, grid[i][j] <= 10^4

import "fmt"
import "sort"

func minOperations(grid [][]int, x int) int {
    flat := make([]int, 0, len(grid) * len(grid[0]))
    for _, row := range grid { // 将 grid 转成一维数组
        for _, v := range row {
            if v % x != grid[0][0] % x { return -1 } // 取余不一样无法平衡
            flat = append(flat, v)
        }
    }
    sort.Ints(flat)
    if len(flat) <= 1 { return 0 }
    mid := len(flat) / 2
    min := func (x, y int) int { if x < y { return x; }; return y; }
    abs := func(x int) int { if x < 0 { return -x; }; return x; }
    count := func(x, target int, flat []int) int {
        res := 0
        for _, v := range flat { res += abs(target - v) / x }
        return res
    }
    return min(count(x, flat[mid], flat), count(x, flat[mid - 1], flat))
}

func main() {
    // Example 1:
    // <img src="https://assets.leetcode.com/uploads/2021/09/21/gridtxt.png" />
    // Input: grid = [[2,4],[6,8]], x = 2
    // Output: 4
    // Explanation: We can make every element equal to 4 by doing the following: 
    // - Add x to 2 once.
    // - Subtract x from 6 once.
    // - Subtract x from 8 twice.
    // A total of 4 operations were used.
    fmt.Println(minOperations([][]int{{2,4},{6,8}}, 2)) // 4
    // Example 2:
    // <img src="https://assets.leetcode.com/uploads/2021/09/21/gridtxt-1.png" />
    // Input: grid = [[1,5],[2,3]], x = 1
    // Output: 5
    // Explanation: We can make every element equal to 3.
    fmt.Println(minOperations([][]int{{1,5},{2,3}}, 1)) // 5
    // Example 3:
    // <img src="https://assets.leetcode.com/uploads/2021/09/21/gridtxt-2.png" />
    // Input: grid = [[1,2],[3,4]], x = 2
    // Output: -1
    // Explanation: It is impossible to make every element equal.
    fmt.Println(minOperations([][]int{{1,2},{3,4}}, 2)) // -1
}