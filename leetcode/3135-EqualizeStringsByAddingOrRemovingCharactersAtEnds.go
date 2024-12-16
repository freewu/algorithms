package main

// 3135. Equalize Strings by Adding or Removing Characters at Ends
// Given two strings initial and target, 
// your task is to modify initial by performing a series of operations to make it equal to target.

// In one operation, you can add or remove one character only at the beginning or the end of the string initial.

// Return the minimum number of operations required to transform initial into target.

// Example 1:
// Input: initial = "abcde", target = "cdef"
// Output: 3
// Explanation:
// Remove 'a' and 'b' from the beginning of initial, then add 'f' to the end.

// Example 2:
// Input: initial = "axxy", target = "yabx"
// Output: 6
// Explanation:
// Operation	Resulting String
// Add 'y' to the beginning	"yaxxy"
// Remove from end	"yaxx"
// Remove from end	"yax"
// Remove from end	"ya"
// Add 'b' to the end	"yab"
// Add 'x' to the end	"yabx"

// Example 3:
// Input: initial = "xyz", target = "xyz"
// Output: 0
// Explanation:
// No operations are needed as the strings are already equal.

// Constraints:
//     1 <= initial.length, target.length <= 1000
//     initial and target consist only of lowercase English letters.

import "fmt"

func minOperations(initial string, target string) int {
    m, n, mx := len(initial), len(target), 0
    facts := make([][]int, m + 1)
    for i := range facts {
        facts[i] = make([]int, n + 1)
    }
    max := func (x, y int) int { if x > y { return x; }; return y; }
    for i, a := range initial {
        for j, b := range target {
            if a == b { // 计算最长公共字串的长度
                facts[i + 1][j + 1] = facts[i][j] + 1
                mx = max(mx, facts[i + 1][j + 1])
            }
        }
    }
    return m + n - 2 * mx
}

func main() {
    // Example 1:
    // Input: initial = "abcde", target = "cdef"
    // Output: 3
    // Explanation:
    // Remove 'a' and 'b' from the beginning of initial, then add 'f' to the end.
    fmt.Println(minOperations("abcde", "cdef")) // 3
    // Example 2:
    // Input: initial = "axxy", target = "yabx"
    // Output: 6
    // Explanation:
    // Operation	Resulting String
    // Add 'y' to the beginning	"yaxxy"
    // Remove from end	"yaxx"
    // Remove from end	"yax"
    // Remove from end	"ya"
    // Add 'b' to the end	"yab"
    // Add 'x' to the end	"yabx"
    fmt.Println(minOperations("axxy", "yabx")) // 6
    // Example 3:
    // Input: initial = "xyz", target = "xyz"
    // Output: 0
    // Explanation:
    // No operations are needed as the strings are already equal.
    fmt.Println(minOperations("xyz", "xyz")) // 0
}