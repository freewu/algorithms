package main

// 3648. Minimum Sensors to Cover Grid
// You are given n × m grid and an integer k.

// A sensor placed on cell (r, c) covers all cells whose Chebyshev distance from (r, c) is at most k.

// The Chebyshev distance between two cells (r1, c1) and (r2, c2) is max(|r1 − r2|,|c1 − c2|).

// Your task is to return the minimum number of sensors required to cover every cell of the grid.

// Example 1:
// Input: n = 5, m = 5, k = 1
// Output: 4
// Explanation:
// Placing sensors at positions (0, 3), (1, 0), (3, 3), and (4, 1) ensures every cell in the grid is covered. Thus, the answer is 4.

// Example 2:
// Input: n = 2, m = 2, k = 2
// Output: 1
// Explanation:
// With k = 2, a single sensor can cover the entire 2 * 2 grid regardless of its position. Thus, the answer is 1.

// Constraints:
//     1 <= n <= 10^3
//     1 <= m <= 10^3
//     0 <= k <= 10^3

import "fmt"

func minSensors(n int, m int, k int) int {
    size := k * 2 + 1
    return ((n - 1) / size + 1) * ((m - 1) / size + 1)
}

func main() {
    // Example 1:
    // Input: n = 5, m = 5, k = 1
    // Output: 4
    // Explanation:
    // Placing sensors at positions (0, 3), (1, 0), (3, 3), and (4, 1) ensures every cell in the grid is covered. Thus, the answer is 4.
    fmt.Println(minSensors(5, 5, 1)) // 4
    // Example 2:
    // Input: n = 2, m = 2, k = 2
    // Output: 1
    // Explanation:
    // With k = 2, a single sensor can cover the entire 2 * 2 grid regardless of its position. Thus, the answer is 1.
    fmt.Println(minSensors(2, 2, 2)) // 1
    
    fmt.Println(minSensors(1, 1, 0)) // 1
    fmt.Println(minSensors(1000, 1, 0)) // 1000
    fmt.Println(minSensors(1, 1000, 0)) // 1000
    fmt.Println(minSensors(1, 1, 1000)) // 1
    fmt.Println(minSensors(1000, 1000, 0)) // 1000000
    fmt.Println(minSensors(1, 1000, 1000)) // 1
    fmt.Println(minSensors(1000, 1, 1000)) // 1
    fmt.Println(minSensors(1000, 1000, 1000)) // 1
}