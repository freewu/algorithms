package main 

// 266. Palindrome Permutation
// Given a string s, return true if a permutation of the string could form a palindrome and false otherwise.

// Example 1:
// Input: s = "code"
// Output: false

// Example 2:
// Input: s = "aab"
// Output: true

// Example 3:
// Input: s = "carerac"
// Output: true

// Constraints:
//     1 <= s.length <= 5000
//     s consists of only lowercase English letters.

import "fmt"

func canPermutePalindrome(s string) bool {
    odd, m := 0, make([]int,26)
    for i := 0; i < len(s); i++ { // 记录每个字符出现频次
        m[(s[i] - 'a')]++
    }
    for _, v := range m {
        if v % 2 == 1 {
            odd++ 
            if odd > 1 { // 出现在了两个奇数,不能成为回文
                return false
            }
        }
    }
    return true
}

func main() {
    // Example 1:
    // Input: s = "code"
    // Output: false
    fmt.Println(canPermutePalindrome("code")) // false
    // Example 2:
    // Input: s = "aab"
    // Output: true
    fmt.Println(canPermutePalindrome("aab")) // true
    // Example 3:
    // Input: s = "carerac"
    // Output: true
    fmt.Println(canPermutePalindrome("carerac")) // true
}