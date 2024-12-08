package main

// 2912. Number of Ways to Reach Destination in the Grid
// You are given two integers n and m which represent the size of a 1-indexed grid. 
// You are also given an integer k, a 1-indexed integer array source and a 1-indexed integer array dest, 
// where source and dest are in the form [x, y] representing a cell on the given grid.

// You can move through the grid in the following way:
//     You can go from cell [x1, y1] to cell [x2, y2] if either x1 == x2 or y1 == y2.
//     Note that you can't move to the cell you are already in e.g. x1 == x2 and y1 == y2.

// Return the number of ways you can reach dest from source by moving through the grid exactly k times.

// Since the answer may be very large, return it modulo 10^9 + 7.

// Example 1:
// Input: n = 3, m = 2, k = 2, source = [1,1], dest = [2,2]
// Output: 2
// Explanation: There are 2 possible sequences of reaching [2,2] from [1,1]:
// - [1,1] -> [1,2] -> [2,2]
// - [1,1] -> [2,1] -> [2,2]

// Example 2:
// Input: n = 3, m = 4, k = 3, source = [1,2], dest = [2,3]
// Output: 9
// Explanation: There are 9 possible sequences of reaching [2,3] from [1,2]:
// - [1,2] -> [1,1] -> [1,3] -> [2,3]
// - [1,2] -> [1,1] -> [2,1] -> [2,3]
// - [1,2] -> [1,3] -> [3,3] -> [2,3]
// - [1,2] -> [1,4] -> [1,3] -> [2,3]
// - [1,2] -> [1,4] -> [2,4] -> [2,3]
// - [1,2] -> [2,2] -> [2,1] -> [2,3]
// - [1,2] -> [2,2] -> [2,4] -> [2,3]
// - [1,2] -> [3,2] -> [2,2] -> [2,3]
// - [1,2] -> [3,2] -> [3,3] -> [2,3]

// Constraints:
//     2 <= n, m <= 10^9
//     1 <= k <= 10^5
//     source.length == dest.length == 2
//     1 <= source[1], dest[1] <= n
//     1 <= source[2], dest[2] <= m

import "fmt"

func numberOfWays(n int, m int, k int, source []int, dest []int) int {
    mod := 1_000_000_007
    facts := []int{1, 0, 0, 0}
    for i := 0; i < k; i++ {
        g := make([]int, 4)
        g[0] = ((n - 1) * facts[1] + (m - 1) * facts[2]) % mod
        g[1] = (facts[0] + (n - 2) * facts[1] + (m - 1) * facts[3]) % mod
        g[2] = (facts[0] + (m-  2) * facts[2] + (n - 1) * facts[3]) % mod
        g[3] = (facts[1] + facts[2] + (n-2) * facts[3] + (m-2) * facts[3]) % mod
        facts = g
    }
    if source[0] == dest[0] {
        if source[1] == dest[1] { return facts[0] }
        return facts[2]
    }
    if source[1] == dest[1] { return facts[1] }
    return facts[3]
}

func main() {
    // Example 1:
    // Input: n = 3, m = 2, k = 2, source = [1,1], dest = [2,2]
    // Output: 2
    // Explanation: There are 2 possible sequences of reaching [2,2] from [1,1]:
    // - [1,1] -> [1,2] -> [2,2]
    // - [1,1] -> [2,1] -> [2,2]
    fmt.Println(numberOfWays(3, 2, 2, []int{1,1}, []int{2,2})) // 2
    // Example 2:
    // Input: n = 3, m = 4, k = 3, source = [1,2], dest = [2,3]
    // Output: 9
    // Explanation: There are 9 possible sequences of reaching [2,3] from [1,2]:
    // - [1,2] -> [1,1] -> [1,3] -> [2,3]
    // - [1,2] -> [1,1] -> [2,1] -> [2,3]
    // - [1,2] -> [1,3] -> [3,3] -> [2,3]
    // - [1,2] -> [1,4] -> [1,3] -> [2,3]
    // - [1,2] -> [1,4] -> [2,4] -> [2,3]
    // - [1,2] -> [2,2] -> [2,1] -> [2,3]
    // - [1,2] -> [2,2] -> [2,4] -> [2,3]
    // - [1,2] -> [3,2] -> [2,2] -> [2,3]
    // - [1,2] -> [3,2] -> [3,3] -> [2,3]
    fmt.Println(numberOfWays(3, 4, 3, []int{1,2}, []int{2,3})) // 9
}