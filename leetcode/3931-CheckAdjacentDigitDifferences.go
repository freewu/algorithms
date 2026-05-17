package main

// 3931. Check Adjacent Digit Differences
// You are given a string s consisting of digits.

// Return true if the absolute difference between every pair of adjacent digits is at most 2, otherwise return false.

// The absolute difference between a and b is defined as abs(a - b).

// Example 1:
// Input: s = "132"
// Output: true
// Explanation:
// The absolute difference between digits at s[0] and s[1] is abs(1 - 3) = 2.
// The absolute difference between digits at s[1] and s[2] is abs(3 - 2) = 1.
// Since both differences are at most 2, the answer is true.

// Example 2:
// Input: s = "129"
// Output: false
// Explanation:
// The absolute difference between digits at s[0] and s[1] is abs(1 - 2) = 1.
// The absolute difference between digits at s[1] and s[2] is abs(2 - 9) = 7, which is greater than 2.
// Therefore, the answer is false.

// Constraints:
//     2 <= s.length <= 100
//     s consists only of digits.

import "fmt"

func isAdjacentDiffAtMostTwo(s string) bool {
    abs := func(x int) int { if x < 0 { return -x; }; return x; }
    for i := 0; i < len(s)-1; i++ {
        if abs(int(s[i]) - int(s[i+1])) > 2 {
            return false
        }
    }
    return true
}

func main() {
    // Example 1:
    // Input: s = "132"
    // Output: true
    // Explanation:
    // The absolute difference between digits at s[0] and s[1] is abs(1 - 3) = 2.
    // The absolute difference between digits at s[1] and s[2] is abs(3 - 2) = 1.
    // Since both differences are at most 2, the answer is true.
    fmt.Println(isAdjacentDiffAtMostTwo("132")) // true
    // Example 2:
    // Input: s = "129"
    // Output: false
    // Explanation:
    // The absolute difference between digits at s[0] and s[1] is abs(1 - 2) = 1.
    // The absolute difference between digits at s[1] and s[2] is abs(2 - 9) = 7, which is greater than 2.
    // Therefore, the answer is false.
    fmt.Println(isAdjacentDiffAtMostTwo("129")) // false

    fmt.Println(isAdjacentDiffAtMostTwo("123456789")) // false
    fmt.Println(isAdjacentDiffAtMostTwo("987654321")) // false
}