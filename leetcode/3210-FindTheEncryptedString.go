package main

// 3210. Find the Encrypted String
// You are given a string s and an integer k. 
// Encrypt the string using the following algorithm:
//     For each character c in s, replace c with the kth character after c in the string (in a cyclic manner).

// Return the encrypted string.

// Example 1:
// Input: s = "dart", k = 3
// Output: "tdar"
// Explanation:
// For i = 0, the 3rd character after 'd' is 't'.
// For i = 1, the 3rd character after 'a' is 'd'.
// For i = 2, the 3rd character after 'r' is 'a'.
// For i = 3, the 3rd character after 't' is 'r'.

// Example 2:
// Input: s = "aaa", k = 1
// Output: "aaa"
// Explanation:
// As all the characters are the same, the encrypted string will also be the same.

// Constraints:
//     1 <= s.length <= 100
//     1 <= k <= 10^4
//     s consists only of lowercase English letters.

import "fmt"

func getEncryptedString(s string, k int) string {
    n := len(s)
    bytes := make([]byte, n)
    k = k % n
    for i := 0; i < n; i++ {
        bytes[i] = s[(i + k) % n]
    }
    return string(bytes)
}

func getEncryptedString1(s string, k int) string {
    res := []byte{}
    for i := range s {
        res = append(res, s[(i + k) % len(s)])
    }
    return string(res)
}

func main() {
    // Example 1:
    // Input: s = "dart", k = 3
    // Output: "tdar"
    // Explanation:
    // For i = 0, the 3rd character after 'd' is 't'.
    // For i = 1, the 3rd character after 'a' is 'd'.
    // For i = 2, the 3rd character after 'r' is 'a'.
    // For i = 3, the 3rd character after 't' is 'r'.
    fmt.Println(getEncryptedString("dart", 3)) // "tdar"
    // Example 2:
    // Input: s = "aaa", k = 1
    // Output: "aaa"
    // Explanation:
    // As all the characters are the same, the encrypted string will also be the same.
    fmt.Println(getEncryptedString("aaa", 1)) // "aaa"

    fmt.Println(getEncryptedString("bluefrog", 3)) // "efrogblu"
    fmt.Println(getEncryptedString("leetcode", 3)) // "tcodelee"

    fmt.Println(getEncryptedString1("dart", 3)) // "tdar"
    fmt.Println(getEncryptedString1("aaa", 1)) // "aaa"
    fmt.Println(getEncryptedString1("bluefrog", 3)) // "efrogblu"
    fmt.Println(getEncryptedString1("leetcode", 3)) // "tcodelee"
}