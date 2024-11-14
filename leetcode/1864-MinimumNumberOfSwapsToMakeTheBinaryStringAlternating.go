package main

// 1864. Minimum Number of Swaps to Make the Binary String Alternating
// Given a binary string s, return the minimum number of character swaps to make it alternating, or -1 if it is impossible.

// The string is called alternating if no two adjacent characters are equal. 
// For example, the strings "010" and "1010" are alternating, while the string "0100" is not.

// Any two characters may be swapped, even if they are not adjacent.

// Example 1:
// Input: s = "111000"
// Output: 1
// Explanation: Swap positions 1 and 4: "111000" -> "101010"
// The string is now alternating.

// Example 2:
// Input: s = "010"
// Output: 0
// Explanation: The string is already alternating, no swaps are needed.

// Example 3:
// Input: s = "1110"
// Output: -1

// Constraints:
//     1 <= s.length <= 1000
//     s[i] is either '0' or '1'.

import "fmt"

func minSwaps(s string) int {
    count1 := 0
    for i := 0; i < len(s); i++ {
        if s[i] == '1' { count1++ }
    }
    abs := func(x int) int { if x < 0 { return -x; }; return x; }
    min := func (x, y int) int { if x < y { return x; }; return y; }
    count0, miss0, miss1 := len(s) - count1, 0, 0
    if abs(count1 - count0) > 1 { return -1 }
    for i := 0; i < len(s); i += 2 {
        if s[i] == '1' {
            miss0++
        } else {
            miss1++
        }
    }
    if count0 == count1 { return min(miss0, miss1) }
    if count0 <  count1 { return miss1 }
    return miss0
}

func main() {
    // Example 1:
    // Input: s = "111000"
    // Output: 1
    // Explanation: Swap positions 1 and 4: "111000" -> "101010"
    // The string is now alternating.
    fmt.Println(minSwaps("111000")) // 1
    // Example 2:
    // Input: s = "010"
    // Output: 0
    // Explanation: The string is already alternating, no swaps are needed.
    fmt.Println(minSwaps("010")) // 0
    // Example 3:
    // Input: s = "1110"
    // Output: -1
    fmt.Println(minSwaps("1110")) // -1
}