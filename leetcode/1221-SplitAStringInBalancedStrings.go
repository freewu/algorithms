package main

// 1221. Split a String in Balanced Strings
// Balanced strings are those that have an equal quantity of 'L' and 'R' characters.

// Given a balanced string s, split it into some number of substrings such that:
//     Each substring is balanced.

// Return the maximum number of balanced strings you can obtain.

// Example 1:
// Input: s = "RLRRLLRLRL"
// Output: 4
// Explanation: s can be split into "RL", "RRLL", "RL", "RL", each substring contains same number of 'L' and 'R'.

// Example 2:
// Input: s = "RLRRRLLRLL"
// Output: 2
// Explanation: s can be split into "RL", "RRRLLRLL", each substring contains same number of 'L' and 'R'.
// Note that s cannot be split into "RL", "RR", "RL", "LR", "LL", because the 2nd and 5th substrings are not balanced.

// Example 3:
// Input: s = "LLLLRRRR"
// Output: 1
// Explanation: s can be split into "LLLLRRRR".

// Constraints:
//     2 <= s.length <= 1000
//     s[i] is either 'L' or 'R'.
//     s is a balanced string.

import "fmt"

func balancedStringSplit(s string) int {
    balance, res := 0, 0
    for i, ch := range s {
        if ch == 'L' { balance++ }
        if ch == 'R' { balance-- }
        if i != 0 && balance == 0 { // 能够抵消的算一次
            res++
        }
    }
    return res
}

func main() {
    // Example 1:
    // Input: s = "RLRRLLRLRL"
    // Output: 4
    // Explanation: s can be split into "RL", "RRLL", "RL", "RL", each substring contains same number of 'L' and 'R'.
    fmt.Println(balancedStringSplit("RLRRLLRLRL")) // 4
    // Example 2:
    // Input: s = "RLRRRLLRLL"
    // Output: 2
    // Explanation: s can be split into "RL", "RRRLLRLL", each substring contains same number of 'L' and 'R'.
    // Note that s cannot be split into "RL", "RR", "RL", "LR", "LL", because the 2nd and 5th substrings are not balanced.
    fmt.Println(balancedStringSplit("RLRRRLLRLL")) // 2
    // Example 3:
    // Input: s = "LLLLRRRR"
    // Output: 1
    // Explanation: s can be split into "LLLLRRRR".
    fmt.Println(balancedStringSplit("LLLLRRRR")) // 1
}