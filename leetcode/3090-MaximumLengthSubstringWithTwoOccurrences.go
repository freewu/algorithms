package main

// 3090. Maximum Length Substring With Two Occurrences
// Given a string s, return the maximum length of a substring such that it contains at most two occurrences of each character.

// Example 1:
// Input: s = "bcbbbcba"
// Output: 4
// Explanation:
// The following substring has a length of 4 and contains at most two occurrences of each character: "bcbbbcba".

// Example 2:
// Input: s = "aaaa"
// Output: 2
// Explanation:
// The following substring has a length of 2 and contains at most two occurrences of each character: "aaaa".

// Constraints:
//     2 <= s.length <= 100
//     s consists only of lowercase English letters.

import "fmt"

// Brute force
func maximumLengthSubstring(s string) int {
    res, n := 0, len(s)
    max := func (x, y int) int { if x > y { return x; }; return y; }
    for i := 0; i < n - 1; i++ {
        mp := make([]int, 26)
        mp[s[i] - 'a'] = 1
        for j := i + 1; j < n; j++ {
            if mp[s[j] - 'a'] == 0 {
                mp[s[j] - 'a'] = 1
                res = max(res, j - i + 1)
            } else if mp[s[j] - 'a'] == 1 {
                mp[s[j] - 'a'] = 2
                res = max(res, j - i + 1 )
            } else if mp[s[j] - 'a'] == 2 {
                break
            }
        }
    }
    return res
}

// Sliding window
func maximumLengthSubstring1(s string) int {
    res, start := 2, 0
    mp := make(map[byte]int)
    max := func (x, y int) int { if x > y { return x; }; return y; }
    for i := 0; i < len(s); i++ {
        mp[s[i]]++
        for mp[s[i]] > 2 {
            mp[s[start]]--
            start++
        }
        res = max(res, i - start + 1)
    }
    return res
}

func maximumLengthSubstring2(s string) int {
    arr, mp := []byte(s), make(map[byte]int)
    res, l, r, n := -1, 0, 0, len(s)
    for r < n {
        if _, ok := mp[arr[r]]; !ok {
            mp[arr[r]] = 1
        } else {
            mp[arr[r]]++
        }
        for mp[arr[r]] >= 3 {
            mp[arr[l]]--
            l++
        }
        if r - l + 1 > res {
            res = r - l + 1
        }
        r++
    }
    return res
}

func main() {
    // Example 1:
    // Input: s = "bcbbbcba"
    // Output: 4
    // Explanation:
    // The following substring has a length of 4 and contains at most two occurrences of each character: "bcbbbcba".
    fmt.Println(maximumLengthSubstring("bcbbbcba")) // 4
    // Example 2:
    // Input: s = "aaaa"
    // Output: 2
    // Explanation:
    // The following substring has a length of 2 and contains at most two occurrences of each character: "aaaa".
    fmt.Println(maximumLengthSubstring("aaaa")) // 2

    fmt.Println(maximumLengthSubstring("bluefrog")) // 8
    fmt.Println(maximumLengthSubstring("leetcode")) // 7

    fmt.Println(maximumLengthSubstring1("bcbbbcba")) // 4
    fmt.Println(maximumLengthSubstring1("aaaa")) // 2
    fmt.Println(maximumLengthSubstring1("bluefrog")) // 8
    fmt.Println(maximumLengthSubstring1("leetcode")) // 7

    fmt.Println(maximumLengthSubstring2("bcbbbcba")) // 4
    fmt.Println(maximumLengthSubstring2("aaaa")) // 2
    fmt.Println(maximumLengthSubstring2("bluefrog")) // 8
    fmt.Println(maximumLengthSubstring2("leetcode")) // 7
}