package main

// 2609. Find the Longest Balanced Substring of a Binary String
// You are given a binary string s consisting only of zeroes and ones.

// A substring of s is considered balanced if all zeroes are before ones and the number of zeroes is equal to the number of ones inside the substring. 
// Notice that the empty substring is considered a balanced substring.

// Return the length of the longest balanced substring of s.

// A substring is a contiguous sequence of characters within a string.

// Example 1:
// Input: s = "01000111"
// Output: 6
// Explanation: The longest balanced substring is "000111", which has length 6.

// Example 2:
// Input: s = "00111"
// Output: 4
// Explanation: The longest balanced substring is "0011", which has length 4. 

// Example 3:
// Input: s = "111"
// Output: 0
// Explanation: There is no balanced substring except the empty substring, so the answer is 0.

// Constraints:
//     1 <= s.length <= 50
//     '0' <= s[i] <= '1'

import "fmt"

func findTheLongestBalancedSubstring(s string) int {
    res, n := 0, len(s)
    min := func (x, y int) int { if x < y { return x; }; return y; }
    max := func (x, y int) int { if x > y { return x; }; return y; }
    for i := 0; i < n; i++ {
        j := i
        for ; j < n; j++ {
            if s[j] != '0' { break }
        }
        k := j
        for ; k < min(n, i + 2 * (j - i)); k++ {
            if s[k] != '1' { break }
        }
        res = max(res, 2 * min(j - i, k - j))
    }
    return res
}

func findTheLongestBalancedSubstring1(s string) int {
    res, zeros, ones, n := 0, 0, 0,  len(s)
    min := func (x, y int) int { if x < y { return x; }; return y; }
    max := func (x, y int) int { if x > y { return x; }; return y; }
    if n == 0 { return 0 }
    for i := 0; i < n; i++ {
        if s[i] == '0' {
            if ones > 0 {
                zeros, ones = 0, 0
            }
            zeros++
        } else {
            ones++
            res = max(res, 2 * min(zeros, ones))
        }
    }
    return res
}

func main() {
    // Example 1:
    // Input: s = "01000111"
    // Output: 6
    // Explanation: The longest balanced substring is "000111", which has length 6.
    fmt.Println(findTheLongestBalancedSubstring("01000111")) // 6
    // Example 2:
    // Input: s = "00111"
    // Output: 4
    // Explanation: The longest balanced substring is "0011", which has length 4. 
    fmt.Println(findTheLongestBalancedSubstring("00111")) // 4
    // Example 3:
    // Input: s = "111"
    // Output: 0
    // Explanation: There is no balanced substring except the empty substring, so the answer is 0.
    fmt.Println(findTheLongestBalancedSubstring("111")) // 0

    fmt.Println(findTheLongestBalancedSubstring1("01000111")) // 6
    fmt.Println(findTheLongestBalancedSubstring1("00111")) // 4
    fmt.Println(findTheLongestBalancedSubstring1("111")) // 0
}