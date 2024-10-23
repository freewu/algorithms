package main

// 3258. Count Substrings That Satisfy K-Constraint I
// You are given a binary string s and an integer k.

// A binary string satisfies the k-constraint if either of the following conditions holds:
//     The number of 0's in the string is at most k.
//     The number of 1's in the string is at most k.

// Return an integer denoting the number of substrings of s that satisfy the k-constraint.

// Example 1:
// Input: s = "10101", k = 1
// Output: 12
// Explanation:
// Every substring of s except the substrings "1010", "10101", and "0101" satisfies the k-constraint.

// Example 2:
// Input: s = "1010101", k = 2
// Output: 25
// Explanation:
// Every substring of s except the substrings with a length greater than 5 satisfies the k-constraint.

// Example 3:
// Input: s = "11111", k = 1
// Output: 15
// Explanation:
// All substrings of s satisfy the k-constraint.

// Constraints:
//     1 <= s.length <= 50
//     1 <= k <= s.length
//     s[i] is either '0' or '1'.

import "fmt"

// Sliding Window
func countKConstraintSubstrings(s string, k int) int {
    res, count0, count1, left := 0, 0, 0, -1
    for right := 0; right < len(s); right++ {
        if s[right] == '1' {
            count1++
        } else {
            count0++
        }
        for left <= right && count0 > k && count1> k {
            left++
            if s[left] == '1' {
                count1--
            } else {
                count0--
            }
        }
        res += right - left
    }
    return res
}

func main() {
    // Example 1:
    // Input: s = "10101", k = 1
    // Output: 12
    // Explanation:
    // Every substring of s except the substrings "1010", "10101", and "0101" satisfies the k-constraint.
    fmt.Println(countKConstraintSubstrings("10101", 1)) // 12
    // Example 2:
    // Input: s = "1010101", k = 2
    // Output: 25
    // Explanation:
    // Every substring of s except the substrings with a length greater than 5 satisfies the k-constraint.
    fmt.Println(countKConstraintSubstrings("1010101", 2)) // 25
    // Example 3:
    // Input: s = "11111", k = 1
    // Output: 15
    // Explanation:
    // All substrings of s satisfy the k-constraint.
    fmt.Println(countKConstraintSubstrings("11111", 1)) // 15
}