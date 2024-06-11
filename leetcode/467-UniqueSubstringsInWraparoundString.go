package main

// 467. Unique Substrings in Wraparound String
// We define the string base to be the infinite wraparound string of "abcdefghijklmnopqrstuvwxyz", so base will look like this:
//     "...zabcdefghijklmnopqrstuvwxyzabcdefghijklmnopqrstuvwxyzabcd....".

// Given a string s, return the number of unique non-empty substrings of s are present in base.

// Example 1:
// Input: s = "a"
// Output: 1
// Explanation: Only the substring "a" of s is in base.

// Example 2:
// Input: s = "cac"
// Output: 2
// Explanation: There are two substrings ("a", "c") of s in base.

// Example 3:
// Input: s = "zab"
// Output: 6
// Explanation: There are six substrings ("z", "a", "b", "za", "ab", and "zab") of s in base.

// Constraints:
//     1 <= s.length <= 10^5
//     s consists of lowercase English letters.

import "fmt"

func findSubstringInWraproundString(s string) int {
    res, a, b, eachMaxLen := 0, 1, 0, make([]int, 26)
    eachMaxLen[int(s[0]-'a')] = 1
    for i := 1; i < len(s); i++ {
        if s[i]-s[i-1] == byte(1) || s[i-1]-s[i] == byte(25) { // 连续时，当前字符结尾子串数量 = 上一字符结尾子串数量 + 1
            b = a + 1
        } else {
            b = 1
        }
        idx := int(s[i] - 'a')
        if b > eachMaxLen[idx] {
            eachMaxLen[idx] = b
        }
        a = b
    }
    for i := 0; i < 26; i++ {
        res += eachMaxLen[i]
    }
    return res
}

func main() {
    // Example 1:
    // Input: s = "a"
    // Output: 1
    // Explanation: Only the substring "a" of s is in base.
    fmt.Println(findSubstringInWraproundString("a")) // 1
    // Example 2:
    // Input: s = "cac"
    // Output: 2
    // Explanation: There are two substrings ("a", "c") of s in base.
    fmt.Println(findSubstringInWraproundString("cac")) // 2
    // Example 3:
    // Input: s = "zab"
    // Output: 6
    // Explanation: There are six substrings ("z", "a", "b", "za", "ab", and "zab") of s in base.
    fmt.Println(findSubstringInWraproundString("zab")) // 6
}