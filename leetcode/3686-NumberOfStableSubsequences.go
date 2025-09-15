package main

// 3686. Number of Stable Subsequences
// You are given an integer array nums.

// A subsequence is stable if it does not contain three consecutive elements with the same parity when the subsequence is read in order (i.e., consecutive inside the subsequence).

// Return the number of stable subsequences.

// Since the answer may be too large, return it modulo 10^9 + 7.

// Example 1:
// Input: nums = [1,3,5]
// Output: 6
// Explanation:
// Stable subsequences are [1], [3], [5], [1, 3], [1, 5], and [3, 5].
// Subsequence [1, 3, 5] is not stable because it contains three consecutive odd numbers. Thus, the answer is 6.

// Example 2:
// Input: nums = [2,3,4,2]
// Output: 14
// Explanation:
// The only subsequence that is not stable is [2, 4, 2], which contains three consecutive even numbers.
// All other subsequences are stable. Thus, the answer is 14.

// Constraints:
//     1 <= nums.length <= 10^5
//     1 <= nums[i] <= 10​​​​​​​^5

import "fmt"

func countStableSubsequences(nums []int) int {
    const mod = 1_000_000_007
    var dp1_0, dp2_0, dp1_1, dp2_1 int64
    for _, v := range nums {
        p := v & 1
        ndp1_0, ndp2_0, ndp1_1, ndp2_1 := dp1_0, dp2_0, dp1_1, dp2_1
        if p == 0 {
            ndp1_0 = (ndp1_0 + 1) % mod
            ndp1_0 = (ndp1_0 + dp1_1 + dp2_1) % mod
            ndp2_0 = (ndp2_0 + dp1_0) % mod
        } else {
            ndp1_1 = (ndp1_1 + 1) % mod
            ndp1_1 = (ndp1_1 + dp1_0 + dp2_0) % mod
            ndp2_1 = (ndp2_1 + dp1_1) % mod
        }
        dp1_0, dp2_0, dp1_1, dp2_1 = ndp1_0, ndp2_0, ndp1_1, ndp2_1
    }
    return int((dp1_0 + dp2_0 + dp1_1 + dp2_1) % mod)
}

func main() {
    // Example 1:
    // Input: nums = [1,3,5]
    // Output: 6
    // Explanation:
    // Stable subsequences are [1], [3], [5], [1, 3], [1, 5], and [3, 5].
    // Subsequence [1, 3, 5] is not stable because it contains three consecutive odd numbers. Thus, the answer is 6.
    fmt.Println(countStableSubsequences([]int{1,3,5})) // 6
    // Example 2:
    // Input: nums = [2,3,4,2]
    // Output: 14
    // Explanation:
    // The only subsequence that is not stable is [2, 4, 2], which contains three consecutive even numbers.
    // All other subsequences are stable. Thus, the answer is 14.
    fmt.Println(countStableSubsequences([]int{2,3,4,2})) // 14

    fmt.Println(countStableSubsequences([]int{1,2,3,4,5,6,7,8,9})) // 419
    fmt.Println(countStableSubsequences([]int{9,8,7,6,5,4,3,2,1})) // 419
}