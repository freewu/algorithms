package main

// 3130. Find All Possible Stable Binary Arrays II
// You are given 3 positive integers zero, one, and limit.
// A binary array arr is called stable if:
//     The number of occurrences of 0 in arr is exactly zero.
//     The number of occurrences of 1 in arr is exactly one.
//     Each subarray of arr with a size greater than limit must contain both 0 and 1.

// Return the total number of stable binary arrays.
// Since the answer may be very large, return it modulo 10^9 + 7.

// Example 1:
// Input: zero = 1, one = 1, limit = 2
// Output: 2
// Explanation:
// The two possible stable binary arrays are [1,0] and [0,1].

// Example 2:
// Input: zero = 1, one = 2, limit = 1
// Output: 1
// Explanation:
// The only possible stable binary array is [1,0,1].

// Example 3:
// Input: zero = 3, one = 3, limit = 2
// Output: 14
// Explanation:
// All the possible stable binary arrays are [0,0,1,0,1,1], [0,0,1,1,0,1], [0,1,0,0,1,1], [0,1,0,1,0,1], [0,1,0,1,1,0], [0,1,1,0,0,1], [0,1,1,0,1,0], [1,0,0,1,0,1], [1,0,0,1,1,0], [1,0,1,0,0,1], [1,0,1,0,1,0], [1,0,1,1,0,0], [1,1,0,0,1,0], and [1,1,0,1,0,0].

// Constraints:
//     1 <= zero, one, limit <= 1000

import "fmt"

func numberOfStableArrays(zero, one, limit int) int {
    mod := 1_000_000_007
    dp := make([][][]int, one + 1)
    for i := range dp {
        dp[i] = make([][]int, zero+1)
        for j := range dp[i] {
            dp[i][j] = make([]int, 2)
        }
    }
    prefixDPZero, prefixDPOne := make([][]int, one+1), make([][]int, one+1)
    for i := range prefixDPZero {
        prefixDPZero[i], prefixDPOne[i] = make([]int, zero + 1), make([]int, zero + 1)
    }
    for i := 0; i <= one; i++ {
        if i <= limit {
            dp[i][0][1] = 1
        } else {
            dp[i][0][1] = 0
        }
        prefixDPOne[i][0] = dp[i][0][1]
    }
    for i := 0; i <= zero; i++ {
        if i <= limit {
            dp[0][i][0] = 1
        } else {
            dp[0][i][0] = 0
        }
        prefixDPZero[0][i] = dp[0][i][0]
    }

    for i := 1; i <= one; i++ {
        for j := 1; j <= zero; j++ {
            dp[i][j][1] = prefixDPZero[i-1][j] % mod
            dp[i][j][0] = prefixDPOne[i][j-1] % mod

            if i-limit-1 >= 0 {
                dp[i][j][1] = (dp[i][j][1] - prefixDPZero[i-limit-1][j] + mod) % mod
            }
            if j-limit-1 >= 0 {
                dp[i][j][0] = (dp[i][j][0] - prefixDPOne[i][j-limit-1] + mod) % mod
            }
            prefixDPZero[i][j] = (prefixDPZero[i-1][j] + dp[i][j][0]) % mod
            prefixDPOne[i][j] = (prefixDPOne[i][j-1] + dp[i][j][1]) % mod
        }
    }
    return (dp[one][zero][0] + dp[one][zero][1]) % mod
}

func numberOfStableArrays1(zero int, one int, limit int) int {
    mod := 1_000_000_007
    dp := make([][][2]int, zero + 1)
    for i := range dp {
        dp[i] = make([][2]int, one + 1)
    }
    min := func (x, y int) int { if x < y { return x; }; return y; }
    d := min(zero, limit)
    for i := 0; i <= d; i++ {
        dp[i][0][0] = 1
    }
    d = min(one, limit)
    for i := 0; i <= d; i++ {
        dp[0][i][1] = 1
    }
    for i := 1; i <= zero; i++ {
        for j := 1; j <= one; j++ {
            d = 0
            if i > limit {
                d = dp[i-limit-1][j][1]
            }
            dp[i][j][0] = (dp[i-1][j][0] + dp[i-1][j][1] - d + mod) % mod
            d = 0
            if j > limit {
                d = dp[i][j-limit-1][0]
            }
            dp[i][j][1] = (dp[i][j-1][0] + dp[i][j-1][1] - d + mod) % mod
        }
    }
    return (dp[zero][one][0] + dp[zero][one][1]) % mod
}

func main() {
    // Example 1:
    // Input: zero = 1, one = 1, limit = 2
    // Output: 2
    // Explanation:
    // The two possible stable binary arrays are [1,0] and [0,1].
    fmt.Println(numberOfStableArrays(1,1,2)) // 2
    // Example 2:
    // Input: zero = 1, one = 2, limit = 1
    // Output: 1
    // Explanation:
    // The only possible stable binary array is [1,0,1].
    fmt.Println(numberOfStableArrays(1,2,1)) // 1
    // Example 3:
    // Input: zero = 3, one = 3, limit = 2
    // Output: 14
    // Explanation:
    // All the possible stable binary arrays are [0,0,1,0,1,1], [0,0,1,1,0,1], [0,1,0,0,1,1], [0,1,0,1,0,1], [0,1,0,1,1,0], [0,1,1,0,0,1], [0,1,1,0,1,0], [1,0,0,1,0,1], [1,0,0,1,1,0], [1,0,1,0,0,1], [1,0,1,0,1,0], [1,0,1,1,0,0], [1,1,0,0,1,0], and [1,1,0,1,0,0].
    fmt.Println(numberOfStableArrays(3,3,2)) // 14

    fmt.Println(numberOfStableArrays1(1,1,2)) // 2
    fmt.Println(numberOfStableArrays1(1,2,1)) // 1
    fmt.Println(numberOfStableArrays1(3,3,2)) // 14
}