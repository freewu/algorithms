package main

// 3325. Count Substrings With K-Frequency Characters I
// Given a string s and an integer k, 
// return the total number of substrings of s where at least one character appears at least k times.

// Example 1:
// Input: s = "abacb", k = 2
// Output: 4
// Explanation:
// The valid substrings are:
// "aba" (character 'a' appears 2 times).
// "abac" (character 'a' appears 2 times).
// "abacb" (character 'a' appears 2 times).
// "bacb" (character 'b' appears 2 times).

// Example 2:
// Input: s = "abcde", k = 1
// Output: 15
// Explanation:
// All substrings are valid because every character appears at least once.

// Constraints:
//     1 <= s.length <= 3000
//     1 <= k <= s.length
//     s consists only of lowercase English letters.

import "fmt"

func numberOfSubstrings(s string, k int) int {
    res, n := 0, len(s)
    check := func(freq []int, k int) bool {
        for _, v := range freq {
            if v >= k { return true }
        }
        return false
    }
    for i := 0; i < n; i++ {
        freq := make([]int, 26)
        for j := i; j < n; j++ {
            freq[s[j]-'a']++
            if check(freq, k) {
                res++
            }
        }
    }
    return res
}

func numberOfSubstrings1(s string, k int) int {
    freq := make([]int, 26)
    res, count, j := 0, 0, 0
    for i := 0; i < len(s); i++ {
        freq[s[i] - 'a']++
        if freq[s[i] - 'a'] == k {
            count++
        }
        for count >= 1 {
            if freq[s[j] - 'a'] == k {
                count--
            }
            freq[s[j] - 'a']--
            j++
        }
        res += j
    }
    return res
}

func main() {
    // Example 1:
    // Input: s = "abacb", k = 2
    // Output: 4
    // Explanation:
    // The valid substrings are:
    // "aba" (character 'a' appears 2 times).
    // "abac" (character 'a' appears 2 times).
    // "abacb" (character 'a' appears 2 times).
    // "bacb" (character 'b' appears 2 times).
    fmt.Println(numberOfSubstrings("abacb", 2)) // 4
    // Example 2:
    // Input: s = "abcde", k = 1
    // Output: 15
    // Explanation:
    // All substrings are valid because every character appears at least once.
    fmt.Println(numberOfSubstrings("abcde", 1)) // 15

    fmt.Println(numberOfSubstrings1("abacb", 2)) // 4
    fmt.Println(numberOfSubstrings1("abcde", 1)) // 15
}