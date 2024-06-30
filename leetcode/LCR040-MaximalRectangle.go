package main

// LCR 040. 最大矩形
// 给定一个由 0 和 1 组成的矩阵 matrix ，找出只包含 1 的最大矩形，并返回其面积。
// 注意：此题 matrix 输入格式为一维 01 字符串数组。

// 示例 1：
// <img src="https://assets.leetcode.com/uploads/2020/09/14/maximal.jpg" />
// 输入：matrix = ["10100","10111","11111","10010"]
// 输出：6
// 解释：最大矩形如上图所示。

// 示例 2：
// 输入：matrix = []
// 输出：0

// 示例 3：
// 输入：matrix = ["0"]
// 输出：0

// 示例 4：
// 输入：matrix = ["1"]
// 输出：1

// 示例 5：
// 输入：matrix = ["00"]
// 输出：0
 
// 提示：
//     rows == matrix.length
//     cols == matrix[0].length
//     0 <= row, cols <= 200
//     matrix[i][j] 为 '0' 或 '1'

import "fmt"

// stack
func maximalRectangle(matrix []string) int {
    if len(matrix) == 0 {
        return 0
    }
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
    // Explanation: The maximal rectangle is shown in the above picture.
    fmt.Println(maximalRectangle([]string{"10100","10111","11111","10010"})) // 6
    fmt.Println(maximalRectangle([]string{"0"})) // 0
    fmt.Println(maximalRectangle([]string{"1"})) // 1
    fmt.Println(maximalRectangle([]string{})) // 0
}