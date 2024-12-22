package main

// 2272. Substring With Largest Variance
// The variance of a string is defined as the largest difference between the number of occurrences of any 2 characters present in the string. Note the two characters may or may not be the same.

// Given a string s consisting of lowercase English letters only, return the largest variance possible among all substrings of s.

// A substring is a contiguous sequence of characters within a string.

// Example 1:
// Input: s = "aababbb"
// Output: 3
// Explanation:
// All possible variances along with their respective substrings are listed below:
// - Variance 0 for substrings "a", "aa", "ab", "abab", "aababb", "ba", "b", "bb", and "bbb".
// - Variance 1 for substrings "aab", "aba", "abb", "aabab", "ababb", "aababbb", and "bab".
// - Variance 2 for substrings "aaba", "ababbb", "abbb", and "babb".
// - Variance 3 for substring "babbb".
// Since the largest possible variance is 3, we return it.

// Example 2:
// Input: s = "abcde"
// Output: 0
// Explanation:
// No letter occurs more than once in s, so the variance of every substring is 0.

// Constraints:
//     1 <= s.length <= 10^4
//     s consists of lowercase English letters.

import "fmt"

func largestVariance(s string) int {
    mp := make([]int, 26)
    for _, v := range s {
        mp[v - 'a'] = 1
    }
    mx := 0
    max := func (x, y int) int { if x > y { return x; }; return y; }
    search := func(a, b int) int {
        res, count, valid, start := 0, 0, false, false
        for _, v := range s {
            val := int(v - 'a')
            if val == a { count++  }
            if val == b {
                valid = true
                if start && count >= 0 {
                    start = false
                } else if count <= 0 {
                    start, count = true, -1
                } else {
                    count--
                }
            }
            if valid {
                res = max(res, count)
            }
        }
        return res
    }
    for a, _ := range mp {
        for b, _ := range mp {
            if a == b { continue }
            mx = max(mx, search(a, b))
        }
    }
    return mx
}

func largestVariance1(s string) int {
    res, n := 0, len(s)
    max := func (x, y int) int { if x > y { return x; }; return y; }
    for a := 'a'; a <= 'z'; a++ {
        for b := 'a'; b <= 'z'; b++ {
            if a == b { continue }
            f := [2]int{0, -n}
            for _, c := range s {
                if c == a {
                    f[0]++
                    f[1]++
                } else if c == b {
                    f[1] = max(f[1]-1, f[0]-1)
                    f[0] = 0
                }
                res = max(res, f[1])
            }
        }
    }
    return res
}

func main() {
    // Example 1:
    // Input: s = "aababbb"
    // Output: 3
    // Explanation:
    // All possible variances along with their respective substrings are listed below:
    // - Variance 0 for substrings "a", "aa", "ab", "abab", "aababb", "ba", "b", "bb", and "bbb".
    // - Variance 1 for substrings "aab", "aba", "abb", "aabab", "ababb", "aababbb", and "bab".
    // - Variance 2 for substrings "aaba", "ababbb", "abbb", and "babb".
    // - Variance 3 for substring "babbb".
    // Since the largest possible variance is 3, we return it.
    fmt.Println(largestVariance("aababbb")) // 3
    // Example 2:
    // Input: s = "abcde"
    // Output: 0
    // Explanation:
    // No letter occurs more than once in s, so the variance of every substring is 0.
    fmt.Println(largestVariance("abcde")) // 0

    fmt.Println(largestVariance1("aababbb")) // 3
    fmt.Println(largestVariance1("abcde")) // 0
}