package main

// 3714. Longest Balanced Substring II
// You are given a string s consisting only of the characters 'a', 'b', and 'c'.

// A substring of s is called balanced if all distinct characters in the substring appear the same number of times.

// Return the length of the longest balanced substring of s.

// Example 1:
// Input: s = "abbac"
// Output: 4
// Explanation:
// The longest balanced substring is "abba" because both distinct characters 'a' and 'b' each appear exactly 2 times.

// Example 2:
// Input: s = "aabcc"
// Output: 3
// Explanation:
// The longest balanced substring is "abc" because all distinct characters 'a', 'b' and 'c' each appear exactly 1 time.

// Example 3:
// Input: s = "aba"
// Output: 2
// Explanation:
// One of the longest balanced substrings is "ab" because both distinct characters 'a' and 'b' each appear exactly 1 time. Another longest balanced substring is "ba".

// Constraints:
//     1 <= s.length <= 10^5
//     s contains only the characters 'a', 'b', and 'c'.

import "fmt"

func longestBalanced(s string) int {
    res, n := 0, len(s)
    max := func (x, y int) int { if x > y { return x; }; return y; } 
    // 一种字母
    for i := 0; i < n; {
        start := i
        for i++; i < n && s[i] == s[i-1]; i++ {
        }
        res = max(res, i - start)
    }
    // 两种字母
    f := func(x, y byte) {
        for i := 0; i < n; i++ {
            // 前缀和数组的首项是 0，位置相当于在 i-1
            pos := map[int]int{0: i - 1}
            d := 0 // x 的个数减去 y 的个数
            for ; i < n && (s[i] == x || s[i] == y); i++ {
                if s[i] == x {
                    d++
                } else {
                    d--
                }
                if j, ok := pos[d]; ok {
                    res = max(res, i-j)
                } else {
                    pos[d] = i
                }
            }
        }
    }
    f('a', 'b')
    f('a', 'c')
    f('b', 'c')
    // 三种字母
    type Pair struct{ diffAB, diffBC int }
    // 前缀和数组的首项是 0，位置相当于在 -1
    pos, count := map[Pair]int{{}: -1}, [3]int{}
    for i, b := range s {
        count[b-'a']++
        p := Pair{count[0] - count[1], count[1] - count[2]}
        if j, ok := pos[p]; ok {
            res = max(res, i-j) 
        } else {
            pos[p] = i
        }
    }
    return res
}

func main() {
    // Example 1:
    // Input: s = "abbac"
    // Output: 4
    // Explanation:
    // The longest balanced substring is "abba" because both distinct characters 'a' and 'b' each appear exactly 2 times.
    fmt.Println(longestBalanced("abbac")) // 4
    // Example 2:
    // Input: s = "aabcc"
    // Output: 3
    // Explanation:
    // The longest balanced substring is "abc" because all distinct characters 'a', 'b' and 'c' each appear exactly 1 time.
    fmt.Println(longestBalanced("aabcc")) // 3
    // Example 3:
    // Input: s = "aba"
    // Output: 2
    // Explanation:
    // One of the longest balanced substrings is "ab" because both distinct characters 'a' and 'b' each appear exactly 1 time. Another longest balanced substring is "ba".
    fmt.Println(longestBalanced("aba")) // 2

    fmt.Println(longestBalanced("aaaaaaaaa")) // 9
    fmt.Println(longestBalanced("bbbbbbbbb")) // 9
    fmt.Println(longestBalanced("ccccccccc")) // 9
    fmt.Println(longestBalanced("abcabcabc")) // 9
    fmt.Println(longestBalanced("aaabbbccc")) // 9
    fmt.Println(longestBalanced("cccbbbaaa")) // 9
    fmt.Println(longestBalanced("bbbcccaaa")) // 9
}