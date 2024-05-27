package main

// 395. Longest Substring with At Least K Repeating Characters
// Given a string s and an integer k, 
// return the length of the longest substring of s such 
// that the frequency of each character in this substring is greater than or equal to k.

// if no such substring exists, return 0.

// Example 1: 
// Input: s = "aaabb", k = 3
// Output: 3
// Explanation: The longest substring is "aaa", as 'a' is repeated 3 times.

// Example 2:
// Input: s = "ababbc", k = 2
// Output: 5

// Explanation: The longest substring is "ababb", as 'a' is repeated 2 times and 'b' is repeated 3 times.

// Constraints:
//     1 <= s.length <= 10^4
//     s consists of only lowercase English letters.
//     1 <= k <= 10^5

import "fmt"

func longestSubstring(s string, k int) int {
    res, n, m, tmpIdx := 0, len(s), make(map[byte]int), 0
    for _, r := range s {
        if _, ok := m[byte(r)]; !ok {
            m[byte(r)] = tmpIdx
            tmpIdx++
        }
    }
    max := func (x, y int) int { if x > y { return x; }; return y; }
    for t := 0; t < tmpIdx; t++ {
        set :=  make([]int, tmpIdx)
        unique, fit, i, j := 0, 0, 0, 0
        for ; j < n; j++ {
            idx := m[s[j]]
            set[idx]++
            if set[idx] == 1 {
                unique++
            }
            if set[idx] == k {
                fit++
            }
            for unique > t + 1  {
                idx := m[s[i]]
                set[idx]--
                if set[idx] == 0 {
                    unique--
                }
                if set[idx] == k-1 {
                    fit--
                }
                i++
            }
            if unique == t + 1 && unique == fit {
                res = max(res, j - i + 1)
            }
        }
    }
    return res
}

func longestSubstring1(s string, k int) int {
    res, sMap, split := 0, make(map[byte]int), make(map[byte]struct{}, 0)
    for i := 0; i < len(s); i++ {
        sMap[s[i]]++
    }
    for c, num := range sMap {
        if num < k {
            split[c] = struct{}{}
        }
    }
    if len(split) == 0 {
        return len(s)
    }
    splitStr, start := make([]string, 0), 0
    for i := 0; i < len(s); i++ {
        if _, ok := split[s[i]]; ok {
            if s[start:i] != "" {
                splitStr = append(splitStr, s[start:i])
            }
                start = i + 1
        }
    }
    if s[start:] != "" {
        splitStr = append(splitStr, s[start:])
    }
    max := func (x, y int) int { if x > y { return x; }; return y; }
    for _, v := range splitStr {
        res = max(res, longestSubstring1(v, k))
    }
    return res
}

func main() {
    // Example 1: 
    // Input: s = "aaabb", k = 3
    // Output: 3
    // Explanation: The longest substring is "aaa", as 'a' is repeated 3 times.
    fmt.Println(longestSubstring("aaabb", 3)) // 3
    // Example 2:
    // Input: s = "ababbc", k = 2
    // Output: 5
    fmt.Println(longestSubstring("ababbc", 2)) // 5

    fmt.Println(longestSubstring1("aaabb", 3)) // 3
    fmt.Println(longestSubstring1("ababbc", 2)) // 5
}