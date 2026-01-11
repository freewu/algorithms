package main

// 85. Maximal Rectangle
// Given a rows x cols binary matrix filled with 0's and 1's, 
// find the largest rectangle containing only 1's and return its area.

// Example 1:
// <img src="https://assets.leetcode.com/uploads/2020/09/14/maximal.jpg" />
// Input: matrix = [["1","0","1","0","0"],["1","0","1","1","1"],["1","1","1","1","1"],["1","0","0","1","0"]]
// Output: 6
// Explanation: The maximal rectangle is shown in the above picture.

// Example 2:
// Input: matrix = [["0"]]
// Output: 0

// Example 3:
// Input: matrix = [["1"]]
// Output: 1

// Constraints:
//     rows == matrix.length
//     cols == matrix[i].length
//     1 <= row, cols <= 200
//     matrix[i][j] is '0' or '1'.

import "fmt"

// stack
func maximalRectangle(matrix [][]byte) int {
    res, heights := 0, make([]int, len(matrix[0]) + 1)
    heights[len(heights)-1] = -1
    max := func (x, y int) int { if x > y { return x; }; return y; }
    for _, row := range matrix {
        for i := range row {
            if row[i] == '1' {
                heights[i]++
            } else {
                heights[i] = 0
            }
        }
        stack := []int{}
        for i, currentHeight := range heights {
            for len(stack) > 0 && heights[stack[len(stack)-1]] > currentHeight {
                prev := heights[stack[len(stack)-1]]
                stack = stack[:len(stack)-1]
                width := i
                if len(stack) > 0 {
                    width = i - stack[len(stack)-1] - 1
                }
                res = max(res, prev * width)
            }
            stack = append(stack, i)
        } 
    }
    return res
}

func main() {
    // Example 1:
    // <img src="https://assets.leetcode.com/uploads/2020/09/14/maximal.jpg" />
    // Input: matrix = [["1","0","1","0","0"],["1","0","1","1","1"],["1","1","1","1","1"],["1","0","0","1","0"]]
    // Output: 6
    // Explanation: The maximal rectangle is shown in the above picture.
    fmt.Println(maximalRectangle([][]byte{{'1','0','1','0','0'},{'1','0','1','1','1'},{'1','1','1','1','1'},{'1','0','0','1','0'}})) // 6
    // Example 2:
    // Input: matrix = [["0"]]
    // Output: 0
    fmt.Println(maximalRectangle([][]byte{{'0'}})) // 0
    // Example 3:
    // Input: matrix = [["1"]]
    // Output: 1
    fmt.Println(maximalRectangle([][]byte{{'1'}})) // 1
}