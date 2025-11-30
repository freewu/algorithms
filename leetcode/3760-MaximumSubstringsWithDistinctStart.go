package main

// 3760. Maximum Substrings With Distinct Start
// You are given a string s consisting of lowercase English letters.

// Return an integer denoting the maximum number of substrings you can split s into such that each substring starts with a distinct character 
// (i.e., no two substrings start with the same character).

// Example 1:
// Input: s = "abab"
// Output: 2
// Explanation:
// Split "abab" into "a" and "bab".
// Each substring starts with a distinct character i.e 'a' and 'b'. Thus, the answer is 2.

// Example 2:
// Input: s = "abcd"
// Output: 4
// Explanation:
// Split "abcd" into "a", "b", "c", and "d".
// Each substring starts with a distinct character. Thus, the answer is 4.

// Example 3:
// Input: s = "aaaa"
// Output: 1
// Explanation:
// All characters in "aaaa" are 'a'.
// Only one substring can start with 'a'. Thus, the answer is 1.
 
// Constraints:
//     1 <= s.length <= 10^5
//     s consists of lowercase English letters.

import "fmt"
import "math/bits"

func maxDistinct(s string) int {
    res, mp := 0, make([]bool,26)
    for _, c := range s {
        c -= 'a'
        if !mp[c] {
            mp[c] = true
            res++
        }
    }
    return res  
}

func maxDistinct1(s string) int {
    set := 0
    for _, c := range s {
        set |= 1 << (c - 'a')
    }
    return bits.OnesCount(uint(set))
}

func main() {
    // Example 1:
    // Input: s = "abab"
    // Output: 2
    // Explanation:
    // Split "abab" into "a" and "bab".
    // Each substring starts with a distinct character i.e 'a' and 'b'. Thus, the answer is 2.
    fmt.Println(maxDistinct("abab")) // 2
    // Example 2:
    // Input: s = "abcd"
    // Output: 4
    // Explanation:
    // Split "abcd" into "a", "b", "c", and "d".
    // Each substring starts with a distinct character. Thus, the answer is 4.
    fmt.Println(maxDistinct("abcd")) // 4
    // Example 3:
    // Input: s = "aaaa"
    // Output: 1
    // Explanation:
    // All characters in "aaaa" are 'a'.
    // Only one substring can start with 'a'. Thus, the answer is 1.
    fmt.Println(maxDistinct("aaaa")) // 1

    fmt.Println(maxDistinct("leetcode")) // 6
    fmt.Println(maxDistinct("bluefrog")) // 8

    fmt.Println(maxDistinct1("abab")) // 2
    fmt.Println(maxDistinct1("abcd")) // 4
    fmt.Println(maxDistinct1("aaaa")) // 1
    fmt.Println(maxDistinct1("leetcode")) // 6
    fmt.Println(maxDistinct1("bluefrog")) // 8
}