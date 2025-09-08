package main

// 3675. Minimum Operations to Transform String
// You are given a string s consisting only of lowercase English letters.

// You can perform the following operation any number of times (including zero):
//     Choose any character c in the string and replace every occurrence of c with the next lowercase letter in the English alphabet.

// Return the minimum number of operations required to transform s into a string consisting of only 'a' characters.

// Note: Consider the alphabet as circular, thus 'a' comes after 'z'.

// Example 1:
// Input: s = "yz"
// Output: 2
// Explanation:
// Change 'y' to 'z' to get "zz".
// Change 'z' to 'a' to get "aa".
// Thus, the answer is 2.

// Example 2:
// Input: s = "a"
// Output: 0
// Explanation:
// The string "a" only consists of 'a'​​​​​​​ characters. Thus, the answer is 0.

// Constraints:
//     1 <= s.length <= 5 * 10^5
//     s consists only of lowercase English letters.

import "fmt"

func minOperations(s string) int {
    res := 0
    for _, v := range s {
        if v != rune('a') {
            if 26 - int(byte(v - 'a')) > res {
                res  = 26 - int(byte(v - 'a'))
            }
        }
    }
    return res
}

func main() {
    // Example 1:
    // Input: s = "yz"
    // Output: 2
    // Explanation:
    // Change 'y' to 'z' to get "zz".
    // Change 'z' to 'a' to get "aa".
    // Thus, the answer is 2.
    fmt.Println(minOperations("yz")) // 2
    // Example 2:
    // Input: s = "a"
    // Output: 0
    // Explanation:
    // The string "a" only consists of 'a'​​​​​​​ characters. Thus, the answer is 0.
    fmt.Println(minOperations("a")) // 0

    fmt.Println(minOperations("leetcode")) // 24
    fmt.Println(minOperations("bluefrog")) // 25
}