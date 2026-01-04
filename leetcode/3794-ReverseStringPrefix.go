package main

// 3794. Reverse String Prefix
// You are given a string s and an integer k.

// Reverse the first k characters of s and return the resulting string.

// Example 1:
// Input: s = "abcd", k = 2
// Output: "bacd"
// Explanation:​​​​​​​
// The first k = 2 characters "ab" are reversed to "ba". The final resulting string is "bacd".

// Example 2:
// Input: s = "xyz", k = 3
// Output: "zyx"
// Explanation:
// The first k = 3 characters "xyz" are reversed to "zyx". The final resulting string is "zyx".

// Example 3:
// Input: s = "hey", k = 1
// Output: "hey"
// Explanation:
// The first k = 1 character "h" remains unchanged on reversal. The final resulting string is "hey".

// Constraints:
//     1 <= s.length <= 100
//     s consists of lowercase English letters.
//     1 <= k <= s.length

import "fmt"

func reversePrefix(s string, k int) string {
    res, arr := "", s[:k]
    for i := len(arr) - 1; i >= 0; i-- {
        res += string(arr[i])
    }
    for i := k; i < len(s); i++ {
        res += string(s[i])
    }
    return res
}

func reversePrefix1(s string, k int) string {
    res := make([]byte,len(s))
    for i := k-1; i >= 0; i-- {
        res[k-1-i] = s[i]
    }
    for i := k; i < len(s); i++ {
        res[i] = s[i]
    }
    return string(res)
}

func main() {
    // Example 1:
    // Input: s = "abcd", k = 2
    // Output: "bacd"
    // Explanation:​​​​​​​
    // The first k = 2 characters "ab" are reversed to "ba". The final resulting string is "bacd".
    fmt.Println(reversePrefix("abcd", 2)) // "bacd"
    // Example 2:
    // Input: s = "xyz", k = 3
    // Output: "zyx"
    // Explanation:
    // The first k = 3 characters "xyz" are reversed to "zyx". The final resulting string is "zyx".
    fmt.Println(reversePrefix("xyz", 3)) // "zyx"
    // Example 3:
    // Input: s = "hey", k = 1
    // Output: "hey"
    // Explanation:
    // The first k = 1 character "h" remains unchanged on reversal. The final resulting string is "hey".
    fmt.Println(reversePrefix("hey", 1)) // "hey"

    fmt.Println(reversePrefix("bluefrog", 2)) // "lbuefrog"
    fmt.Println(reversePrefix("leetcode", 2)) // "eletcode"

    fmt.Println(reversePrefix1("abcd", 2)) // "bacd"
    fmt.Println(reversePrefix1("xyz", 3)) // "zyx"
    fmt.Println(reversePrefix1("hey", 1)) // "hey"
    fmt.Println(reversePrefix1("bluefrog", 2)) // "lbuefrog"
    fmt.Println(reversePrefix1("leetcode", 2)) // "eletcode"
}