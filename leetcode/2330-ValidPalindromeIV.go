package main

// 2330. Valid Palindrome IV
// You are given a 0-indexed string s consisting of only lowercase English letters. 
// In one operation, you can change any character of s to any other character.

// Return true if you can make s a palindrome after performing exactly one or two operations, or return false otherwise.

// Example 1:
// Input: s = "abcdba"
// Output: true
// Explanation: One way to make s a palindrome using 1 operation is:
// - Change s[2] to 'd'. Now, s = "abddba".
// One operation could be performed to make s a palindrome so return true.

// Example 2:
// Input: s = "aa"
// Output: true
// Explanation: One way to make s a palindrome using 2 operations is:
// - Change s[0] to 'b'. Now, s = "ba".
// - Change s[1] to 'b'. Now, s = "bb".
// Two operations could be performed to make s a palindrome so return true.

// Example 3:
// Input: s = "abcdef"
// Output: false
// Explanation: It is not possible to make s a palindrome using one or two operations so return false.

// Constraints:
//     1 <= s.length <= 10^5
//     s consists only of lowercase English letters.

import "fmt"

func makePalindrome(s string) bool {
    count, n := 0, len(s)
    for i := 0; i < n / 2; i++ {
        if s[i] != s[n - i - 1] {
            count++
        }
    }
    return count <= 2 // 不能超过2个以上
}

func main() {
    // Example 1:
    // Input: s = "abcdba"
    // Output: true
    // Explanation: One way to make s a palindrome using 1 operation is:
    // - Change s[2] to 'd'. Now, s = "abddba".
    // One operation could be performed to make s a palindrome so return true.
    fmt.Println(makePalindrome("abcdba")) // true
    // Example 2:
    // Input: s = "aa"
    // Output: true
    // Explanation: One way to make s a palindrome using 2 operations is:
    // - Change s[0] to 'b'. Now, s = "ba".
    // - Change s[1] to 'b'. Now, s = "bb".
    // Two operations could be performed to make s a palindrome so return true.
    fmt.Println(makePalindrome("aa")) // true
    // Example 3:
    // Input: s = "abcdef"
    // Output: false
    // Explanation: It is not possible to make s a palindrome using one or two operations so return false.
    fmt.Println(makePalindrome("abcdef")) // false
}