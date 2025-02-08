package main

// 面试题 17.23. Max Black Square LCCI
// Imagine you have a square matrix, where each cell (pixel) is either black 
// or white Design an algorithm to find the maximum subsquare such that all four borders are filled with black pixels.

// Return an array [r, c, size], where r, c are the row number and the column number of the subsquare's upper left corner respectively, and size is the side length of the subsquare. 
// If there are more than one answers, return the one that has smallest r. 
// If there are more than one answers that have the same r, return the one that has smallest c. 
// If there's no answer, return an empty array.

// Example 1:
// Input:
// [
//    [1,0,1],
//    [0,0,1],
//    [0,0,1]
// ]
// Output: [1,0,2]
// Explanation: 0 represents black, and 1 represents white, bold elements in the input is the answer.

// Example 2:
// Input:
// [
//    [0,1,1],
//    [1,0,1],
//    [1,1,0]
// ]
// Output: [0,0,1]

// Note:
//     matrix.length == matrix[0].length <= 200

import "fmt"

func findSquare(matrix [][]int) []int {
    n := len(matrix)
    left, up := make([][]int, n + 1), make([][]int, n + 1)
    for i := range left {
        left[i], up[i] = make([]int, n + 1), make([]int, n + 1)
    }
    r, c, size := 0, 0, 0
    min := func (x, y int) int { if x < y { return x; }; return y; }
    for i := 1; i <= n; i++ {
        for j := 1; j <= n; j++ {
            if matrix[i-1][j-1] == 0 {
                left[i][j], up[i][j] = left[i][j-1] + 1,  up[i-1][j] + 1
                border := min(left[i][j], up[i][j])
                for left[i-border+1][j] < border || up[i][j-border+1] < border {
                    border--
                }
                if border > size {
                    r, c, size  = i - border, j - border, border
                }
            }
        }
    }
    if size > 0 {
        return []int{ r, c, size }
    }
    return []int{}
}

func main() {
    // Example 1:
    // Input:
    // [
    //    [1,0,1],
    //    [0,0,1],
    //    [0,0,1]
    // ]
    // Output: [1,0,2]
    // Explanation: 0 represents black, and 1 represents white, bold elements in the input is the answer.
    fmt.Println(findSquare([][]int{{1,0,1},{0,0,1},{0,0,1}})) // [1,0,2]
    // Example 2:
    // Input:
    // [
    //    [0,1,1],
    //    [1,0,1],
    //    [1,1,0]
    // ]
    // Output: [0,0,1]
    fmt.Println(findSquare([][]int{{0,1,1},{1,0,1},{1,1,0}})) // [0,0,1]
}