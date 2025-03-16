package main

// 1963. Minimum Number of Swaps to Make the String Balanced
// You are given a 0-indexed string s of even length n. 
// The string consists of exactly n / 2 opening brackets '[' and n / 2 closing brackets ']'.

// A string is called balanced if and only if:
//     It is the empty string, or
//     It can be written as AB, where both A and B are balanced strings, or
//     It can be written as [C], where C is a balanced string.

// You may swap the brackets at any two indices any number of times.

// Return the minimum number of swaps to make s balanced.

// Example 1:
// Input: s = "][]["
// Output: 1
// Explanation: You can make the string balanced by swapping index 0 with index 3.
// The resulting string is "[[]]".

// Example 2:
// Input: s = "]]][[["
// Output: 2
// Explanation: You can do the following to make the string balanced:
// - Swap index 0 with index 4. s = "[]][][".
// - Swap index 1 with index 5. s = "[[][]]".
// The resulting string is "[[][]]".

// Example 3:
// Input: s = "[]"
// Output: 0
// Explanation: The string is already balanced.

// Constraints:
//     n == s.length
//     2 <= n <= 10^6
//     n is even.
//     s[i] is either '[' or ']'.
//     The number of opening brackets '[' equals n / 2, and the number of closing brackets ']' equals n / 2.

import "fmt"

func minSwaps(s string) int {
    close, maxClose := 0, 0
    max := func (x, y int) int { if x > y { return x; }; return y; }
    for _, c := range s {
        if c == '[' {
            close--
        } else {
            close++ // 累加闭合的数量
        }
        maxClose = max(close, maxClose)
    }
    return maxClose / 2 + maxClose % 2
}

func minSwaps1(s string) int {
    right, left, res  := 0, 0, 0
    for _, v := range s {
        if v == '[' {
            left++
        } else if right + 1 <= left {
            right++
        } else {
            res++
        }
    }
    if res == 0 {
        return 0
    }
    return (res - 1) / 2 + 1
}

func main() {
    // Example 1:
    // Input: s = "][]["
    // Output: 1
    // Explanation: You can make the string balanced by swapping index 0 with index 3.
    // The resulting string is "[[]]".
    fmt.Println(minSwaps("][][")) // 1
    // Example 2:
    // Input: s = "]]][[["
    // Output: 2
    // Explanation: You can do the following to make the string balanced:
    // - Swap index 0 with index 4. s = "[]][][".
    // - Swap index 1 with index 5. s = "[[][]]".
    // The resulting string is "[[][]]".
    fmt.Println(minSwaps("]]][[[")) // 2
    // Example 3:
    // Input: s = "[]"
    // Output: 0
    // Explanation: The string is already balanced.
    fmt.Println(minSwaps("[]")) // 0

    fmt.Println(minSwaps1("][][")) // 1
    fmt.Println(minSwaps1("]]][[[")) // 2
    fmt.Println(minSwaps1("[]")) // 0
}