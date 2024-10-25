package main

// 2209. Minimum White Tiles After Covering With Carpets
// You are given a 0-indexed binary string floor, which represents the colors of tiles on a floor:
//     1. floor[i] = '0' denotes that the ith tile of the floor is colored black.
//     2. On the other hand, floor[i] = '1' denotes that the ith tile of the floor is colored white.

// You are also given numCarpets and carpetLen. 
// You have numCarpets black carpets, each of length carpetLen tiles. 
// Cover the tiles with the given carpets such that the number of white tiles still visible is minimum. 
// Carpets may overlap one another.

// Return the minimum number of white tiles still visible.

// Example 1:
// <img src="https://assets.leetcode.com/uploads/2022/02/10/ex1-1.png" />
// Input: floor = "10110101", numCarpets = 2, carpetLen = 2
// Output: 2
// Explanation: 
// The figure above shows one way of covering the tiles with the carpets such that only 2 white tiles are visible.
// No other way of covering the tiles with the carpets can leave less than 2 white tiles visible.

// Example 2:
// <img src="https://assets.leetcode.com/uploads/2022/02/10/ex2.png" />
// Input: floor = "11111", numCarpets = 2, carpetLen = 3
// Output: 0
// Explanation: 
// The figure above shows one way of covering the tiles with the carpets such that no white tiles are visible.
// Note that the carpets are able to overlap one another.

// Constraints:
//     1 <= carpetLen <= floor.length <= 1000
//     floor[i] is either '0' or '1'.
//     1 <= numCarpets <= 1000

import "fmt"

func minimumWhiteTiles(floor string, numCarpets int, carpetLen int) int {
    n := len(floor)
    dp := make([][]int, n + 1)
    for i := range dp {
        dp[i] = make([]int, numCarpets + 1)
    }
    min := func (x, y int) int { if x < y { return x; }; return y; }
    max := func (x, y int) int { if x > y { return x; }; return y; }
    for i := 1; i <= n; i++ {
        for j := 0; j <= numCarpets; j++ {
            jump := dp[i - 1][j] + int(floor[i - 1] - '0')
            cover := 1000
            if j > 0 {
                cover = dp[max(i - carpetLen, 0)][j - 1]
            }
            dp[i][j] = min(cover, jump)
        }
    }
    return dp[n][numCarpets]
}

func minimumWhiteTiles1(floor string, numCarpets int, carpetLen int) int {
    n := len(floor)
    if numCarpets * carpetLen >= n { return 0 }
    f, nf := make([]int, n), make([]int, n)
    if floor[0] == '1' {
        f[0] = 1
    }
    for i := 1; i < n; i++ {
        f[i] = f[i-1]
        if floor[i] == '1' {
            f[i] += 1
        }
    }
    min := func (x, y int) int { if x < y { return x; }; return y; }
    for i := 1; i <= numCarpets; i++ {
        for j := 0; j < i * carpetLen; j++ {
            nf[j] = 0
        }
        for j := i * carpetLen; j < n; j++ {
            if floor[j] == '1' {
                nf[j] = min(nf[j-1]+1, f[j-carpetLen])
            } else {
                nf[j] = min(nf[j-1], f[j-carpetLen])
            }
        }
        f, nf = nf, f
    }
    return f[n-1]
}

func main() {
    // Example 1:
    // <img src="https://assets.leetcode.com/uploads/2022/02/10/ex1-1.png" />
    // Input: floor = "10110101", numCarpets = 2, carpetLen = 2
    // Output: 2
    // Explanation: 
    // The figure above shows one way of covering the tiles with the carpets such that only 2 white tiles are visible.
    // No other way of covering the tiles with the carpets can leave less than 2 white tiles visible.
    fmt.Println(minimumWhiteTiles("10110101", 2, 2)) // 2
    // Example 2:
    // <img src="https://assets.leetcode.com/uploads/2022/02/10/ex2.png" />
    // Input: floor = "11111", numCarpets = 2, carpetLen = 3
    // Output: 0
    // Explanation: 
    // The figure above shows one way of covering the tiles with the carpets such that no white tiles are visible.
    // Note that the carpets are able to overlap one another.
    fmt.Println(minimumWhiteTiles("11111", 2, 3)) // 0

    fmt.Println(minimumWhiteTiles1("10110101", 2, 2)) // 2
    fmt.Println(minimumWhiteTiles1("11111", 2, 3)) // 0
}