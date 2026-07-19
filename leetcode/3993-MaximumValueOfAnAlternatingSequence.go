package main

// 3993. Maximum Value of an Alternating Sequence
// You are given three integers n, s, and m.

// A sequence seq of integers of length n is considered valid if:
//     1. seq[0] = s.
//     2. The sequence is alternating, meaning that either:
//         2.1 seq[0] > seq[1] < seq[2] > ..., or
//         2.2 seq[0] < seq[1] > seq[2] < ....
//     3. For every adjacent pair, |seq[i] - seq[i - 1]| <= m.

// A sequence of length 1 is considered alternating.

// Return the maximum possible element that can appear in any valid sequence.

// Example 1:
// Input: n = 4, s = 3, m = 5
// Output: 12
// Explanation:
// One valid sequence is [3, 8, 7, 12].
// The maximum element in the sequence is 12.

// Example 2:
// Input: n = 2, s = 4, m = 3
// Output: 7
// Explanation:
// One valid sequence is [4, 7].
// The maximum element in the sequence is 7.

// Constraints:
//     1 <= n, s <= 10^9
//     1 <= m <= 10^5

import "fmt"

func maximumValue(n int, s int, m int) int64 {
    if n == 1 {
        return int64(s)
    }
    return int64(s + (m - 1) * (n / 2) + 1)
}

func main() {
    // Example 1:
    // Input: n = 4, s = 3, m = 5
    // Output: 12
    // Explanation:
    // One valid sequence is [3, 8, 7, 12].
    // The maximum element in the sequence is 12.
    fmt.Println(maximumValue(4, 3, 5)) // 12
    // Example 2:
    // Input: n = 2, s = 4, m = 3
    // Output: 7
    // Explanation:
    // One valid sequence is [4, 7].
    // The maximum element in the sequence is 7.
    fmt.Println(maximumValue(2, 4, 3)) // 7

    fmt.Println(maximumValue(1, 1, 1)) // 1
    fmt.Println(maximumValue(1, 1_000_000_000, 100_000)) // 1000000000
    fmt.Println(maximumValue(1_000_000_000, 1, 100_000)) // 49999500000002
    fmt.Println(maximumValue(1_000_000_000, 1_000_000_000, 1)) // 1000000001
    fmt.Println(maximumValue(1, 1, 100_000)) // 1
    fmt.Println(maximumValue(1, 1_000_000_000, 1)) // 1000000000
    fmt.Println(maximumValue(1_000_000_000, 1, 1)) // 2
    fmt.Println(maximumValue(1_000_000_000, 1_000_000_000, 100_000)) // 50000500000001
}