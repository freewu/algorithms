package main

// 2481. Minimum Cuts to Divide a Circle
// A valid cut in a circle can be:
//     1. A cut that is represented by a straight line 
//        that touches two points on the edge of the circle and passes through its center, or
//     2. A cut that is represented by a straight line 
//        that touches one point on the edge of the circle and its center.

// Some valid and invalid cuts are shown in the figures below.
// <img src="https://assets.leetcode.com/uploads/2022/10/29/alldrawio.png" />

// Given the integer n, return the minimum number of cuts needed to divide a circle into n equal slices.

// Example 1:
// <img src="https://assets.leetcode.com/uploads/2022/10/24/11drawio.png" />
// Input: n = 4
// Output: 2
// Explanation: 
// The above figure shows how cutting the circle twice through the middle divides it into 4 equal slices.

// Example 2:
// <img src="https://assets.leetcode.com/uploads/2022/10/24/22drawio.png" />
// Input: n = 3
// Output: 3
// Explanation:
// At least 3 cuts are needed to divide the circle into 3 equal slices. 
// It can be shown that less than 3 cuts cannot result in 3 slices of equal size and shape.
// Also note that the first cut will not divide the circle into distinct parts.

// Constraints:
//     1 <= n <= 100

import "fmt"

func numberOfCuts(n int) int {
    if n == 1 { return 0 }
    if n % 2 == 0 { return n / 2 }
    return n
}

func main() {
    // Example 1:
    // <img src="https://assets.leetcode.com/uploads/2022/10/24/11drawio.png" />
    // Input: n = 4
    // Output: 2
    // Explanation: 
    // The above figure shows how cutting the circle twice through the middle divides it into 4 equal slices.
    fmt.Println(numberOfCuts(4)) // 2
    // Example 2:
    // <img src="https://assets.leetcode.com/uploads/2022/10/24/22drawio.png" />
    // Input: n = 3
    // Output: 3
    // Explanation:
    // At least 3 cuts are needed to divide the circle into 3 equal slices. 
    // It can be shown that less than 3 cuts cannot result in 3 slices of equal size and shape.
    // Also note that the first cut will not divide the circle into distinct parts.
    fmt.Println(numberOfCuts(3)) // 3

    fmt.Println(numberOfCuts(1)) // 0
    fmt.Println(numberOfCuts(8)) // 4
    fmt.Println(numberOfCuts(99)) // 99
    fmt.Println(numberOfCuts(100)) // 50
}