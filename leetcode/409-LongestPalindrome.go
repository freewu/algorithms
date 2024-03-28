package main

// 409. Longest Palindrome
// Given a string s which consists of lowercase or uppercase letters, 
// return the length of the longest palindrome that can be built with those letters.

// Letters are case sensitive, for example, "Aa" is not considered a palindrome here.

// Example 1:
// Input: s = "abccccdd"
// Output: 7
// Explanation: One longest palindrome that can be built is "dccaccd", whose length is 7.

// Example 2:
// Input: s = "a"
// Output: 1
// Explanation: The longest palindrome that can be built is "a", whose length is 1.
 
// Constraints:
//     1 <= s.length <= 2000
//     s consists of lowercase and/or uppercase English letters only.

import "fmt"

func longestPalindrome(s string) int {
    m := make(map[rune]int)
    for _, x := range s {
        m[x]++
    }
    res,odd := 0, 0
    for _, v := range m {
        rem := v % 2
        res += v - rem // 只累加偶数个
        if rem == 1 { // 出现奇数个 可以有一个放中间 +1 
            odd = 1
        }
    }
    return res + odd
}

func main() {
    fmt.Println(longestPalindrome("abccccdd")) // 7
    fmt.Println(longestPalindrome("a")) // 1
}