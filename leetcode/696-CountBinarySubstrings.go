package main

// 696. Count Binary Substrings
// Given a binary string s, return the number of non-empty substrings that have the same number of 0's and 1's, 
// and all the 0's and all the 1's in these substrings are grouped consecutively.

// Substrings that occur multiple times are counted the number of times they occur.

// Example 1:
// Input: s = "00110011"
// Output: 6
// Explanation: There are 6 substrings that have equal number of consecutive 1's and 0's: "0011", "01", "1100", "10", "0011", and "01".
// Notice that some of these substrings repeat and are counted the number of times they occur.
// Also, "00110011" is not a valid substring because all the 0's (and 1's) are not grouped together.

// Example 2:
// Input: s = "10101"
// Output: 4
// Explanation: There are 4 substrings: "10", "01", "10", "01" that have equal number of consecutive 1's and 0's.

// Constraints:
//     1 <= s.length <= 10^5
//     s[i] is either '0' or '1'.

import "fmt"

func countBinarySubstrings(s string) int {
    res := 0
    check := func(l, r int) {
        count := 1
        for l >= 1 && r < len(s) - 1 {
            if s[l-1] == s[l] && s[r+1] == s[r] { // 0011 | 1100
                count++
                l--
                r++
            } else {
                break
            }
        } 
        res += count
    }
    for i := 1; i < len(s); i++ {
        if s[i] != s[i-1] {
            check(i-1, i)
        }
    }
    return res
}

func countBinarySubstrings1(s string) int {
    count, prev, curr := 0, 0, 1
    min := func (x, y int) int { if x < y { return x; }; return y; }
    for i := 1; i < len(s); i++ {
        if s[i - 1] == s[i] {
            curr++
        } else {
            count += min(curr, prev)
            prev = curr
            curr = 1
        }
    }
    return count + min(curr, prev)
}

func main() {
    // Example 1:
    // Input: s = "00110011"
    // Output: 6
    // Explanation: There are 6 substrings that have equal number of consecutive 1's and 0's: "0011", "01", "1100", "10", "0011", and "01".
    // Notice that some of these substrings repeat and are counted the number of times they occur.
    // Also, "00110011" is not a valid substring because all the 0's (and 1's) are not grouped together.
    fmt.Println(countBinarySubstrings("00110011")) // 6
    // Example 2:
    // Input: s = "10101"
    // Output: 4
    // Explanation: There are 4 substrings: "10", "01", "10", "01" that have equal number of consecutive 1's and 0's.
    fmt.Println(countBinarySubstrings("10101")) // 4

    fmt.Println(countBinarySubstrings("1111111111")) // 0
    fmt.Println(countBinarySubstrings("0000000000")) // 0
    fmt.Println(countBinarySubstrings("1111100000")) // 5
    fmt.Println(countBinarySubstrings("0000011111")) // 5
    fmt.Println(countBinarySubstrings("1010101010")) // 9
    fmt.Println(countBinarySubstrings("0101010101")) // 9

    fmt.Println(countBinarySubstrings1("00110011")) // 6
    fmt.Println(countBinarySubstrings1("10101")) // 4
    fmt.Println(countBinarySubstrings1("1111111111")) // 0
    fmt.Println(countBinarySubstrings1("0000000000")) // 0
    fmt.Println(countBinarySubstrings1("1111100000")) // 5
    fmt.Println(countBinarySubstrings1("0000011111")) // 5
    fmt.Println(countBinarySubstrings1("1010101010")) // 9
    fmt.Println(countBinarySubstrings1("0101010101")) // 9
}