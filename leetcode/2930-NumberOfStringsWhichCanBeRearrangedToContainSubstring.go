package main

// 2930. Number of Strings Which Can Be Rearranged to Contain Substring
// You are given an integer n.

// A string s is called good if it contains only lowercase English characters 
// and it is possible to rearrange the characters of s such that the new string contains "leet" as a substring.

// For example:
//     1. The string "lteer" is good because we can rearrange it to form "leetr" .
//     2. "letl" is not good because we cannot rearrange it to contain "leet" as a substring.

// Return the total number of good strings of length n.

// Since the answer may be large, return it modulo 10^9 + 7.

// A substring is a contiguous sequence of characters within a string.

// Example 1:
// Input: n = 4
// Output: 12
// Explanation: The 12 strings which can be rearranged to have "leet" as a substring are: "eelt", "eetl", "elet", "elte", "etel", "etle", "leet", "lete", "ltee", "teel", "tele", and "tlee".

// Example 2:
// Input: n = 10
// Output: 83943898
// Explanation: The number of strings with length 10 which can be rearranged to have "leet" as a substring is 526083947580. Hence the answer is 526083947580 % (10^9 + 7) = 83943898.

// Constraints:
//     1 <= n <= 10^5

import "fmt"

func stringCount(n int) int {
    mod := 1_000_000_007
    pow := func (a, b int) int {
        res := 1
        for ; b > 0; b >>= 1 {
            if (b & 1) == 1 {
                res = (res * a) % mod
            }
            a = ( a * a ) % mod
        }
        return res
    }
    res := (pow(26, n) - (75 + n) * pow(25, n - 1) + (72 + 2 * n) * pow(24, n - 1) - (23 + n) * pow(23, n - 1)) % mod
    return (res + mod) % mod
}

func main() {
    // Example 1:
    // Input: n = 4
    // Output: 12
    // Explanation: The 12 strings which can be rearranged to have "leet" as a substring are: "eelt", "eetl", "elet", "elte", "etel", "etle", "leet", "lete", "ltee", "teel", "tele", and "tlee".
    fmt.Println(stringCount(4)) // 12
    // Example 2:
    // Input: n = 10
    // Output: 83943898
    // Explanation: The number of strings with length 10 which can be rearranged to have "leet" as a substring is 526083947580. Hence the answer is 526083947580 % (10^9 + 7) = 83943898.
    fmt.Println(stringCount(10)) // 83943898

    fmt.Println(stringCount(1)) // 0
    fmt.Println(stringCount(2)) // 0
    fmt.Println(stringCount(3)) // 0
    fmt.Println(stringCount(8)) // 295164156
    fmt.Println(stringCount(999)) // 939688677
    fmt.Println(stringCount(1000)) // 452229772
    fmt.Println(stringCount(1024)) // 804789060
    fmt.Println(stringCount(99_999)) // 795274594
    fmt.Println(stringCount(100_000)) // 778066325
}