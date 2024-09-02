package main

// 1240. Tiling a Rectangle with the Fewest Squares
// Given a rectangle of size n x m, return the minimum number of integer-sided squares that tile the rectangle.

// Example 1:
// <img src="https://assets.leetcode.com/uploads/2019/10/17/sample_11_1592.png" />
// Input: n = 2, m = 3
// Output: 3
// Explanation: 3 squares are necessary to cover the rectangle.
// 2 (squares of 1x1)
// 1 (square of 2x2)

// Example 2:
// <img src="https://assets.leetcode.com/uploads/2019/10/17/sample_22_1592.png" />
// Input: n = 5, m = 8
// Output: 5

// Example 3:
// <img src="https://assets.leetcode.com/uploads/2019/10/17/sample_33_1592.png" />
// Input: n = 11, m = 13
// Output: 6

// Constraints:    
//     1 <= n, m <= 13

import "fmt"

func tilingRectangle(n int, m int) int {
    if m == n { // 正方形只需要一块
        return 1 
    }
    min := func(arr []int) (int, int) {
        area, index := 0, 0
        for i := 0; i < len(arr); i++ {
            if i == 0 || arr[i] < area {
                area = arr[i]
                index = i
            }
        }
        return area, index
    }
    res, arr := n * m,  make([]int, m)
    var backtrack func(int) 
    backtrack = func(squareCount int) {
        if squareCount >= res { return }
        area, start := min(arr)
        if area == n {
            if res > squareCount { res = squareCount }
            return
        }
        end := start
        for end < m && arr[end] == arr[start] && end - start+1 <= n - area {
            end += 1
        }
        for i := end-1; i >= start; i-- {
            side := i - start + 1
            for j := start; j <= i; j++ {
                arr[j] += side
            }
            backtrack(squareCount + 1)
            for j := start; j <= i; j++ {
                arr[j] -= side
            }
        }
    }
    backtrack(0)
    return res
}

func main() {
    // Example 1:
    // <img src="https://assets.leetcode.com/uploads/2019/10/17/sample_11_1592.png" />
    // Input: n = 2, m = 3
    // Output: 3
    // Explanation: 3 squares are necessary to cover the rectangle.
    // 2 (squares of 1x1)
    // 1 (square of 2x2)
    fmt.Println(tilingRectangle(2, 3)) // 3 
    // Example 2:
    // <img src="https://assets.leetcode.com/uploads/2019/10/17/sample_22_1592.png" />
    // Input: n = 5, m = 8
    // Output: 5
    fmt.Println(tilingRectangle(5, 8)) // 5
    // Example 3:
    // <img src="https://assets.leetcode.com/uploads/2019/10/17/sample_33_1592.png" />
    // Input: n = 11, m = 13
    // Output: 6
    fmt.Println(tilingRectangle(11, 13)) // 6
}