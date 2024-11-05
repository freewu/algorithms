package main

// 1758. Minimum Changes To Make Alternating Binary String
// You are given a string s consisting only of the characters '0' and '1'. 
// In one operation, you can change any '0' to '1' or vice versa.

// The string is called alternating if no two adjacent characters are equal. 
// For example, the string "010" is alternating, while the string "0100" is not.

// Return the minimum number of operations needed to make s alternating.

// Example 1:
// Input: s = "0100"
// Output: 1
// Explanation: If you change the last character to '1', s will be "0101", which is alternating.

// Example 2:
// Input: s = "10"
// Output: 0
// Explanation: s is already alternating.

// Example 3:
// Input: s = "1111"
// Output: 2
// Explanation: You need two operations to reach "0101" or "1010".

// Constraints:
//     1 <= s.length <= 10^4
//     s[i] is either '0' or '1'.

import "fmt"

func minOperations(s string) int {
    swaps := 0
    for i := 0; i < len(s); i++ {
        if i % 2 == 0 {
            if s[i] != '0' {
                swaps++
            }
        } else {
            if s[i] != '1' {
                swaps++
            }
        }
    }
    min := func (x, y int) int { if x < y { return x; }; return y; }
    return min(swaps, len(s) - swaps)
}

func minOperations1(s string) int {
    swaps := 0
    for i := 0; i < len(s); i++ {
        if (i % 2 == 0 && s[i] != '0') || (i % 2 == 1 && s[i] != '1') {
            swaps++
        }
    }
    min := func (x, y int) int { if x < y { return x; }; return y; }
    return min(swaps, len(s) - swaps)
}

func minOperations2(s string) int {
    n, count := len(s), 0
    for i, ch := range s {
        if ch == '1' && i & 1 == 0 {
            count++
        } else if ch == '0' && i & 1 != 0 {
            count++
        }
    }
    return min(count, n - count)
}

func main() {
    // Example 1:
    // Input: s = "0100"
    // Output: 1
    // Explanation: If you change the last character to '1', s will be "0101", which is alternating.
    fmt.Println(minOperations("0100")) // 1
    // Example 2:
    // Input: s = "10"
    // Output: 0
    // Explanation: s is already alternating.
    fmt.Println(minOperations("10")) // 0
    // Example 3:
    // Input: s = "1111"
    // Output: 2
    // Explanation: You need two operations to reach "0101" or "1010".
    fmt.Println(minOperations("1111")) // 2

    fmt.Println(minOperations1("0100")) // 1
    fmt.Println(minOperations1("10")) // 0
    fmt.Println(minOperations1("1111")) // 2

    fmt.Println(minOperations2("0100")) // 1
    fmt.Println(minOperations2("10")) // 0
    fmt.Println(minOperations2("1111")) // 2
}