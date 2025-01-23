package main

// 2579. Count Total Number of Colored Cells
// There exists an infinitely large two-dimensional grid of uncolored unit cells. 
// You are given a positive integer n, indicating that you must do the following routine for n minutes:
//     1. At the first minute, color any arbitrary unit cell blue.
//     2. Every minute thereafter, color blue every uncolored cell that touches a blue cell.

// Below is a pictorial representation of the state of the grid after minutes 1, 2, and 3.

// Return the number of colored cells at the end of n minutes.

// Example 1:
// Input: n = 1
// Output: 1
// Explanation: After 1 minute, there is only 1 blue cell, so we return 1.

// Example 2:
// Input: n = 2
// Output: 5
// Explanation: After 2 minutes, there are 4 colored cells on the boundary and 1 in the center, so we return 5. 

// Constraints:
//     1 <= n <= 10^5

import "fmt"

func coloredCells(n int) int64 {
    res := 1
    for i := 2; i <= n; i++ {
        res += 4 * (i - 1)
    }
    return int64(res)
}

func coloredCells1(n int) int64 {
    return int64(n * 2) * int64(n) - int64(2 * (n - 1) + 1)
}

func main() {
    // Example 1:
    // Input: n = 1
    // Output: 1
    // Explanation: After 1 minute, there is only 1 blue cell, so we return 1.
    fmt.Println(coloredCells(1)) // 1
    // Example 2:
    // Input: n = 2
    // Output: 5
    // Explanation: After 2 minutes, there are 4 colored cells on the boundary and 1 in the center, so we return 5. 
    fmt.Println(coloredCells(2)) // 5

    fmt.Println(coloredCells(3)) // 13
    fmt.Println(coloredCells(4)) // 25
    fmt.Println(coloredCells(8)) // 113
    fmt.Println(coloredCells(64)) // 8065
    fmt.Println(coloredCells(100)) // 19801
    fmt.Println(coloredCells(999)) // 1994005
    fmt.Println(coloredCells(1000)) // 1998001
    fmt.Println(coloredCells(9999)) // 199940005
    fmt.Println(coloredCells(10000)) // 199980001

    fmt.Println(coloredCells1(1)) // 1
    fmt.Println(coloredCells1(2)) // 5
    fmt.Println(coloredCells1(3)) // 13
    fmt.Println(coloredCells1(4)) // 25
    fmt.Println(coloredCells1(8)) // 113
    fmt.Println(coloredCells1(64)) // 8065
    fmt.Println(coloredCells1(100)) // 19801
    fmt.Println(coloredCells1(999)) // 1994005
    fmt.Println(coloredCells1(1000)) // 1998001
    fmt.Println(coloredCells1(9999)) // 199940005
    fmt.Println(coloredCells1(10000)) // 199980001
}