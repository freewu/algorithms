package main

// 4000. Largest Integer With Given Digit Sum
// You are given two non-negative integers n and s.

// Return the largest integer that has at most n digits and whose sum of digits is s.
// If no such integer exists, return -1.

// Example 1:
// Input: n = 2, s = 9
// Output: 90
// Explanation:
// The largest integer with at most 2 digits that has a sum of digits of 9 is 90.

// Example 2:
// Input: n = 2, s = 19
// Output: -1
// Explanation:
// There is no integer with at most 2 digits that has a sum of digits of 19, so the answer is -1.

// Example 3:
// Input: n = 5, s = 0
// Output: 0
// Explanation:
// The only non-negative integer whose digits sum to 0 is 0.

// Constraints:
//     1 <= n <= 5
//     0 <= s <= 100

import "fmt"

func largestInteger(n int, s int) int {
    if s == 0 { return 0 }
    if s > 9 * n { return -1 }
    res := 0
    for ; n > 0; n-- {
        if s >= 9 {
            res += 9
            s -= 9
        } else {
            res += s
            s = 0
        }
        res *= 10
    }
    return res / 10
}

func largestInteger1(n int, s int) int {
    res := 0
    for n > 0 {
        res = res * 10 + min(9, s)
        s -= min(9, s)
        n--
    }
    if s > 0 {
        return -1
    }
    return res
}

func main() {
    // Example 1:
    // Input: n = 2, s = 9
    // Output: 90
    // Explanation:
    // The largest integer with at most 2 digits that has a sum of digits of 9 is 90.
    fmt.Println(largestInteger(2, 9)) // 90
    // Example 2:
    // Input: n = 2, s = 19
    // Output: -1
    // Explanation:
    // There is no integer with at most 2 digits that has a sum of digits of 19, so the answer is -1.
    fmt.Println(largestInteger(2, 19)) // -1
    // Example 3:
    // Input: n = 5, s = 0
    // Output: 0
    // Explanation:
    // The only non-negative integer whose digits sum to 0 is 0.
    fmt.Println(largestInteger(5, 0)) // 0

    fmt.Println(largestInteger(1, 0)) // 0
    fmt.Println(largestInteger(1, 100)) // -1
    fmt.Println(largestInteger(5, 0)) // 0
    fmt.Println(largestInteger(5, 100)) // -1

    fmt.Println(largestInteger1(2, 9)) // 90
    fmt.Println(largestInteger1(2, 19)) // -1
    fmt.Println(largestInteger1(5, 0)) // 0
    fmt.Println(largestInteger1(1, 0)) // 0
    fmt.Println(largestInteger1(1, 100)) // -1
    fmt.Println(largestInteger1(5, 0)) // 0
    fmt.Println(largestInteger1(5, 100)) // -1
}