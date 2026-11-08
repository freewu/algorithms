package main

// 4102. Count Valid Substrings I
// You are given a string s consisting only of the characters 'a', 'b', and 'c'.

// A string is called valid if it does not contain "abc" as a substring.

// Return the number of valid substrings in s. 
// Substrings with different starting or ending indices are counted separately.

// A substring is a contiguous non-empty sequence of characters within a string.

// Example 1:
// Input: s = "abac"
// Output: 10
// Explanation:
// The following substrings don't contain "abc"​​​​​​​ as a substring:
//     s[0..0] = "a"
//     s[0..1] = "ab"
//     s[0..2] = "aba"
//     s[0..3] = "abac"​​​​​​​
//     s[1..1] = "b"
//     s[1..2] = "ba"
//     s[1..3] = "bac"
//     s[2..2] = "a"
//     s[2..3] = "ac"
//     s[3..3] = "c"
// Thus, the answer is 10.

// Example 2:
// Input: s = "abc"
// Output: 5
// Explanation:
// The following substrings don't contain "abc"​​​​​​​ as a substring:
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
// All substrings in s don't contain "abc" as a substring​​​​​​​. Thus, the answer is 3.

// Constraints:
//     1 <= s.length <= 10^{5}​​​​​​​
//     s only consists of 'a', 'b', and 'c' characters.

import "fmt"

func countValidSubstrings(s string) int64 {
    const BAD = 'a' << 16 | 'b'<<8 | 'c'
    const MASK = 1 << 24 - 1
    res, last, hash := int64(0), int64(-1), uint32(0)
    for i := range int64(len(s)) {
        hash = hash<<8 | uint32(s[i])
        if hash & MASK == BAD {
            last = i - 2
        }
        res += i - last
    }
    return res
}

func countValidSubstrings1(s string) int64 {
    res, left := 0 , -1 // left+1 是有效子串的最小左端点
    for i := range s {
        if i >= 2 && s[i-2:i+1] == "abc" {
            left = i - 2
        }
        res += i - left
    }
    return int64(res)
}

func main() {
    // Example 1:
    // Input: s = "abac"
    // Output: 10
    // Explanation:
    // The following substrings don't contain "abc"​​​​​​​ as a substring:
    //     s[0..0] = "a"
    //     s[0..1] = "ab"
    //     s[0..2] = "aba"
    //     s[0..3] = "abac"​​​​​​​
    //     s[1..1] = "b"
    //     s[1..2] = "ba"
    //     s[1..3] = "bac"
    //     s[2..2] = "a"
    //     s[2..3] = "ac"
    //     s[3..3] = "c"
    // Thus, the answer is 10.
    fmt.Println(countValidSubstrings("abac")) // 10
    // Example 2:
    // Input: s = "abc"
    // Output: 5
    // Explanation:
    // The following substrings don't contain "abc"​​​​​​​ as a substring:
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
    // All substrings in s don't contain "abc" as a substring​​​​​​​. Thus, the answer is 3.
    fmt.Println(countValidSubstrings("ac")) // 3
    
    fmt.Println(countValidSubstrings("aaaaaaaaa")) // 25
    fmt.Println(countValidSubstrings("bbbbbbbbb")) // 45
    fmt.Println(countValidSubstrings("ccccccccc")) // 45
    fmt.Println(countValidSubstrings("abcabcabc")) // 23
    fmt.Println(countValidSubstrings("aaabbbccc")) // 45
    fmt.Println(countValidSubstrings("cbacbacba")) // 45
    fmt.Println(countValidSubstrings("cccbbbaaa")) // 45

    fmt.Println(countValidSubstrings1("abac")) // 10
    fmt.Println(countValidSubstrings1("abc")) // 5
    fmt.Println(countValidSubstrings1("ac")) // 3
    fmt.Println(countValidSubstrings1("aaaaaaaaa")) // 25
    fmt.Println(countValidSubstrings1("bbbbbbbbb")) // 45
    fmt.Println(countValidSubstrings1("ccccccccc")) // 45
    fmt.Println(countValidSubstrings1("abcabcabc")) // 23
    fmt.Println(countValidSubstrings1("aaabbbccc")) // 45
    fmt.Println(countValidSubstrings1("cbacbacba")) // 45
    fmt.Println(countValidSubstrings1("cccbbbaaa")) // 45
}