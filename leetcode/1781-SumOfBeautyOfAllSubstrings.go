package main

// 1781. Sum of Beauty of All Substrings
// The beauty of a string is the difference in frequencies between the most frequent and least frequent characters.
//     For example, the beauty of "abaacc" is 3 - 1 = 2.

// Given a string s, return the sum of beauty of all of its substrings.

// Example 1:
// Input: s = "aabcb"
// Output: 5
// Explanation: The substrings with non-zero beauty are ["aab","aabc","aabcb","abcb","bcb"], each with beauty equal to 1.

// Example 2:
// Input: s = "aabcbaa"
// Output: 17

// Constraints:
//     1 <= s.length <= 500
//     s consists of only lowercase English letters.

import "fmt"

func beautySum(s string) int {
    res, n, inf := 0, len(s), 1 << 31
    for i := 0; i < n; i++ {
        freq := make(map[byte]int)
        for j := i; j < n; j++ {
            freq[s[j]]++
            mn, mx := inf, -inf
            for _, v := range freq {
                if v < mn { mn = v }
                if v > mx { mx = v
                }
            }
            res += (mx - mn)
        }
    }
    return res
}

func beautySum1(s string) int {
    res, n, inf := 0, len(s), 1 << 31
    min := func (x, y int) int { if x < y { return x; }; return y; }
    max := func (x, y int) int { if x > y { return x; }; return y; }
    for i := 0; i < n; i++ {
        freq, mn, mx := [26]int{}, 0, 0
        for j := i; j < n; j++ {
            v := s[j] - 'a'
            freq[v]++
            mx, mn = max(mx, freq[v]), inf
            for k := 0; k < 26; k++ {
                if freq[k] > 0 {
                    mn = min(mn, freq[k])
                }
            }
            res += (mx - mn)
        }
    }
    return res
}

func main() {
    // Example 1:
    // Input: s = "aabcb"
    // Output: 5
    // Explanation: The substrings with non-zero beauty are ["aab","aabc","aabcb","abcb","bcb"], each with beauty equal to 1.
    fmt.Println(beautySum("aabcb")) // 5
    // Example 2:
    // Input: s = "aabcbaa"
    // Output: 17
    fmt.Println(beautySum("aabcbaa")) // 17

    fmt.Println(beautySum1("aabcb")) // 5
    fmt.Println(beautySum1("aabcbaa")) // 17
}