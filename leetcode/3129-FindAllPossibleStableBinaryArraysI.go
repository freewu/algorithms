package main

// 3129. Find All Possible Stable Binary Arrays I
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
// The two possible stable binary arrays are [1,0] and [0,1], as both arrays have a single 0 and a single 1, and no subarray has a length greater than 2.

// Example 2:
// Input: zero = 1, one = 2, limit = 1
// Output: 1
// Explanation:
// The only possible stable binary array is [1,0,1].
// Note that the binary arrays [1,1,0] and [0,1,1] have subarrays of length 2 with identical elements, hence, they are not stable.

// Example 3:
// Input: zero = 3, one = 3, limit = 2
// Output: 14
// Explanation:
// All the possible stable binary arrays are [0,0,1,0,1,1], [0,0,1,1,0,1], [0,1,0,0,1,1], [0,1,0,1,0,1], [0,1,0,1,1,0], [0,1,1,0,0,1], [0,1,1,0,1,0], [1,0,0,1,0,1], [1,0,0,1,1,0], [1,0,1,0,0,1], [1,0,1,0,1,0], [1,0,1,1,0,0], [1,1,0,0,1,0], and [1,1,0,1,0,0].

// Constraints:
//     1 <= zero, one, limit <= 200

import "fmt"

// func numberOfStableArrays(zero int, one int, limit int) int {
//     mod := 1_000_000_007
//     dp := make([zero + 1][one + 1][3][limit + 1]int)
//     var countStableArrays func(int zeroes, int ones, int prev, int prevCount, int limit) int
//     countStableArrays = func(int zeroes, int ones, int prev, int prevCount, int limit) int {
//         if zeroes == 0 && ones == 0 { return 1 }
//         if dp[zeroes][ones][prev][prevCount] != 0 {
//             return dp[zeroes][ones][prev][prevCount]
//         }
//         countHere := 0
//         if zeroes > 0 {
//             if prev != 0 || prevCount + 1 <= limit {
//                 if prev == 0 {
//                     countHere = countStableArrays(zeroes - 1, ones, 0, prevCount + 1, limit) % mod
//                 } else {
//                     countHere = countStableArrays(zeroes - 1, ones, 0, 1, limit) % mod
//                 }
//             }
//         }
//         if ones > 0 {
//             if prev != 1 || prevCount + 1 <= limit {
//                 if prev == 1 {
//                     countHere = (countHere + countStableArrays(zeroes, ones - 1, 1, prevCount + 1, limit)) % mod
//                 } else {
//                     countHere = (countHere + countStableArrays(zeroes, ones - 1, 1, 1, limit)) % mod
//                 }
//             }
//         }
//         dp[zeroes][ones][prev][prevCount] = countHere
//         return countHere
//     }
//     return countStableArrays(zero, one, 2, 0, limit)
// }

func numberOfStableArrays(zero int, one int, limit int) int {
    dp, mod := make([][][2]int, zero + 1), int(1e9 + 7)
    min := func (x, y int) int { if x < y { return x; }; return y; }
    for i := 0; i <= zero; i++ {
        dp[i] = make([][2]int, one + 1)
    }
    for i := 0; i <= min(zero, limit); i++ {
        dp[i][0][0] = 1
    }
    for j := 0; j <= min(one, limit); j++ {
        dp[0][j][1] = 1
    }
    for i := 1; i <= zero; i++ {
        for j := 1; j <= one; j++ {
            if i > limit {
                dp[i][j][0] = dp[i - 1][j][0] + dp[i - 1][j][1] - dp[i - limit - 1][j][1]
            } else {
                dp[i][j][0] = dp[i - 1][j][0] + dp[i - 1][j][1]
            }
            dp[i][j][0] = (dp[i][j][0] % mod + mod) % mod
            if j > limit {
                dp[i][j][1] = dp[i][j - 1][1] + dp[i][j - 1][0] - dp[i][j - limit - 1][0]
            } else {
                dp[i][j][1] = dp[i][j - 1][1] + dp[i][j - 1][0]
            }
            dp[i][j][1] = (dp[i][j][1] % mod + mod) % mod
        }
    }
    return (dp[zero][one][0] + dp[zero][one][1]) % mod
}

func main() {
    // Example 1:
    // Input: zero = 1, one = 1, limit = 2
    // Output: 2
    // Explanation:
    // The two possible stable binary arrays are [1,0] and [0,1], as both arrays have a single 0 and a single 1, and no subarray has a length greater than 2.
    fmt.Println(numberOfStableArrays(1,1,2)) // 2
    // Example 2:
    // Input: zero = 1, one = 2, limit = 1
    // Output: 1
    // Explanation:
    // The only possible stable binary array is [1,0,1].
    // Note that the binary arrays [1,1,0] and [0,1,1] have subarrays of length 2 with identical elements, hence, they are not stable.
    fmt.Println(numberOfStableArrays(1,2,1)) // 1
    // Example 3:
    // Input: zero = 3, one = 3, limit = 2
    // Output: 14
    // Explanation:
    // All the possible stable binary arrays are [0,0,1,0,1,1], [0,0,1,1,0,1], [0,1,0,0,1,1], [0,1,0,1,0,1], [0,1,0,1,1,0], [0,1,1,0,0,1], [0,1,1,0,1,0], [1,0,0,1,0,1], [1,0,0,1,1,0], [1,0,1,0,0,1], [1,0,1,0,1,0], [1,0,1,1,0,0], [1,1,0,0,1,0], and [1,1,0,1,0,0].
    fmt.Println(numberOfStableArrays(3,3,2)) // 14
}