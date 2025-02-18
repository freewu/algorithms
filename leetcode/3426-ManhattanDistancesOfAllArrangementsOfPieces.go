package main

// 3426. Manhattan Distances of All Arrangements of Pieces
// You are given three integers m, n, and k.

// There is a rectangular grid of size m × n containing k identical pieces. 
// Return the sum of Manhattan distances between every pair of pieces over all valid arrangements of pieces.

// A valid arrangement is a placement of all k pieces on the grid with at most one piece per cell.

// Since the answer may be very large, return it modulo 10^9 + 7.

// The Manhattan Distance between two cells (xi, yi) and (xj, yj) is |xi - xj| + |yi - yj|.

// Example 1:
// Input: m = 2, n = 2, k = 2
// Output: 8
// Explanation:
// The valid arrangements of pieces on the board are:
// <img src="https://assets.leetcode.com/uploads/2024/12/25/untitled-diagramdrawio.png" />
// In the first 4 arrangements, the Manhattan distance between the two pieces is 1.
// In the last 2 arrangements, the Manhattan distance between the two pieces is 2.
// Thus, the total Manhattan distance across all valid arrangements is 1 + 1 + 1 + 1 + 2 + 2 = 8.

// Example 2:
// Input: m = 1, n = 4, k = 3
// Output: 20
// Explanation:
// The valid arrangements of pieces on the board are:
// <img src="https://assets.leetcode.com/uploads/2024/12/25/4040example2drawio.png" />
// The first and last arrangements have a total Manhattan distance of 1 + 1 + 2 = 4.
// The middle two arrangements have a total Manhattan distance of 1 + 2 + 3 = 6.
// The total Manhattan distance between all pairs of pieces across all arrangements is 4 + 6 + 6 + 4 = 20.

// Constraints:
//     1 <= m, n <= 10^5
//     2 <= m * n <= 10^5
//     2 <= k <= m * n

import "fmt"

const mod = 1_000_000_007
const mx = 100_000

var f [mx]int    // f[i] = i!
var invF [mx]int // invF[i] = i!^-1

func init() {
    pow := func(x, n int) int {
        res := 1
        for ; n > 0; n /= 2 {
            if n%2 > 0 {
                res = res * x % mod
            }
            x = x * x % mod
        }
        return res
    }
    f[0] = 1
    for i := 1; i < mx; i++ {
        f[i] = f[i-1] * i % mod
    }
    invF[mx-1] = pow(f[mx-1], mod-2)
    for i := mx - 1; i > 0; i-- {
        invF[i-1] = invF[i] * i % mod
    }
}

func distanceSum(m, n, k int) int {
    comb := func(n, m int) int { return f[n] * invF[m] % mod * invF[n-m] % mod }
    return (m * n * (m*(n*n-1) + n*(m*m-1))) / 6 % mod * comb(m*n-2, k-2) % mod
}

func main() {
    // Example 1:
    // Input: m = 2, n = 2, k = 2
    // Output: 8
    // Explanation:
    // The valid arrangements of pieces on the board are:
    // <img src="https://assets.leetcode.com/uploads/2024/12/25/untitled-diagramdrawio.png" />
    // In the first 4 arrangements, the Manhattan distance between the two pieces is 1.
    // In the last 2 arrangements, the Manhattan distance between the two pieces is 2.
    // Thus, the total Manhattan distance across all valid arrangements is 1 + 1 + 1 + 1 + 2 + 2 = 8.
    fmt.Println(distanceSum(2, 2, 2)) // 8
    // Example 2:
    // Input: m = 1, n = 4, k = 3
    // Output: 20
    // Explanation:
    // The valid arrangements of pieces on the board are:
    // <img src="https://assets.leetcode.com/uploads/2024/12/25/4040example2drawio.png" />
    // The first and last arrangements have a total Manhattan distance of 1 + 1 + 2 = 4.
    // The middle two arrangements have a total Manhattan distance of 1 + 2 + 3 = 6.
    // The total Manhattan distance between all pairs of pieces across all arrangements is 4 + 6 + 6 + 4 = 20.
    fmt.Println(distanceSum(1, 4, 3)) // 20

    fmt.Println(distanceSum(500, 200, 35762)) // 610532901
    fmt.Println(distanceSum(1, 2, 2)) // 1
    fmt.Println(distanceSum(2, 1, 2)) // 1
    fmt.Println(distanceSum(100_000, 1, 2)) // 665483338
    fmt.Println(distanceSum(100_000, 1, 100_000)) // 665483338
    fmt.Println(distanceSum(1, 100_000, 2)) // 665483338
    fmt.Println(distanceSum(1, 100_000, 100_000)) // 665483338
}


