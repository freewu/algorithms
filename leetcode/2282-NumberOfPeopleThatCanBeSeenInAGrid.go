package main

// 2282. Number of People That Can Be Seen in a Grid
// You are given an m x n 0-indexed 2D array of positive integers heights 
// where heights[i][j] is the height of the person standing at position (i, j).

// A person standing at position (row1, col1) can see a person standing at position (row2, col2) if:
//     1. The person at (row2, col2) is to the right or below the person at (row1, col1). 
//        More formally, this means that either row1 == row2 and col1 < col2 or row1 < row2 and col1 == col2.
//     2. Everyone in between them is shorter than both of them.

// Return an m x n 2D array of integers answer where answer[i][j] is the number of people that the person at position (i, j) can see.

// Example 1:
// <img src="https://assets.leetcode.com/uploads/2022/05/24/image-20220524180458-1.png" />
// Input: heights = [[3,1,4,2,5]]
// Output: [[2,1,2,1,0]]
// Explanation:
// - The person at (0, 0) can see the people at (0, 1) and (0, 2).
//   Note that he cannot see the person at (0, 4) because the person at (0, 2) is taller than him.
// - The person at (0, 1) can see the person at (0, 2).
// - The person at (0, 2) can see the people at (0, 3) and (0, 4).
// - The person at (0, 3) can see the person at (0, 4).
// - The person at (0, 4) cannot see anybody.

// Example 2:
// <img src="https://assets.leetcode.com/uploads/2022/05/23/image-20220523113533-2.png" />
// Input: heights = [[5,1],[3,1],[4,1]]
// Output: [[3,1],[2,1],[1,0]]
// Explanation:
// - The person at (0, 0) can see the people at (0, 1), (1, 0) and (2, 0).
// - The person at (0, 1) can see the person at (1, 1).
// - The person at (1, 0) can see the people at (1, 1) and (2, 0).
// - The person at (1, 1) can see the person at (2, 1).
// - The person at (2, 0) can see the person at (2, 1).
// - The person at (2, 1) cannot see anybody.

// Constraints:
//     1 <= heights.length <= 400
//     1 <= heights[i].length <= 400
//     1 <= heights[i][j] <= 10^5

import "fmt"

// 单调栈
func seePeople(heights [][]int) [][]int {
    m, n := len(heights), len(heights[0])
    res := make([][]int, m)
    for i := 0; i < m; i++ {
        res[i] = make([]int, n)
    }
    for i := 0; i < m; i++ {
        stack := []int{}
        for j := n - 1; j >= 0; j-- {
            for len(stack) > 0 {
                res[i][j]++
                if stack[len(stack)-1] == heights[i][j] {
                    stack = stack[:len(stack)-1]
                    break
                }
                if stack[len(stack)-1] > heights[i][j] {
                    break
                }
                stack = stack[:len(stack)-1]
            }
            stack = append(stack, heights[i][j])
        }
    }
    for j := 0; j < n; j++ {
        stack := []int{}
        for i := m - 1; i >= 0; i-- {
            for len(stack) > 0 {
                res[i][j]++
                if stack[len(stack)-1] == heights[i][j] {
                    stack = stack[:len(stack)-1]
                    break
                }
                if stack[len(stack)-1] > heights[i][j] {
                    break
                }
                stack = stack[:len(stack)-1]
            }
            stack = append(stack, heights[i][j])
        }
    }
    return res
}

func main() {
    // Example 1:
    // <img src="https://assets.leetcode.com/uploads/2022/05/24/image-20220524180458-1.png" />
    // Input: heights = [[3,1,4,2,5]]
    // Output: [[2,1,2,1,0]]
    // Explanation:
    // - The person at (0, 0) can see the people at (0, 1) and (0, 2).
    //   Note that he cannot see the person at (0, 4) because the person at (0, 2) is taller than him.
    // - The person at (0, 1) can see the person at (0, 2).
    // - The person at (0, 2) can see the people at (0, 3) and (0, 4).
    // - The person at (0, 3) can see the person at (0, 4).
    // - The person at (0, 4) cannot see anybody.
    fmt.Println(seePeople([][]int{{3,1,4,2,5}})) // [[2,1,2,1,0]]
    // Example 2:
    // <img src="https://assets.leetcode.com/uploads/2022/05/23/image-20220523113533-2.png" />
    // Input: heights = [[5,1],[3,1],[4,1]]
    // Output: [[3,1],[2,1],[1,0]]
    // Explanation:
    // - The person at (0, 0) can see the people at (0, 1), (1, 0) and (2, 0).
    // - The person at (0, 1) can see the person at (1, 1).
    // - The person at (1, 0) can see the people at (1, 1) and (2, 0).
    // - The person at (1, 1) can see the person at (2, 1).
    // - The person at (2, 0) can see the person at (2, 1).
    // - The person at (2, 1) cannot see anybody.
    fmt.Println(seePeople([][]int{{5,1},{3,1},{4,1}})) // [[3,1],[2,1],[1,0]]
}