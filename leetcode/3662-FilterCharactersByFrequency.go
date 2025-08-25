package main

// 3662. Filter Characters by Frequency
// You are given a string s consisting of lowercase English letters and an integer k.

// Your task is to construct a new string that contains only those characters from s which appear fewer than k times in the entire string. 
// The order of characters in the new string must be the same as their order in s.

// Return the resulting string. If no characters qualify, return an empty string.

// Note: Every occurrence of a character that occurs fewer than k times is kept.

// Example 1:
// Input: s = "aadbbcccca", k = 3
// Output: "dbb"
// Explanation:
// Character frequencies in s:
// 'a' appears 3 times
// 'd' appears 1 time
// 'b' appears 2 times
// 'c' appears 4 times
// Only 'd' and 'b' appear fewer than 3 times. Preserving their order, the result is "dbb".

// Example 2:
// Input: s = "xyz", k = 2
// Output: "xyz"
// Explanation:
// All characters ('x', 'y', 'z') appear exactly once, which is fewer than 2. Thus the whole string is returned.

// Constraints:
//     1 <= s.length <= 100
//     s consists of lowercase English letters.
//     1 <= k <= s.length

import "fmt"

func filterCharacters(s string, k int) string {
    res, mp := make([]byte, 0), make(map[byte]int)
    for i := 0; i < len(s); i++ {
        mp[s[i]]++
    }
    for i := 0; i < len(s); i++ {
        if mp[s[i]] >= k {
            continue
        }
        res = append(res, s[i])
    }
    return string(res)
}

func main() {
    // Example 1:
    // Input: s = "aadbbcccca", k = 3
    // Output: "dbb"
    // Explanation:
    // Character frequencies in s:
    // 'a' appears 3 times
    // 'd' appears 1 time
    // 'b' appears 2 times
    // 'c' appears 4 times
    // Only 'd' and 'b' appear fewer than 3 times. Preserving their order, the result is "dbb".
    fmt.Println(filterCharacters("aadbbcccca", 3)) // dbb
    // Example 2:
    // Input: s = "xyz", k = 2
    // Output: "xyz"
    // Explanation:
    // All characters ('x', 'y', 'z') appear exactly once, which is fewer than 2. Thus the whole string is returned.
    fmt.Println(filterCharacters("xyz", 2)) // xyz

    fmt.Println(filterCharacters("bluefrog", 2)) // bluefrog
    fmt.Println(filterCharacters("leetcode", 2)) // ltcod
}
