package main

// 1328. Break a Palindrome
// Given a palindromic string of lowercase English letters palindrome, 
// replace exactly one character with any lowercase English letter so that the resulting string is not a palindrome and that it is the lexicographically smallest one possible.

// Return the resulting string. If there is no way to replace a character to make it not a palindrome, 
// return an empty string.

// A string a is lexicographically smaller than a string b (of the same length) if in the first position where a and b differ,
// a has a character strictly smaller than the corresponding character in b. 
// For example, "abcc" is lexicographically smaller than "abcd" because the first position they differ is at the fourth character, 
// and 'c' is smaller than 'd'.

// Example 1:
// Input: palindrome = "abccba"
// Output: "aaccba"
// Explanation: There are many ways to make "abccba" not a palindrome, such as "zbccba", "aaccba", and "abacba".
// Of all the ways, "aaccba" is the lexicographically smallest.

// Example 2:
// Input: palindrome = "a"
// Output: ""
// Explanation: There is no way to replace a single character to make "a" not a palindrome, so return an empty string.

// Constraints:
//     1 <= palindrome.length <= 1000
//     palindrome consists of only lowercase English letters.

import "fmt"

func breakPalindrome(palindrome string) string {
    n := len(palindrome)
    if n <= 1 { return "" } // 一个的怎么变都是回文
    arr := []byte( palindrome )
    for i := 0; i < n / 2; i++ {
        if arr[i] != 'a' { // 遇上第一个不是 a 的，改成 a 即可
            arr[i] = 'a'
            return string(arr)
        }
    }
    arr[n-1] = 'b' // 处理全是 a 的情况 把最后一个改为 b 即可
    return string(arr)
}

func main() {
    // Example 1:
    // Input: palindrome = "abccba"
    // Output: "aaccba"
    // Explanation: There are many ways to make "abccba" not a palindrome, such as "zbccba", "aaccba", and "abacba".
    // Of all the ways, "aaccba" is the lexicographically smallest.
    fmt.Println(breakPalindrome("abccba")) // aaccba
    // Example 2:
    // Input: palindrome = "a"
    // Output: ""
    // Explanation: There is no way to replace a single character to make "a" not a palindrome, so return an empty string.
    fmt.Println(breakPalindrome("a")) // ""

    fmt.Println(breakPalindrome("aaaa")) // "aaab"
}