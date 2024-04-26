package main 

// 2639. Find the Width of Columns of a Grid
// You are given a 0-indexed m x n integer matrix grid. The width of a column is the maximum length of its integers.
//     For example, if grid = [[-10], [3], [12]], the width of the only column is 3 since -10 is of length 3.

// Return an integer array ans of size n where ans[i] is the width of the ith column.
// The length of an integer x with len digits is equal to len if x is non-negative, and len + 1 otherwise.

// Example 1:
// Input: grid = [[1],[22],[333]]
// Output: [3]
// Explanation: In the 0th column, 333 is of length 3.

// Example 2:
// Input: grid = [[-15,1,3],[15,7,12],[5,6,-2]]
// Output: [3,1,2]
// Explanation: 
// In the 0th column, only -15 is of length 3.
// In the 1st column, all integers are of length 1. 
// In the 2nd column, both 12 and -2 are of length 2.
 
// Constraints:
//     m == grid.length
//     n == grid[i].length
//     1 <= m, n <= 100
//     -10^9 <= grid[r][c] <= 10^9

import "fmt"

// O(m * n) 暴力
func findColumnWidth(grid [][]int) []int {
    res := make([]int, len(grid[0]))
    for _, s := range grid {
        for i, n := range s {
            c := 0
            if n <= 0 { // include minus sign in length
                c = 1
            }
            for n != 0 { // calculate length of n
                n /= 10
                c++
            }
            if res[i] < c { // remember max length of n
                res[i] = c
            }
        }
    }
    return res
}

func findColumnWidth1(grid [][]int) []int {
    res := make([]int, len(grid[0]))
    min := func (x, y int) int { if x < y { return x; }; return y; }
    max := func (x, y int) int { if x > y { return x; }; return y; }
    for j := range grid[0] {
        mn, mx := 0, 0
        for _, row := range grid { // 找出每行最大&最小的值出来
            mn = min(mn, row[j])
            mx = max(mx, row[j])
        }
        mx = max(mx/10, -mn) // 最大值 / 10 最小值取决对值(负数要算一个长度) 
        l := 1
        for ; mx > 0; mx /= 10 {
            l++
        }
        res[j] = max(res[j], l)
    }
    return res
}

func main() {
    // Example 1:
    // Input: grid = [[1],[22],[333]]
    // Output: [3]
    // Explanation: In the 0th column, 333 is of length 3.
    fmt.Println(findColumnWidth([][]int{{1},{22},{333}})) // 3
    // Example 2:
    // Input: grid = [[-15,1,3],[15,7,12],[5,6,-2]]
    // Output: [3,1,2]
    // Explanation: 
    // In the 0th column, only -15 is of length 3.
    // In the 1st column, all integers are of length 1. 
    // In the 2nd column, both 12 and -2 are of length 2.
    fmt.Println(findColumnWidth([][]int{{-15,1,3},{15,7,12},{5,6,-2}})) // [3,1,2]

    fmt.Println(findColumnWidth1([][]int{{1},{22},{333}})) // 3
    fmt.Println(findColumnWidth1([][]int{{-15,1,3},{15,7,12},{5,6,-2}})) // [3,1,2]
}