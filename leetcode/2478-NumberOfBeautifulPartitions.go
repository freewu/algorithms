package main

// 2478. Number of Beautiful Partitions
// You are given a string s that consists of the digits '1' to '9' and two integers k and minLength.

// A partition of s is called beautiful if:
//     1. s is partitioned into k non-intersecting substrings.
//     2. Each substring has a length of at least minLength.
//     3. Each substring starts with a prime digit and ends with a non-prime digit. 
//        Prime digits are '2', '3', '5', and '7', and the rest of the digits are non-prime.

// Return the number of beautiful partitions of s. 
// Since the answer may be very large, return it modulo 10^9 + 7.

// A substring is a contiguous sequence of characters within a string.

// Example 1:
// Input: s = "23542185131", k = 3, minLength = 2
// Output: 3
// Explanation: There exists three ways to create a beautiful partition:
// "2354 | 218 | 5131"
// "2354 | 21851 | 31"
// "2354218 | 51 | 31"

// Example 2:
// Input: s = "23542185131", k = 3, minLength = 3
// Output: 1
// Explanation: There exists one way to create a beautiful partition: "2354 | 218 | 5131".

// Example 3:
// Input: s = "3312958", k = 3, minLength = 1
// Output: 1
// Explanation: There exists one way to create a beautiful partition: "331 | 29 | 58".

// Constraints:
//     1 <= k, minLength <= s.length <= 1000
//     s consists of the digits '1' to '9'.

import "fmt"

func beautifulPartitions(s string, k int, minLength int) int {
    prime := func(c byte) bool { return c == '2' || c == '3' || c == '5' || c == '7' }
    n, mod := len(s), 1_000_000_007
    if !prime(s[0]) || prime(s[n-1]) { return 0 }
    f, g := make([][]int, n + 1), make([][]int, n + 1)
    for i := range f {
        f[i], g[i] = make([]int, k + 1), make([]int, k + 1)
    }
    f[0][0], g[0][0] = 1, 1
    for i := 1; i <= n; i++ {
        if i >= minLength && !prime(s[i-1]) && (i == n || prime(s[i])) {
            for j := 1; j <= k; j++ {
                f[i][j] = g[i-minLength][j-1]
            }
        }
        for j := 0; j <= k; j++ {
            g[i][j] = (g[i-1][j] + f[i][j]) % mod
        }
    }
    return f[n][k]
}

func main() {
    // Example 1:
    // Input: s = "23542185131", k = 3, minLength = 2
    // Output: 3
    // Explanation: There exists three ways to create a beautiful partition:
    // "2354 | 218 | 5131"
    // "2354 | 21851 | 31"
    // "2354218 | 51 | 31"
    fmt.Println(beautifulPartitions("23542185131", 3, 2)) // 2
    // Example 2:
    // Input: s = "23542185131", k = 3, minLength = 3
    // Output: 1
    // Explanation: There exists one way to create a beautiful partition: "2354 | 218 | 5131".
    fmt.Println(beautifulPartitions("23542185131", 3, 3)) // 1
    // Example 3:
    // Input: s = "3312958", k = 3, minLength = 1
    // Output: 1
    // Explanation: There exists one way to create a beautiful partition: "331 | 29 | 58".
    fmt.Println(beautifulPartitions("3312958", 3, 1)) // 1
}