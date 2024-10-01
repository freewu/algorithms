package main

// 1400. Construct K Palindrome Strings
// Given a string s and an integer k, 
// return true if you can use all the characters in s to construct k palindrome strings or false otherwise.

// Example 1:
// Input: s = "annabelle", k = 2
// Output: true
// Explanation: You can construct two palindromes using all characters in s.
// Some possible constructions "anna" + "elble", "anbna" + "elle", "anellena" + "b"

// Example 2:
// Input: s = "leetcode", k = 3
// Output: false
// Explanation: It is impossible to construct 3 palindromes using all the characters of s.

// Example 3:
// Input: s = "true", k = 4
// Output: true
// Explanation: The only possible solution is to put each character in a separate string.

// Constraints:
//     1 <= s.length <= 10^5
//     s consists of lowercase English letters.
//     1 <= k <= 10^5

import "fmt"

func canConstruct(s string, k int) bool {
    n := len(s)
    if n <= k { return n == k }
    mp := make(map[byte]int)
    for i := 0; i < n; i++ {
        mp[s[i]]++
    }
    odd := 0
    for _, v := range mp {
        if v % 2 == 1 { // 累加奇数
            odd++
        }
    }
    if odd > k { return false } // 奇数个数大于 k 没办法放中间
    return true
}

func canConstruct1(s string, k int) bool {
    if len(s) <= k { return len(s) == k }
    mp := [26]int{}
    for _, ch := range s {
        mp[ch - 'a']++
    }
    odd := 0
    for _, v := range mp {
        if v & 1 == 1 {
            odd++
        }
    }
    return odd <= k
}

func main() {
    // Example 1:
    // Input: s = "annabelle", k = 2
    // Output: true
    // Explanation: You can construct two palindromes using all characters in s.
    // Some possible constructions "anna" + "elble", "anbna" + "elle", "anellena" + "b"
    fmt.Println(canConstruct("annabelle", 2)) // true
    // Example 2:
    // Input: s = "leetcode", k = 3
    // Output: false
    // Explanation: It is impossible to construct 3 palindromes using all the characters of s.
    fmt.Println(canConstruct("leetcode", 3)) // false
    // Example 3:
    // Input: s = "true", k = 4
    // Output: true
    // Explanation: The only possible solution is to put each character in a separate string.
    fmt.Println(canConstruct("true", 4)) // true

    fmt.Println(canConstruct1("annabelle", 2)) // true
    fmt.Println(canConstruct1("leetcode", 3)) // false
    fmt.Println(canConstruct1("true", 4)) // true
}