package main

// 3922. Minimum Flips to Make Binary String Coherent
// You are given a binary string s.

// A string is considered coherent if it does not contain "011" or "110" as subsequences.

// In one operation, you can flip any character in s ('0' to '1' or '1' to '0').

// Return an integer denoting the minimum number of modifications required to make s coherent.

// A subsequence is a string that can be derived from another string by deleting some or no characters without changing the order of the remaining characters.

// Example 1:
// Input: s = "1010"
// Output: 1
// Explanation:
// Flip s[0] to get "0010", which contains no "011" or "110" subsequences.

// Example 2:
// Input: s = "0110"
// Output: 1
// Explanation:
// Flip s[1] to get "0010", removing all forbidden subsequences "011" and "110".

// Example 3:
// Input: s = "1000"
// Output: 0
// Explanation:
// The string already has no "011" or "110" subsequences, so no flips are needed.

// Constraints:
//     1 <= s.length <= 10^5
//     s[i] is either '0' or '1'.

import "fmt"
import "strings"

func minFlips(s string) int {
    n := len(s)
    zero := strings.Count(s, "0") // 0 的数量
    one := n - zero - 1
    if s[0] == '1' && s[n - 1] == '1' {
        one--
    }
    return min(zero, max(one, 0))   
}

func main() {
    // Example 1:
    // Input: s = "1010"
    // Output: 1
    // Explanation:
    // Flip s[0] to get "0010", which contains no "011" or "110" subsequences.
    fmt.Println(minFlips("1010")) // 1
    // Example 2:
    // Input: s = "0110"
    // Output: 1
    // Explanation:
    // Flip s[1] to get "0010", removing all forbidden subsequences "011" and "110".
    fmt.Println(minFlips("0110")) // 1  
    // Example 3:
    // Input: s = "1000"
    // Output: 0
    // Explanation:
    // The string already has no "011" or "110" subsequences, so no flips are needed.
    fmt.Println(minFlips("1000")) // 0

    fmt.Println(minFlips("0000000000")) // 0
    fmt.Println(minFlips("1111111111")) // 0
    fmt.Println(minFlips("0000011111")) // 4
    fmt.Println(minFlips("1111100000")) // 4
    fmt.Println(minFlips("1010101010")) // 4
    fmt.Println(minFlips("0101010101")) // 4
}