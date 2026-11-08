package main

// 4104. Count Valid Substrings II
// You are given a string s consisting only of the characters 'a', 'b', and 'c'.

// A string is called valid if it does not contain "abc" as a subsequence.

// Return the number of valid substrings in s. 
// Substrings with different starting or ending indices are counted separately.

// A substring is a contiguous non-empty sequence of characters within a string.

// A subsequence is a string that can be derived from another string by deleting some or no characters without changing the order of the remaining characters.

// Example 1:
// Input: s = "abac"
// Output: 9
// Explanation:
// The following substrings don't contain "abc"​​​​​​​ as a subsequence:
//     s[0..0] = "a"
//     s[0..1] = "ab"
//     s[0..2] = "aba"
//     s[1..1] = "b"
//     s[1..2] = "ba"
//     s[1..3] = "bac"
//     s[2..2] = "a"
//     s[2..3] = "ac"
//     s[3..3] = "c"
// Thus, the answer is 9.

// Example 2:
// Input: s = "abc"
// Output: 5
// Explanation:
// The following substrings don't contain "abc"​​​​​​​ as a subsequence:
//     s[0..0] = "a"
//     s[0..1] = "ab"
//     s[1..1] = "b"
//     s[1..2] = "bc"
//     s[2..2] = "c"​​​​​​​
// Thus, the answer is 5.

// Example 3:
// Input: s = "ac"
// Output: 3
// Explanation:
// All substrings in s don't contain "abc"​​​​​​​ as a subsequence. Thus, the answer is 3.

// Constraints:
//     1 <= s.length <= 10^{5}​​​​​​​
//     s only consists of 'a', 'b', and 'c' characters.

import "fmt"

func countValidSubstrings(s string) int64 {
    res, n := int64(0), int64(len(s))
    state := [4]int64{n, n, n, 0}
    for i := n - 1; i >= 0; i-- {
        state[3] = i
        c := s[i] - 'a'
        state[c] = state[c+1]
        res += state[0] - i
    }
    return res
}

func countValidSubstrings1(s string) int64 {
    res, empty, a, b := int64(0), int64(0), int64(0), int64(0)
    for _, v := range s {
        if v == 'a' {
            a += empty + 1
            empty = 0
        } else if v == 'b' {
            b += a
            empty++
            a = 0
        } else if v == 'c' {
            empty++
            b = 0
        }
        res += empty + a + b
    }
    return res
}

func main() {
    // Example 1:
    // Input: s = "abac"
    // Output: 9
    // Explanation:
    // The following substrings don't contain "abc"​​​​​​​ as a subsequence:
    //     s[0..0] = "a"
    //     s[0..1] = "ab"
    //     s[0..2] = "aba"
    //     s[1..1] = "b"
    //     s[1..2] = "ba"
    //     s[1..3] = "bac"
    //     s[2..2] = "a"
    //     s[2..3] = "ac"
    //     s[3..3] = "c"
    // Thus, the answer is 9.
    fmt.Println(countValidSubstrings("abac")) // 9
    // Example 2:
    // Input: s = "abc"
    // Output: 5
    // Explanation:
    // The following substrings don't contain "abc"​​​​​​​ as a subsequence:
    //     s[0..0] = "a"
    //     s[0..1] = "ab"
    //     s[1..1] = "b"
    //     s[1..2] = "bc"
    //     s[2..2] = "c"​​​​​​​
    // Thus, the answer is 5.
    fmt.Println(countValidSubstrings("abc")) // 5
    // Example 3:
    // Input: s = "ac"
    // Output: 3
    // Explanation:
    // All substrings in s don't contain "abc"​​​​​​​ as a subsequence. Thus, the answer is 3.
    fmt.Println(countValidSubstrings("ac")) // 3
    
    fmt.Println(countValidSubstrings("aaaaaaaaa")) // 25
    fmt.Println(countValidSubstrings("bbbbbbbbb")) // 45
    fmt.Println(countValidSubstrings("ccccccccc")) // 45
    fmt.Println(countValidSubstrings("abcabcabc")) // 23
    fmt.Println(countValidSubstrings("aaabbbccc")) // 36
    fmt.Println(countValidSubstrings("cbacbacba")) // 36
    fmt.Println(countValidSubstrings("cccbbbaaa")) // 45

    fmt.Println(countValidSubstrings1("abac")) // 9
    fmt.Println(countValidSubstrings1("abc")) // 5
    fmt.Println(countValidSubstrings1("ac")) // 3
    fmt.Println(countValidSubstrings1("aaaaaaaaa")) // 25
    fmt.Println(countValidSubstrings1("bbbbbbbbb")) // 45
    fmt.Println(countValidSubstrings1("ccccccccc")) // 45
    fmt.Println(countValidSubstrings1("abcabcabc")) // 23
    fmt.Println(countValidSubstrings1("aaabbbccc")) // 36
    fmt.Println(countValidSubstrings1("cbacbacba")) // 36
    fmt.Println(countValidSubstrings1("cccbbbaaa")) // 45
}